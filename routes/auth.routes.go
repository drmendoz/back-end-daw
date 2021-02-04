package routes

import (
	"github.com/drmendoz/backend-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func authRoutes(r *gin.RouterGroup) {
	auth := r.Group("auth")

	auth.POST("login", controllers.LoginCliente)

}
