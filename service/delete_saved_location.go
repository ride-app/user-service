package service

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/user-service/api/gen/ride/rider/v1alpha1"
)

func (service *UserServiceServer) DeleteSavedLocation(ctx context.Context,
	req *connect.Request[pb.DeleteSavedLocationRequest]) (*connect.Response[pb.DeleteSavedLocationResponse], error) {
	if err := req.Msg.Validate(); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("name cannot be empty"))
	}

	if _, err := service.savedlocationrepository.DeleteSavedLocation(ctx, strings.Split(req.Msg.Name, "/")[1], strings.Split(req.Msg.Name, "/")[3]); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteSavedLocationResponse{}), nil
}
