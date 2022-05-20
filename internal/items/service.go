package items

import (
	"context"

	grpc_clients "rodrigues.igor.com/attic/internal/grpc"
	"rodrigues.igor.com/attic/pb"
)

type Service interface {
	Create(item *Item) (*Item, error)
	Update(item *Item) (*Item, error)
	Get(item *Item) (*Item, error)
	List() ([]Item, error)
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

func (s *service) Create(item *Item) (*Item, error) {
	users, err := s.userGRPC.Get(context.Background(), &pb.UserGetRequest{
		ID: int64(item.GetUserID()),
	})
	if err != nil {
		return nil, err
	}

	pe, err := s.physicalEnvironmentGRPC.Get(context.Background(), &pb.PhysicalEnvironmentGetRequest{
		ID: int64(item.PhysicalEnvironmentID),
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

	return s.repository.Insert(item)
}

func (s *service) Update(item *Item) (*Item, error) {
	return s.repository.Update(item)
}

func (s *service) Get(item *Item) (*Item, error) {
	return s.repository.Get(item)
}

func (s *service) List() ([]Item, error) {
	return s.repository.List()
}
