package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//GinRouter -> Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

//SetupRoutes : all the routes are defined here
func SetupRoutes() GinRouter {
	httpRouter := gin.Default()

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "GOlang FX Up and Running"})
	})

	return GinRouter{
		Gin: httpRouter,
	}
}
