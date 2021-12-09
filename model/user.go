package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


const (
	passWordCost = 12
)

//User 用户表
type User struct {
	gorm.Model
	UserName   string `gorm:"column:username;type:char(255);index" json:"username"`
	Password   string `gorm:"column:password;type:char(255)" json:"password"`
	Group      string `gorm:"column:group;type:char(32)" json:"group"`
	ExpireTime int    `gorm:"column:expire_time;type:int(16)" json:"expire_time"`
}

func (User) TableName() string {
	return "user"
}

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

//ChangePassword 修改密码
func (user *User) ChangePassword(newPassword string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(newPassword), passWordCost)
	if err != nil {
		return err
	}
	return DB.Model(&user).Update("password", string(bytes)).Error
}

//GetUser 查询用户
func GetUser(uid uint) (existUser *User, err error) {
	existUser = &User{}
	err = DB.Where("uid = ?", uid).First(existUser).Error
	if err != nil {
		return nil, err
	}
	return
}

type LoginedUserInfo struct {
	UID        uint      `json:"uid"`
	Username   string    `json:"username"`
	CreateAt   time.Time `json:"create_at"`
	Group      string    `json:"group"`
	ExpireTime int       `gorm:"column:expire_time" json:"expire_time"`
}

//ToLoginedUserInfo User Model 转换为 用户登录后信息
func (u *User) ToLoginedUserInfo() *LoginedUserInfo {
	l := LoginedUserInfo{
		UID:        u.ID,
		Username:   u.UserName,
		CreateAt:   u.CreatedAt,
		Group:      u.Group,
		ExpireTime: u.ExpireTime,
	}
	return &l
}

