package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/Parthh191/backendtask/internal/logger"
	"time"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		if c.Request.Method=="OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func LoggingMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()
		duration:=time.Since(startTime)
		statusCode:=c.Writer.Status()
		method:=c.Request.Method
		path:=c.Request.URL.Path

		logger.Info("[%s] %s %s - %d (%v)", method, path, c.ClientIP(), statusCode, duration)
	}
}

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err:=range c.Errors {
			if err.Type==gin.ErrorTypeBind {
				c.JSON(400, gin.H{"error": "Invalid request format"})
				return
			}
		}
	}
}
