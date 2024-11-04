package loginhandler

import (
	logincontroller "gin_go_learn/internal/controllers/login"
	"gin_go_learn/internal/helper"
	"gin_go_learn/internal/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service logincontroller.Service
}

func NewLoginHandler(service logincontroller.Service) *handler {
	return &handler{
		service: service,
	}

}

func (h *handler) HandleLogin(c *gin.Context) {
	logininput, err := validation.LoginInputValidation(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	loginservice, errLogin := h.service.LoginService(logininput)
	if errLogin != "" {
		switch errLogin {
		case "USER_NOT_FOUND":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errLogin,
			})
			return
		case "DATABASE_ERROR":
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errLogin,
			})
			return
		case "WRONG_PASSWORD":
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": errLogin,
			})
			return
		}
	}
	data := map[string]any{
		"email": loginservice.Email,
		"id":    loginservice.ID,
	}
	tokenjwt, err := helper.NewSign(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenjwt,
	})

}
