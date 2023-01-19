package server

import (
	"context"
	"reservation/dal"
	"reservation/github.com/reservation/api"
	"reservation/user"
)

func (svcImpl *apiServiceImpl) GetSystemInfo(ctx context.Context, req *api.GetSystemInfoReq) (*api.GetSystemInfoResp, error) {
	return dal.GetSystemInfo()
}

func (svcImpl *apiServiceImpl) QueryActivityList(ctx context.Context, req *api.QueryActivityListReq) (*api.QueryActivityListResp, error) {
	return dal.QueryActivityList()
}

func (svcImpl *apiServiceImpl) QueryActivityInfo(ctx context.Context, req *api.QueryActivityInfoReq) (*api.QueryActivityInfoResp, error) {
	return dal.QueryActivityInfo(req)
}

func (svcImpl *apiServiceImpl) JoinActivity(ctx context.Context, req *api.JoinActivityReq) (*api.JoinActivityResp, error) {
	return dal.JoinActivity(req)
}

func (svcImpl *apiServiceImpl) QueryOrderList(ctx context.Context, req *api.QueryOrderListReq) (*api.QueryOrderListResp, error) {
	return dal.QueryOrderList(req)
}

func (svcImpl *apiServiceImpl) EditUser(ctx context.Context, req *api.EditUserReq) (*api.EditUserResp, error) {
	return user.EditUser(req)
}

func (svcImpl *apiServiceImpl) QueryUserInfo(ctx context.Context, req *api.QueryUserInfoReq) (*api.QueryUserInfoResp, error) {
	return user.QueryUserInfo(req)
}

func (svcImpl *apiServiceImpl) Register(ctx context.Context, req *api.RegisterReq) (*api.RegisterResp, error) {
	return user.Register(req)
}

// WXLogin 微信登陆接口
func (svcImpl *apiServiceImpl) WXLogin(ctx context.Context, req *api.WXLoginReq) (*api.WXLoginResp, error) {
	return user.WXLogin(req)
}
