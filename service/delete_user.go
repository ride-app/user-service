package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
	log "github.com/sirupsen/logrus"
)

func (service *UserServiceServer) DeleteUser(ctx context.Context,
	req *connect.Request[pb.DeleteUserRequest]) (*connect.Response[pb.DeleteUserResponse], error) {
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

	if _, err := service.userRepository.DeleteUser(ctx, uid); err != nil {
		log.Error("Failed to delete user: ", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully deleted user")
	return connect.NewResponse(&pb.DeleteUserResponse{}), nil
}
