package groups

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	Group *Group `json:"group,omitempty"`
}
type CreateResponse struct {
	Group *Group `json:"group,omitempty"`
}

func CreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(CreateRequest)
		group, err := svc.Create(request.Group)
		if err != nil {
			return CreateResponse{}, err
		}

		return CreateResponse{
			Group: group,
		}, nil
	}
}

type UpdateRequest struct {
	Group *Group `json:"group,omitempty"`
}
type UpdateResponse struct {
	Group *Group `json:"group,omitempty"`
}

func UpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(UpdateRequest)
		Group, err := svc.Update(request.Group)
		if err != nil {
			return UpdateResponse{}, err
		}

		return UpdateResponse{
			Group: Group,
		}, nil
	}
}

type GetRequest struct {
	Group *Group `json:"group,omitempty"`
}
type GetResponse struct {
	Group *Group
}

func GetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(GetRequest)
		Group, err := svc.Get(request.Group)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{
			Group: Group,
		}, nil
	}
}

type ListRequest struct{}
type ListResponse struct {
	Group []Group `json:"group,omitempty"`
}

func ListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		Groups, err := svc.List()
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{
			Group: Groups,
		}, nil
	}
}
