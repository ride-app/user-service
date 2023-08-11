//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ride-app/user-service/logger"
	savedlocationrepository "github.com/ride-app/user-service/repositories/saved-location"
	userrepository "github.com/ride-app/user-service/repositories/user"
	"github.com/ride-app/user-service/service"
	thirdparty "github.com/ride-app/user-service/third-party"
)

func InitializeService() (*service.UserServiceServer, error) {
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
			service.New,
		),
	)
}
