package service

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
	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	user, err := service.savedlocationrepository.GetSavedLocation(ctx, strings.Split(req.Msg.SavedLocation.Name, "/")[1], strings.Split(req.Msg.SavedLocation.Name, "/")[3])

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if user == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("location not found"))
	}

	updateTime, err := service.savedlocationrepository.UpdateSavedLocation(ctx, req.Msg.SavedLocation)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	req.Msg.SavedLocation.UpdateTime = timestamppb.New(*updateTime)

	res := &pb.UpdateSavedLocationResponse{
		SavedLocation: req.Msg.SavedLocation,
	}

	if err := res.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(res), nil
}
