package getuser

import (
	"errors"
	"gin_go_learn/internal/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetUser(id int) (*models.User, error)
}
type repository struct {
	db *gorm.DB
}

func NewGetUserRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetUser(id int) (*models.User, error) {
	var user models.User
	err := r.db.Model(&user).Omit("password").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("USER_NOT_FOUND")
		}
		return nil, err
	}
	return &user, nil
}
