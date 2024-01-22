package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/user-service/api/ride/rider/v1alpha1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *UserServiceServer) CreateSavedLocation(ctx context.Context,
	req *connect.Request[pb.CreateSavedLocationRequest]) (*connect.Response[pb.CreateSavedLocationResponse], error) {
	log := service.logger.WithField("method", "CreateSavedLocation")

	validator, err := protovalidate.New()
	if err != nil {
		log.WithError(err).Info("Failed to initialize validator")

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if err := validator.Validate(req.Msg); err != nil {
		log.WithError(err).Info("Invalid request")

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

	res := connect.NewResponse(&pb.CreateSavedLocationResponse{
		SavedLocation: req.Msg.SavedLocation,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully created saved location")
	return res, nil
}
