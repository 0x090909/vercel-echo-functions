package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Device struct {
	Name     string `validate:"required"  json:"name"`
	Location string `validate:"required"  json:"location"`
}

func HandlerDevice(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	mainGroup := server.Group("/device")

	mainGroup.GET("/", GETDevices)

	server.ServeHTTP(w, r)
}

// Get device list
//
//	@Summary		Get list of devices
//	@Description	Get list of devices
//	@Tags			Devices
//	@Consumes		multipart/form-data
//	@Produces		json
//	@Param			platform	path		string	true	"platform"	Enums(1, 2)
//	@Success		200			{object}	Device
//	@Failure		500			{object}	nil
//	@Router			/device/{platform} [get]
//	@Security		ApiKeyAuth

func GETDevices(c echo.Context) error {
	return c.HTML(200, "Hello device")
}
