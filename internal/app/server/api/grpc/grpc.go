package grpc

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/middleware"
	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/pb"
	"github.com/webkimru/go-keeper/internal/app/server/config"
	"github.com/webkimru/go-keeper/pkg/grpcserver"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// New returns a new gRPC server.
func New(
	userService UserService,
	keyValueService KeyValueService,
	jwtManager *jwtmanager.JWTManager,
	cfg *config.Config,
	l *logger.Log,
) *grpcserver.Server {
	// user server description
	userServer := NewUserServer(userService)
	// key-value server description
	keyValueServer := NewKeyValueServer(keyValueService)
	interceptor := middleware.NewAuthInterceptor(jwtManager)
	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(middleware.InterceptorLogger(l)), // advanced logging
			interceptor.UnaryAuthInterceptor,                                // custom
		),
	}
	// starting from reusable pkg
	srv, err := grpcserver.New(cfg.GRPC.Address, serverOptions...)
	if err != nil {
		l.Log.Fatal(err)
	}
	// register our services
	pb.RegisterUserServiceServer(srv.Server, userServer)
	pb.RegisterKeyValueServiceServer(srv.Server, keyValueServer)
	// to use some qRPC clients: evans and etc
	reflection.Register(srv.Server)

	srv.Start(srv.Listen)

	return srv
}
