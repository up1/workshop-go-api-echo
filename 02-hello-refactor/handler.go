package demo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func SayHi(c echo.Context) (err error) {
	res := HelloResponse{Message: "Hello world"}
	return c.JSON(http.StatusOK, res)
}
