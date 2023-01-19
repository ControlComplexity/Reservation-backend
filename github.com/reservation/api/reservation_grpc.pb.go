// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.4.0
// source: proto/reservation.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	// 获取预约系统状态信息.
	GetSystemInfo(ctx context.Context, in *GetSystemInfoReq, opts ...grpc.CallOption) (*GetSystemInfoResp, error)
	//用户注册
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	//微信登录
	WXLogin(ctx context.Context, in *WXLoginReq, opts ...grpc.CallOption) (*WXLoginResp, error)
	//生成MBTI
	GenerateMBTI(ctx context.Context, in *GenerateMBTIReq, opts ...grpc.CallOption) (*GenerateMBTIResp, error)
	//查询某人的MBTI结果
	QueryMBTI(ctx context.Context, in *QueryMBTIReq, opts ...grpc.CallOption) (*QueryMBTIResp, error)
	//查询订单接口
	QueryOrderList(ctx context.Context, in *QueryOrderListReq, opts ...grpc.CallOption) (*QueryOrderListResp, error)
	//查询活动信息
	QueryActivityInfo(ctx context.Context, in *QueryActivityInfoReq, opts ...grpc.CallOption) (*QueryActivityInfoResp, error)
	// 获取全部可参加的活动列表
	QueryActivityList(ctx context.Context, in *QueryActivityListReq, opts ...grpc.CallOption) (*QueryActivityListResp, error)
	// 获取某一天的可参加的活动列表
	QueryActivityListByDay(ctx context.Context, in *QueryActivityListByDayReq, opts ...grpc.CallOption) (*QueryActivityListByDayResp, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) GetSystemInfo(ctx context.Context, in *GetSystemInfoReq, opts ...grpc.CallOption) (*GetSystemInfoResp, error) {
	out := new(GetSystemInfoResp)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetSystemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) WXLogin(ctx context.Context, in *WXLoginReq, opts ...grpc.CallOption) (*WXLoginResp, error) {
	out := new(WXLoginResp)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/WXLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GenerateMBTI(ctx context.Context, in *GenerateMBTIReq, opts ...grpc.CallOption) (*GenerateMBTIResp, error) {
	out := new(GenerateMBTIResp)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GenerateMBTI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) QueryMBTI(ctx context.Context, in *QueryMBTIReq, opts ...grpc.CallOption) (*QueryMBTIResp, error) {
	out := new(QueryMBTIResp)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/QueryMBTI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) QueryOrderList(ctx context.Context, in *QueryOrderListReq, opts ...grpc.CallOption) (*QueryOrderListResp, error) {
	out := new(QueryOrderListResp)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/QueryOrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) QueryActivityInfo(ctx context.Context, in *QueryActivityInfoReq, opts ...grpc.CallOption) (*QueryActivityInfoResp, error) {
	out := new(QueryActivityInfoResp)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/QueryActivityInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) QueryActivityList(ctx context.Context, in *QueryActivityListReq, opts ...grpc.CallOption) (*QueryActivityListResp, error) {
	out := new(QueryActivityListResp)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/QueryActivityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) QueryActivityListByDay(ctx context.Context, in *QueryActivityListByDayReq, opts ...grpc.CallOption) (*QueryActivityListByDayResp, error) {
	out := new(QueryActivityListByDayResp)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/QueryActivityListByDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	// 获取预约系统状态信息.
	GetSystemInfo(context.Context, *GetSystemInfoReq) (*GetSystemInfoResp, error)
	//用户注册
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	//微信登录
	WXLogin(context.Context, *WXLoginReq) (*WXLoginResp, error)
	//生成MBTI
	GenerateMBTI(context.Context, *GenerateMBTIReq) (*GenerateMBTIResp, error)
	//查询某人的MBTI结果
	QueryMBTI(context.Context, *QueryMBTIReq) (*QueryMBTIResp, error)
	//查询订单接口
	QueryOrderList(context.Context, *QueryOrderListReq) (*QueryOrderListResp, error)
	//查询活动信息
	QueryActivityInfo(context.Context, *QueryActivityInfoReq) (*QueryActivityInfoResp, error)
	// 获取全部可参加的活动列表
	QueryActivityList(context.Context, *QueryActivityListReq) (*QueryActivityListResp, error)
	// 获取某一天的可参加的活动列表
	QueryActivityListByDay(context.Context, *QueryActivityListByDayReq) (*QueryActivityListByDayResp, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) GetSystemInfo(context.Context, *GetSystemInfoReq) (*GetSystemInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemInfo not implemented")
}
func (UnimplementedReservationServiceServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedReservationServiceServer) WXLogin(context.Context, *WXLoginReq) (*WXLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WXLogin not implemented")
}
func (UnimplementedReservationServiceServer) GenerateMBTI(context.Context, *GenerateMBTIReq) (*GenerateMBTIResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateMBTI not implemented")
}
func (UnimplementedReservationServiceServer) QueryMBTI(context.Context, *QueryMBTIReq) (*QueryMBTIResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMBTI not implemented")
}
func (UnimplementedReservationServiceServer) QueryOrderList(context.Context, *QueryOrderListReq) (*QueryOrderListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryOrderList not implemented")
}
func (UnimplementedReservationServiceServer) QueryActivityInfo(context.Context, *QueryActivityInfoReq) (*QueryActivityInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryActivityInfo not implemented")
}
func (UnimplementedReservationServiceServer) QueryActivityList(context.Context, *QueryActivityListReq) (*QueryActivityListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryActivityList not implemented")
}
func (UnimplementedReservationServiceServer) QueryActivityListByDay(context.Context, *QueryActivityListByDayReq) (*QueryActivityListByDayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryActivityListByDay not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_GetSystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetSystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetSystemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetSystemInfo(ctx, req.(*GetSystemInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_WXLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WXLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).WXLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/WXLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).WXLogin(ctx, req.(*WXLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GenerateMBTI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateMBTIReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GenerateMBTI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GenerateMBTI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GenerateMBTI(ctx, req.(*GenerateMBTIReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_QueryMBTI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMBTIReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).QueryMBTI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/QueryMBTI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).QueryMBTI(ctx, req.(*QueryMBTIReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_QueryOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOrderListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).QueryOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/QueryOrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).QueryOrderList(ctx, req.(*QueryOrderListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_QueryActivityInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActivityInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).QueryActivityInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/QueryActivityInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).QueryActivityInfo(ctx, req.(*QueryActivityInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_QueryActivityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActivityListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).QueryActivityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/QueryActivityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).QueryActivityList(ctx, req.(*QueryActivityListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_QueryActivityListByDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActivityListByDayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).QueryActivityListByDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/QueryActivityListByDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).QueryActivityListByDay(ctx, req.(*QueryActivityListByDayReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystemInfo",
			Handler:    _ReservationService_GetSystemInfo_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _ReservationService_Register_Handler,
		},
		{
			MethodName: "WXLogin",
			Handler:    _ReservationService_WXLogin_Handler,
		},
		{
			MethodName: "GenerateMBTI",
			Handler:    _ReservationService_GenerateMBTI_Handler,
		},
		{
			MethodName: "QueryMBTI",
			Handler:    _ReservationService_QueryMBTI_Handler,
		},
		{
			MethodName: "QueryOrderList",
			Handler:    _ReservationService_QueryOrderList_Handler,
		},
		{
			MethodName: "QueryActivityInfo",
			Handler:    _ReservationService_QueryActivityInfo_Handler,
		},
		{
			MethodName: "QueryActivityList",
			Handler:    _ReservationService_QueryActivityList_Handler,
		},
		{
			MethodName: "QueryActivityListByDay",
			Handler:    _ReservationService_QueryActivityListByDay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reservation.proto",
}
