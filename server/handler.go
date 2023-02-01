package server

import (
	"context"
	"google.golang.org/grpc/metadata"
	"reservation/constant"
	"reservation/dal"
	"reservation/github.com/reservation/api"
	"reservation/user"
)

// GetSystemInfo 获取系统信息
func (svcImpl *apiServiceImpl) GetSystemInfo(ctx context.Context, req *api.GetSystemInfoReq) (*api.GetSystemInfoResp, error) {
	return dal.GetSystemInfo()
}

// QueryActivityList 查询活动列表
func (svcImpl *apiServiceImpl) QueryActivityList(ctx context.Context, req *api.QueryActivityListReq) (*api.QueryActivityListResp, error) {
	return dal.QueryActivityList()
}

// QueryActivityInfo 查询活动信息
func (svcImpl *apiServiceImpl) QueryActivityInfo(ctx context.Context, req *api.QueryActivityInfoReq) (*api.QueryActivityInfoResp, error) {
	return dal.QueryActivityInfo(req)
}

// JoinActivity 参加活动
func (svcImpl *apiServiceImpl) JoinActivity(ctx context.Context, req *api.JoinActivityReq) (*api.JoinActivityResp, error) {
	return dal.JoinActivity(req)
}

// QueryOrderList 查询订单信息
func (svcImpl *apiServiceImpl) QueryOrderList(ctx context.Context, req *api.QueryOrderListReq) (*api.QueryOrderListResp, error) {
	return dal.QueryOrderList(req)
}

// EditUser 编辑用户
func (svcImpl *apiServiceImpl) EditUser(ctx context.Context, req *api.EditUserReq) (*api.EditUserResp, error) {
	return user.EditUser(req)
}

// QueryUserInfo 查询用户信息
func (svcImpl *apiServiceImpl) QueryUserInfo(ctx context.Context, req *api.QueryUserInfoReq) (*api.QueryUserInfoResp, error) {
	var token string
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		res := headers.Get(constant.TOKEN)
		if res != nil && len(res) > 0 {
			token = res[0]
		} else {
			return &api.QueryUserInfoResp{
				Success:  false,
				ErrorMsg: "token 不存在",
			}, nil
		}
	}
	return user.QueryUserInfo(token)
}

// Register 注册
func (svcImpl *apiServiceImpl) Register(ctx context.Context, req *api.RegisterReq) (*api.RegisterResp, error) {
	return user.Register(req)
}

// WXLogin 微信登陆接口
func (svcImpl *apiServiceImpl) WXLogin(ctx context.Context, req *api.WXLoginReq) (*api.WXLoginResp, error) {
	return user.WXLogin(req)
}

// QuerySwipers 轮播图查询接口
func (svcImpl *apiServiceImpl) QuerySwipers(ctx context.Context, req *api.QuerySwipersReq) (*api.QuerySwipersResp, error) {
	return dal.QuerySwipers(req)
}
