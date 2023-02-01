package main

import (
	"flag"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"reservation/config"
	"reservation/message"
	"reservation/server"
	"syscall"
)

var (
	configs    = &config.Config{}
	logger     *zap.Logger
	configPath = flag.String("config", "config/reservation.yml", "reservation-service config file")
)

func InitLogger() {
	// 通过 zap.NewProduction() 创建一个 logger
	logger, _ = zap.NewProduction()
}

func main() {
	InitLogger()
	logger.Info("=== reservation-service ===")
	err := config.ReadConfig(*configPath, configs)
	ch := make(chan string, 0)

	if err != nil {
		panic(err)
	}
	logger.Info("GRPC server starting at " + configs.GRPCPort)
	go func() {
		if err := server.RunGRPCServer(*configs, ch); err != nil {
			log.Fatal("GRPC server start failed: " + err.Error())
		}
	}()
	logger.Info("HTTP server starting at " + configs.HTTPPort)

	go func() {
		if err := server.RunHTTPServer(configs); err != nil {
			log.Fatal("HTTP server start failed: " + err.Error())
		}
	}()
	//发送WebSocket消息
	message.Send(ch)
	// Wait for quit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	sig := <-c
	logger.Info("Received signal " + sig.String())
	os.Exit(0)
}
