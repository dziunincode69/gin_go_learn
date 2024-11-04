package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	gin.SetMode(os.Getenv("MODE"))
	app := gin.New()
	app.GET("/", func(c *gin.Context) {
		ua := c.GetHeader("User-Agent")
		if ua != "insomnia/10.0.0" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Access Denied"})
			return
		}
		resp := gin.H{"message": "Hello World"}
		c.JSON(http.StatusOK, resp)

	})
	err := app.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to bind ", err)
	}

}
