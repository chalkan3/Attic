package physical_environment

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	PhysicalEnvironment *PhysicalEnvironment `json:"physical_environment,omitempty"`
}
type CreateResponse struct {
	PhysicalEnvironment *PhysicalEnvironment `json:"physical_environment,omitempty"`
}

func CreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(CreateRequest)
		physicalEnvironment, err := svc.Create(request.PhysicalEnvironment)
		if err != nil {
			return CreateResponse{}, err
		}

		return CreateResponse{
			PhysicalEnvironment: physicalEnvironment,
		}, nil
	}
}

type UpdateRequest struct {
	PhysicalEnvironment *PhysicalEnvironment `json:"physical_environment,omitempty"`
}
type UpdateResponse struct {
	PhysicalEnvironment *PhysicalEnvironment `json:"physical_environment,omitempty"`
}

func UpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(UpdateRequest)
		physicalEnvironment, err := svc.Update(request.PhysicalEnvironment)
		if err != nil {
			return UpdateResponse{}, err
		}

		return UpdateResponse{
			PhysicalEnvironment: physicalEnvironment,
		}, nil
	}
}

type GetRequest struct {
	PhysicalEnvironment *PhysicalEnvironment `json:"physical_environment,omitempty"`
}
type GetResponse struct {
	PhysicalEnvironment *PhysicalEnvironment
}

func GetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(GetRequest)
		physicalEnvironment, err := svc.Get(request.PhysicalEnvironment)
		if err != nil {
			return GetResponse{}, err
		}

		return GetResponse{
			PhysicalEnvironment: physicalEnvironment,
		}, nil
	}
}

type ListRequest struct{}
type ListResponse struct {
	PhysicalEnvironment []PhysicalEnvironment `json:"physical_environment,omitempty"`
}

func ListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		physicalEnvironments, err := svc.List()
		if err != nil {
			return ListResponse{}, err
		}

		return ListResponse{
			PhysicalEnvironment: physicalEnvironments,
		}, nil
	}
}
