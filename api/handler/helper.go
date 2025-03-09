package handler

import "github.com/gin-gonic/gin"

func errorJSON(c *gin.Context, code int, message string) {
	c.JSON(code, map[string]interface{}{
		"error": message,
	})
}

func responseJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, map[string]interface{}{
		"data": data,
	})
}
