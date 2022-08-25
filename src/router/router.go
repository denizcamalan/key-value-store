package router

import (
	"github.com/denizcamalan/key-value-store/controller"
	"github.com/gin-gonic/gin"
)

// NewRoutes router global
func NewRoutes() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api")

	// register router from each controller service
	RoutesPost(v1)

	return router
}

func RoutesPost(rg *gin.RouterGroup) {
	keys := rg.Group("/keys")

	keys.POST("/", controller.CreateData)
	keys.GET("/:id", controller.GetDataById)
	keys.HEAD("/:id", controller.CheckIfExist)
	keys.PUT("/:id", controller.UpdateData)
	keys.DELETE("/:id",controller.DeleteDataByID)
	keys.GET("/",controller.GetAll)
	keys.DELETE("/", controller.DeleteAll)
}