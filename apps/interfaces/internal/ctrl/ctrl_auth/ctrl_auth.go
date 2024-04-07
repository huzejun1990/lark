// @Author huzejun 2024/3/7 1:33:00
package ctrl_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lark/apps/interfaces/internal/dto/dto_auth"
	"lark/apps/interfaces/internal/service/svc_auth"
)

type AuthCtrl struct {
	authService svc_auth.AuthService
}

func NewAuthCtrl(authService svc_auth.AuthService) *AuthCtrl {
	return &AuthCtrl{authService: authService}
}

func (ctrl AuthCtrl) SignIn(ctx *gin.Context) {
	fmt.Println("访问了SignIn Api")
	ctx.SecureJSON(200, "调用API成功！")
}

func (ctrl AuthCtrl) SignUp(ctx *gin.Context) {
	var (
		req dto_auth.SignUpReq
		err error
	)
	if err = ctx.ShouldBind(&req); err != nil {
		ctx.SecureJSON(602, "调参数错误！")
		return
	}

	fmt.Println(req.Password)
	ctx.SecureJSON(200, "调用api成功")

}
