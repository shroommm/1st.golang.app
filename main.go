package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		res := Response{
			Message: "Hello, World!",
		}

		return c.JSON(http.StatusOK, res)
	})

	e.Start("localhost:8080")
}
