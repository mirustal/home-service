package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gateway/internal/swagger"
	"gateway/pkg/config"
	"gateway/pkg/logger"
	pbhome "gateway/pkg/pb/api/home"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cfg, err := config.LoadConfig("config", "yaml")
	if err != nil {
		log.Fatal("fail load config: %v", err)
	}

	logger := logger.LogInit(cfg.ModeLog)

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(headerName string) (string, bool) {
			switch headerName {
			case "Authorization":
				return "authorization", true
			default:
				return runtime.DefaultHeaderMatcher(headerName)
			}
		}),
	)

	fmt.Println(*cfg, "\n", *cfg.DB, "\n", *cfg.GRPC)
	err = pb.RegisterAuthHandlerFromEndpoint(ctx, mux, cfg.GRPC.PortAuth, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		logger.Log.Error("failed to register auth service: %v", err)
	}

	err = pbhome.RegisterHouseServiceHandlerFromEndpoint(ctx, mux, cfg.GRPC.PortHomes, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		logger.Log.Error("failed to register personalinfo service: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		logger.Log.Info("Serving gRPC-Gateway on http://0.0.0.0:8080")

		httpServer := &http.Server{
			Addr:    ":8080",
			Handler: middleware.CorsMiddleware(loggingMiddleware(logger.Log, mux)),
		}

		httpServer.ListenAndServe()
	}()

	go swagger.InitSwagger(&wg, logger.Log)

	wg.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	logger.Log.Info("recieved signal", <-c)

	logger.Log.Info("auth serviec stop")

}

func loggingMiddleware(log *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Info("request",
			"host", r.Host,
			"method", r.Method,
			"url", r.URL,
			"remote_addr", r.RemoteAddr,
			"user_agent", r.UserAgent(),
			"duration", time.Since(start),
		)
	})
}
