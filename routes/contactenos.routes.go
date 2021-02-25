package routes

import (
	"github.com/drmendoz/backend-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func contactenosRoutes(r *gin.RouterGroup) {
	prod := r.Group("contactenos")
	//adm.Use(middlewares.AuthMiddleWare())

	prod.POST("", controllers.HandleContactenos)

}
