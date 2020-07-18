package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index default page
func Index(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, "<h1>It works!</h1>")
}
