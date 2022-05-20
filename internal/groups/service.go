package groups

import (
	"context"

	grpc_clients "rodrigues.igor.com/attic/internal/grpc"
	"rodrigues.igor.com/attic/pb"
)

type Service interface {
	Create(group *Group) (*Group, error)
	Update(group *Group) (*Group, error)
	Get(group *Group) (*Group, error)
	List() ([]Group, error)
}

type service struct {
	repository              Repository
	userGRPC                grpc_clients.GRPCClient
	physicalEnvironmentGRPC grpc_clients.GRPCClient
}

func NewService(repository Repository, grpc *grpc_clients.ClientGRPC) Service {
	return &service{
		repository:              repository,
		userGRPC:                grpc.Create(grpc_clients.Users),
		physicalEnvironmentGRPC: grpc.Create(grpc_clients.PhysicalEnvironment),
	}
}

func (s *service) Create(group *Group) (*Group, error) {
	users, err := s.userGRPC.Get(context.Background(), &pb.UserGetRequest{
		ID: int64(group.GetUserID()),
	})
	if err != nil {
		return nil, err
	}

	pe, err := s.physicalEnvironmentGRPC.Get(context.Background(), &pb.PhysicalEnvironmentGetRequest{
		ID: int64(group.PhysicalEnvironmentID),
	})
	if err != nil {
		return nil, ErrUserIDNotExist
	}

	user := users.(*pb.UserResponse)
	physicalEnvironment := pe.(*pb.PhysicalEnvironmentResponse)

	if user.GetID() == 0 {
		return nil, ErrUserIDNotExist
	}

	if physicalEnvironment.GetID() == 0 {
		return nil, ErrPeIDNotExist
	}

	if user.GetID() != physicalEnvironment.GetID() {
		return nil, ErrUserDontMatch
	}

	return s.repository.Insert(group)
}

func (s *service) Update(group *Group) (*Group, error) {
	return s.repository.Update(group)
}

func (s *service) Get(group *Group) (*Group, error) {
	return s.repository.Get(group)
}

func (s *service) List() ([]Group, error) {
	return s.repository.List()
}
