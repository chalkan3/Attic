package users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type UsersCreateRequest struct {
	User *User `json:"user,omitempty"`
}
type UsersCreateResponse struct {
	User *User `json:"user,omitempty"`
}

func CreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(UsersCreateRequest)
		user, err := svc.Create(request.User)
		if err != nil {
			return UsersCreateResponse{}, err
		}

		return UsersCreateResponse{
			User: user,
		}, nil
	}
}

type UsersUpdateRequest struct {
	User *User `json:"user,omitempty"`
}
type UsersUpdateResponse struct {
	User *User `json:"user,omitempty"`
}

func UpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(UsersUpdateRequest)
		user, err := svc.Update(request.User)
		if err != nil {
			return UsersUpdateResponse{}, err
		}

		return UsersUpdateResponse{
			User: user,
		}, nil
	}
}

type UsersGetRequest struct {
	User *User
}
type UsersGetResponse struct {
	User *User
}

func GetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(UsersGetRequest)
		user, err := svc.Get(request.User)
		if err != nil {
			return UsersGetResponse{}, err
		}

		return UsersGetResponse{
			User: user,
		}, nil
	}
}

type UsersListRequest struct{}
type UsersListResponse struct {
	Users []User
}

func ListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		users, err := svc.List()
		if err != nil {
			return UsersListResponse{}, err
		}

		return UsersListResponse{
			Users: users,
		}, nil
	}
}
