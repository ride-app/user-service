package interceptors

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt/v5"
)

func NewAuthInterceptor(ctx context.Context) (*connect.UnaryInterceptorFunc, error) {
	jwksURI := "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

	options := keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(jwksURI, options)

	if err != nil {
		return nil, err
	}

	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Header().Get("authorization") == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}

			if req.Header().Get("authorization")[:7] != "Bearer " {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("invalid token"),
				)
			}
			token := req.Header().Get("authorization")[7:]

			log.Printf("Token: %s", token)

			_, err := jwt.Parse(token, jwks.Keyfunc)

			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New(err.Error()),
				)
			}

			return next(ctx, req)
		})
	}

	interceptorFunc := connect.UnaryInterceptorFunc(interceptor)

	return &interceptorFunc, nil
}
