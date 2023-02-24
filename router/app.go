package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)

	userRouter := r.Group("/user")
	{
		userRouter.POST("/getUserInfo", service.User{}.GetUser)
		userRouter.GET("/getUserList", service.User{}.GetUserList)
		userRouter.POST("/createUser", service.User{}.CreateUser)
		userRouter.PUT("/updateUser", service.User{}.UpdateUser)
		userRouter.DELETE("/deleteUser", service.User{}.DeleteUser)
	}

	return r
}
