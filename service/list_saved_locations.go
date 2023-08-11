package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
)

func (service *UserServiceServer) ListSavedLocations(ctx context.Context,
	req *connect.Request[pb.ListSavedLocationsRequest]) (*connect.Response[pb.ListSavedLocationsResponse], error) {
	log := service.logger.WithField("method", "ListSavedLocations")

	if err := req.Msg.Validate(); err != nil {
		log.Info("Invalid request")
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

	res := &pb.ListSavedLocationsResponse{
		SavedLocations: locations,
	}

	if err := res.Validate(); err != nil {
		log.WithError(err).Error("Failed to validate response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	log.Info("Successfully listed saved locations")
	return connect.NewResponse(res), nil
}
