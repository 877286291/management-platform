package service

import (
	"encoding/base64"
	"errors"
	"gorm.io/gorm"
	"management-platform/model"
	"management-platform/model/request"
	"management-platform/util"
)
import "management-platform/db"

func Login(loginRequest request.LoginRequest) (bool, error) {
	var systemUser model.SystemUser
	result, err := util.DesEncrypt([]byte(loginRequest.LoginPassword))
	if err != nil {
		return false, err
	}
	loginPassword := base64.StdEncoding.EncodeToString(result)
	err = db.Conn.Where("login_name=? and login_password=?", loginRequest.LoginName, loginPassword).Take(&systemUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil

}
