package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	server.GET("/", func(c echo.Context) error {
		return c.HTML(200, "Hello index")
	})

	server.GET("/endpoint1", func(c echo.Context) error {
		return c.HTML(200, "Hello endpoint1")
	})

	server.GET("/endpoint2", func(c echo.Context) error {
		return c.HTML(200, "Hello endpoint1")
	})

	server.ServeHTTP(w, r)
}
