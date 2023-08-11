package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1/riderv1alpha1connect"
	"github.com/ride-app/user-service/config"
	"github.com/ride-app/user-service/di"
	"github.com/ride-app/user-service/interceptors"
	"github.com/ride-app/user-service/logger"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	err := cleanenv.ReadEnv(&config.Env)

	log := logger.New()

	if err != nil {
		log.WithError(err).Fatal("Failed to read environment variables")
	}

	service, err := di.InitializeService()

	if err != nil {
		log.WithError(err).Fatal("Failed to initialize service")
	}

	log.Info("Service Initialized")

	// Create a context that, when cancelled, ends the JWKS background refresh goroutine.
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	authInterceptor, err := interceptors.NewAuthInterceptor(ctx)

	if err != nil {
		log.WithError(err).Fatal("Failed to initialize auth interceptor")
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
