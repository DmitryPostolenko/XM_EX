package handlers

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"

	"github.com/DmitryPostolenko/XM_EX/internal/ipapi"
	"github.com/DmitryPostolenko/XM_EX/internal/jwt"
)

type AccessToken struct {
	Token string `json:"token"`
}

func checkAuthorization(c echo.Context, token string) bool {
	data, err := ipapi.GetIpData()
	if err != nil {
		log.Println(err)
	}
	if data.CountryName == "Cyprus" {
		return true
	}

	accessDetails, err := jwt.ExtractTokenMetadata(token)
	if err != nil {
		log.Println(err)
		return false
	}

	ctx := context.Background()
	redisClient, _ := c.Get("redis").(*redis.Client)
	_, err = redisClient.Get(ctx, accessDetails.AccessUuid).Result()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
