package main

import (
	"gin_go_learn/config"
	"gin_go_learn/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.InitializeENV()
}

func main() {
	port := os.Getenv("PORT")
	gin.SetMode(os.Getenv("MODE"))
	app := SetupRouter()
	err := app.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to bind ", err)
	}

}

func SetupRouter() *gin.Engine {
	db := config.ConnectDB()
	router := gin.Default()
	routes.InitUserRoutes(db, router)
	return router
}
