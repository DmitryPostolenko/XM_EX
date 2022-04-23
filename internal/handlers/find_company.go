package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

// FindCompany
// Use curl:
// curl -v GET http://localhost:8080/v0.9/company?field=code&value=23323
func FindCompany(c echo.Context) error {
	db, _ := c.Get("db").(*bun.DB)
	compRep := repository.GetCompaniesRepository(db)

	field := c.QueryParam("field")
	value := c.QueryParam("value")

	resp, ok := compRep.FindCompany(field, value)
	if ok != true {
		return handleError(nil, "No companies found", http.StatusBadRequest)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	enc := json.NewEncoder(c.Response())
	err := enc.Encode(resp)
	if err != nil {
		errMsg := "Internal server error. Failed to encode listCompaniesResponse: "
		return handleError(err, errMsg, http.StatusInternalServerError)
	}
	return nil
}
