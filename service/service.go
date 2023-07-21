package service

import (
	slr "github.com/ride-app/user-service/repositories/saved-location"
	tr "github.com/ride-app/user-service/repositories/token"
	er "github.com/ride-app/user-service/repositories/user"
)

type UserServiceServer struct {
	userRepository          er.UserRepository
	savedlocationrepository slr.SavedLocationRepository
	tokenRepository         tr.TokenRepository
}

func New(
	userRepository er.UserRepository,
	savedlocationrepository slr.SavedLocationRepository,
	tokenRepository tr.TokenRepository,
) *UserServiceServer {
	return &UserServiceServer{
		userRepository:          userRepository,
		savedlocationrepository: savedlocationrepository,
		tokenRepository:         tokenRepository,
	}
}
