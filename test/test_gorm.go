package main

import (
	"ginchat/models"
	"ginchat/utils"
	"time"
)

func main() {
	// Migrate the schema
	utils.DB.AutoMigrate(&models.UserBasic{})
	user := &models.UserBasic{LoginTime: time.Now(), LogoutTime: time.Now(), HeartbeatTime: time.Now()}
	user.Name = "测试"
	// Create
	utils.DB.Create(user)

	// Read
	// var models.UserBasic models.UserBasic
	utils.DB.First(user, 1) // find models.UserBasic with integer primary key
	// utils.DB.First(user, "code = ?", "D42") // find models.UserBasic with code D42

	// Update - update models.UserBasic's price to 200
	utils.DB.Model(user).Update("PassWord", "123456")
	// Update - update multiple fields
	// utils.DB.Model(user).Updates(&models.UserBasic{Name: "测试"}) // non-zero fields
	// utils.DB.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete models.UserBasic
	// utils.DB.Delete(user, 1)
}
