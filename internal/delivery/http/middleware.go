package http

import "github.com/gin-gonic/gin"

func optionMiddleware(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		c.Writer.WriteHeader(200)
		return
	}
}

func corsMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
