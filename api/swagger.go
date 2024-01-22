package api

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	_ "vercel-echo-functions/docs"
)

//	@title			Vercel OpenAPI Spec
//	@version		1.0
//	@description	Vercel OpenAPI Spec
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		https://vercel-echo-functions.vercel.app/
//	@BasePath	/

//	@securitydefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func HandlerSwagger(w http.ResponseWriter, r *http.Request) {
	server := echo.New()

	mainGroup := server.Group("/swagger")

	mainGroup.GET("/*", echoSwagger.WrapHandler)

	server.ServeHTTP(w, r)
}
