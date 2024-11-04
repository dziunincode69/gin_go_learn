package routes

import (
	getuser "gin_go_learn/internal/controllers/get_user"
	"gin_go_learn/internal/controllers/login"
	"gin_go_learn/internal/controllers/register"
	loginhandler "gin_go_learn/internal/handlers/login"
	registerhandler "gin_go_learn/internal/handlers/register"
	getuserhandler "gin_go_learn/internal/handlers/user"
	"gin_go_learn/internal/helper"
	"gin_go_learn/internal/middlewares"
	"gin_go_learn/internal/models"
	"net/http"
	"strconv"

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
	protected := apigroup.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user/:id", GetUserHandler.HandleGetUser)
	protected.DELETE("/user/:id", func(c *gin.Context) {
		idparam := c.Param("id")
		idparamint, err := strconv.Atoi(idparam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		val, exist := c.Get("user")
		if !exist {
			c.JSON(400, gin.H{
				"error": "data not exist",
			})
			return
		}
		valtostruct := val.(*helper.MyAppClaims)
		if idparamint == valtostruct.ID {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "cannot delete your account",
			})
			return
		}
		var user models.User
		db.Unscoped().Delete(&user, idparamint)
		c.JSON(200, gin.H{
			"status": "success",
		})
	})

}
