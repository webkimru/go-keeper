package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	apigrpc "github.com/webkimru/go-keeper/internal/app/server/api/grpc"
	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/middleware"
	pb "github.com/webkimru/go-keeper/internal/app/server/api/grpc/proto"
	"github.com/webkimru/go-keeper/internal/app/server/config"
	"github.com/webkimru/go-keeper/internal/app/server/repository/store/inmemory"
	"github.com/webkimru/go-keeper/internal/app/server/service"
	"github.com/webkimru/go-keeper/pkg/grpcserver"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// Run runs application.
func Run(cfg *config.Config) {
	l, err := logger.NewZap(cfg.Log.Level)
	if err != nil {
		log.Fatal(err)
	}

	l.Log.Infoln("Starting configuration:",
		"APP_SECRET_KEY", cfg.SecretKey,
		"APP_TOKEN_EXP", cfg.TokenExp,
		"LOG_LEVEL", cfg.Log.Level,
		"GRPC_ADDRESS", cfg.GRPC.Address,
	)

	// App DI init
	db := inmemory.NewStorage()
	userService := service.NewUserService(db)
	jwtManager := jwtmanager.New(cfg.SecretKey, cfg.TokenExp)
	userServer := apigrpc.NewUserServer(userService, jwtManager)

	// gRPC server
	l.Log.Infof("Starting gRPC server on %s", cfg.GRPC.Address)
	interceptor := middleware.NewAuthInterceptor(jwtManager)
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.UnaryAuthInterceptor),
	}
	grpcServer, err := grpcserver.New(cfg.GRPC.Address, serverOptions...)
	pb.RegisterUserServiceServer(grpcServer.Reg(), userServer)
	reflection.Register(grpcServer.Reg())

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case s := <-interrupt:
		l.Log.Info("Got signal: " + s.String())

	case err = <-grpcServer.Notify():
		l.Log.Errorf("grpcServer.Notify: %v", err)
	}

	// Shutdown
	grpcServer.Shutdown()
}
