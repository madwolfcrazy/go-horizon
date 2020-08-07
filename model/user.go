package model

import "golang.org/x/crypto/bcrypt"

const (
	passWordCost = 12
)

//User 用户表
type User struct {
	UID      uint   `gorm:"column:uid;primary_key;auto_increment:true"`
	UserName string `gorm:"column:username"`
	Password string
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
