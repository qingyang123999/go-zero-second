package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:32" json:"username"`
	Password string `gorm:"size:64" json:"password"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (User) TableName() string {
	return "user"
}
