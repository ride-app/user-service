package apihandlers

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *UserServiceServer) CreateSavedLocation(ctx context.Context,
	req *connect.Request[pb.CreateSavedLocationRequest]) (*connect.Response[pb.CreateSavedLocationResponse], error) {
	log := service.logger.WithField("method", "CreateSavedLocation")

	if err := req.Msg.Validate(); err != nil {
		log.Info("Invalid request")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.SavedLocation.Name, "/")[1]
	log.Debug("uid: ", uid)
	log.Debug("Request header uid: ", req.Header().Get("uid"))

	if uid != req.Header().Get("uid") {
		log.Info("Permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	createTime, err := service.savedlocationrepository.CreateSavedLocation(ctx, req.Msg.SavedLocation, log)

	if err != nil {
		log.WithError(err).Error("Failed to create saved location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	req.Msg.SavedLocation.CreateTime = timestamppb.New(*createTime)
	req.Msg.SavedLocation.UpdateTime = timestamppb.New(*createTime)

	res := &pb.CreateSavedLocationResponse{
		SavedLocation: req.Msg.SavedLocation,
	}

	if err := res.Validate(); err != nil {
		log.WithError(err).Error("Failed to validate response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully created saved location")
	return connect.NewResponse(res), nil
}
