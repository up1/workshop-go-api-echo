package demo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func SayHi(m Message) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		res := HelloResponse{Message: m.Data}
		return c.JSON(http.StatusOK, res)
	}
}
