package service

import (
	"encoding/base64"
	"management-platform/model/request"
	"management-platform/repository"
	"management-platform/utils"
)

type IAuthService interface {
	Login(loginRequest request.LoginRequest) (bool, error)
}
type AuthService struct {
	UserRepo repository.ISysUserRepo `inject:""`
}

func (service *AuthService) Login(loginRequest request.LoginRequest) (bool, error) {
	result, err := utils.DesEncrypt([]byte(loginRequest.LoginPassword))
	if err != nil {
		return false, err
	}
	loginPassword := base64.StdEncoding.EncodeToString(result)
	row, err := service.UserRepo.CheckUser(loginRequest.LoginName, loginPassword)
	if err != nil {
		return false, err
	} else if !row {
		return false, nil
	}
	return true, nil
}
