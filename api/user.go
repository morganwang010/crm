package api

import (
	"context"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"gogin/model"
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
	users := []model.User{}
	// db.Where(&user).First(&user)
	// result = db.Table("dr.user").Select(`*`)
	// result := db.Table("dr.user").Select(`*`).Where("employee_number = ?", "wangjw30w").Find(&users)
	result := model.DB.Table("dr.user").Select(`*`).Find(&users)
	// db.Find(&user)
	if result.Error != nil {
		// 处理错误
		fmt.Println("user.RoleName")
	}
	for _, user := range users {
		fmt.Println(user.EmployeeNumber)
		fmt.Println(user.RoleName)
	}
}
func FindAllContact(c *gin.Context) {
	m := &model.ContactModel{} // 创建 ContactModel 实例

	ctx := context.Background() // 创建上下文对象

	contacts, err := m.FindAllContact(ctx) // 调用 FindAllContact 函数

	if err != nil {
		// 处理错误
		fmt.Println("user.RoleName")
	}
	for _, contact := range contacts {
		fmt.Println(contact.Company)
		fmt.Println(contact.Focal)
	}
}

func FindAllContact2(c *gin.Context) {

	var contacts []model.Contact
	result := model.DB.Table("contact").Select(`*`).Find(&contacts)
	if result.Error != nil {
		// 处理错误
		fmt.Println("user.RoleName")
	}
	for _, contact := range contacts {
		fmt.Println(contact.Company)
		fmt.Println(contact.Focal)
	}
	// return result
}
