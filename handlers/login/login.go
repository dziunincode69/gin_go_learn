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
	loginservice, err := h.service.LoginService(logininput)
	if err != nil {
		errmsg := err.Error()
		switch errmsg {
		case "USER_NOT_FOUND":
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errmsg,
			})
			return
		case "WRONG_PASSWORD":
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": errmsg,
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errmsg,
			})
			return
		}
	}
	data := map[string]any{
		"email":    loginservice.Email,
		"id":       loginservice.ID,
		"is_admin": loginservice.IsAdmin,
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
