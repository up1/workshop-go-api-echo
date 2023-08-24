package demo_test

import (
	"demo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandlerSayHi(t *testing.T) {
	// Arrange
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/hello")

	// Assertions
	var messageJSON = `{"message":"Hello world"}`
	if assert.NoError(t, demo.SayHi(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, messageJSON, strings.TrimSpace(rec.Body.String()))
	}
}
