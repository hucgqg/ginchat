package main

import (
	"ginchat/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(172.17.73.120:33006)/ginchat?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})
	user := &models.UserBasic{LoginTime: time.Now(), LogoutTime: time.Now(), HeartbeatTime: time.Now()}
	user.Name = "测试"
	// Create
	db.Create(user)

	// Read
	// var models.UserBasic models.UserBasic
	db.First(user, 1) // find models.UserBasic with integer primary key
	// db.First(user, "code = ?", "D42") // find models.UserBasic with code D42

	// Update - update models.UserBasic's price to 200
	db.Model(user).Update("PassWord", "123456")
	// Update - update multiple fields
	// db.Model(user).Updates(&models.UserBasic{Name: "测试"}) // non-zero fields
	// db.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete models.UserBasic
	// db.Delete(user, 1)
}
