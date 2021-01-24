package routes

import "github.com/gin-gonic/gin"

func authRoutes(r *gin.RouterGroup) {
	auth := r.Group("auth")

	auth.GET("/login", func(c *gin.Context) {
		c.String(200, "logg")
	})

}
