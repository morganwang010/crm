package api

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"gogin/serializer"
)

// GetCaptchaId 获取验证码ID
func GetCaptchaId(c *gin.Context) {
	//验证码长度
	captchaId := captcha.NewLen(4)
	c.JSON(200, serializer.Response{
		Status: 200,
		Data:   captchaId,
		Msg:    "验证码ID",
		Error:  "无",
	})
}

func GetUserbyName(c *gin.Context) {
	//验证码长度
	captchaId := captcha.NewLen(4)
	c.JSON(200, serializer.Response{
		Status: 200,
		Data:   captchaId,
		Msg:    "验证码ID",
		Error:  "无",
	})
}
