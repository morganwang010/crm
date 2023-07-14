package routes

import (
	"github.com/gin-gonic/gin"
	"gogin/api"
)

// 路由函数，用于管理各个路由的注册
func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("captchaId", api.GetCaptchaId)
	}
	return r
}
