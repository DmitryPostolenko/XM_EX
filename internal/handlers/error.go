package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
)

func handleError(err error, errMsg string, statusCode int) *echo.HTTPError {
	log.Printf(errMsg+"%v\n", err)
	return echo.NewHTTPError(statusCode, errMsg)
}
