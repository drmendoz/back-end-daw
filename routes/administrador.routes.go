package routes

import (
	"github.com/drmendoz/backend-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

type Administrador struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func administradorRoutes(r *gin.RouterGroup) {
	adm := r.Group("administradores")
	//adm.Use(middlewares.AuthMiddleWare())
	adm.GET("", controllers.GetAdministradores)
	adm.POST("", controllers.CreateAdministrador)
	adm.PUT("/:id", controllers.UpdateAdministrador)
	adm.GET("/:id", controllers.GetAdministradorPorId)
}
