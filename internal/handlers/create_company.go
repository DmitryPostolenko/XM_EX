package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/models"
	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

// CreateCompanyResponse defines model for CreateCompanyResponse.
type CreateCompanyResponse struct {
	Id          string `json:"id"`
	CompanyName string `json:"company_name"`
}

// CreateCompany
// Use curl:
// curl -v POST http://localhost:8080/v0.9/company/ -H 'Content-Type: application/json' -d '{"name":"my_company","code":"23323","country":"Ukraine","website":"https://something.com","phone":"23323"}'
func CreateCompany(c echo.Context) error {

	// Binding request data
	companyRequest, err := bindCompanyData(c)
	if err != nil {
		return err
	}

	if checkAuthorization(c, companyRequest.Token) == false {
		errMsg := "Only access from Cyprus or for authorized users allowed"
		return handleError(nil, errMsg, http.StatusInternalServerError)
	}

	//Getting DB instance
	db, _ := c.Get("db").(*bun.DB)

	//Getting company repo
	compRep := repository.GetCompaniesRepository(db)

	// Checking if company name exists
	if _, ok := compRep.FindCompany("name", companyRequest.Name); !ok {
		errMsg := "Company already exists: " + companyRequest.Name
		return handleError(nil, errMsg, http.StatusBadRequest)
	}

	// Creating company id
	companyUuid, err := uuid.NewV4()
	if err != nil {
		errMsg := "Internal server error. Failed to create company id: "
		return handleError(err, errMsg, http.StatusInternalServerError)
	}

	//Setting up company
	company := models.Company{
		Id:      companyUuid.String(),
		Name:    companyRequest.Name,
		Code:    companyRequest.Code,
		Country: companyRequest.Country,
		Website: companyRequest.Website,
		Phone:   companyRequest.Phone,
	}

	// Saving company
	err = compRep.SaveCompany(company)
	if err != nil {
		errMsg := "Internal server error. Saving company failed! "
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
