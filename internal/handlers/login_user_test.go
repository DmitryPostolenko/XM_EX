package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/repository"
	td "github.com/DmitryPostolenko/XM_EX/test_data"

	"github.com/DmitryPostolenko/XM_EX/internal/models"
)

var userUuid = uuid.NewV4()

var user = models.User{
	Id:       userUuid.String(),
	Username: "my_test_user",
	Password: "$2a$14$Nlv5V7Iu6ZHpks0zwoD9zejiCZNfwhG9SWKv2jTjBIAl0eazWN/nm",
}

func addTestUser(dbc *bun.DB) error {
	userRep := repository.GetUsersRepository(dbc)
	err := userRep.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func TestLoginUser(t *testing.T) {
	dbConnect := td.DBConnection()
	redisConnect := td.RedisConnection()

	delUser(dbConnect, t)

	err := addTestUser(dbConnect)
	if err != nil {
		t.Fatalf("Saving test user failed")
	}

	e := echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/v1/user/login", bytes.NewReader(td.TestDataM))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)

	c.Set("db", dbConnect)
	c.Set("redis", redisConnect)

	errLogin := LoginUser(c)

	u := new(LoginUserResponse)
	dec := json.NewDecoder(rec.Body)
	err = dec.Decode(&u)
	if err != nil {
		t.Fatalf("Error while decoding CreateUserResponse: %v", err)
	}
	if assert.NoError(t, errLogin) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	if u.AccessToken == "" {
		t.Fatalf("Wrong LoginUserResponse")
	}

	delUser(dbConnect, t)
}
