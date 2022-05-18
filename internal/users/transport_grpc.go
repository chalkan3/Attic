package users

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"rodrigues.igor.com/attic/pb"
)

type gRPCServer struct {
	create gt.Handler
	update gt.Handler
	get    gt.Handler
	list   gt.Handler

	pb.UnimplementedUsersServer
}

func NewGRPCServer(svc Service, logger log.Logger) pb.UsersServer {
	return &gRPCServer{
		create: gt.NewServer(
			CreateEndpoint(svc),
			decodeCreateGRPCRequest,
			encodeCreateGRPCResponse,
		),

		update: gt.NewServer(
			UpdateEndpoint(svc),
			decodeUpdateGRPCRequest,
			encodeUpdateGRPCResponse,
		),

		get: gt.NewServer(
			GetEndpoint(svc),
			decodGetGRPCRequest,
			encodGetGRPCResponse,
		),

		list: gt.NewServer(
			ListEndpoint(svc),
			decodListGRPCRequest,
			encodListGRPCResponse,
		),
	}

}

func (s *gRPCServer) Create(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

func (s *gRPCServer) Update(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	_, resp, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

func (s *gRPCServer) Get(ctx context.Context, req *pb.UserGetRequest) (*pb.UserResponse, error) {
	_, resp, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponse), nil
}

func (s *gRPCServer) List(ctx context.Context, req *pb.UserRequestEmpty) (*pb.UserResponseList, error) {
	_, resp, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.UserResponseList), nil
}

// CREATE
func decodeCreateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserRequest)
	return UsersCreateRequest{
		User: NewUser().SetName(req.Name).SetRole(req.Role).SetActive(req.Active),
	}, nil

}

func encodeCreateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(UsersCreateResponse)
	return &pb.UserResponse{
		ID:     resp.User.ID,
		Name:   resp.User.Name,
		Active: resp.User.Active,
		Role:   resp.User.Role,
	}, nil
}

// UPDATE
func decodeUpdateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserRequest)
	return UsersUpdateRequest{
		User: NewUser().SetName(req.Name).SetRole(req.Role).SetActive(req.Active),
	}, nil

}
func encodeUpdateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(UsersUpdateResponse)
	return &pb.UserResponse{
		ID:     resp.User.ID,
		Name:   resp.User.Name,
		Active: resp.User.Active,
		Role:   resp.User.Role,
	}, nil
}

// GET

func decodGetGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserGetRequest)
	return UsersGetRequest{
		User: NewUser().SetID(req.GetID()),
	}, nil

}
func encodGetGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(UsersGetResponse)
	return &pb.UserResponse{
		ID:     resp.User.ID,
		Name:   resp.User.Name,
		Active: resp.User.Active,
		Role:   resp.User.Role,
	}, nil
}

// List

func decodListGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	return UsersListRequest{}, nil

}
func encodListGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(UsersListResponse)
	var listResponse []*pb.UserResponse

	for _, user := range resp.Users {
		listResponse = append(listResponse, &pb.UserResponse{
			ID:     user.ID,
			Name:   user.Name,
			Active: user.Active,
			Role:   user.Role,
		})
	}

	return &pb.UserResponseList{
		User: listResponse,
	}, nil
}
