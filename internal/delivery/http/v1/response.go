package v1

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}

func newGoodResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{Message: message})
}
