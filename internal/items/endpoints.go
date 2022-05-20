package items

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	Item *Item `json:"item,omitempty"`
}
type CreateResponse struct {
	Item *Item `json:"item,omitempty"`
}

func CreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(CreateRequest)
		item, err := svc.Create(request.Item)
		if err != nil {
			return CreateResponse{}, err
		}

		return CreateResponse{
			Item: item,
		}, nil
	}
}

type UpdateRequest struct {
	Item *Item `json:"item,omitempty"`
}
type UpdateResponse struct {
	Item *Item `json:"item,omitempty"`
}

func UpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(UpdateRequest)
		Item, err := svc.Update(request.Item)
		if err != nil {
			return UpdateResponse{}, err
		}

		return UpdateResponse{
			Item: Item,
		}, nil
	}
}

type GetRequest struct {
	Item *Item `json:"item,omitempty"`
}
type GetResponse struct {
	Item *Item
}

func GetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(GetRequest)
		Item, err := svc.Get(request.Item)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{
			Item: Item,
		}, nil
	}
}

type ListRequest struct{}
type ListResponse struct {
	Item []Item `json:"item,omitempty"`
}

func ListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		Items, err := svc.List()
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{
			Item: Items,
		}, nil
	}
}
