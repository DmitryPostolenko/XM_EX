package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBindUserData(t *testing.T) {
	e := echo.New()

	testData := map[string]string{"user_name": "my_login22133", "password": "my_password"}
	testDataM, _ := json.Marshal(testData)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/user", bytes.NewReader(testDataM))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)

	userRequest, err := bindUserData(c)

	if assert.NoError(t, err) {
		assert.Equal(t, userRequest.Password, testData["password"])
		assert.Equal(t, userRequest.UserName, testData["user_name"])
	}
}

func TestBindUserDataInvalidUserName(t *testing.T) {
	tests := []struct {
		name     string
		testData map[string]string
		err      string
	}{
		{
			name:     "UserName validation",
			testData: map[string]string{"user_name": "my", "password": "my_password"},
			err:      "UserName validation failed",
		},
		{
			name:     "Password validation",
			testData: map[string]string{"user_name": "my_login22133", "password": "my"},
			err:      "Password validation failed",
		},
		{
			name:     "Empty file",
			testData: map[string]string{},
			err:      "Loaded empty file",
		},
	}

	e := echo.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDataM, _ := json.Marshal(tt.testData)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(echo.POST, "/user", bytes.NewReader(testDataM))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			c := e.NewContext(req, rec)

			_, err := bindUserData(c)
			if err == nil {
				t.Fatal(tt.err)
			}
		})
	}
}
