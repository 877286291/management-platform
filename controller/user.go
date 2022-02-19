package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"management-platform/model"
	"management-platform/model/response"
	"management-platform/service"
	"net/http"
)

func SystemUserRouterRegister(routerGroup *gin.RouterGroup) {
	systemUserRouter := routerGroup.Group("/users")
	{
		systemUserRouter.GET("/", getSystemUserListHandler)
		systemUserRouter.GET("/:id", getSystemUserInfoByIdHandler)
		systemUserRouter.POST("/add", addSystemUserHandler)
	}
}

// 通过ID获取系统用户信息
func getSystemUserInfoByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	zap.L().Info(fmt.Sprintf("查询系统用户ID:%s", id))
	systemUser, err := service.GetSystemUserInfoById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithMsg("获取用户信息失败").WithErrMsg(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithMsg("获取用户信息成功").WithData(systemUser))
}

// 获取系统用户列表
func getSystemUserListHandler(ctx *gin.Context) {
	wrapper := model.NewSystemUser()
	if err := ctx.ShouldBindQuery(wrapper); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请求参数错误").WithErrMsg(err))
		return
	}
	systemUserList, err := service.GetUserList(wrapper)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithMsg("获取用户列表失败").WithErrMsg(err))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithMsg("获取用户列表成功").WithData(systemUserList))
}

// 添加用户
func addSystemUserHandler(ctx *gin.Context) {
	systemUser := model.NewSystemUser()
	if err := ctx.ShouldBindJSON(systemUser); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请求参数错误").WithErrMsg(err))
		return
	}
	result, err := service.AddSystemUser(systemUser)
	if err != nil {
		ctx.JSON(http.StatusOK, response.OK.WithMsg("用户添加失败").WithErrMsg(err))
		return
	}
	if result {
		ctx.JSON(http.StatusOK, response.OK.WithMsg("用户添加成功"))
		return
	}
	ctx.JSON(http.StatusOK, response.OK.WithMsg("用户添加失败"))
}
