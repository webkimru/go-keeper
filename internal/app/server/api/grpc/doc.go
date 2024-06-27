/*
Package grpc describes the handlers and provides starting gRPC server.

It's a wrapper over the grpc package google.golang.org/grpc.
The grpc package inits like this:

	server := grpc.New(userService, keyValueService, jwtManager, cfg, l)

	And has business-logic arguments as interfaces:

	userService 		works with service.UserService, handles Add, Find methods for login and register user
	keyValueService 	works with service.KeyValueService, handles CRUD for key-value data

	Authentication argument:

	jwtManager			generates a token and get the user ID form it by jwtmanager.JWTManager

	Configuration and logging argument

	cfg					contains all application settings by config.Config
	l 					creates all logs by logger.Log, include all requests and responses

Initial function contains interceptors that works as middleware.
See more details in the grpcserver package: https://github.com/webkimru/go-keeper/pkg/grpcserver
*/
package grpc
