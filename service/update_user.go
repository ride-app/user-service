package service

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *UserServiceServer) UpdateUser(ctx context.Context,
	req *connect.Request[pb.UpdateUserRequest]) (*connect.Response[pb.UpdateUserResponse], error) {
	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	user, err := service.userRepository.GetUser(ctx, req.Msg.User.Name)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if user == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
	}

	updateTime, err := service.userRepository.UpdateUser(ctx, req.Msg.User)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	req.Msg.User.UpdateTime = timestamppb.New(*updateTime)

	res := &pb.UpdateUserResponse{
		User: req.Msg.User,
	}

	if err := res.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(res), nil
}
