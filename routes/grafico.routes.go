package routes

import (
	"github.com/drmendoz/backend-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func graficoRoutes(r *gin.RouterGroup) {
	prod := r.Group("graficos")
	//adm.Use(middlewares.AuthMiddleWare())

	prod.GET("ciudades", controllers.GetCiudadesGrafico)
	prod.GET("productos", controllers.GetProductosGrafico)
	prod.GET("clientes", controllers.GetClientesGrafico)

}
