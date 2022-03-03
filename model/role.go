package model

type SystemRole struct {
	ID         string `gorm:"column:id" db:"id" json:"id,omitempty" form:"id"`
	Type       string `gorm:"column:type" db:"type" json:"type,omitempty" form:"type"`
	Name       string `gorm:"column:name" db:"name" json:"name,omitempty" form:"name"`
	ShortName  string `gorm:"column:short_name" db:"short_name" json:"short_name,omitempty" form:"short_name"`
	Status     string `gorm:"column:status" db:"status" json:"status,omitempty" form:"status"`
	OrderBy    string `gorm:"column:order_by" db:"order_by" json:"order_by,omitempty" form:"order_by"`
	Remark     string `gorm:"column:remark" db:"remark" json:"remark,omitempty" form:"remark"`
	CreateTime string `gorm:"column:create_time" db:"create_time" json:"create_time,omitempty" form:"create_time"`
	CreateUser string `gorm:"column:create_user" db:"create_user" json:"create_user,omitempty" form:"create_user"`
	UpdateUser string `gorm:"column:update_user" db:"update_user" json:"update_user,omitempty" form:"update_user"`
	UpdateTime string `gorm:"column:update_time" db:"update_time" json:"update_time,omitempty" form:"update_time"`
}

func (s SystemRole) TableName() string {
	return "system_role"
}
