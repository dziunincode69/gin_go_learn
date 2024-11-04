package routes

import (
	loginhandler "gin_go_learn/handlers/login"
	"gin_go_learn/internal/controllers/login"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, router *gin.Engine) {

	LoginRepository := login.NewLoginRepository(db)
	LoginService := login.NewLoginService(LoginRepository)
	LoginHandler := loginhandler.NewLoginHandler(LoginService)

	apigroup := router.Group("/api")
	apigroup.POST("/login", LoginHandler.HandleLogin)
}
