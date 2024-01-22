# Vercel with go-echo and functions

This project is a template that uses echo web framework integrated in vercel functions.

You can leverage all functionalities of echo with individual functions as well.

The structure is as follows:

Each file has to implement a function with the following header:

<code>func HandlerFunction(w http.ResponseWriter, r *http.Request)</code> 

The handler creates a new instance of echo and setsup all the nedded routes.
```
{
    server := echo.New()
    mainGroup := server.Group("/device")
    mainGroup.GET("/", func (c echo.Context) error {
        return c.HTML(200, "Hello device")
    })
    server.ServeHTTP(w, r)
}
```

Its important the function has a root prefix so that it can match the route defined in <code>vercel.json</code>

### Example

#### vercel.json

```
{
  "routes": [
    { "src": "/device/(.*)", "dest": "/api/device.go" }
  ]
}
```

#### api/device.go

```
package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandlerDevice(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	mainGroup := server.Group("/device")
	mainGroup.GET("/", func (c echo.Context) error {
        return c.HTML(200, "Hello device")
    })

	server.ServeHTTP(w, r)
}

```
