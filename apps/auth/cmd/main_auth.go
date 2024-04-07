// @Author huzejun 2024/3/24 21:54:00
package main

import (
	"lark/apps/auth/dig"
	"lark/apps/auth/internal/server"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	dig.Invoke(func(srv server.Server) {
		srv.Run()
	})

	wg.Add(1)
	wg.Wait()
}
