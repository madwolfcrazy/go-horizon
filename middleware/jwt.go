package middleware

import (
	"errors"
	"time"
	"ymz465/go-horizon/model"
	"ymz465/go-horizon/serializer"
	"ymz465/go-horizon/services"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// MakeJwtAuthMiddleware the jwt middleware
func MakeJwtAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	identityKey := "Username"
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm: "hostlaskdflakjdf",
		Key:   []byte("secret key %78&f75r.!"),
		// 过期时间
		Timeout:     time.Hour * 24 * 3,
		MaxRefresh:  time.Hour * 23 * 20,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.LoginedUserInfo); ok {
				return jwt.MapClaims{
					identityKey:  v.Username,
					"Group":      v.Group,
					"ExpireTime": v.ExpireTime,
					"UID":        v.UID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			c.Set("UID", claims["UID"])
			c.Set("username", claims[identityKey])
			c.Set("Group", claims["Group"])
			return &model.User{
				UserName: claims[identityKey].(string),
			}
		},
		//验证登录
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginService services.LoginService
			if err := c.ShouldBind(&loginService); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			//
			if loginedUser, err := loginService.Login(); err != nil {
				return nil, errors.New("权限验证错误")
			} else {
				return loginedUser, nil
			}
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			if code == 200 {
				code = 0
			}
			c.SetCookie("jwt", token, 86400, "/", "*", false, true)
			c.JSON(200, gin.H{
				"code": code,
				"data": gin.H{
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				},
				"err": "",
				"msg": "",
			})
		},
		//鉴权
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(200, serializer.ErrorResponse(errors.New("权限验证失败")))
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
