package service

import (
	slr "github.com/ride-app/user-service/repositories/saved-location"
	er "github.com/ride-app/user-service/repositories/user"
)

type UserServiceServer struct {
	userRepository          er.UserRepository
	savedlocationrepository slr.SavedLocationRepository
}

func New(
	userRepository er.UserRepository,
	savedlocationrepository slr.SavedLocationRepository,
) *UserServiceServer {
	return &UserServiceServer{
		userRepository:          userRepository,
		savedlocationrepository: savedlocationrepository,
	}
}
