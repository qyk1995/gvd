package api

import (
	"errors"
	"fmt"

	"example.com/m/initialize"
	"example.com/m/utils"
	uuid "github.com/satori/go.uuid"

	"example.com/m/model"

	"gorm.io/gorm"
)

type UserService struct{}

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	db := initialize.Gormmysql()

	if nil == db {
		return fmt.Errorf("db not init"), nil
	}
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password)) //md5 解密
	err = db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user

}

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	db := initialize.Gormmysql()
	if !db.Migrator().HasTable(&user) {
		db.AutoMigrate(&user)
	}
	if !errors.Is(db.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = db.Create(&u).Error
	return err, u
}
