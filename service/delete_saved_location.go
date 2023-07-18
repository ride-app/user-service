package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
	log "github.com/sirupsen/logrus"
)

func (service *UserServiceServer) DeleteSavedLocation(ctx context.Context,
	req *connect.Request[pb.DeleteSavedLocationRequest]) (*connect.Response[pb.DeleteSavedLocationResponse], error) {
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

	if _, err := service.savedlocationrepository.DeleteSavedLocation(ctx, uid, strings.Split(req.Msg.Name, "/")[3]); err != nil {
		log.Error("Failed to delete saved location: ", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteSavedLocationResponse{}), nil
}
