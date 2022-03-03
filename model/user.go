package model

type SystemUser struct {
	ID               string `gorm:"column:id" database:"id" json:"id,omitempty" form:"id"`                                                 //  主键id
	Type             string `gorm:"column:type" database:"type" json:"type,omitempty" form:"type"`                                         //  类型
	LoginName        string `gorm:"column:login_name" database:"login_name" json:"login_name,omitempty" form:"login_name"`                 //  登录名称
	LoginPassword    string `gorm:"column:login_password" database:"login_password" json:"login_password,omitempty" form:"login_password"` //  登录密码
	Name             string `gorm:"column:name" database:"name" json:"name,omitempty" form:"name"`                                         //  姓名
	Sex              string `gorm:"column:sex" database:"sex" json:"sex,omitempty" form:"sex"`                                             //  性别
	Phone            string `gorm:"column:phone" database:"phone" json:"phone,omitempty" form:"phone"`                                     //  手机号
	Email            string `gorm:"column:email" database:"email" json:"email,omitempty" form:"email"`                                     //  电子邮箱
	BirthDate        string `gorm:"column:birth_date" database:"birth_date" json:"birth_date,omitempty" form:"birth_date"`                 //  出生日期
	LiveAddress      string `gorm:"column:live_address" database:"live_address" json:"live_address,omitempty" form:"live_address"`         //  居住地
	HeadAddress      string `gorm:"column:head_address" database:"head_address" json:"head_address,omitempty" form:"head_address"`         //  头像地址
	LastLoginTime    string `gorm:"column:last_login_time" database:"last_login_time" json:"last_login_time,omitempty" form:"last_login_time"`
	Browser          string `gorm:"column:browser" database:"browser" json:"browser,omitempty" form:"browser"`
	Os               string `gorm:"column:os" database:"os" json:"os,omitempty" form:"os"`
	Ipaddr           string `gorm:"column:ipaddr" database:"ipaddr" json:"ipaddr,omitempty" form:"ipaddr"`
	IpRealAddr       string `gorm:"column:ip_real_addr" database:"ip_real_addr" json:"ip_real_addr,omitempty" form:"ip_real_addr"`
	Status           string `gorm:"column:status" database:"status" json:"status,omitempty" form:"status"`                                 //  状态
	OrderBy          string `gorm:"column:order_by" database:"order_by" json:"order_by,omitempty" form:"order_by"`                         //  排序
	OpenId           string `gorm:"column:open_id" database:"open_id" json:"open_id,omitempty" form:"open_id"`                             //  微信id
	InitPassword     string `gorm:"column:init_password" database:"init_password" json:"init_password,omitempty" form:"init_password"`     //  是否初始密码
	PwdErrorNum      string `gorm:"column:pwd_error_num" database:"pwd_error_num" json:"pwd_error_num,omitempty" form:"pwd_error_num"`     //  密码输入错误次数
	PwdErrorTime     string `gorm:"column:pwd_error_time" database:"pwd_error_time" json:"pwd_error_time,omitempty" form:"pwd_error_time"` //  密码输入错误时间
	ThemeId          string `gorm:"column:theme_id" database:"theme_id" json:"theme_id,omitempty" form:"theme_id"`
	DepartLeader     string `gorm:"column:depart_leader" database:"depart_leader" json:"depart_leader,omitempty" form:"depart_leader"`
	DirectLeader     string `gorm:"column:direct_leader" database:"direct_leader" json:"direct_leader,omitempty" form:"direct_leader"`
	BranchLeader     string `gorm:"column:branch_leader" database:"branch_leader" json:"branch_leader,omitempty" form:"branch_leader"`
	Token            string `gorm:"column:token" database:"token" json:"token,omitempty" form:"token"`
	TokenRefreshTime string `gorm:"column:token_refresh_time" database:"token_refresh_time" json:"token_refresh_time,omitempty" form:"token_refresh_time"`
	Remark           string `gorm:"column:remark" database:"remark" json:"remark,omitempty" form:"remark"`                     //  备注
	CreateTime       string `gorm:"column:create_time" database:"create_time" json:"create_time,omitempty" form:"create_time"` //  创建时间
	CreateUser       string `gorm:"column:create_user" database:"create_user" json:"create_user,omitempty" form:"create_user"` //  创建人
	UpdateUser       string `gorm:"column:update_user" database:"update_user" json:"update_user,omitempty" form:"update_user"` //  修改人
	UpdateTime       string `gorm:"column:update_time" database:"update_time" json:"update_time,omitempty" form:"update_time"` //  修改时间
	BirthAddress     string `gorm:"column:birth_address" database:"birth_address" json:"birth_address,omitempty" form:"birth_address"`
	Motto            string `gorm:"column:motto" database:"motto" json:"motto,omitempty" form:"motto"`
}

func (s SystemUser) TableName() string {
	return "system_user"
}
func NewSystemUser() *SystemUser {
	return new(SystemUser)
}
