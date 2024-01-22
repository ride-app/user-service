//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ride-app/go/pkg/logger"
	"github.com/ride-app/user-service/config"
	apihandlers "github.com/ride-app/user-service/internal/api-handlers"
	savedlocationrepository "github.com/ride-app/user-service/internal/repositories/saved-location"
	userrepository "github.com/ride-app/user-service/internal/repositories/user"
	thirdparty "github.com/ride-app/user-service/third-party"
)

func InitializeService(logger logger.Logger, config *config.Config) (*apihandlers.UserServiceServer, error) {
	panic(
		wire.Build(
			thirdparty.NewFirebaseApp,
			userrepository.NewFirebaseUserRepository,
			savedlocationrepository.NewFirebaseSavedLocationRepository,
			wire.Bind(
				new(userrepository.UserRepository),
				new(*userrepository.FirebaseImpl),
			),
			wire.Bind(
				new(savedlocationrepository.SavedLocationRepository),
				new(*savedlocationrepository.FirebaseImpl),
			),
			apihandlers.New,
		),
	)
}
