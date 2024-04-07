// @Author huzejun 2024/3/24 21:53:00
package service

import (
	"context"
	"fmt"
	"lark/pkg/proto/pb_auth"
)

type AuthService interface {
	SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error)
	SignIn(ctx context.Context, req *pb_auth.SignInReq) (resp *pb_auth.SignInResp, err error)
	RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error)
	SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, err error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) SignUp(ctx context.Context, req *pb_auth.SignUpReq) (resp *pb_auth.SignUpResp, err error) {
	resp = new(pb_auth.SignUpResp)
	resp.Msg = "注册成功！"
	fmt.Println(req.Nickname, req.Mobile)
	return
}
func (s *authService) SignIn(ctx context.Context, req *pb_auth.SignInReq) (resp *pb_auth.SignInResp, err error) {
	return
}
func (s *authService) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error) {
	return
}
func (s *authService) SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, err error) {
	return
}
