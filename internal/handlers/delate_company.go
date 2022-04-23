package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

//DeleteCompanyRespponse defines model for DeleteCompanyRespponse.
type DeleteCompanyRespponse struct {
	Msg string `json:"msg"`
}

// ListCompanies
// Use curl:
// curl -v GET http://localhost:8080/v0.9/company/list
func DeleteCompany(c echo.Context) error {
	db, _ := c.Get("db").(*bun.DB)
	compRep := repository.GetCompaniesRepository(db)

	companyId := c.Param("id")

	//fmt.Println("companyId" + companyId)
	//fmt.Println(companyId)
	ok := compRep.DeleteCompany(companyId)
	if ok != true {
		return handleError(nil, "No companies found", http.StatusBadRequest)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	deleteCompanyRespponse := &DeleteCompanyRespponse{
		Msg: "Success",
	}

	// Encoding response
	enc := json.NewEncoder(c.Response())
	err := enc.Encode(deleteCompanyRespponse)
	if err != nil {
		errMsg := "Failed to encode LogoutUserResponse: "
		return handleError(err, errMsg, http.StatusBadRequest)
	}

	return nil
}
