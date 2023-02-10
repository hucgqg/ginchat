package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

func InitMySQL() {
	m := viper.GetStringMapString("mysql")
	DB, Err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", m["user"], m["password"], m["host"], m["prot"], m["database"])), &gorm.Config{})
	if Err != nil {
		panic("failed to connect database")
	}
}
