package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	// 生产环境设置
	// gin.SetMode(gin.ReleaseMode)

	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()

	r := router.Router()

	r.Run()
}
