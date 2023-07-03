package service

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
)

func (service *UserServiceServer) GetUser(ctx context.Context,
	req *connect.Request[pb.GetUserRequest]) (*connect.Response[pb.GetUserResponse], error) {

	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("name cannot be empty"))
	}

	user, err := service.userRepository.GetUser(ctx, req.Msg.Name)

	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if user == nil {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
	}

	res := &pb.GetUserResponse{
		User: user,
	}

	if err := res.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(res), nil
}
