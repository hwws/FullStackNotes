// internal/auth/handler.go
package auth

import (
	"Reps/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// 在这里添加与数据库或其他身份验证逻辑的集成
	// 例如，检查用户名和密码是否有效

	// 假设有一个名为 "validUser" 的用户
	validUser := "validUser"
	validPasswordHash := "$2a$14$8qUd52Sk5v7c4gkUH7jkq.BVWVdOfz64tDWVhfojCWsTWr9b7eLw6"

	if creds.Username == validUser && CheckPasswordHash(creds.Password, validPasswordHash) {
		// 用户身份验证成功
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Welcome, %s!", creds.Username)})
	} else {
		// 用户身份验证失败
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
