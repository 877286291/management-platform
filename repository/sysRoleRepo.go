package repository

type ISysRoleRepo interface {
}
type SysRoleRepo struct {
	BaseRepo *BaseRepo `inject:""`
}
