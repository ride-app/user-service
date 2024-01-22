package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	pb "github.com/ride-app/user-service/api/ride/rider/v1alpha1"
)

func (service *UserServiceServer) DeleteSavedLocation(ctx context.Context,
	req *connect.Request[pb.DeleteSavedLocationRequest]) (*connect.Response[pb.DeleteSavedLocationResponse], error) {
	log := service.logger.WithField("method", "DeleteSavedLocation")

	if err := req.Msg.Validate(); err != nil {
		log.Info("Invalid request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Name, "/")[1]
	log.Debug("uid: ", uid)

	if uid != req.Header().Get("uid") {
		log.Info("Permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	if _, err := service.savedlocationrepository.DeleteSavedLocation(ctx, uid, strings.Split(req.Msg.Name, "/")[3], log); err != nil {
		log.WithError(err).Error("Failed to delete saved location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteSavedLocationResponse{}), nil
}
