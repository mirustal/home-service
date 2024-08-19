package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	middleware "gateway/internal/http-server"
	"gateway/internal/swagger"
	"gateway/pkg/config"
	"gateway/pkg/logger"
	pbauth "gateway/pkg/pb/api/auth"
	pbhome "gateway/pkg/pb/api/home"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var requestCounter uint64 

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cfg, err := config.LoadConfig("config", "yaml")
	if err != nil {
		log.Fatal("fail load config: %w", err)
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
		runtime.WithErrorHandler(CustomErrorHandler),
	)

	err = pbauth.RegisterAuthHandlerFromEndpoint(ctx, mux, cfg.GRPC.PortAuth, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		logger.Log.Error("failed to register auth service: %w", err)
	}

	err = pbhome.RegisterHouseServiceHandlerFromEndpoint(ctx, mux, cfg.GRPC.PortHomes, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		logger.Log.Error("failed to register personalinfo service: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		logger.Log.Info("Serving gRPC-Gateway on http://0.0.0.0:8080")

		httpServer := &http.Server{
			Addr:    ":8080",
			Handler: middleware.CorsMiddleware(middleware.LoggingMiddleware(logger.Log, mux)),
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

func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	grpcStatus, ok := status.FromError(err)
	if !ok {
		runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
		return
	}


	requestID := incrementRequestCounter()

	code := grpcStatus.Code()
	errorMessage := grpcStatus.Message()

	httpStatus := runtime.HTTPStatusFromCode(code)
	w.Header().Set("Content-Type", "application/json")
	
	if httpStatus >= 500 && httpStatus < 600 {
		customError := CustomErrorResponse{
			Message:   errorMessage,  // Сообщение об ошибке из gRPC
			RequestID: requestID,     // Идентификатор запроса
			Code:      int(code),     // Код ошибки
		}

		if code == codes.Unavailable {
			w.Header().Set("Retry-After", "5") 
		}


		w.WriteHeader(httpStatus)
		json.NewEncoder(w).Encode(customError)
	} else {
		w.WriteHeader(httpStatus)
		w.Write([]byte(errorMessage))
	}
}

type CustomErrorResponse struct {
	Message   string `json:"message"`
	RequestID int `json:"request_id,omitempty"`
	Code      int    `json:"code"`
}

func incrementRequestCounter() int {
	newID := atomic.AddUint64(&requestCounter, 1)
	return int(newID)
}
