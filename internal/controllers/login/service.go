package login

import "gin_go_learn/internal/models"

type Service interface {
	LoginService(input *LoginInput) (*models.User, error)
}

type service struct {
	repository Repository
}

func NewLoginService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) LoginService(input *LoginInput) (*models.User, error) {
	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}
	check, err := s.repository.LoginRepository(&user)
	return check, err
}
