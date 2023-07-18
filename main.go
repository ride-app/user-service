package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1/riderv1alpha1connect"
	"github.com/ride-app/user-service/config"
	"github.com/ride-app/user-service/di"
	"github.com/ride-app/user-service/interceptors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	log "github.com/sirupsen/logrus"
)

func main() {
	service, err := di.InitializeService()

	if err != nil {
		log.Fatalf("Failed to initialize service: %v", err)
	}

	log.Info("Service Initialized")

	// Create a context that, when cancelled, ends the JWKS background refresh goroutine.
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	authInterceptor, err := interceptors.NewAuthInterceptor(ctx)

	if err != nil {
		log.Fatalf("Failed to initialize auth interceptor: %v", err)
	}

	connectInterceptors := connect.WithInterceptors(authInterceptor)

	path, handler := riderv1alpha1connect.NewUserServiceHandler(service, connectInterceptors)
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	panic(http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", config.Env.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	))
}

func init() {
	log.SetReportCaller(true)

	log.SetFormatter(&log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "severity",
			log.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	})

	log.SetLevel(log.InfoLevel)

	err := cleanenv.ReadEnv(&config.Env)

	if config.Env.Debug {
		log.SetLevel(log.DebugLevel)
	}

	if err != nil {
		log.Warnf("Could not load config: %v", err)
	}
}
