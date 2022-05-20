package grpc

import (
	"context"

	connector "google.golang.org/grpc"
	"rodrigues.igor.com/attic/pb"
)

type UserClient struct {
	client pb.UsersClient
}

func NewUserGRPC() *UserClient {
	return &UserClient{}
}

func (uc *UserClient) Raise(addr string) GRPCClient {
	cc, err := NewClient(addr).Dial()
	if err != nil {
		panic(err)
	}
	uc.client = pb.NewUsersClient(cc)
	return uc
}

func (uc *UserClient) Create(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.UserRequest)
	return uc.client.Create(ctx, bind)
}
func (uc *UserClient) Update(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.UserRequest)
	return uc.client.Update(ctx, bind)
}
func (uc *UserClient) Get(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.UserGetRequest)
	return uc.client.Get(ctx, bind)
}
func (uc *UserClient) List(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error) {
	bind := in.(*pb.UserRequestEmpty)
	return uc.client.List(ctx, bind)
}
