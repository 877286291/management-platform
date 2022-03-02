package repository

import (
	"management-platform/model"
)

type ISysUserRepo interface {
	GetUserById(id string) (*model.SystemUser, error)
	GetUserByLoginName(loginName string) (*model.SystemUser, error)
	CheckUser(loginName, password string) (bool, error)
	InsertUser(user *model.SystemUser) (bool, error)
	UpdateUser(id string, user *model.SystemUser) (bool, error)
	DeleteUser(id string) (bool, error)
	ListUser(page, size int, total *int64, where interface{}) ([]*model.SystemUser, error)
	EnableUser(userId int32) (bool, error)
	//SaveUserRole(userId, roleId int32) (bool, error)
	//DelUserRole(userId int32) (bool, error)
}
type SysUserRepo struct {
	BaseRepo *BaseRepo `inject:""`
}

func (repo *SysUserRepo) GetUserById(id string) (*model.SystemUser, error) {
	var user model.SystemUser
	if err := repo.BaseRepo.First(id, &user, "id,type,login_name,name,sex,phone,last_login_time"); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *SysUserRepo) GetUserByLoginName(loginName string) (*model.SystemUser, error) {
	var user model.SystemUser
	where := model.SystemUser{LoginName: loginName}
	if err := repo.BaseRepo.First(&where, &user, "*"); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *SysUserRepo) CheckUser(loginName, password string) (bool, error) {
	var count int64
	if err := repo.BaseRepo.Conn.DB().Table("system_user").Where("login_name = ? and login_password = ?", loginName, password).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *SysUserRepo) InsertUser(user *model.SystemUser) (bool, error) {
	if err := repo.BaseRepo.Create(user); err != nil {

		return false, err
	}
	return true, nil
}

func (repo *SysUserRepo) UpdateUser(id string, user *model.SystemUser) (bool, error) {
	if err := repo.BaseRepo.Conn.DB().Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *SysUserRepo) DeleteUser(id string) (bool, error) {
	user := model.SystemUser{}
	where := &model.SystemUser{ID: id}
	if count, err := repo.BaseRepo.DeleteByWhere(&user, where); err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (repo *SysUserRepo) ListUser(page, size int, total *int64, where interface{}) ([]*model.SystemUser, error) {
	var users []*model.SystemUser
	if err := repo.BaseRepo.GetPages(&model.SystemUser{}, &users, "id,type,login_name,name,sex,phone,last_login_time", page, size, total, where); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *SysUserRepo) EnableUser(userId int32) (bool, error) {
	if err := repo.BaseRepo.Conn.DB().Exec("update system_user set is_enable = abs(is_enable - 1) where user_id = ?", userId).Error; err != nil {
		return false, err
	}
	return true, nil
}

//func (repo *SysUserRepo) SaveUserRole(userId, roleId int32) (bool, error) {
//	if err := repo.BaseRepo.Conn.Exec("insert into t_sys_user_role(user_id, role_id) values (?, ?)", userId, roleId).Error; err != nil {
//		return false, err
//	}
//	return true, nil
//}
//
//func (repo *SysUserRepo) DelUserRole(userId int32) (bool, error) {
//	if err := repo.BaseRepo.Conn.Exec("delete from t_sys_user_role where user_id = ?", userId).Error; err != nil {
//		return false, err
//	}
//	return true, nil
//}
