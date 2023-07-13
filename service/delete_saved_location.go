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

	uid := strings.Split(req.Msg.Name, "/")[1]

	if uid != req.Header().Get("uid") {
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	if _, err := service.savedlocationrepository.DeleteSavedLocation(ctx, uid, strings.Split(req.Msg.Name, "/")[3]); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&pb.DeleteSavedLocationResponse{}), nil
}
