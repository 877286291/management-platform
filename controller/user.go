package controller

import (
	"github.com/gin-gonic/gin"
	"management-platform/model"
	"management-platform/model/response"
	"management-platform/service"
	"net/http"
	"strconv"
)

type SysUser struct {
	UserService service.ISysUserService `inject:""`
}

// GetSystemUserInfoById 通过ID获取系统用户信息
func (sysUser *SysUser) GetSystemUserInfoById(ctx *gin.Context) {
	id := ctx.Param("id")
	systemUser, err := sysUser.UserService.GetSysUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithMsg("获取用户信息失败").WithErrMsg(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithMsg("获取用户信息成功").WithData(systemUser))
}

// GetSystemUserList 获取系统用户列表
func (sysUser *SysUser) GetSystemUserList(ctx *gin.Context) {
	where := model.NewSystemUser()
	if err := ctx.ShouldBindQuery(where); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请求参数错误").WithErrMsg(err))
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
	systemUserList, err := sysUser.UserService.GetSysUserList(page, size, total, where)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithMsg("获取用户列表失败").WithErrMsg(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithMsg("获取用户列表成功").WithData(systemUserList))
}

// AddSystemUser 添加用户
func (sysUser *SysUser) AddSystemUser(ctx *gin.Context) {
	systemUser := model.NewSystemUser()
	if err := ctx.ShouldBindJSON(systemUser); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请求参数错误").WithErrMsg(err))
		return
	}
	result, err := sysUser.UserService.AddSystemUser(systemUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithMsg("用户添加失败").WithErrMsg(err))
		return
	}
	if result {
		ctx.JSON(http.StatusOK, response.OK.WithMsg("用户添加成功"))
		return
	}
	ctx.JSON(http.StatusOK, response.Err.WithMsg("用户添加失败"))
}
func (sysUser *SysUser) DeleteSystemUser(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := sysUser.UserService.DeleteSystemUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.OK.WithMsg("用户删除成功"))
		return
	}
	if result {
		ctx.JSON(http.StatusOK, response.OK.WithMsg("用户删除成功"))
		return
	}
	ctx.JSON(http.StatusOK, response.Err.WithMsg("用户删除失败"))
}

func (sysUser *SysUser) UpdateSystemUser(ctx *gin.Context) {
	id := ctx.Param("id")
	systemUser := model.NewSystemUser()
	err := ctx.ShouldBindJSON(systemUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请检查请求参数").WithErrMsg(err))
		return
	}
	result, err := sysUser.UserService.UpdateSystemUser(id, systemUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.OK.WithMsg("更新用户失败").WithErrMsg(err))
		return
	}
	if result {
		ctx.JSON(http.StatusOK, response.OK.WithMsg("更新用户成功"))
		return
	}
	ctx.JSON(http.StatusOK, response.Err.WithMsg("更新用户失败"))
}
