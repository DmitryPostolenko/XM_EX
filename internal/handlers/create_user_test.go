package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"

	td "github.com/DmitryPostolenko/XM_EX/test_data"

	"github.com/DmitryPostolenko/XM_EX/internal/models"
)

func delUser(dbConnect *bun.DB, t *testing.T) {
	ctx := context.Background()
	user := new(models.User)

	_, err := dbConnect.NewDelete().Model(user).Where("username = ?", td.TestData["userName"]).Exec(ctx)
	if err != nil {
		t.Fatalf("Error while deleting test user: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	dbConnect := td.DBConnection()

	delUser(dbConnect, t)

	e := echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/v1/user", bytes.NewReader(td.TestDataM))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)

	c.Set("db", dbConnect)

	errCreate := CreateUser(c)

	u := new(CreateUserResponse)
	dec := json.NewDecoder(rec.Body)
	err := dec.Decode(&u)
	if err != nil {
		t.Fatalf("Error while decoding CreateUserResponse: %v", err)
	}

	if assert.NoError(t, errCreate) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, u.UserName, td.TestData["userName"])
	}

	delUser(dbConnect, t)
}
