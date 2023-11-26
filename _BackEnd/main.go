package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/generate-timestamp", func(c *gin.Context) {
		timestamp := generateRandomTimestamp()
		print(timestamp)
		c.JSON(http.StatusOK, gin.H{"timestamp": timestamp})
	})

	r.Run(":8080") // 启动Gin服务器，监听端口8080
}

func generateRandomTimestamp() int64 {
	rand.Seed(time.Now().UnixNano())
	min := time.Date(2000, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	return rand.Int63n(max-min) + min
}
