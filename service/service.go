package service

import (
	er "github.com/ride-app/entity-service/repositories/entity"
)

type EntityServiceServer struct {
	entityRepository er.EntityRepository
}

func New(
	entityRepository er.EntityRepository,
) *EntityServiceServer {
	return &EntityServiceServer{
		entityRepository: entityRepository,
	}
}
