package service

import (
	"awesomeProject/db"
	"gorm.io/gorm"
	"time"
)

var connection *gorm.DB = db.GetConnection()

type UserInfo struct {
	OpenID          string    `json:"open_id"`
	AvatarURL       string    `json:"avatar_url"`
	IsManager       db.MyBool `json:"is_manager"`
	LatestLoginDate time.Time `json:"latest_login_date"`
	Nickname        string    `json:"nickname"`
	RealName        string    `json:"real_name"`
	UnionID         string    `json:"union_id"`
	Username        string    `json:"username"`
}

func (UserInfo) TableName() string { return "rs_user_info" }

func QueryUserById(userId string) (UserInfo, error) {
	var userInfo UserInfo
	result := connection.Where("open_id = ?", userId).First(&userInfo)
	return userInfo, result.Error
}

func FindManager() ([]UserInfo, error) {
	var user []UserInfo
	result := connection.Where("is_manager = ?", true).Find(&user)
	return user, result.Error
}

func QueryAllUserInfo() ([]UserInfo, error) {
	var user []UserInfo
	result := connection.Find(&user)
	return user, result.Error
}
