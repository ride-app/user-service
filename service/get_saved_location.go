package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
)

func (service *UserServiceServer) GetSavedLocation(ctx context.Context,
	req *connect.Request[pb.GetSavedLocationRequest]) (*connect.Response[pb.GetSavedLocationResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Name, "/")[1]

	if uid != req.Header().Get("uid") {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	location, err := service.savedlocationrepository.GetSavedLocation(ctx, uid, strings.Split(req.Msg.Name, "/")[3])

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if location == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("location not found"))
	}

	res := &pb.GetSavedLocationResponse{
		SavedLocation: location,
	}

	if err := res.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(res), nil
}
