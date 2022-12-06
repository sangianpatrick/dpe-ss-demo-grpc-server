package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/config"
	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/database"
	customerpb "github.com/sangianpatrick/dpe-ss-demo-grpc-server/pb/customer"
	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/repository"
	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.GetConfig()

	logger := logrus.New()
	logger.SetFormatter(cfg.LogFormatter)
	logger.SetReportCaller(true)

	db := database.GetDatabase()

	accountRepository := repository.NewAccountRepository(logger, db)
	customerService := service.NewCustomerService(cfg.Application.Location, logger, accountRepository)

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),
				grpc_recovery.UnaryServerInterceptor(),
			),
		),
		grpc.ChainStreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_logrus.StreamServerInterceptor(logrus.NewEntry(logger)),
				grpc_recovery.StreamServerInterceptor(),
			),
		),
	)
	customerpb.RegisterCustomerServer(server, customerService)
	reflection.Register(server)

	listener, _ := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Application.Port))

	go func() {
		logger.Infof("application run on port %d", cfg.Application.Port)
		server.Serve(listener)
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm

	server.GracefulStop()
	db.Close()

	logger.Info("shutdown application")
}
