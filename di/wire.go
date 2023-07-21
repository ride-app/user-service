//go:build wireinject

package di

import (
	"github.com/google/wire"
	savedlocationrepository "github.com/ride-app/user-service/repositories/saved-location"
	tokenrepository "github.com/ride-app/user-service/repositories/token"
	userrepository "github.com/ride-app/user-service/repositories/user"
	"github.com/ride-app/user-service/service"
	thirdparty "github.com/ride-app/user-service/third-party"
)

func InitializeService() (*service.UserServiceServer, error) {
	panic(
		wire.Build(
			thirdparty.NewFirebaseApp,
			userrepository.NewFirebaseUserRepository,
			savedlocationrepository.NewFirebaseSavedLocationRepository,
			tokenrepository.NewFirebaseTokenRepository,
			wire.Bind(
				new(userrepository.UserRepository),
				new(*userrepository.FirebaseImpl),
			),
			wire.Bind(
				new(savedlocationrepository.SavedLocationRepository),
				new(*savedlocationrepository.FirebaseImpl),
			),
			wire.Bind(
				new(tokenrepository.TokenRepository),
				new(*tokenrepository.FirebaseImpl),
			),
			service.New,
		),
	)
}
