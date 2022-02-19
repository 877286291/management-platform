package service

import (
	"errors"
	"gorm.io/gorm"
	"management-platform/db"
	"management-platform/model"
	"management-platform/util"
	"time"
)

func GetUserList(wrapper *model.SystemUser) ([]model.SystemUser, error) {
	var systemUserList []model.SystemUser
	err := db.Conn.Select("id,type,login_name,name,sex,phone,last_login_time").Where(wrapper).Find(&systemUserList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []model.SystemUser{}, nil
	} else if err != nil {
		return nil, err
	}
	return systemUserList, nil
}
func GetSystemUserInfoById(id string) (model.SystemUser, error) {
	var systemUser model.SystemUser
	if err := db.Conn.Select("id,type,login_name,name,sex,phone,last_login_time").Find(&systemUser, "id=?", id).Error; err != nil {
		return systemUser, err
	}
	return systemUser, nil
}
func AddSystemUser(systemUser *model.SystemUser) (bool, error) {
	//id := uuid.New().String()
	//systemUser.ID = id
	encodePassword, err := util.DesEncrypt([]byte(systemUser.LoginPassword))
	if err != nil {
		return false, err
	}
	systemUser.LoginPassword = string(encodePassword)
	systemUser.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	result := db.Conn.Create(systemUser)
	if result.RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}
