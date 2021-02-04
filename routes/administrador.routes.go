package routes

import (
	"github.com/drmendoz/backend-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func administradorRoutes(r *gin.RouterGroup) {
	admin := r.Group("administradores")
	//adm.Use(middlewares.AuthMiddleWare())

	admin.GET("", controllers.GetAdministradores)
	admin.POST("", controllers.CreateAdministrador)
	admin.PUT("/:id", controllers.UpdateAdministrador)
	admin.GET("/:id", controllers.GetAdministradorPorId)
	admin.DELETE("/:id", controllers.DeleteAdministrador)

}
