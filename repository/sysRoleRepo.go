package repository

import "management-platform/model"

type ISysRoleRepo interface {
	ListRole(page, size int, total *int64, where interface{}) ([]*model.SystemRole, error)
	InsertRole(*model.SystemRole) (bool, error)
	DeleteRole(id string) (bool, error)
	UpdateRole(id string, role *model.SystemRole) (bool, error)
}
type SysRoleRepo struct {
	BaseRepo *BaseRepo `inject:""`
}

func (repo *SysRoleRepo) ListRole(page, size int, total *int64, where interface{}) ([]*model.SystemRole, error) {
	var roles []*model.SystemRole
	err := repo.BaseRepo.GetPages(&model.SystemRole{}, &roles, "*", page, size, total, where)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
func (repo *SysRoleRepo) InsertRole(role *model.SystemRole) (bool, error) {
	err := repo.BaseRepo.Create(role)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (repo *SysRoleRepo) DeleteRole(id string) (bool, error) {
	role := new(model.SystemRole)
	where := &model.SystemRole{ID: id}
	count, err := repo.BaseRepo.DeleteByWhere(role, where)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (repo *SysRoleRepo) UpdateRole(id string, role *model.SystemRole) (bool, error) {
	if err := repo.BaseRepo.Conn.DB().Model(&role).Where("id = ?", id).Updates(role).Error; err != nil {
		return false, err
	}
	return true, nil
}
