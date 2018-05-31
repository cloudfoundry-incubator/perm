package api

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/perm/pkg/api/logging"
	"code.cloudfoundry.org/perm/pkg/api/protos"
	"code.cloudfoundry.org/perm/pkg/api/repos"
	"code.cloudfoundry.org/perm/pkg/api/rpc"
	"code.cloudfoundry.org/perm/pkg/permauth"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

type Server struct {
	logger         lager.Logger
	securityLogger rpc.SecurityLogger
	server         *grpc.Server
}

type Store interface {
	repos.PermissionRepo
	repos.RoleRepo
}

func NewServer(store Store, opts ...ServerOption) *Server {
	config := &serverConfig{
		logger:         &emptyLogger{},
		securityLogger: &emptySecurityLogger{},
	}

	for _, opt := range opts {
		opt(config)
	}

	logger := config.logger

	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			grpcErr := status.Errorf(codes.Internal, "%s", p)
			logger.Error(internal, grpcErr)
			return grpcErr
		}),
	}
	unaryServerInterceptors := []grpc.UnaryServerInterceptor{
		grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
	}

	if config.oidcProvider != nil {
		unaryServerInterceptors = append(unaryServerInterceptors, permauth.ServerInterceptor(config.oidcProvider))
	}

	unaryMiddleware := grpc_middleware.ChainUnaryServer(unaryServerInterceptors...)
	streamMiddleware := grpc_middleware.ChainStreamServer(grpc_recovery.StreamServerInterceptor(recoveryOpts...))

	unaryInterceptor := grpc.UnaryInterceptor(unaryMiddleware)
	streamInterceptor := grpc.StreamInterceptor(streamMiddleware)

	serverOpts := []grpc.ServerOption{
		grpc.KeepaliveParams(config.keepalive),
		unaryInterceptor,
		streamInterceptor,
	}

	if config.credentials != nil {
		serverOpts = append(serverOpts, grpc.Creds(config.credentials))
	}

	server := grpc.NewServer(serverOpts...)

	securityLogger := config.securityLogger

	roleServiceServer := rpc.NewRoleServiceServer(logger, securityLogger, store)
	protos.RegisterRoleServiceServer(server, roleServiceServer)

	permissionServiceServer := rpc.NewPermissionServiceServer(logger, securityLogger, store)
	protos.RegisterPermissionServiceServer(server, permissionServiceServer)

	return &Server{
		logger:         logger,
		securityLogger: config.securityLogger,
		server:         server,
	}
}

func (s *Server) Serve(listener net.Listener) error {
	err := s.server.Serve(listener)

	switch err {
	case nil:
		return nil
	case grpc.ErrServerStopped:
		return ErrServerStopped
	default:
		return ErrServerFailedToStart
	}
}

func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}

func (s *Server) Stop() {
	s.server.Stop()
}

type ServerOption func(*serverConfig)

func WithLogger(logger lager.Logger) ServerOption {
	return func(o *serverConfig) {
		o.logger = logger
	}
}

func WithSecurityLogger(logger rpc.SecurityLogger) ServerOption {
	return func(o *serverConfig) {
		o.securityLogger = logger
	}
}

func WithTLSConfig(config *tls.Config) ServerOption {
	return func(o *serverConfig) {
		o.credentials = credentials.NewTLS(config)
	}
}

func WithMaxConnectionIdle(duration time.Duration) ServerOption {
	return func(o *serverConfig) {
		o.keepalive.MaxConnectionIdle = duration
	}
}

func WithOIDCProvider(provider permauth.OIDCProvider) ServerOption {
	return func(o *serverConfig) {
		o.oidcProvider = provider
	}
}

type serverConfig struct {
	logger         lager.Logger
	securityLogger rpc.SecurityLogger

	credentials credentials.TransportCredentials
	keepalive   keepalive.ServerParameters

	oidcProvider permauth.OIDCProvider
}

type emptyLogger struct{}

func (l *emptyLogger) RegisterSink(lager.Sink) {}

func (l *emptyLogger) SessionName() string {
	return ""
}

func (l *emptyLogger) Session(string, ...lager.Data) lager.Logger {
	return l
}

func (l *emptyLogger) WithData(lager.Data) lager.Logger {
	return l
}

func (l *emptyLogger) Debug(string, ...lager.Data) {}

func (l *emptyLogger) Info(string, ...lager.Data) {}

func (l *emptyLogger) Error(string, error, ...lager.Data) {}

func (l *emptyLogger) Fatal(string, error, ...lager.Data) {}

type emptySecurityLogger struct{}

func (l *emptySecurityLogger) Log(context.Context, string, string, ...logging.CustomExtension) {}
