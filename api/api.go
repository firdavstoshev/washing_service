package api

import (
	"net/http"

	"github.com/firdavstoshev/washing_service/api/handler"

	"github.com/gin-gonic/gin"
)

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

	return router
}
