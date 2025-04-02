package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 定义一个简单的模型
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
}

func main() {
	// 数据库连接
	dsn := "host=localhost user=dlx_admin password=fuckyou365 dbname=mis port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	log.Println("数据库连接成功")

	// 自动迁移模型
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
	log.Println("数据库迁移成功")

	// 使用 gin.New() 创建没有默认中间件的实例
	r := gin.New()

	// 手动添加 Logger 和 Recovery 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 启动服务
	r.Run(":8080") // 默认监听 8080 端口
}
