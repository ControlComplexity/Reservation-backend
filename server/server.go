package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
	"net"
	"net/http"
	"reservation/config"
	"reservation/github.com/reservation/api"
)

type apiServiceImpl struct {
	api.UnimplementedReservationServiceServer
	config config.Config
}

type Validator interface {
	Validate() error
}

func InitLogger() {
	// 通过 zap.NewProduction() 创建一个 logger
	logger, _ = zap.NewProduction()
}

var logger *zap.Logger

func RunGRPCServer(config config.Config) error {
	fmt.Print("version 34")
	InitLogger()
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if p, ok := req.(Validator); ok {
			if err := p.Validate(); err != nil {
				// 参数没有通过验证时;返回一个错误不继续向下执行
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		// 通过验证后继续向下执行
		return handler(ctx, req)
	}

	// Init listener
	grpcPort := config.GRPCPort
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.Error("Run(). Failed to listen " + grpcPort + err.Error())
		return err
	}
	defer func() {
		if err := listener.Close(); err != nil {
			logger.Error("Run(). Failed to close " + grpcPort + err.Error())
		}
	}()
	options := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(math.MaxUint32),
		grpc.UnaryInterceptor(
			interceptor,
		),
	}
	resServer, err := NewReservationServer(config)
	if err != nil {
		logger.Error("Run(). Create res server failed:" + err.Error())
		return err
	}
	s := grpc.NewServer(options...)
	api.RegisterReservationServiceServer(s, resServer)
	// Run grpc server
	if err := s.Serve(listener); err != nil {
		log.Fatal("Run(). Failed to serve: " + err.Error())
		return err
	}
	return nil
}

func RunHTTPServer(config *config.Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		"localhost"+config.GRPCPort,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxUint32)),
	)
	if err != nil {
		log.Fatal("start http server failed" + err.Error())
	}
	gateway := runtime.NewServeMux()
	err = api.RegisterReservationServiceHandler(context.Background(), gateway, conn)
	if err != nil {
		log.Fatal("start http server failed, " + err.Error())
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                                                         // 指明哪些请求源被允许访问资源，值可以为 "*"，"null"，或者单个源地址。
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")                              //对于预请求来说，指明了哪些头信息可以用于实际的请求中。
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")                                                                       //对于预请求来说，哪些请求方式可以用于实际的请求。
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type") //对于预请求来说，指明哪些头信息可以安全的暴露给 CORS API 规范的 API
		w.Header().Set("Access-Control-Allow-Credentials", "true")                                                                                 //指明当请求中省略 creadentials 标识时响应是否暴露。对于预请求来说，它表明实际的请求中可以包含用户凭证。
		//放行所有OPTIONS方法
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		gateway.ServeHTTP(w, r)
	}))
	err = http.ListenAndServe(config.HTTPPort, nil)
	if err != nil {
		log.Fatal("RunHttpServer().start HTTP server failed. " + err.Error())
		return err
	}

	return nil
}

// NewReservationServer creates an instance of DataManagerServiceServer
func NewReservationServer(config config.Config) (api.ReservationServiceServer, error) {
	svc := &apiServiceImpl{
		config: config,
	}
	return svc, nil
}
