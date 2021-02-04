package routes

import (
	"github.com/drmendoz/backend-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func productoRoutes(r *gin.RouterGroup) {
	prod := r.Group("productos")
	//adm.Use(middlewares.AuthMiddleWare())

	prod.GET("", controllers.GetProductos)
	prod.POST("", controllers.CreateProducto)
	prod.PUT("/:id", controllers.UpdateProducto)
	prod.GET("/:id", controllers.GetProductoPorId)
	prod.DELETE("/:id", controllers.DeleteProdcuto)

}
