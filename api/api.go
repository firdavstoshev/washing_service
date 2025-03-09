package api

import (
	"net/http"

	"github.com/firdavstoshev/washing_service/api/handler"
	_ "github.com/firdavstoshev/washing_service/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Washing Service API
// @version 1.0
// @description This is a sample API for washing services.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @schemes http
// @host localhost:8080
func SetupRoutes(h *handler.Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method not allowed",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found",
		})
	})

	router.GET("/services", h.GetServices)
	router.POST("/order", h.CreateOrder)
	router.POST("/order-price", h.OrderPrice)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
