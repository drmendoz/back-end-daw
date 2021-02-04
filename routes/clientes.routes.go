package routes

import (
	"github.com/drmendoz/backend-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func clientesRoutes(r *gin.RouterGroup) {
	prod := r.Group("clientes")
	//adm.Use(middlewares.AuthMiddleWare())

	prod.GET("", controllers.GetCliente)
	prod.POST("", controllers.CreateCliente)
	prod.PUT("/:id", controllers.UpdateCliente)
	prod.GET("/:id", controllers.GetClientePorId)
	prod.DELETE("/:id", controllers.DeleteCliente)

}
