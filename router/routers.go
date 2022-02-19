package router

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"management-platform/config"
	"management-platform/controller"
	"management-platform/router/middleware"
	"time"
)

func InitRouter() error {
	router := gin.New()
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(middleware.Cors())
	v1 := router.Group("/v1")
	{
		controller.AuthRouterRegister(v1)
		controller.SystemUserRouterRegister(v1)
	}
	err := router.Run(fmt.Sprintf(":%d", config.Config.Server.Port))
	if err != nil {
		return err
	}
	return nil
}
