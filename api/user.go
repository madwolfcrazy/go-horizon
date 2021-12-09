package api

import (
	"ymz465/go-horizon/serializer"
	"ymz465/go-horizon/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Chpwd 修改密码
func Chpwd(c *gin.Context) {
	var chPWDService services.ChangePassword
	if err := c.ShouldBind(&chPWDService); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}
	chPWDService.UID = GetUID(c)
	if err := chPWDService.ChangePassword(); err != nil {
		c.JSON(200, serializer.ErrorResponse(err))
		return
	}
	// 登出系统，要求用户重新登录系统
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	//
	c.JSON(200, serializer.CommonOK())
}

//Me 当前登录人员信息
func Me(c *gin.Context) {
	UID := GetUID(c)
	Username := GetUsername(c)
	//
	c.JSON(200, serializer.Common(gin.H{"uid": UID, "username": Username}))
}
