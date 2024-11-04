package register

import "gin_go_learn/internal/models"

type Service interface {
	RegisterService(input *RegisterInput) (*models.User, error)
}

type service struct {
	repository Repository
}

func NewRegisterService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) RegisterService(input *RegisterInput) (*models.User, error) {
	userdata := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	user, err := s.repository.RegisterRepository(userdata)
	if err != nil {
		return nil, err
	}
	return user, nil
}
