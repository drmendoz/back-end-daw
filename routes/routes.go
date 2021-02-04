package routes

import (
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	setLogging()
	R = gin.Default()
	setLoggingMiddleware()
	router := R.Group("/api/v1")
	authRoutes(router)
	administradorRoutes(router)
	productoRoutes(router)
	clientesRoutes(router)
	pedidoRoutes(router)
}
