package grpc

import (
	"context"

	connector "google.golang.org/grpc"
)

type ClientType int

const (
	Users ClientType = iota
	PhysicalEnvironment
	Groups
)

const (
	UserAddr                = "localhost:50051"
	PhysicalEnvironmentAddr = "localhost:50052"
	GroupsAddr              = "localhost:50053"
)

type ClientANDAddr struct {
	Client GRPCClient
	Addr   string
}

func NewClientANDAddr(client GRPCClient, addr string) *ClientANDAddr {
	return &ClientANDAddr{
		Client: client,
		Addr:   addr,
	}
}

type GRPCClient interface {
	Raise(addr string) GRPCClient
	Create(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error)
	Update(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error)
	Get(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error)
	List(ctx context.Context, in interface{}, conectors ...connector.CallOption) (interface{}, error)
}

type ClientGRPC struct {
	clients map[ClientType]*ClientANDAddr
}

func NewClientGRPC() *ClientGRPC {
	return &ClientGRPC{
		clients: map[ClientType]*ClientANDAddr{
			Users:               NewClientANDAddr(NewUserGRPC(), UserAddr),
			PhysicalEnvironment: NewClientANDAddr(NewPhysicalEnvironmentClientGRPC(), PhysicalEnvironmentAddr),
			Groups:              NewClientANDAddr(NewGroupsGRPC(), GroupsAddr),
		},
	}
}

func (cg *ClientGRPC) Create(ct ClientType) GRPCClient {
	clientAndAddr := cg.clients[ct]
	return clientAndAddr.Client.Raise(clientAndAddr.Addr)
}
