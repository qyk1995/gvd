package route

import (
	"example.com/m/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func InitUserRouter() (route *gin.Engine) {
	router := gin.Default()
	v1 := router.Group("user")
	baseApi := service.BaseApi{}
	{
		v1.POST("user_register", baseApi.Register)
		v1.POST("user_login", baseApi.Login)

	}
	return router

}
