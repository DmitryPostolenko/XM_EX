package service

import (
	"github.com/labstack/echo/v4"

	"github.com/DmitryPostolenko/XM_EX/internal/configuration"
	"github.com/DmitryPostolenko/XM_EX/internal/handlers"
)

func initRoutes(e *echo.Echo, c *configuration.Configuration) {
	// Users endpoint
	e.POST("/"+c.Api.Version+"/user/register", handlers.CreateUser)
	e.POST("/"+c.Api.Version+"/user/login", handlers.LoginUser)
	e.POST("/"+c.Api.Version+"/user/logout", handlers.LogoutUser)

	//Companies
	//e.POST("/"+c.Api.Version+"/company/add", handlers.CreateCompany)
	//e.GET("/"+c.Api.Version+"/company/list", handlers.ListCompanies)
	//e.GET("/"+c.Api.Version+"/company/find/:id", handlers.FindCompany)
	//e.PUT("/"+c.Api.Version+"/company/update/:id", handlers.UpdateCompany)
	//e.DELETE("/"+c.Api.Version+"/company/delete/:id", handlers.DeleteCompany)
}
