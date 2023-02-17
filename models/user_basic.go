package models

import (
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name"`
	PassWord      string    `json:"password"`
	Phone         string    `json:"phone" valid:"matches(^1[3-9]\\d{9})"`
	Email         string    `json:"email" valid:"email"`
	Identity      string    `json:"identity"`
	ClientIP      string    `json:"client_ip"`
	ClientPort    string    `json:"client_port"`
	LoginTime     time.Time `json:"login_time"`
	LogoutTime    time.Time `json:"logout_time"`
	HeartbeatTime time.Time `json:"heartbeat_time"`
	IsLogout      bool      `json:"is_logout"`
	DeviceInfo    string    `json:"device_info"`
	Salt          string    `json:"salt"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}

func CreateUser(user *UserBasic) *gorm.DB {
	user.LoginTime, user.LogoutTime, user.HeartbeatTime = time.Now(), time.Now(), time.Now()
	return utils.DB.Create(&user)
}

func DeleteUser(user *UserBasic) *gorm.DB {
	return utils.DB.Where("id=?", user.ID).Delete(&user)
}

func UpdateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{
		Name:     user.Name,
		PassWord: user.PassWord,
		Phone:    user.Phone,
		Email:    user.Email,
	})
}
