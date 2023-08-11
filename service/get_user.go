package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
)

func (service *UserServiceServer) GetUser(ctx context.Context,
	req *connect.Request[pb.GetUserRequest]) (*connect.Response[pb.GetUserResponse], error) {
	log := service.logger.WithField("method", "GetUser")

	if err := req.Msg.Validate(); err != nil {
		log.Info("Invalid request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Name, "/")[1]

	log.Debug("uid: ", uid)
	log.Debug("Request header uid: ", req.Header().Get("uid"))

	if uid != req.Header().Get("uid") {
		log.Info("Permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	user, err := service.userRepository.GetUser(ctx, uid, log)

	if err != nil {
		log.WithError(err).Error("Failed to get user")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if user == nil {
		log.Info("User not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
	}

	res := &pb.GetUserResponse{
		User: user,
	}

	if err := res.Validate(); err != nil {
		log.WithError(err).Error("Failed to validate response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully retrieved user")
	return connect.NewResponse(res), nil
}
