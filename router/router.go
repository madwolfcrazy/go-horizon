package router

import (
	"net/http"
	"ymz465/go-horizon/api"
	"ymz465/go-horizon/config"
	"ymz465/go-horizon/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	//add some middleware
	sessionSecret := viper.GetString("server.sessionsecret")
	r.Use(middleware.Session(sessionSecret))
	//r.Use(middleware.MarkTraceID())
	r.Use(middleware.Cors())
	//
	r.GET("/", api.Index)
	r.Static("/html/", config.HTMLDir)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//
	//r.Static("/html/", "./html")
	apiV1 := r.Group("/api/v1")
	apiV1.GET("/", api.Index)
	//
	jwtObject, err := middleware.MakeJwtAuthMiddleware()
	if err != nil {
		panic("jwt auth error pleace check")
	}
	apiV1.POST("/login", jwtObject.LoginHandler)
	// 需要登录保护的
	authed := apiV1.Group("")
	authed.Use(jwtObject.MiddlewareFunc())

	{
		//
		authed.POST("chpwd", api.Chpwd)
	}

	return r
}
