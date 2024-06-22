/*
Package middleware is a collection of gRPC middleware:

	auth	controls authorization by info.FullMethod and set context.Value to userID from a user token
	logger	logging all gRPC requests and responses

# How to use:

	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			 // advanced logging with http://github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging
			logging.UnaryServerInterceptor(middleware.InterceptorLogger(l)),
			// custom auth
			interceptor.UnaryAuthInterceptor, // interceptor := middleware.NewAuthInterceptor(jwtManager)
		),
	}
	srv, err := grpcserver.New(cfg.GRPC.Address, serverOptions...)
*/
package middleware
