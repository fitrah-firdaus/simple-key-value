package configuration

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

type Config interface {
	Get(key string) string
	Set(key string, value string) error
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func (config *configImpl) Set(key string, value string) error {
	return os.Setenv(key, value)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return &configImpl{}
}
