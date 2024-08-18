package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	homegrpc "home-service/internal/app/grpc/home"
	authclient "home-service/internal/client/auth-service"
	"home-service/internal/modules/home"
	"home-service/pkg/config"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	cfg        *config.GRPCConfig
}

func New(log *slog.Logger, authClient authclient.AuthClient, home *home.Home, cfg *config.GRPCConfig) *App {
	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryLoggingInterceptor(log)),
		grpc.StreamInterceptor(streamLoggingInterceptor(log)),
	)

	homegrpc.Register(gRPCServer, home, authClient)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		cfg:        cfg,
	}
}

func (a *App) Run() error {
	const op = "app.grpc.run"
	log := a.log.With(
		slog.String("where", op),
		slog.Int("port", a.cfg.Port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "app.grpc.stop"

	a.log.With(slog.String("op", op)).Info("stopped grpc server")
	a.gRPCServer.GracefulStop()
}

func unaryLoggingInterceptor(log *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		duration := time.Since(start)

		st, _ := status.FromError(err)

		log.Info("gRPC request",
			"method", info.FullMethod,
			"duration", duration,
			"error", st.Err(),
		)

		return resp, err
	}
}

func streamLoggingInterceptor(log *slog.Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		start := time.Now()
		err := handler(srv, ss)
		duration := time.Since(start)

		st, _ := status.FromError(err)

		log.Info("gRPC stream request",
			"method", info.FullMethod,
			"duration", duration,
			"error", st.Err(),
		)

		return err
	}
}
