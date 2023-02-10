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
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Identity      string    `json:"identity"`
	ClientIP      string    `json:"client_ip"`
	ClientPort    string    `json:"client_port"`
	LoginTime     time.Time `json:"login_time"`
	LogoutTime    time.Time `json:"logout_time"`
	HeartbeatTime time.Time `json:"heartbeat_time"`
	IsLogout      bool      `json:"is_logout"`
	DeviceInfo    string    `json:"device_info"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}
