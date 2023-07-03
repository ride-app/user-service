package service

import (
	"context"
	"errors"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
)

func (service *UserServiceServer) DeleteUser(ctx context.Context,
	req *connect.Request[pb.DeleteUserRequest]) (*connect.Response[pb.DeleteUserResponse], error) {
	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("name cannot be empty"))
	}

	if _, err := service.userRepository.DeleteUser(ctx, req.Msg.Name); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteUserResponse{}), nil
}
