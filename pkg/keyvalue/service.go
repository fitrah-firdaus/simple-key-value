package keyvalue

import "simple-key-value/pkg/entities"

type Service interface {
	CreateKey(value *entities.KeyValue) (*entities.KeyValue, error)
	GetKey(key string) (*entities.KeyValue, error)
	UpdateKey(value *entities.KeyValue) (*entities.KeyValue, error)
	DeleteKey(key string) error
}

type service struct {
	repository Repository
}

func (s *service) CreateKey(value *entities.KeyValue) (*entities.KeyValue, error) {
	return s.repository.CreateKey(value)
}

func (s *service) GetKey(key string) (*entities.KeyValue, error) {
	return s.repository.GetKey(key)
}

func (s *service) UpdateKey(value *entities.KeyValue) (*entities.KeyValue, error) {
	return s.repository.UpdateKey(value)
}

func (s *service) DeleteKey(key string) error {
	return s.repository.DeleteKey(key)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
