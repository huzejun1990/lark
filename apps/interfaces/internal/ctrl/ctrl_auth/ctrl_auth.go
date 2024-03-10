// @Author huzejun 2024/3/7 1:33:00
package ctrl_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthCtrl struct {
}

func NewAuthCtrl() *AuthCtrl {
	return &AuthCtrl{}
}

func (ctrl AuthCtrl) SignIn(ctx *gin.Context) {
	fmt.Println("访问了SignIn Api")
	ctx.SecureJSON(200, "调用API成功！")
}
