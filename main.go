package main

import (
	"fmt"
	"net/http"
	"reflect"

	"2nd.app/config"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	env := config.EnvConfig{}
	err = envconfig.Process("", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	v := reflect.ValueOf(env)
	typeOfEnv := v.Type()

	for i := range v.NumField() {
		fmt.Printf("%s: %v\n", typeOfEnv.Field(i).Name, v.Field(i).Interface())
	}

	db, err := sqlx.Connect(env.DBDriver, env.DataSourceName())
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

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

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", env.AppHost, env.AppPort)))

}
