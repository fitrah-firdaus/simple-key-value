package test

import (
	"context"
	"github.com/fitrah-firdaus/simple-key-value/api/routes"
	"github.com/fitrah-firdaus/simple-key-value/configuration"
	"github.com/fitrah-firdaus/simple-key-value/pkg/keyvalue"
	"github.com/fitrah-firdaus/simple-key-value/pkg/keyvalue/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var app *fiber.App

type TestSuite struct {
	suite.Suite
	mySQLTestContainer *MySQLTestContainer
	mongoContainer     *MongoTestContainer
	redisContainer     *RedisTestContainer
	server             httptest.Server
}

func (s *TestSuite) SetupSuite() {
	var err error
	s.mySQLTestContainer, err = NewMySQLTestContainer()
	s.NoError(err)

	s.mongoContainer, err = NewMongoTestContainer()
	s.NoError(err)

	s.redisContainer, err = NewRedisTestContainer()
	s.NoError(err)

	appInit := configuration.NewAppInit()
	config := configuration.New(".env.test")

	_ = config.Set("MONGO_URI", s.mongoContainer.GetURI())
	_ = config.Set("REDIS_URI", s.redisContainer.GetURI())
	_ = config.Set("MYSQL_URI", s.mySQLTestContainer.GetURI())

	//collection := appInit.InitMongoDB(config)
	//db := appInit.InitMySQL(config)
	db := appInit.InitGormMySQL(config)
	redisCache := appInit.InitRedis(config)

	//keyValueRepository := repository.NewMySQLRepo(db)
	keyValueRepository := repository.NewGormRepository(db)
	keyValueService := keyvalue.NewService(keyValueRepository, redisCache)

	app = appInit.InitFiberApp()
	api := app.Group("/api")
	routes.KeyValueRouter(api, keyValueService)

	//	s.server = httptest.NewServer(NewRouter())
}

func (s *TestSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	s.Require().NoError(s.mongoContainer.Terminate(ctx))
	s.Require().NoError(s.redisContainer.Terminate(ctx))
	s.Require().NoError(app.ShutdownWithContext(ctx))
}

func TestSuite_Run(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) Test_001_CreateOrUpdateKey() {

	req := httptest.NewRequest("POST", "/api/kv", strings.NewReader(`{"key":"test","value":"test value"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *TestSuite) Test_002_GetKey() {

	req := httptest.NewRequest("GET", "/api/kv/test", nil)
	req.Header.Set("Accept", "application/json")

	resp, err := app.Test(req)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
}

func (s *TestSuite) Test_003_DeleteKey() {
	req := httptest.NewRequest("DELETE", "/api/kv/test", nil)
	req.Header.Set("Accept", "application/json")

	resp, err := app.Test(req)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)

	req = httptest.NewRequest("GET", "/api/kv/test", nil)
	req.Header.Set("Accept", "application/json")

	resp, err = app.Test(req)
	s.Equal(http.StatusNotFound, resp.StatusCode)
}
