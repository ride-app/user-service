package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *UserServiceServer) CreateSavedLocation(ctx context.Context,
	req *connect.Request[pb.CreateSavedLocationRequest]) (*connect.Response[pb.CreateSavedLocationResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	createTime, err := service.savedlocationrepository.CreateSavedLocation(ctx, req.Msg.SavedLocation)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	req.Msg.SavedLocation.CreateTime = timestamppb.New(*createTime)
	req.Msg.SavedLocation.UpdateTime = timestamppb.New(*createTime)

	res := &pb.CreateSavedLocationResponse{
		SavedLocation: req.Msg.SavedLocation,
	}

	if err := res.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(res), nil
}
