package getuserhandler

import (
	"fmt"
	getuser "gin_go_learn/internal/controllers/get_user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getuser.Service
}

func NewGetUserHandler(service getuser.Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) HandleGetUser(c *gin.Context) {
	idparam := c.Param("id")
	idparamtoint, err := strconv.Atoi(idparam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if idparam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "parameter id required",
		})
		return
	}
	user, err := h.service.GetUserService(idparamtoint)
	fmt.Println(user)

}
