package validation

import (
	"gin_go_learn/internal/controllers/register"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RegisterInputValidation(c *gin.Context) (*register.RegisterInput, error) {
	var registerinput register.RegisterInput
	validate := validator.New()
	err := c.ShouldBindJSON(&registerinput)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(&registerinput)
	if err != nil {
		return nil, err
	}
	return &registerinput, nil
}
