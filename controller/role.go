package controller

import (
	"github.com/gin-gonic/gin"
	"management-platform/model"
	"management-platform/model/response"
	"management-platform/service"
	"net/http"
	"strconv"
)

type SysRole struct {
	RoleService service.ISysRoleService `inject:""`
}

func (service *SysRole) GetRoleList(ctx *gin.Context) {
	where := new(model.SystemRole)
	if err := ctx.ShouldBindQuery(where); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请检查请求参数").WithErrMsg(err))
		return
	}
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请求参数错误").WithErrMsg(err))
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请求参数错误").WithErrMsg(err))
		return
	}
	total := new(int64)
	roles, err := service.RoleService.GetRoleList(page, size, total, where)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithMsg("获取角色列表失败").WithErrMsg(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithMsg("获取角色列表成功").WithData(roles).WithPagination(&page, &size, total))
}
func (service *SysRole) AddSystemRole(ctx *gin.Context) {
	systemRole := new(model.SystemRole)
	if err := ctx.ShouldBindJSON(systemRole); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请求参数错误").WithErrMsg(err))
		return
	}
	result, err := service.RoleService.AddSystemRole(systemRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithMsg("添加角色失败").WithErrMsg(err))
		return
	}
	if result {
		ctx.JSON(http.StatusOK, response.OK.WithMsg("添加角色成功"))
		return
	}
	ctx.JSON(http.StatusOK, response.Err.WithMsg("添加角色失败"))
}

func (service *SysRole) DeleteSystemRole(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := service.RoleService.DeleteSystemRole(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithMsg("删除角色失败").WithErrMsg(err))
		return
	}
	if result {
		ctx.JSON(http.StatusOK, response.OK.WithMsg("删除角色成功"))
		return
	}
	ctx.JSON(http.StatusOK, response.Err.WithMsg("删除角色失败"))
}

func (service *SysRole) UpdateSystemRole(ctx *gin.Context) {
	id := ctx.Param("id")
	systemRole := new(model.SystemRole)

	if err := ctx.ShouldBindJSON(systemRole); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请求参数错误").WithErrMsg(err))
		return
	}
	result, err := service.RoleService.UpdateSystemRole(id, systemRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithMsg("更新角色失败"))
		return
	}
	if result {
		ctx.JSON(http.StatusOK, response.OK.WithMsg("更新角色成功"))
		return
	}
	ctx.JSON(http.StatusOK, response.Err.WithMsg("更新角色失败"))
}
