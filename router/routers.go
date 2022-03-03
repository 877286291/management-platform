package router

import (
	"github.com/facebookgo/inject"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"management-platform/controller"
	"management-platform/database"
	"management-platform/repository"
	"management-platform/router/middleware"
	"management-platform/service"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(middleware.Cors())
	Inject(router)
	return router
}
func Inject(router *gin.Engine) {
	var conn database.Mysql
	var auth controller.Auth
	var user controller.SysUser
	var role controller.SysRole
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &conn},
		&inject.Object{Value: &repository.BaseRepo{}},
		&inject.Object{Value: &auth},
		&inject.Object{Value: &service.AuthService{}},
		&inject.Object{Value: &user},
		&inject.Object{Value: &service.SysUserService{}},
		&inject.Object{Value: &repository.SysUserRepo{}},
		&inject.Object{Value: &role},
		&inject.Object{Value: &service.SysRoleService{}},
		&inject.Object{Value: &repository.SysRoleRepo{}},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}
	if err := conn.Connect(); err != nil {
		log.Fatal("database fatal:", err)
	}
	v1 := router.Group("/v1")
	authRouter := v1.Group("/auth")
	{
		authRouter.POST("/login", auth.Login)
	}
	systemUserRouter := v1.Group("/users")
	{
		systemUserRouter.GET("/", user.GetSystemUserList)
		systemUserRouter.GET("/:id", user.GetSystemUserInfoById)
		systemUserRouter.POST("/", user.AddSystemUser)
		systemUserRouter.DELETE("/:id", user.DeleteSystemUser)
		systemUserRouter.PUT("/:id", user.UpdateSystemUser)
	}
	systemRoleRouter := v1.Group("/roles")
	{
		systemRoleRouter.GET("/", role.GetRoleList)
		systemRoleRouter.POST("/", role.AddSystemRole)
		systemRoleRouter.DELETE("/:id", role.DeleteSystemRole)
		systemRoleRouter.PUT("/:id", role.UpdateSystemRole)
	}
}
