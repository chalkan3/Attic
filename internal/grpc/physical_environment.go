package grpc

import (
	"context"

	connector "google.golang.org/grpc"
	"rodrigues.igor.com/attic/pb"
)

type PhysicalEnvironmentClient struct {
	client pb.PhysicalEnvironmentClient
}

func NewPhysicalEnvironmentClientGRPC() *PhysicalEnvironmentClient {
	return &PhysicalEnvironmentClient{}
}

func (uc *PhysicalEnvironmentClient) Raise(addr string) GRPCClient {
	cc, err := NewClient(addr).Dial()
	if err != nil {
		panic(err)
	}
	uc.client = pb.NewPhysicalEnvironmentClient(cc)
	return uc
}

func (uc *PhysicalEnvironmentClient) Create(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.PhysicalEnvironmentRequest)
	return uc.client.Create(ctx, bind)
}
func (uc *PhysicalEnvironmentClient) Update(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.PhysicalEnvironmentRequest)
	return uc.client.Update(ctx, bind)
}
func (uc *PhysicalEnvironmentClient) Get(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.PhysicalEnvironmentGetRequest)
	return uc.client.Get(ctx, bind)
}
func (uc *PhysicalEnvironmentClient) List(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.PhysicalEnvironmentRequestEmpty)
	return uc.client.List(ctx, bind)
}
