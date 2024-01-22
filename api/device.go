package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandlerDevice(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	mainGroup := server.Group("/device")
	mainGroup.GET("/", func(c echo.Context) error {
		return c.HTML(200, "Hello device")
	})

	server.ServeHTTP(w, r)
}
