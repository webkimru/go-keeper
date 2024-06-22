// Package grpcserver implements an grPC server.
package grpcserver

import (
	"net"

	"google.golang.org/grpc"
)

// Server contains a new gRPC server.
type Server struct {
	Server *grpc.Server
	Listen net.Listener
	notify chan error
}

// New returns a new gRPC server.
func New(address string, option ...grpc.ServerOption) (*Server, error) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	s := &Server{
		Listen: listen,
		Server: grpc.NewServer(option...),
		notify: make(chan error, 1),
	}

	return s, nil
}

// Start turns on gRPC server.
func (s *Server) Start(listen net.Listener) {
	go func() {
		s.notify <- s.Server.Serve(listen)
		close(s.notify)
	}()
}

// Notify returns an error.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown turns off gRPC server.
func (s *Server) Shutdown() {
	s.Server.GracefulStop()
}
