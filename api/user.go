package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandlerUser(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	server.GET("/user/", func(c echo.Context) error {
		return c.HTML(200, "Hello user")
	})

	server.ServeHTTP(w, r)
}
