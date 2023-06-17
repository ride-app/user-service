package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	pb "github.com/ride-app/entity-service/api/gen/ride/entity/v1alpha1"
)

func (service *EntityServiceServer) CreateEntity(ctx context.Context,
	req *connect.Request[pb.CreateEntityRequest]) (*connect.Response[pb.CreateEntityResponse], error) {

	return connect.NewResponse(&pb.CreateEntityResponse{}), nil
}
