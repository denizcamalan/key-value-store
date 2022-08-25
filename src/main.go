package main

import (
	"log"
	"os"

	"github.com/denizcamalan/key-value-store/configuration"
	"github.com/denizcamalan/key-value-store/docs"
	"github.com/denizcamalan/key-value-store/router"
	"github.com/spf13/viper"
	swgFiles "github.com/swaggo/files"
	swgGin "github.com/swaggo/gin-swagger"
)

func init() {

	os.Setenv("APP_ENVIRONMENT", "STAGING")

	// read config environment
	configuration.ReadConfig()

}
func main() {
	
	defer configuration.NewDatabase().Close()

	port := viper.GetString("PORT")


	docs.SwaggerInfo.Title = "Swagger Service key-value-store"
	docs.SwaggerInfo.Description = "This is service API documentation."
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Setup router
	router := router.NewRoutes()
	url := swgGin.URL("http://localhost:" + port + "/swagger/doc.json")
	router.GET("/swagger/*any", swgGin.WrapHandler(swgFiles.Handler, url))

	log.Fatal(router.Run(":" + port))
}
