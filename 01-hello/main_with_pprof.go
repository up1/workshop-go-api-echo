package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "net/http/pprof"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", func(c echo.Context) error {
		res := HelloResponse{Message: "Hello world"}
		return c.JSON(http.StatusOK, res)
	})

	e.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))

	e.Logger.Fatal(e.Start(":1323"))
}
