// @Author huzejun 2024/3/24 21:52:00
package server

import "lark/apps/auth/internal/server/auth"

type Server struct {
	authServer auth.AuthServer
}

func NewServer(authServer auth.AuthServer) *Server {
	return &Server{authServer: authServer}
}

func (s *Server) Run() {
	s.authServer.Run()
}
