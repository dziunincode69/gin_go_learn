package middlewares

import (
	"gin_go_learn/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token := c.GetHeader("access_token")
		if access_token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		claims, err := helper.ParseToken(access_token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		check, err := helper.CheckToken(claims)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		if !check.IsAdmin {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "only admin can access this page",
			})
			c.Abort()
			return
		}
		c.Set("user", check)
		c.Next()
	}
}
