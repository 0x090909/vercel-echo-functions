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

	server.GET("/user_all", func(c echo.Context) error {
		return c.HTML(200, "Hello all")
	})

	server.GET("/user_gone", func(c echo.Context) error {
		return c.HTML(200, "Hello gone")
	})

	server.ServeHTTP(w, r)
}
