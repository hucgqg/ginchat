package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name"`
	PassWord      string    `json:"-"`
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
	Salt          string    `json:"-"`
}

func (table *UserBasic) TableName() string { return "user_basic" }

func FindUserBynameAndPwd(name, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name=? AND pass_word=?", name, password).First(&user)

	// token加密
	str := fmt.Sprintf("%d", time.Now().Nanosecond())
	token := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id=?", user.ID).Update("identity=?", token)
	return user
}

func FindUserByname(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name=?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("phone=?", phone).First(&user)
	return user
}

func FindUserByEmail(email string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("phone=?", email).First(&user)
	return user
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
