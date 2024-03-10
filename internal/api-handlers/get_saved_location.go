package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/user-service/api/ride/rider/v1alpha1"
)

func (service *UserServiceServer) GetSavedLocation(ctx context.Context,
	req *connect.Request[pb.GetSavedLocationRequest],
) (*connect.Response[pb.GetSavedLocationResponse], error) {
	log := service.logger.WithField("method", "GetSavedLocation")

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

	location, err := service.savedlocationrepository.GetSavedLocation(
		ctx,
		uid,
		strings.Split(req.Msg.Name, "/")[3],
		log,
	)
	if err != nil {
		log.WithError(err).Error("Failed to get saved location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if location == nil {
		log.Info("Saved location not found")
		return nil, connect.NewError(connect.CodeNotFound, errors.New("location not found"))
	}

	res := connect.NewResponse(&pb.GetSavedLocationResponse{
		SavedLocation: location,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully retrieved saved location")
	return res, nil
}
