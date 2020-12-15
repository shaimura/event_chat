package model

import (
	"github.com/jinzhu/gorm"
)

// gorm.ModelはID, CreatedAt, UpdatedAt, DeletedAtをフィールドに持つ構造体
// "gorm.Model"のIDはuint型で登録されている

type User struct {
	gorm.Model
	Username    string `gorm:"not null;size:40" validata:"required"`
	Password    string `gorm:"not null"`
	Email       string `gorm:"not null" validate:"required"`
	Usertoken   string `gorm:"unique;not null"`
	Avatarimage string `gorm:"default:'unkown.png'"`
}

type Accesstoken struct {
	UserID         uint   `gorm:"not null"`
	Username       string `gorm:"not null"`
	Accesstoken    string `gorm:"not null"`
	Expirationdata int64  `gorm:"not null"` // 1970/01/01 00:00:00 からの経過ミリ秒で保存(有効時間)
}

// ユーザーチャットのモデル
type Userchatroom struct {
	gorm.Model
	Firstuserid  uint
	Seconduserid uint
}

type UserMessage struct {
	gorm.Model
	UserchatroomID uint
	UserID         uint
	Message        string
}
