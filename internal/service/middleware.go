package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/DmitryPostolenko/XM_EX/internal/configuration"
)

func initMiddleware(e *echo.Echo, c *configuration.Configuration) {
	// Configurating and connecting to database
	dbConnect := setUpDB(c)

	// Configurating and connecting to redis
	redisClient := setUpRedis(c)

	// Echo middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(bodyDumpHandler))

	// Data source middleware
	e.Use(dataSourceMiddleware("db", dbConnect))
	e.Use(dataSourceMiddleware("redis", redisClient))
}

// Middleware function to transfer sql dataStore to handlers
func dataSourceMiddleware(dataType string, dataStore interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(dataType, dataStore)
			return next(c)
		}
	}
}

// Middleware to log requests and responses
func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("\nRequest Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}
