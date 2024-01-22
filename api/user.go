package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	server.GET("/user/", func(c echo.Context) error {
		return c.HTML(200, "Hello index")
	})

	server.GET("/user/all", func(c echo.Context) error {
		return c.HTML(200, "Hello all")
	})

	server.GET("/user/gone", func(c echo.Context) error {
		return c.HTML(200, "Hello gone")
	})

	server.ServeHTTP(w, r)
}
