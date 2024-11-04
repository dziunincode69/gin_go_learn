package login

import (
	"errors"
	"gin_go_learn/internal/helper"
	"gin_go_learn/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(input *models.User) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) LoginRepository(input *models.User) (*models.User, error) {
	var user models.User
	err := r.db.Model(&user).Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("USER_NOT_FOUND")
		}
		return nil, err
	}
	hashedpwd := user.Password
	checkpwd := helper.CheckPasswordHash(input.Password, hashedpwd)
	if !checkpwd {
		return nil, errors.New("WRONG_PASSWORD")
	}
	return &user, nil
}
