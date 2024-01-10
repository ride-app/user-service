package apihandlers

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *UserServiceServer) UpdateSavedLocation(ctx context.Context,
	req *connect.Request[pb.UpdateSavedLocationRequest]) (*connect.Response[pb.UpdateSavedLocationResponse], error) {
	log := service.logger.WithField("method", "UpdateSavedLocation")

	if err := req.Msg.Validate(); err != nil {
		log.Info("Invalid request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.SavedLocation.Name, "/")[1]
	log.Debug("uid: ", uid)

	if uid != req.Header().Get("uid") {
		log.Info("Permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	user, err := service.savedlocationrepository.GetSavedLocation(ctx, uid, strings.Split(req.Msg.SavedLocation.Name, "/")[3], log)

	if err != nil {
		log.WithError(err).Error("Failed to get saved location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if user == nil {
		log.Info("Location not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("location not found"))
	}

	updateTime, err := service.savedlocationrepository.UpdateSavedLocation(ctx, req.Msg.SavedLocation, log)

	if err != nil {
		log.WithError(err).Error("Failed to update saved location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	req.Msg.SavedLocation.UpdateTime = timestamppb.New(*updateTime)

	res := &pb.UpdateSavedLocationResponse{
		SavedLocation: req.Msg.SavedLocation,
	}

	if err := res.Validate(); err != nil {
		log.WithError(err).Error("Failed to validate response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully updated saved location")
	return connect.NewResponse(res), nil
}
