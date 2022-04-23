package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

// ListCompanies
// Use curl:
// curl -v GET http://localhost:8080/v0.9/company/
func ListCompanies(c echo.Context) error {
	db, _ := c.Get("db").(*bun.DB)
	compRep := repository.GetCompaniesRepository(db)

	token := c.QueryParam("token")

	fmt.Println("ListCompanies token")
	fmt.Println(token)

	if checkAuthorization(c, token) == false {
		errMsg := "Only access from Cyprus or for authorized users allowed"
		return handleError(nil, errMsg, http.StatusInternalServerError)
	}

	resp, ok := compRep.ListCompanies()
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
