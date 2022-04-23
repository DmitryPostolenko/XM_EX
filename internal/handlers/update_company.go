package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/models"
	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

// UpdateCompanyResponse defines model for UpdateCompanyResponse.
type UpdateCompanyResponse struct {
	Id          string `json:"id"`
	CompanyName string `json:"company_name"`
}

// UpdateCompanyResponse
// Use curl:
// curl -v PUT http://localhost:8080/v0.9/company/ -H 'Content-Type: application/json' -d '{"id":"6e52a033-c6e9-4305-51a6-59f00bb108b3","name":"my_company","code":"23323","country":"Ukraine","website":"https://something.com","phone":"23323"}'
func UpdateCompany(c echo.Context) error {
	// Binding request data
	companyRequest, err := bindCompanyData(c)
	if err != nil {
		return err
	}

	//Getting DB instance
	db, _ := c.Get("db").(*bun.DB)

	//Getting company repo
	compRep := repository.GetCompaniesRepository(db)

	//Setting up company
	company := models.Company{
		Id:      companyRequest.Id,
		Name:    companyRequest.Name,
		Code:    companyRequest.Code,
		Country: companyRequest.Country,
		Website: companyRequest.Website,
		Phone:   companyRequest.Phone,
	}

	// Updating company
	ok := compRep.UpdateCompany(company)
	if !ok {
		errMsg := "Internal server error. Updating company failed! "
		return handleError(err, errMsg, http.StatusInternalServerError)
	}

	// Setting up response
	createCompanyResponse := CreateCompanyResponse{
		Id:          company.Id,
		CompanyName: company.Name,
	}

	// Setting up headers
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	// Encoding and sending response
	enc := json.NewEncoder(c.Response())
	err = enc.Encode(createCompanyResponse)
	if err != nil {
		errMsg := "Internal server error. Failed to encode createCompanyResponse: "
		return handleError(err, errMsg, http.StatusInternalServerError)
	}
	return nil
}
