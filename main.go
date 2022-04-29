package main

import (
	"example.com/m/route"
)

func main() {
	// router := gin.Default()
	// v1 := router.Group("user")
	// baseApi := service.BaseApi{}
	// {
	// 	v1.POST("user_register", baseApi.Register)
	// 	// userRouterWithoutRecord.POST("user_login", baseApi.Login) //

	// }
	// router.Run(":8080")
	rout := route.InitUserRouter()
	rout.Run(":8088")
}
