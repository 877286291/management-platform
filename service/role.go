package service

import (
	"github.com/google/uuid"
	"management-platform/model"
	"management-platform/repository"
	"time"
)

type ISysRoleService interface {
	GetRoleList(page, size int, total *int64, where interface{}) ([]*model.SystemRole, error)
	AddSystemRole(systemRole *model.SystemRole) (bool, error)
	DeleteSystemRole(id string) (bool, error)
	UpdateSystemRole(id string, systemRole *model.SystemRole) (bool, error)
}
type SysRoleService struct {
	RoleRepo repository.ISysRoleRepo `inject:""`
}

func (s *SysRoleService) GetRoleList(page, size int, total *int64, where interface{}) ([]*model.SystemRole, error) {
	roles, err := s.RoleRepo.ListRole(page, size, total, where)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
func (s *SysRoleService) AddSystemRole(systemRole *model.SystemRole) (bool, error) {
	systemRole.ID = uuid.New().String()
	systemRole.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	// TODO add create user
	//systemRole.CreateUser = ""
	role, err := s.RoleRepo.InsertRole(systemRole)
	if err != nil {
		return false, err
	}
	return role, nil
}
func (s *SysRoleService) DeleteSystemRole(id string) (bool, error) {
	result, err := s.RoleRepo.DeleteRole(id)
	if err != nil {
		return false, err
	}
	return result, nil
}
func (s *SysRoleService) UpdateSystemRole(id string, systemRole *model.SystemRole) (bool, error) {
	systemRole.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	// TODO add update user
	result, err := s.RoleRepo.UpdateRole(id, systemRole)
	if err != nil {
		return false, err
	}
	return result, nil
}
