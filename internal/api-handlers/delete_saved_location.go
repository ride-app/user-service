package apihandlers

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/user-service/api/ride/rider/v1alpha1"
)

func (service *UserServiceServer) DeleteSavedLocation(ctx context.Context,
	req *connect.Request[pb.DeleteSavedLocationRequest],
) (*connect.Response[pb.DeleteSavedLocationResponse], error) {
	log := service.logger.WithField("method", "DeleteSavedLocation")

	validator, err := protovalidate.New()
	if err != nil {
		log.WithError(err).Info("Failed to initialize validator")

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if err := validator.Validate(req.Msg); err != nil {
		log.WithError(err).Info("Invalid request")

		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	uid := strings.Split(req.Msg.Name, "/")[1]
	log.Debug("uid: ", uid)

	if uid != req.Header().Get("uid") {
		log.Info("Permission denied")
		return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
	}

	if _, err := service.savedlocationrepository.DeleteSavedLocation(ctx, uid, strings.Split(req.Msg.Name, "/")[3], log); err != nil {
		log.WithError(err).Error("Failed to delete saved location")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&pb.DeleteSavedLocationResponse{})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return res, nil
}
