package grpcserver

import (
	"google.golang.org/grpc"
	"net"
)

// Server contains a new gRPC server.
type Server struct {
	grpc.Server
	server *grpc.Server
	notify chan error
}

// New returns a new gRPC server.
func New(address string) (*Server, error) {
	server := grpc.NewServer()

	listen, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	s := &Server{server: server}
	s.Start(listen)

	return s, nil
}

// Start turns on gRPC server.
func (s *Server) Start(listen net.Listener) {
	go func() {
		s.notify <- s.server.Serve(listen)
		close(s.notify)
	}()
}

// Notify returns an error.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown turns off gRPC server.
func (s *Server) Shutdown() {
	s.server.GracefulStop()
}
