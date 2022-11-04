package jwt_mod

import (
	"fmt"
	"github.com/bravedu/brave-go-factory/app/controller"
	"github.com/bravedu/brave-go-factory/config"
	"github.com/bravedu/brave-go-factory/constants"
	"github.com/bravedu/brave-go-factory/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	EnvDev  = "dev"
	EnvTest = "test"
)

type JwtMod struct {
	Env     string
	UserKey string
}

func NewJwtMod(env string, userKey string) *JwtMod {
	obj := &JwtMod{
		Env:     env,
		UserKey: "uid",
	}
	if userKey != "" {
		obj.UserKey = userKey
	}
	return obj
}

func (j *JwtMod) JwtAuthGinHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if j.Env == EnvDev || j.Env == EnvTest {
			if devUid := c.Query(fmt.Sprintf("%s_%s", EnvDev, j.UserKey)); devUid != "" {
				c.Set(j.UserKey, devUid)
				c.Next()
				return
			}
		}
		v := c.Request.Header.Get("Authorization")
		if s := strings.Split(v, " "); len(s) == 2 && s[0] == "Bearer" {
			token := s[1]
			j.JwtAuthDecrypt(token, c)
		} else {
			c.JSON(http.StatusUnauthorized, controller.WriteResponse(constants.RequestTokenFail, nil))
			c.Abort()
		}
	}
}

func (j *JwtMod) JwtAuthDecrypt(token string, c *gin.Context) {
	jwt := util.NewJWT(config.Conf.YamlDao.Jwt.JwtSecret)
	claims, err := jwt.ParseToken(token)
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			//用旧token换新token；超过刷新token有效期，要求前端重新登录
			c.JSON(http.StatusUnauthorized, controller.WriteResponse(constants.RequestTokenFail, nil))
			c.Abort()
		} else {
			c.JSON(http.StatusInternalServerError, controller.WriteResponse(http.StatusInternalServerError, err))
			c.Abort()
		}
	} else {
		uid := claims.UID
		c.Set("uid", uid)
		c.Next()
	}
}
