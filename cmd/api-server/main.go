package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ride-app/go/pkg/logger"
	"github.com/ride-app/user-service/api/ride/rider/v1alpha1/riderv1alpha1connect"
	"github.com/ride-app/user-service/config"
	"github.com/ride-app/user-service/internal/api-handlers/interceptors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	config, err := config.New()

	log := logger.New(!config.Production, config.LogDebug)

	if err != nil {
		log.WithError(err).Fatal("Failed to read environment variables")
	}

	service, err := InitializeService(log, config)

	if err != nil {
		log.WithError(err).Fatal("Failed to initialize service")
	}

	log.Info("Service Initialized")

	// Create a context that, when cancelled, ends the JWKS background refresh goroutine.
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	authInterceptor, err := interceptors.NewAuthInterceptor(ctx, log)

	if err != nil {
		log.WithError(err).Fatal("Failed to initialize auth interceptor")
	}

	connectInterceptors := connect.WithInterceptors(authInterceptor)

	path, handler := riderv1alpha1connect.NewUserServiceHandler(service, connect.WithInterceptors(authInterceptor))
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	panic(http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", config.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	))
}
