// @Author huzejun 2024/2/28 16:00:00
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lark/pkg/common/xjwt"
)

func main() {
	engine := gin.Default()
	engine.Use(JwtAuth(), Test())
	engine.GET("hello", hello)
	engine.Run(":9166")
	/*	engine := xgin.NewGinServer()
		engine.Engine.GET("hello", hello)
		engine.Run(9166)*/
}

func hello(c *gin.Context) {
	fmt.Println("访问了hello这个api")
	c.SecureJSON(0, "访问成功！")
}

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := xjwt.ParseFromCookie(ctx)
		if err != nil {
			ctx.Abort()
			ctx.SecureJSON(601, "token 验证失败")
			return
		}
		fmt.Println("token 验证成功", token)

	}
}

func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("测试中间件")
	}
}
