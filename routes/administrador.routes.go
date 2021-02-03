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

	//Administradores
	admin := adm.Group("administradores")
	admin.GET("", controllers.GetAdministradores)
	admin.POST("", controllers.CreateAdministrador)
	admin.PUT("/:id", controllers.UpdateAdministrador)
	admin.GET("/:id", controllers.GetAdministradorPorId)
	admin.DELETE("/:id", controllers.DeleteAdministrador)

	//Categorias
	cat := adm.Group("categorias")
	cat.GET("", controllers.GetCategorias)
	cat.POST("", controllers.CreateCategoria)
	cat.GET(":id", controllers.GetCategoriaPorId)
	cat.DELETE(":id", controllers.DeleteCategoria)
	cat.PUT(":id", controllers.UpdateCategoria)

	afi := adm.Group("afiliados")
	afi.GET("", controllers.GetAfiliados)
	afi.POST("", controllers.CreateAfiliados)
	afi.GET(":id", controllers.GetAfiliadoPorId)
	afi.DELETE(":id", controllers.DeleteAfiliado)
	afi.PUT(":id", controllers.UpdateAfiliado)

}
