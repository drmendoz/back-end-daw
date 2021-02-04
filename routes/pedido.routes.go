package routes

import (
	"github.com/drmendoz/backend-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func pedidoRoutes(r *gin.RouterGroup) {
	admin := r.Group("pedidos")
	//adm.Use(middlewares.AuthMiddleWare())

	admin.GET("", controllers.GetPedidos)
	admin.POST("", controllers.CreatePedido)
	admin.PUT("/:id", controllers.UpdatePedido)
	admin.GET("/:id", controllers.GetPedidoPorId)
	admin.DELETE("/:id", controllers.DeletePedido)

}
