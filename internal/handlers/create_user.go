package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/uptrace/bun"

	"github.com/DmitryPostolenko/XM_EX/internal/hasher"
	"github.com/DmitryPostolenko/XM_EX/internal/models"
	"github.com/DmitryPostolenko/XM_EX/internal/repository"
)

// CreateUserResponse defines model for CreateUserResponse.
type CreateUserResponse struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
}

// CreateUser
// Use curl:
// curl -v POST http://localhost:8080/v0.9/user/register -H 'Content-Type: application/json' -d '{"userName":"my_login","password":"my_password"}'
func CreateUser(c echo.Context) error {
	// Binding request data
	userRequest, err := bindUserData(c)
	if err != nil {
		return err
	}

	//Getting DB instance
	db, _ := c.Get("db").(*bun.DB)

	//Getting user repo
	userRep := repository.GetUsersRepository(db)

	// Checking if username exists
	if _, ok := userRep.GetUserByField("username", userRequest.UserName); ok {
		errMsg := "Username already exists: " + userRequest.UserName
		return handleError(nil, errMsg, http.StatusBadRequest)
	}

	// Creating user id
	userUuid, err := uuid.NewV4()
	if err != nil {
		errMsg := "Internal server error. Failed to create user id: "
		return handleError(err, errMsg, http.StatusInternalServerError)
	}

	// Hashing password
	hashedPass, err := hasher.HashPassword(userRequest.Password)
	if err != nil {
		errMsg := "Internal server error. Failed to hash password: "
		return handleError(err, errMsg, http.StatusInternalServerError)
	}

	// Setting up User
	user := models.User{
		Id:       userUuid.String(),
		Username: userRequest.UserName,
		Password: hashedPass,
	}

	// Saving user
	err = userRep.SaveUser(user)
	if err != nil {
		errMsg := "Internal server error. Saving user failed! "
		return handleError(err, errMsg, http.StatusInternalServerError)
	}

	// Setting up response
	createUserResponse := CreateUserResponse{
		Id:       user.Id,
		UserName: userRequest.UserName,
	}

	// Setting up headers
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	// Encoding and sending response
	enc := json.NewEncoder(c.Response())
	err = enc.Encode(createUserResponse)
	if err != nil {
		errMsg := "Internal server error. Failed to encode createUserResponse: "
		return handleError(err, errMsg, http.StatusInternalServerError)
	}
	return nil
}
