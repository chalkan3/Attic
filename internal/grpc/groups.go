package grpc

import (
	"context"

	connector "google.golang.org/grpc"
	"rodrigues.igor.com/attic/pb"
)

type GroupsClient struct {
	client pb.GroupClient
}

func NewGroupsGRPC() *GroupsClient {
	return &GroupsClient{}
}

func (uc *GroupsClient) Raise(addr string) GRPCClient {
	cc, err := NewClient(addr).Dial()
	if err != nil {
		panic(err)
	}
	uc.client = pb.NewGroupClient(cc)
	return uc
}

func (uc *GroupsClient) Create(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.GroupCreateRequest)
	return uc.client.Create(ctx, bind)
}
func (uc *GroupsClient) Update(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.GroupUpdateRequest)
	return uc.client.Update(ctx, bind)
}
func (uc *GroupsClient) Get(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.GroupGetRequest)
	return uc.client.Get(ctx, bind)
}
func (uc *GroupsClient) List(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.GroupListRequest)
	return uc.client.List(ctx, bind)
}
