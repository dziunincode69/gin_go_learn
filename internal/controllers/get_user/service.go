package getuser

import "gin_go_learn/internal/models"

type Service interface {
	GetUserService(id int) (*models.User, error)
}

type service struct {
	repository Repository
}

func NewGetUserService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetUserService(id int) (*models.User, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
