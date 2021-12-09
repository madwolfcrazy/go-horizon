package api

import (
	"strconv"
	"ymz465/go-horizon/serializer"

	"github.com/gin-gonic/gin"
)

//GetPaginationFromURL 从URL获取分页信息
func GetPaginationFromURL(c *gin.Context) (page uint, pagesize uint) {
	pageInt, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	pagesizeInt, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		pagesize = 20
	}
	return uint(pageInt), uint(pagesizeInt)
}

//GetIntParamFromURL  从URL 获取param int类型参数
func GetIntParamFromURL(c *gin.Context, name string) (num int) {
	num, _ = strconv.Atoi(c.Param(name))
	return
}

//GetUintParamFromURL  从URL 获取param uint类型参数
func GetUintParamFromURL(c *gin.Context, name string) uint {
	num, _ := strconv.ParseUint(c.Param(name), 10, 64)
	return uint(num)
}

//GetIntQueryFromURL  从URL 获取query int类型参数
func GetIntQueryFromURL(c *gin.Context, name string) (num int) {
	num, _ = strconv.Atoi(c.Query(name))
	return
}

//GetUintQueryFromURL  从URL 获取query uint类型参数
func GetUintQueryFromURL(c *gin.Context, name string) uint {
	num, _ := strconv.ParseUint(c.Query(name), 10, 64)
	return uint(num)
}

//GetGroupID 获取session Group ID
func GetGroupID(c *gin.Context) uint {
	GIDi, ok := c.Get("GroupID")
	if !ok {
		return 0
	}
	if GID, err := GIDi.(float64); !err {
		return 0
	} else {
		return uint(GID)
	}
}

//GetUID 获取session UID
func GetUID(c *gin.Context) uint {
	UIDi, ok := c.Get("UID")
	if !ok {
		return 0
	}
	if UID, err := UIDi.(float64); !err {
		return 0
	} else {
		return uint(UID)
	}

}

//GetUsername 获取session Username
func GetUsername(c *gin.Context) string {
	username, ok := c.Get("username")
	if !ok {
		return ""
	}
	return username.(string)
}

//ReturnError 返回错误
func ReturnError(c *gin.Context, err error) {
	c.JSON(200, serializer.ErrorResponse(err))
	return
}
