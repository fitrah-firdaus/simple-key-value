package keyvalue

import "simple-key-value/pkg/entities"

type Service interface {
	CreateOrUpdateKey(value *entities.KeyValue) (*entities.KeyValue, error)
	GetKey(key string) (*entities.KeyValue, error)
	DeleteKey(key string) error
}

type service struct {
	repository Repository
}

func (s *service) CreateOrUpdateKey(value *entities.KeyValue) (*entities.KeyValue, error) {
	return s.repository.CreateOrUpdateKey(value)
}

func (s *service) GetKey(key string) (*entities.KeyValue, error) {
	return s.repository.GetKey(key)
}

func (s *service) DeleteKey(key string) error {
	return s.repository.DeleteKey(key)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
