// @Author huzejun 2024/2/28 21:47:00
package main

import (
	"fmt"
	"lark/pkg/common/xjwt"
)

func main() {
	token, err := xjwt.CreateToken(1, 1, true, 3600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println(token)
	fmt.Println(token.Token)
}
