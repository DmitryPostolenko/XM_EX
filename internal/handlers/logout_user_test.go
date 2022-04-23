package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/DmitryPostolenko/XM_EX/internal/jwt"
	td "github.com/DmitryPostolenko/XM_EX/test_data"
)

func TestLogoutUser(t *testing.T) {
	ctx := context.Background()
	redisConnect := td.RedisConnection()

	uid := "5238126c-2ceb-4987-4bbc-71772f4fd752"
	token, err := jwt.CreateToken(uid)
	if err != nil {
		t.Fatalf("CreateToken failed")
	}

	TestData := map[string]string{"token": token.AccessToken}
	TestDataM, _ := json.Marshal(TestData)
	log.Println(token)
	at := time.Unix(token.Expires, 0)
	now := time.Now()
	saveErr := redisConnect.Set(ctx, token.AccessUuid, uid, at.Sub(now)).Err()
	if saveErr != nil {
		t.Fatalf("Saving token failed")
	}

	e := echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/v1/user/logout", bytes.NewReader(TestDataM))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)
	c.Set("redis", redisConnect)

	errLogin := LogoutUser(c)

	logoutUserResponse := new(LogoutUserResponse)
	dec := json.NewDecoder(rec.Body)
	err = dec.Decode(&logoutUserResponse)
	if err != nil {
		t.Fatalf("Error while decoding CreateUserResponse: %v", err)
	}
	if assert.NoError(t, errLogin) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	if logoutUserResponse.Msg != "Success" {
		t.Fatalf("Wrong LoginUserResponse")
	}
}
