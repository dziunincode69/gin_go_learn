package routes

import (
	getuser "gin_go_learn/internal/controllers/get_user"
	"gin_go_learn/internal/controllers/login"
	"gin_go_learn/internal/controllers/register"
	loginhandler "gin_go_learn/internal/handlers/login"
	registerhandler "gin_go_learn/internal/handlers/register"
	getuserhandler "gin_go_learn/internal/handlers/user"
	"gin_go_learn/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, router *gin.Engine) {
	LoginRepository := login.NewLoginRepository(db)
	LoginService := login.NewLoginService(LoginRepository)
	LoginHandler := loginhandler.NewLoginHandler(LoginService)

	GetUserRepository := getuser.NewGetUserRepository(db)
	GetUserService := getuser.NewGetUserService(GetUserRepository)
	GetUserHandler := getuserhandler.NewGetUserHandler(GetUserService)

	RegisterUserRepository := register.NewRegisterRepository(db)
	RegisterUserService := register.NewRegisterService(RegisterUserRepository)
	RegisterUserHandler := registerhandler.NewRegisterHandler(RegisterUserService)

	apigroup := router.Group("/api")
	apigroup.POST("/login", LoginHandler.HandleLogin)
	apigroup.POST("/register", RegisterUserHandler.RegisterHandler)
	protected := apigroup.Use(middlewares.JwtAuthMiddleware(), GetUserHandler.HandleGetUser)
	protected.GET("/user/:id")
}
