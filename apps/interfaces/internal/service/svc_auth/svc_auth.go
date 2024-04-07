// @Author huzejun 2024/3/7 1:36:00
package svc_auth

import (
	"lark/apps/interfaces/internal/config"
	"lark/apps/interfaces/internal/dto/dto_auth"
)

type AuthService interface {
}

type authService struct {
	conf *config.Config
}

func NewAuthService(conf *config.Config) AuthService {
	return &authService{conf: conf}
}

func (s *authService) SignUp(req *dto_auth.SignUpReq) {

}
