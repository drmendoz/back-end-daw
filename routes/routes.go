package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/drmendoz/backend-gin-gorm/controllers"
)

var R *gin.Engine

func init() {
	setLogging()

	R = gin.Default()
	R.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))
	setLoggingMiddleware()
	router := R.Group("/api/v1")

	router.GET("resenasAdmin", controllers.GetResenasAdmin)
	authRoutes(router)
	administradorRoutes(router)
	productoRoutes(router)
	clientesRoutes(router)
	pedidoRoutes(router)
	resenaRoutes(router)
	contactenosRoutes(router)
	graficoRoutes(router)
}
