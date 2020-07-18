package middleware

import (
	"ymz465/go-horizon/helper"

	"github.com/gin-gonic/gin"
)

//MarkTraceID 标记trace id
func MarkTraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		existsVal := c.GetHeader("Trace-ID")
		if existsVal == "" {
			c.Header("Trace-ID", helper.RandString(32))
		}
		// before request
		c.Next()
	}
}
