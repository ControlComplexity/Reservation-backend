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

func (svcImpl *apiServiceImpl) QueryActivityListByDay(ctx context.Context, req *api.QueryActivityListByDayReq) (*api.QueryActivityListByDayResp, error) {
	return dal.QueryActivityListByDay(req)
}

func (svcImpl *apiServiceImpl) QueryOrderList(ctx context.Context, req *api.QueryOrderListReq) (*api.QueryOrderListResp, error) {
	return dal.QueryOrderList(req)
}

func (svcImpl *apiServiceImpl) Register(ctx context.Context, req *api.RegisterReq) (*api.RegisterResp, error) {
	return dal.Register(req)
}

// WXLogin 微信登陆接口
func (svcImpl *apiServiceImpl) WXLogin(ctx context.Context, req *api.WXLoginReq) (*api.WXLoginResp, error) {
	return user.WXLogin(req)
}
