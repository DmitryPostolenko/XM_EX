package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

//DeleteCompanyResponse defines model for DeleteCompanyResponse.
type DeleteCompanyResponse struct {
	Msg string `json:"msg"`
}

// DeleteCompany
// Use curl:
// curl -v DELETE http://localhost:8080/v0.9/company/86e9860c-d11b-4317-7625-c95ee3db87c7
func DeleteCompany(c echo.Context) error {
	db, _ := c.Get("db").(*bun.DB)
	compRep := repository.GetCompaniesRepository(db)

	companyId := c.Param("id")

	ok := compRep.DeleteCompany(companyId)
	if ok != true {
		return handleError(nil, "No companies found", http.StatusBadRequest)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	deleteCompanyResponse := &DeleteCompanyResponse{
		Msg: "Success",
	}

	// Encoding response
	enc := json.NewEncoder(c.Response())
	err := enc.Encode(deleteCompanyResponse)
	if err != nil {
		errMsg := "Failed to encode LogoutUserResponse: "
		return handleError(err, errMsg, http.StatusBadRequest)
	}

	return nil
}
