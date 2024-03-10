// @Author huzejun 2024/3/6 18:22:00
package main

import "lark/apps/interfaces/internal/server"

func main() {
	s := server.NewServer()
	s.Run()
}
