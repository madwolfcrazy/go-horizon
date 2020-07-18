package router

import (
	"net/http"
	"ymz465/go-horizon/api"
	"ymz465/go-horizon/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//
	//add some middleware
	sessionSecret := viper.GetString("server.sessionsecret")
	r.Use(middleware.Session(sessionSecret))
	r.Use(middleware.MarkTraceID())
	//
	r.GET("/", api.Index)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	apiV1 := r.Group("/api/v1")
	apiV1.GET("/", api.Index)

	return r
}
