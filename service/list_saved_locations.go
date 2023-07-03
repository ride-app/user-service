package service

import (
	"context"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
)

func (service *UserServiceServer) ListSavedLocations(ctx context.Context,
	req *connect.Request[pb.ListSavedLocationsRequest]) (*connect.Response[pb.ListSavedLocationsResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	locations, err := service.savedlocationrepository.GetSavedLocations(ctx, strings.Split(req.Msg.Parent, "/")[1])

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := &pb.ListSavedLocationsResponse{
		SavedLocations: locations,
	}

	if err := res.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(res), nil
}
