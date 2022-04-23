package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// CompanyRequest defines model for CompanyRequest.
type CompanyRequest struct {
	Id      string `json:"id"`
	Name    string `json:"name" validate:"required,min=1,max=100"`
	Code    string `json:"code" validate:"required,min=1,max=100"`
	Country string `json:"country" validate:"required,min=1,max=20"`
	Website string `json:"website" validate:"required,min=1,max=100"`
	Phone   string `json:"phone" validate:"required,min=1,max=16"`
}

func bindCompanyData(c echo.Context) (*CompanyRequest, error) {
	cd := new(CompanyRequest)

	// Binding data
	err := c.Bind(cd)
	if err != nil {
		errMsg := "Error during request body decoding: "
		return cd, handleError(err, errMsg, http.StatusBadRequest)
	}

	// Validating data
	v := validator.New()
	err = v.Struct(cd)
	if err != nil {
		errMsg := "All fields are required (name, code, country, website, phone): "
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
		}
		return cd, echo.NewHTTPError(http.StatusBadRequest, errMsg+err.Error())
	}

	return cd, nil
}
