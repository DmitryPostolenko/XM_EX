package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/hasher"
	"github.com/DmitryPostolenko/XM_EX/internal/jwt"
	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

// LoginUserResponse defines model for LoginUserResponse.
type LoginUserResponse struct {
	AccessToken string `json:"url"`
}

// LoginUser
// Use curl:
// curl -v POST http://localhost:8080/v1/user/login -H 'Content-Type: application/json' -d '{"userName":"my_login","password":"my_password"}'
func LoginUser(c echo.Context) error {
	db, _ := c.Get("db").(*bun.DB)
	userRep := repository.GetUsersRepository(db)
	client, _ := c.Get("redis").(*redis.Client)

	errorMsg := "Invalid username/password"

	// Binding request data
	userRequest, err := bindUserData(c)
	if err != nil {
		return err
	}

	//Checking if username exists
	if user, ok := userRep.GetUserByField("username", userRequest.UserName); ok {
		match := hasher.CheckPasswordHash(userRequest.Password, user.Password)
		if !match {
			return handleError(err, errorMsg, http.StatusBadRequest)
		}

		loginUserResponse, err := prepareLoginResponse(client, user.Id)
		setLoginResponseHeaders(c)
		// Encoding response
		enc := json.NewEncoder(c.Response())
		enc.SetEscapeHTML(false)
		err = enc.Encode(loginUserResponse)
		if err != nil {
			errMsg := "Failed to encode UserStorage: "
			return handleError(err, errMsg, http.StatusBadRequest)
		}
		return nil
	} else {
		return handleError(err, errorMsg, http.StatusBadRequest)
	}
}

func prepareLoginResponse(c *redis.Client, uid string) (*LoginUserResponse, error) {
	ctx := context.Background()
	lr := new(LoginUserResponse)
	// Generating token
	token, err := jwt.CreateToken(uid)
	if err != nil {
		errMsg := "Failed to generate token"
		return lr, handleError(err, errMsg, http.StatusBadRequest)
	}

	at := time.Unix(token.Expires, 0)
	now := time.Now()
	saveErr := c.Set(ctx, token.AccessUuid, uid, at.Sub(now)).Err()
	if saveErr != nil {
		errMsg := "CreateAuth failed: "
		return lr, handleError(err, errMsg, http.StatusInternalServerError)
	}

	loginUrl := "wss://dry-river-87369.herokuapp.com/v1/chat/ws.rtm.start?token=" + token.AccessToken
	loginUserResponse := &LoginUserResponse{
		AccessToken: loginUrl,
	}

	return loginUserResponse, nil
}

func setLoginResponseHeaders(c echo.Context) {
	// Setting up headers
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().Header().Set("X-Rate-Limit", strconv.Itoa(360))
	c.Response().Header().Set("X-Expires-After", time.Now().Add(time.Minute*10).Format(time.RFC1123))
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
}
