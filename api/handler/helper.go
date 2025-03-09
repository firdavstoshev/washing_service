package handler

import (
	"github.com/firdavstoshev/washing_service/internal/dto"

	"github.com/gin-gonic/gin"
)

func errorJSON(c *gin.Context, code int, message string) {
	c.JSON(code, dto.ErrorResponse{Error: message})
}

func responseJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, map[string]interface{}{
		"data": data,
	})
}
