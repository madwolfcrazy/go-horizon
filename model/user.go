package model

//User 用户表
type User struct {
	UID      uint   `gorm:"column:uid;primary_key;auto_increment:true"`
	UserName string `gorm:"column:username"`
	Password string
}
