package service

import (
	"github.com/ride-app/user-service/logger"
	slr "github.com/ride-app/user-service/repositories/saved-location"
	er "github.com/ride-app/user-service/repositories/user"
)

type UserServiceServer struct {
	userRepository          er.UserRepository
	savedlocationrepository slr.SavedLocationRepository
	logger                  logger.Logger
}

func New(
	userRepository er.UserRepository,
	savedlocationrepository slr.SavedLocationRepository,
	logger logger.Logger,
) *UserServiceServer {
	return &UserServiceServer{
		userRepository:          userRepository,
		savedlocationrepository: savedlocationrepository,
		logger:                  logger,
	}
}
