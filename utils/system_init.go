package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	// 自定义日志模板，打印Sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\t", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色日志
		},
	)
	m := viper.GetStringMapString("mysql")
	DB, Err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", m["user"], m["password"], m["host"], m["prot"], m["database"])), &gorm.Config{Logger: newLogger})
	if Err != nil {
		panic("failed to connect database")
	}
}
