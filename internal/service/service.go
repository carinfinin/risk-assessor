package service

import "github.com/carinfinin/risk-assessor/internal/model"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateUser(clientData model.ClientData) (model.User, error) {
	return model.User{}, nil
}
