package physical_environment

import (
	"context"
	"fmt"

	grpc_clients "rodrigues.igor.com/attic/internal/grpc"
	"rodrigues.igor.com/attic/pb"
)

type Service interface {
	Create(pe *PhysicalEnvironment) (*PhysicalEnvironment, error)
	Update(pe *PhysicalEnvironment) (*PhysicalEnvironment, error)
	Get(pe *PhysicalEnvironment) (*PhysicalEnvironment, error)
	List() ([]PhysicalEnvironment, error)
}

type service struct {
	repository Repository
	userGRPC   grpc_clients.GRPCClient
}

func NewService(repository Repository, grpc *grpc_clients.ClientGRPC) Service {
	return &service{
		repository: repository,
		userGRPC:   grpc.Create(grpc_clients.Users),
	}
}

func (s *service) Create(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	users, err := s.userGRPC.Get(context.Background(), &pb.UserGetRequest{
		ID: int64(pe.GetUserID()),
	})

	if err != nil {
		return nil, err
	}

	if users.(*pb.UserResponse).GetID() == 0 {
		return nil, ErrUserIDNotExist
	}

	fmt.Println(users)
	return s.repository.Insert(pe)
}

func (s *service) Update(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	return s.repository.Update(pe)
}

func (s *service) Get(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	return s.repository.Get(pe)
}

func (s *service) List() ([]PhysicalEnvironment, error) {
	return s.repository.List()
}
