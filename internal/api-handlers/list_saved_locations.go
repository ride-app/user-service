package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/user-service/api/ride/rider/v1alpha1"
)

func (service *UserServiceServer) ListSavedLocations(ctx context.Context,
	req *connect.Request[pb.ListSavedLocationsRequest]) (*connect.Response[pb.ListSavedLocationsResponse], error) {
	log := service.logger.WithField("method", "ListSavedLocations")

	validator, err := protovalidate.New()
	if err != nil {
		log.WithError(err).Info("Failed to initialize validator")

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if err := validator.Validate(req.Msg); err != nil {
		log.WithError(err).Info("Invalid request")

		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Parent, "/")[1]
	log.Debug("uid: ", uid)

	if uid != req.Header().Get("uid") {
		log.Info("Permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	locations, err := service.savedlocationrepository.GetSavedLocations(ctx, uid, log)

	if err != nil {
		log.WithError(err).Error("Failed to get saved locations")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&pb.ListSavedLocationsResponse{
		SavedLocations: locations,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully listed saved locations")
	return res, nil
}
