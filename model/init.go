package model

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//
)

// DB connection single unit
var DB *gorm.DB

// InitDatabase init function
func InitDatabase(connString string) {
	gormConf := gorm.Config{}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 3, // Slow SQL threshold
			LogLevel:                  logger.Silent,   // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,           // Disable color
		},
	)
	gormConf = gorm.Config{Logger: newLogger}
	if viper.GetString("runmode") == "debug" {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			},
		)
		gormConf = gorm.Config{Logger: newLogger}
	}
	db, err := gorm.Open(mysql.Open(connString), &gormConf)
	// 开发调试模式可打开查看 gorm 日志
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(100)
	//超时
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
	DB = db
	// 设置自动建立表时的字符集
	DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(&User{})
}
