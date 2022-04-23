package test_data

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/DmitryPostolenko/XM_EX/internal/configuration"
	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

var TestDataC = map[string]string{"name": "my_test_company", "code": "23323", "country": "Ukraine", "website": "https://something.com", "phone": "23323"}
var TestDataCM, _ = json.Marshal(TestDataC)
var TestData = map[string]string{"userName": "my_test_user", "password": "my_test_password"}
var TestDataM, _ = json.Marshal(TestData)

func DBConnection() *bun.DB {
	// DB CONFIGURATION AND SETTINGS
	cfg, _ := configuration.Load("../../internal/configuration/configuration.yml")
	pool := repository.Pool{}
	var dbUrl = os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.Host, cfg.DataBase.Port, cfg.DataBase.Name)
	}

	dbc := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dbUrl)))
	pool.DB = bun.NewDB(dbc, pgdialect.New())

	dbRep := repository.InitDatabase(&pool)
	dbConnect := dbRep.GetDB()

	return dbConnect
}

func RedisConnection() *redis.Client {
	ctx := context.Background()

	//var pool = client.Pool{}
	cfg, _ := configuration.Load("../../internal/configuration/configuration.yml")

	var redisUrl = os.Getenv("REDIS_URL")
	if redisUrl == "" {
		redisUrl = fmt.Sprintf("redis://%s:%d",
			cfg.Redis.Host, cfg.Redis.Port)
	}

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}
