package physical_environment

import (
	"context"
	"fmt"

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
	userGRPC   pb.UsersClient
}

func NewService(repository Repository, userGRPC pb.UsersClient) Service {
	return &service{
		repository: repository,
		userGRPC:   userGRPC,
	}
}

func (s *service) Create(pe *PhysicalEnvironment) (*PhysicalEnvironment, error) {
	users, err := s.userGRPC.Get(context.Background(), &pb.UserGetRequest{
		ID: int64(pe.GetUserID()),
	})

	if err != nil {
		return nil, err
	}

	if users.GetID() == 0 {
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
