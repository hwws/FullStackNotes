// cmd/microservice1/main.go
package main

import (
	"Reps/internal/auth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 路由定义
	r.POST("/login", auth.LoginHandler)

	// 启动服务
	fmt.Println("Microservice1 is running on :8080...")
	r.Run(":8080")
}
