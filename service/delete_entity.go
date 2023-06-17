package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/entity-service/api/gen/ride/entity/v1alpha1"
)

func (service *EntityServiceServer) DeleteEntity(ctx context.Context,
	req *connect.Request[pb.DeleteEntityRequest]) (*connect.Response[pb.DeleteEntityResponse], error) {

	return connect.NewResponse(&pb.DeleteEntityResponse{}), nil
}
