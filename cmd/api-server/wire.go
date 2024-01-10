//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ride-app/user-service/internal/utils/logger"
	savedlocationrepository "github.com/ride-app/user-service/internal/repositories/saved-location"
	userrepository "github.com/ride-app/user-service/internal/repositories/user"
	"github.com/ride-app/user-service/internal/api-handlers"
	thirdparty "github.com/ride-app/user-service/third-party"
)

func InitializeService() (*apihandlers.UserServiceServer, error) {
	panic(
		wire.Build(
			logger.New,
			thirdparty.NewFirebaseApp,
			userrepository.NewFirebaseUserRepository,
			savedlocationrepository.NewFirebaseSavedLocationRepository,
			wire.Bind(
				new(logger.Logger),
				new(*logger.LogrusLogger),
			),
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
