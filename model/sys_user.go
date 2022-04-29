package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`    // 用户UUID
	Username string    `json:"userName" gorm:"comment:用户登录名"` // 用户登录名
	Password string    `json:"-"  gorm:"comment:用户登录密码"`      // 用户登录密码
	Phone    string    `json:"phone"  gorm:"comment:用户手机号"`   // 用户手机号
	Email    string    `json:"email"  gorm:"comment:用户邮箱"`    // 用户邮箱
}

type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码

}
