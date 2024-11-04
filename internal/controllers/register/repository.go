package register

import (
	"errors"
	"gin_go_learn/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *models.User) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) RegisterRepository(input *models.User) (*models.User, error) {
	var user models.User
	input.IsAdmin = false
	err := r.db.Where("email = ?", input.Email).FirstOrCreate(&user, input).Error
	if err != nil {
		return nil, errors.New("email already registered")
	}
	return &user, nil
}
