package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"
const MODE = gin.DebugMode

func main() {
	gin.SetMode(MODE)
	app := gin.New()
	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World from gin")
	})
	app.Run(":" + PORT)

}
