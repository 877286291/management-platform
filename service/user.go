package service

import (
	"encoding/base64"
	"github.com/google/uuid"
	"management-platform/model"
	"management-platform/repository"
	"management-platform/utils"
	"time"
)

type ISysUserService interface {
	GetSysUserById(id string) (*model.SystemUser, error)
	GetSysUserList(page, size int, total *int64, where interface{}) ([]*model.SystemUser, error)
	AddSystemUser(systemUser *model.SystemUser) (bool, error)
}
type SysUserService struct {
	UserRepo repository.ISysUserRepo `inject:""`
}

func (service *SysUserService) GetSysUserList(page, size int, total *int64, where interface{}) ([]*model.SystemUser, error) {
	systemUserList, err := service.UserRepo.ListUser(page, size, total, where)
	if err != nil {
		return nil, err
	}
	return systemUserList, nil
}
func (service *SysUserService) GetSysUserById(id string) (*model.SystemUser, error) {
	systemUser, err := service.UserRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return systemUser, nil
}
func (service *SysUserService) AddSystemUser(systemUser *model.SystemUser) (bool, error) {
	id := uuid.New().String()
	systemUser.ID = id
	password, err := utils.DesEncrypt([]byte(systemUser.LoginPassword))
	if err != nil {
		return false, err
	}
	encodePassword := base64.StdEncoding.EncodeToString(password)
	systemUser.LoginPassword = encodePassword
	systemUser.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	user, err := service.UserRepo.InsertUser(systemUser)
	if err != nil {
		return false, err
	}
	if !user {
		return false, nil
	}
	return true, nil
}
