package routes

import (
	"github.com/drmendoz/backend-gin-gorm/controllers"
	"github.com/drmendoz/backend-gin-gorm/middlewares"
	"github.com/gin-gonic/gin"
)

func resenaRoutes(r *gin.RouterGroup) {
	prod := r.Group("resenas")
	prod.Use(middlewares.AuthMiddleWare())

	prod.GET("", controllers.GetResenas)
	prod.POST("", controllers.CreateResena)
	prod.PUT("/:id", controllers.UpdateResena)
	prod.GET("/:id", controllers.GetResenaPorId)
	prod.DELETE("/:id", controllers.DeleteResena)

}
