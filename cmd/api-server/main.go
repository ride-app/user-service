package main

import (
	"fmt"
	"net/http"

	"connectrpc.com/authn"
	"connectrpc.com/connect"
	interceptors "github.com/dragonfish/go/v2/pkg/connect/interceptors"
	middlewares "github.com/dragonfish/go/v2/pkg/connect/middlewares"
	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/ride-app/user-service/api/ride/rider/v1alpha1/v1alpha1connect"
	"github.com/ride-app/user-service/config"
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

	panicInterceptor, err := interceptors.NewPanicInterceptor()

	if err != nil {
		log.WithError(err).Fatal("Failed to initialize auth interceptor")
	}

	connectInterceptors := connect.WithInterceptors(panicInterceptor)

	mux := http.NewServeMux()
	mux.Handle(v1alpha1connect.NewUserServiceHandler(service, connectInterceptors))

	middleware := authn.NewMiddleware(middlewares.FirebaseAuth)
	handler := middleware.Wrap(mux)

	panic(http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", config.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(handler, &http2.Server{}),
	))
}
