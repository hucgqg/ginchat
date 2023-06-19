package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Redis    *redis.Client
	DB       *gorm.DB
	err      error
	redisCtx = context.Background()
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
	})
	pong, err := Redis.Ping(redisCtx).Result()
	if err != nil {
		fmt.Println("connect redis failed:", err)
	} else {
		fmt.Println("connect redis ok:", pong)
	}
}

func InitMySQL() {
	// 自定义日志模板，打印Sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "[SQL] ", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色日志
		},
	)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.prot"),
		viper.GetString("mysql.database"),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
}
