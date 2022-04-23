package handlers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	td "github.com/DmitryPostolenko/XM_EX/test_data"
)

func TestListCompanies(t *testing.T) {
	dbConnect := td.DBConnection()

	e := echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/v0.9/company/list", bytes.NewReader(td.TestDataCM))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)

	c.Set("db", dbConnect)

	err := ListCompanies(c)
	if err != nil {
		t.Fatalf("Error listing companies: %v", err)
	}
}
