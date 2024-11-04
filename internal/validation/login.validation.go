package validation

import (
	logincontroller "gin_go_learn/internal/controllers/login"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func LoginInputValidation(c *gin.Context) (*logincontroller.LoginInput, error) {
	var logininput logincontroller.LoginInput
	validate := validator.New()
	err := c.ShouldBindJSON(&logininput)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(&logininput)
	if err != nil {
		return nil, err
	}
	return &logininput, nil
}
