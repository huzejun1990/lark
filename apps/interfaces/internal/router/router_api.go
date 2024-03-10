// @Author huzejun 2024/3/7 1:35:00
package router

import (
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/ctrl/ctrl_auth"
)

func Register(engine *gin.Engine) {
	openGroup := engine.Group("open")
	registerOpenRouter(openGroup)
	apiGroup := engine.Group("api")
	registerApiRouter(apiGroup)
}

func registerOpenRouter(group *gin.RouterGroup) {
	authGroup := group.Group("auth")
	ctrl := ctrl_auth.NewAuthCtrl()
	authGroup.POST("sign_in", ctrl.SignIn)
}

func registerApiRouter(group *gin.RouterGroup) {

}
