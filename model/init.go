package model

import (
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB connection single unit
var DB *gorm.DB

// InitDatabase init function
func InitDatabase(connString string) {
	DB, err := gorm.Open("mysql", connString)
	// 开发调试模式可打开查看 gorm 日志
	//db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	DB.DB().SetMaxIdleConns(10)
	//打开
	DB.DB().SetMaxOpenConns(100)
	//超时
	DB.DB().SetConnMaxLifetime(time.Minute * 30)
	// 设置自动建立表时的字符集
	DB = DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")

	DB.AutoMigrate(&User{})
}
