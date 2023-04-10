package system

import (
	"ewa_admin_server/model/system"
	"ewa_admin_server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")

	{
		baseRouter.POST("login", func(context *gin.Context) {
			context.JSON(http.StatusOK, "ok")
		})
		baseRouter.POST("register", func(context *gin.Context) {
			var form system.Register
			if err := context.ShouldBindJSON(&form); err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": utils.GetErrorMsg(form, err),
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}

	return baseRouter
}

/**
这段代码定义了一个 `BaseRouter` 结构体和它的一个 `InitBaseRouter` 方法。

该方法接收一个 gin 的 `RouterGroup` 类型参数，创建一个名为 "base" 的路由组，并为其添加一个 `POST` 请求路由 `/login`。
当该路由被请求时，会返回一个 JSON 格式的字符串 "ok"。

该方法返回值类型为 `gin.IRoutes` 接口，因此实际上返回的是创建的 `baseRouter` 对象，可以在其他地方使用该对象以继续往该路由组中添加更多的路由。
*/
