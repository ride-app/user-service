package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/user/v1alpha1"
)

func (service *UserServiceServer) GetSavedLocation(ctx context.Context,
	req *connect.Request[pb.GetSavedLocationRequest]) (*connect.Response[pb.GetSavedLocationResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	location, err := service.savedlocationrepository.GetSavedLocation(ctx, strings.Split(req.Msg.Name, "/")[1], strings.Split(req.Msg.Name, "/")[3])

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
