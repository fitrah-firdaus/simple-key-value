package keyvalue

import (
	"github.com/gofiber/fiber/v2/log"
	"simple-key-value/configuration"
	"simple-key-value/pkg/entities"
)

type Service interface {
	CreateOrUpdateKey(value *entities.KeyValue) (*entities.KeyValue, error)
	GetKey(key string) (*entities.KeyValue, error)
	DeleteKey(key string) error
}

type keyValueService struct {
	repository Repository
	cache      configuration.RedisCache
}

func (s *keyValueService) CreateOrUpdateKey(value *entities.KeyValue) (*entities.KeyValue, error) {
	result, err := s.repository.CreateOrUpdateKey(value)
	err = s.cache.Set(value.Key, value.Value)
	if err != nil {
		log.Error(err)
	}
	return result, err
}

func (s *keyValueService) GetKey(key string) (*entities.KeyValue, error) {
	result, err := s.cache.Get(key)
	if err != nil {
		log.Error(err)
	}
	if result != "" {
		log.Info("result from Cache")
		return &entities.KeyValue{Key: key, Value: result}, err
	}

	resultFromDatabase, err := s.repository.GetKey(key)
	if err != nil {
		log.Error(err)
	}

	if resultFromDatabase != nil {
		log.Info("result from database")
		err = s.cache.Set(resultFromDatabase.Key, resultFromDatabase.Value)
		if err != nil {
			log.Error(err)
		}
	}
	return resultFromDatabase, err
}

func (s *keyValueService) DeleteKey(key string) error {
	_ = s.cache.Remove(key)
	return s.repository.DeleteKey(key)
}

func NewService(r Repository, cache configuration.RedisCache) Service {
	return &keyValueService{
		repository: r,
		cache:      cache,
	}
}
