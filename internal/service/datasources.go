package service

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/DmitryPostolenko/XM_EX/internal/configuration"
	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

// Set up and config DB
func setUpDB(cfg *configuration.Configuration) *bun.DB {
	switch cfg.DataBase.Type {
	case "postgres":
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
	default:
		pool := repository.Pool{}

		dbRep := repository.InitDatabase(&pool)
		dbConnect := dbRep.GetDB()

		return dbConnect
	}
}

// Set up and config redis
func setUpRedis(cfg *configuration.Configuration) *redis.Client {
	ctx := context.Background()
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

// Set up server port configuration
func setUpPort(p string) string {
	// Getting port number from env or setting up from configuration file if env port is empty
	port := os.Getenv("PORT")
	if port == "" {
		port = p
	}
	return port
}
