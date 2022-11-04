package api

import (
	"github.com/bravedu/brave-go-factory/app/controller"
	"github.com/bravedu/brave-go-factory/config"
	"github.com/bravedu/brave-go-factory/docs"
	"github.com/bravedu/brave-go-factory/modules/jwt_mod"
	"os"
)

func registerRoute() {
	authJwt := jwt_mod.NewJwtMod(os.Getenv("APP_ENV"), config.Conf.YamlDao.Jwt.JwtUserKey)
	// v1群组对任何人开放
	docs.SwaggerInfo.BasePath = "/v1"
	v1 := router.Group("/v1")
	{
		v1.GET("hello-lists", controller.HelloList)
	}
	// v2群组使用中间件JwtAuth，需要token权限才能请求到
	docs.SwaggerInfo.BasePath = "/v2"
	v2 := router.Group("/v2", authJwt.JwtAuthGinHandlerFunc())
	{
		v2.POST("hello-v2", controller.HelloV2)
	}
}
