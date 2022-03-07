package model

type SystemUserRole struct {
	ID         string `gorm:"column:id" db:"id" json:"id" form:"id"`
	RoleId     string `gorm:"column:role_id" db:"role_id" json:"role_id" form:"role_id"`
	UserId     string `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	CreateUser string `gorm:"column:create_user" db:"create_user" json:"create_user" form:"create_user"`
	CreateTime string `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
	UpdateTime string `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
}

func (s SystemUserRole) TableName() string {
	return "system_user_role"
}
