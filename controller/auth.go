package controller

import (
	"github.com/gin-gonic/gin"
	"management-platform/model/request"
	"management-platform/model/response"
	"management-platform/service"
	"net/http"
)

type Auth struct {
	AuthService service.IAuthService `inject:""`
}

func (auth *Auth) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err.WithMsg("请求参数错误").WithErrMsg(err))
		return
	}
	isSuccess, err := auth.AuthService.Login(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err.WithErrMsg(err))
		return
	}
	if isSuccess {
		ctx.JSON(http.StatusOK, response.OK.WithMsg("登录成功"))
		return
	}
	ctx.JSON(http.StatusOK, response.Err.WithMsg("登录失败"))
}
