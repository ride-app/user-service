package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
	log "github.com/sirupsen/logrus"
)

func (service *UserServiceServer) GetSavedLocation(ctx context.Context,
	req *connect.Request[pb.GetSavedLocationRequest]) (*connect.Response[pb.GetSavedLocationResponse], error) {

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

	location, err := service.savedlocationrepository.GetSavedLocation(ctx, uid, strings.Split(req.Msg.Name, "/")[3])

	if err != nil {
		log.Error("Failed to get saved location: ", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if location == nil {
		log.Info("Saved location not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("location not found"))
	}

	res := &pb.GetSavedLocationResponse{
		SavedLocation: location,
	}

	if err := res.Validate(); err != nil {
		log.Error("Failed to validate response: ", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully retrieved saved location")
	return connect.NewResponse(res), nil
}
