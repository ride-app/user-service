package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/user-service/api/ride/rider/v1alpha1"
)

func (service *UserServiceServer) GetUser(ctx context.Context,
	req *connect.Request[pb.GetUserRequest]) (*connect.Response[pb.GetUserResponse], error) {
	log := service.logger.WithField("method", "GetUser")

	validator, err := protovalidate.New()
	if err != nil {
		log.WithError(err).Info("Failed to initialize validator")

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if err := validator.Validate(req.Msg); err != nil {
		log.WithError(err).Info("Invalid request")

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

	res := connect.NewResponse(&pb.GetUserResponse{
		User: user,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully retrieved user")
	return res, nil
}
