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

func delCompany(dbConnect *bun.DB, t *testing.T) {
	ctx := context.Background()
	company := new(models.Company)

	_, err := dbConnect.NewDelete().Model(company).Where("name = ?", td.TestDataC["name"]).Exec(ctx)
	if err != nil {
		t.Fatalf("Error while deleting test company: %v", err)
	}
}

func TestCreateCompany(t *testing.T) {
	dbConnect := td.DBConnection()

	delCompany(dbConnect, t)

	e := echo.New()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.POST, "/v0.9/company/add", bytes.NewReader(td.TestDataCM))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := e.NewContext(req, rec)

	c.Set("db", dbConnect)

	errCreate := CreateCompany(c)

	createCompanyResponse := new(CreateCompanyResponse)
	dec := json.NewDecoder(rec.Body)
	err := dec.Decode(&createCompanyResponse)
	if err != nil {
		t.Fatalf("Error while decoding CreateCompanyResponse: %v", err)
	}

	if assert.NoError(t, errCreate) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, createCompanyResponse.CompanyName, td.TestDataC["name"])
	}

	delCompany(dbConnect, t)
}
