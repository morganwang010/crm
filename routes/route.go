package routes

import (
	"github.com/gin-gonic/gin"
	"gogin/api"
	"gogin/model"
	"gogin/services"
)

// 路由函数，用于管理各个路由的注册
func NewRouter() *gin.Engine {
	r := gin.Default()
	cm := model.NewContactModel(model.DB)
	contactService := services.NewContactService(cm)

	contactAPI := api.NewContactAPI(contactService)
	// model.DataBase()

	// contactService := services.NewContactService(model.ContactModel)
	// api2 := api.GetAllContactsHandler(contactService)

	v1 := r.Group("api/v1")
	{
		v1.GET("/captchaId", api.GetCaptchaId)
		v1.GET("/bb", api.GetUserbyName)
		// v1.GET("/dd", api.GetAllContactsHandler)

		contact := v1.Group("/contact")
		{
			contact.GET("", contactAPI.GetAllContactsHandler)
		}
	}

	return r
}
