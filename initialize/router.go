package initialize

import (
	"ewa_admin_server/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			time.Sleep(5 * time.Second)
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	return Router
}
