package registerhandler

import (
	"gin_go_learn/internal/controllers/register"
	"gin_go_learn/internal/helper"
	"gin_go_learn/internal/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service register.Service
}

func NewRegisterHandler(service register.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) RegisterHandler(c *gin.Context) {
	registerinput, err := validation.RegisterInputValidation(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	registerservice, err := h.service.RegisterService(registerinput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := map[string]any{
		"email":    registerservice.Email,
		"id":       registerservice.ID,
		"is_admin": registerservice.IsAdmin,
	}
	tokenjwt, err := helper.NewSign(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":       "success",
		"id":           registerservice.ID,
		"email":        registerservice.Email,
		"access_token": tokenjwt,
	})

}
