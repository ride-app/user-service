package apihandlers

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/ride/rider/v1alpha1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *UserServiceServer) UpdateUser(ctx context.Context,
	req *connect.Request[pb.UpdateUserRequest]) (*connect.Response[pb.UpdateUserResponse], error) {
	log := service.logger.WithField("method", "UpdateUser")

	if err := req.Msg.Validate(); err != nil {
		log.Info("Invalid request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.User.Name, "/")[1]
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

	updateTime, err := service.userRepository.UpdateUser(ctx, req.Msg.User, log)

	if err != nil {
		log.WithError(err).Error("Failed to update user")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	req.Msg.User.UpdateTime = timestamppb.New(*updateTime)

	res := &pb.UpdateUserResponse{
		User: req.Msg.User,
	}

	if err := res.Validate(); err != nil {
		log.WithError(err).Error("Failed to validate response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully updated user")
	return connect.NewResponse(res), nil
}
