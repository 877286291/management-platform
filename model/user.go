package model

type SystemUser struct {
	ID               string `gorm:"column:id;default:-" db:"id" json:"id,omitempty" form:"id"`                                         //  主键id
	Type             string `gorm:"column:type" db:"type" json:"type,omitempty" form:"type"`                                                            //  类型
	LoginName        string `gorm:"column:login_name" db:"login_name" json:"login_name,omitempty" form:"login_name" binding:"required"`                 //  登录名称
	LoginPassword    string `gorm:"column:login_password" db:"login_password" json:"login_password,omitempty" form:"login_password" binding:"required"` //  登录密码
	Name             string `gorm:"column:name" db:"name" json:"name,omitempty" form:"name" binding:"required"`                                         //  姓名
	Sex              string `gorm:"column:sex" db:"sex" json:"sex,omitempty" form:"sex"`                                                                //  性别
	Phone            string `gorm:"column:phone" db:"phone" json:"phone,omitempty" form:"phone"`                                                        //  手机号
	Email            string `gorm:"column:email" db:"email" json:"email,omitempty" form:"email"`                                                        //  电子邮箱
	BirthDate        string `gorm:"column:birth_date" db:"birth_date" json:"birth_date,omitempty" form:"birth_date"`                                    //  出生日期
	LiveAddress      string `gorm:"column:live_address" db:"live_address" json:"live_address,omitempty" form:"live_address"`                            //  居住地
	HeadAddress      string `gorm:"column:head_address" db:"head_address" json:"head_address,omitempty" form:"head_address"`                            //  头像地址
	LastLoginTime    string `gorm:"column:last_login_time" db:"last_login_time" json:"last_login_time,omitempty" form:"last_login_time"`
	Browser          string `gorm:"column:browser" db:"browser" json:"browser,omitempty" form:"browser"`
	Os               string `gorm:"column:os" db:"os" json:"os,omitempty" form:"os"`
	Ipaddr           string `gorm:"column:ipaddr" db:"ipaddr" json:"ipaddr,omitempty" form:"ipaddr"`
	IpRealAddr       string `gorm:"column:ip_real_addr" db:"ip_real_addr" json:"ip_real_addr,omitempty" form:"ip_real_addr"`
	Status           string `gorm:"column:status" db:"status" json:"status,omitempty" form:"status"`                                 //  状态
	OrderBy          string `gorm:"column:order_by" db:"order_by" json:"order_by,omitempty" form:"order_by"`                         //  排序
	OpenId           string `gorm:"column:open_id" db:"open_id" json:"open_id,omitempty" form:"open_id"`                             //  微信id
	InitPassword     string `gorm:"column:init_password" db:"init_password" json:"init_password,omitempty" form:"init_password"`     //  是否初始密码
	PwdErrorNum      string `gorm:"column:pwd_error_num" db:"pwd_error_num" json:"pwd_error_num,omitempty" form:"pwd_error_num"`     //  密码输入错误次数
	PwdErrorTime     string `gorm:"column:pwd_error_time" db:"pwd_error_time" json:"pwd_error_time,omitempty" form:"pwd_error_time"` //  密码输入错误时间
	ThemeId          string `gorm:"column:theme_id" db:"theme_id" json:"theme_id,omitempty" form:"theme_id"`
	DepartLeader     string `gorm:"column:depart_leader" db:"depart_leader" json:"depart_leader,omitempty" form:"depart_leader"`
	DirectLeader     string `gorm:"column:direct_leader" db:"direct_leader" json:"direct_leader,omitempty" form:"direct_leader"`
	BranchLeader     string `gorm:"column:branch_leader" db:"branch_leader" json:"branch_leader,omitempty" form:"branch_leader"`
	Token            string `gorm:"column:token" db:"token" json:"token,omitempty" form:"token"`
	TokenRefreshTime string `gorm:"column:token_refresh_time" db:"token_refresh_time" json:"token_refresh_time,omitempty" form:"token_refresh_time"`
	Remark           string `gorm:"column:remark" db:"remark" json:"remark,omitempty" form:"remark"`                                     //  备注
	CreateTime       string `gorm:"column:create_time" db:"create_time" json:"create_time,omitempty" form:"create_time"`                 //  创建时间
	CreateUser       string `gorm:"column:create_user" db:"create_user" json:"create_user,omitempty" form:"create_user"`                 //  创建人
	CreateOrganize   string `gorm:"column:create_organize" db:"create_organize" json:"create_organize,omitempty" form:"create_organize"` //  创建组织
	UpdateUser       string `gorm:"column:update_user" db:"update_user" json:"update_user,omitempty" form:"update_user"`                 //  修改人
	UpdateTime       string `gorm:"column:update_time" db:"update_time" json:"update_time,omitempty" form:"update_time"`                 //  修改时间
	BirthAddress     string `gorm:"column:birth_address" db:"birth_address" json:"birth_address,omitempty" form:"birth_address"`
	Motto            string `gorm:"column:motto" db:"motto" json:"motto,omitempty" form:"motto"`
}

func (s SystemUser) TableName() string {
	return "system_user"
}
func NewSystemUser() *SystemUser {
	return new(SystemUser)
}
