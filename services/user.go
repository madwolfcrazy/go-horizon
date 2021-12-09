package services

import (
	"errors"
	"ymz465/go-horizon/model"

	"gorm.io/gorm"
)

type CreateUser struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
	ExpireTime int    `json:"expire_time"`
	Group      string `json:"group"`
}

//Create 新建用户
func (c *CreateUser) Create() (*model.User, error) {
	if c.Password != c.RePassword {
		return nil, errors.New("两次密码输入不一致")
	}
	newUser := model.User{}
	// found user exists ?
	err := model.DB.Where("username = ?", c.Username).First(&newUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newUser = model.User{
			UserName:   c.Username,
			ExpireTime: c.ExpireTime,
		}
		newUser.SetPassword(c.Password)
		err := model.DB.Create(&newUser).Error
		if err != nil {
			return nil, err
		}
	}
	newUser.Password = "-"
	return &newUser, nil
}

//DeleteUser 删除用户
type DeleteUser struct {
	UID uint
}

//Delete 删除
func (d *DeleteUser) Delete() error {
	return model.DB.Model(model.User{}).Delete("id=?", d.UID).Error
}

//ChangePassword 修改密码
type ChangePassword struct {
	UID         uint   `json:"uid"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

//ChangePassword 修改密码
func (c *ChangePassword) ChangePassword() error {
	existsUser := model.User{}
	err := model.DB.Where("id = ?", c.UID).First(&existsUser).Error
	if err != nil {
		return err
	}
	if !existsUser.CheckPassword(c.OldPassword) {
		return errors.New("旧密码输入错误")
	}
	return existsUser.ChangePassword(c.NewPassword)
}

//LoginService 登录服务
type LoginService struct {
	Username string `json:"username"` //
	Password string `json:"password"` //
}

//Login 登录
func (l *LoginService) Login() (*model.LoginedUserInfo, error) {
	loginedUser := model.User{}
	err := model.DB.Where("username = ?", l.Username).First(&loginedUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名或密码错误")
	}
	if !loginedUser.CheckPassword(l.Password) {
		return nil, errors.New("用户名或密码错误")
	}
	loginedUser.Password = "-"
	return loginedUser.ToLoginedUserInfo(), nil
}
