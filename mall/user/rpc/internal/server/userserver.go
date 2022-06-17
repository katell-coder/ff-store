// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"ff-store/mall/user/rpc/internal/logic"
	"ff-store/mall/user/rpc/internal/svc"
	"ff-store/mall/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdRequest) (*user.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}
