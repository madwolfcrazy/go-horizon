package router

import (
	"net/http"
	"ymz465/go-horizon/api"
	"ymz465/go-horizon/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	runMode := viper.GetString("runmode")
	if runMode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	//
	//add some middleware
	sessionSecret := viper.GetString("server.sessionsecret")
	r.Use(middleware.Session(sessionSecret))
	r.Use(middleware.MarkTraceID())
	r.Use(middleware.Cors())
	//
	r.GET("/", api.Index)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	apiV1 := r.Group("/api/v1")
	apiV1.GET("/", api.Index)

	return r
}
