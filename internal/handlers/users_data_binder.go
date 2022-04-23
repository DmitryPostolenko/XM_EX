package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

// UserRequest defines model for UserRequest.
type UserRequest struct {
	UserName string `json:"user_name" validate:"required,min=4,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

func bindUserData(c echo.Context) (*UserRequest, error) {
	u := new(UserRequest)

	// Binding data
	err := c.Bind(u)
	if err != nil {
		errMsg := "Error during request body decoding: "
		return u, handleError(err, errMsg, http.StatusBadRequest)
	}

	// Validating data
	v := validator.New()
	err = v.Struct(u)
	if err != nil {
		errMsg := "Minimal username length is 4 symbols, minimal password length is 8 symbols: "
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
		}
		return u, echo.NewHTTPError(http.StatusBadRequest, errMsg+err.Error())
	}

	return u, nil
}
