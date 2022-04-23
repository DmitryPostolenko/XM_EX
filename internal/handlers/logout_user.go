package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"

	"github.com/DmitryPostolenko/XM_EX/internal/jwt"
)

//LogoutUserRequest defines model for LogoutUserRequest.
type LogoutUserRequest struct {
	AccessToken string `json:"token"`
}

//LogoutUserResponse defines model for LoginUserResponse.
type LogoutUserResponse struct {
	Msg string `json:"msg"`
}

// LogoutUser
// Use curl:
// curl -v POST http://localhost:8080/v1/user/logout -H 'Content-Type: application/json' -d '{"userName":"access_token","password":"my_access_token"}'
func LogoutUser(c echo.Context) error {
	ctx := context.Background()

	req := new(LogoutUserRequest)

	// Binding data
	err := c.Bind(req)
	if err != nil {
		errMsg := "Error during request body decoding: "
		return handleError(err, errMsg, http.StatusBadRequest)
	}

	accessDetails, err := jwt.ExtractTokenMetadata(req.AccessToken)
	if err != nil {
		errMsg := "Invalid token: "
		log.Printf(errMsg+"%v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, errMsg+err.Error())
	}

	redisClient, _ := c.Get("redis").(*redis.Client)

	_, err = redisClient.Get(ctx, accessDetails.AccessUuid).Result()
	if err != nil {
		errMsg := "Unauthorized. "
		log.Printf(errMsg+"%v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, errMsg+err.Error())
	}
	redisClient.Del(ctx, accessDetails.AccessUuid)

	logoutUserResponse := &LogoutUserResponse{
		Msg: "Success",
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	// Encoding response
	enc := json.NewEncoder(c.Response())
	err = enc.Encode(logoutUserResponse)
	if err != nil {
		errMsg := "Failed to encode LogoutUserResponse: "
		return handleError(err, errMsg, http.StatusBadRequest)
	}

	return nil
}
