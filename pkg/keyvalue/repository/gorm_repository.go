package repository

import (
	"github.com/fitrah-firdaus/simple-key-value/pkg/entities"
	"gorm.io/gorm"
)

var queryFilter = "keylog =?"

type gormRepository struct {
	DB *gorm.DB
}

func (g gormRepository) CreateOrUpdateKey(value *entities.KeyValue) (*entities.KeyValue, error) {
	result := g.DB.FirstOrCreate(&value, queryFilter, value.Key)
	result.Scan(&value)
	return value, nil
}

func (g gormRepository) GetKey(key string) (*entities.KeyValue, error) {
	var value entities.KeyValue
	result := g.DB.Where(queryFilter, key).First(&value)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &value, nil
}

func (g gormRepository) DeleteKey(key string) error {
	g.DB.Where(queryFilter, key).Delete(&entities.KeyValue{})
	return nil
}

func NewGormRepository(db *gorm.DB) Repository {
	return &gormRepository{
		DB: db,
	}
}
