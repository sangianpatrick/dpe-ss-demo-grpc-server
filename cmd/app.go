package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

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

	server := grpc.NewServer()
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
}
