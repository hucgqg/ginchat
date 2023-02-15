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
	r.GET("/user/getUserList", service.User{}.GetUserList)
	r.POST("/user/createUser", service.User{}.CreateUser)
	r.PUT("/user/updateUser", service.User{}.UpdateUser)
	r.DELETE("/user/deleteUser", service.User{}.DeleteUser)
	return r
}
