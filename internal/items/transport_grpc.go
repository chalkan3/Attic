package groups

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rodrigues.igor.com/attic/pb"
)

type gRPCServer struct {
	create gt.Handler
	update gt.Handler
	get    gt.Handler
	list   gt.Handler

	pb.UnimplementedGroupServer
}

func NewGRPCServer(svc Service, logger log.Logger) pb.GroupServer {
	opt := optionsGRPCServer(logger)
	return &gRPCServer{
		create: gt.NewServer(
			CreateEndpoint(svc),
			decodeCreateGRPCRequest,
			encodeCreateGRPCResponse,
			opt...,
		),

		update: gt.NewServer(
			UpdateEndpoint(svc),
			decodeUpdateGRPCRequest,
			encodeUpdateGRPCResponse,
			opt...,
		),

		get: gt.NewServer(
			GetEndpoint(svc),
			decodeGetGRPCRequest,
			encodeGetGRPCResponse,
			opt...,
		),

		list: gt.NewServer(
			ListEndpoint(svc),
			decodListGRPCRequest,
			encodListGRPCResponse,
			opt...,
		),
	}

}

func (s *gRPCServer) Create(ctx context.Context, req *pb.GroupCreateRequest) (*pb.GroupResponse, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GroupResponse), nil
}

func (s *gRPCServer) Update(ctx context.Context, req *pb.GroupUpdateRequest) (*pb.GroupResponse, error) {
	_, resp, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GroupResponse), nil
}

func (s *gRPCServer) Get(ctx context.Context, req *pb.GroupGetRequest) (*pb.GroupResponse, error) {
	_, resp, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GroupResponse), nil
}

func (s *gRPCServer) List(ctx context.Context, req *pb.GroupListRequest) (*pb.GroupListResponse, error) {
	_, resp, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GroupListResponse), nil
}

// CREATE
func decodeCreateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GroupCreateRequest)
	return CreateRequest{
		Group: NewGroup().
			SetActive(req.GetActive()).
			SetName(req.GetName()).
			SetUserID(int(req.GetUserID())).
			SetPhysicalEnvironmentID(int(req.GetPhysicalEnvironmentID())),
	}, nil

}

func encodeCreateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(CreateResponse)
	return &pb.GroupResponse{
		ID:     int32(resp.Group.GetID()),
		Name:   resp.Group.GetName(),
		Active: resp.Group.GetActive(),
		UserID: int32(resp.Group.GetUserID()),
	}, nil
}

// UPDATE
func decodeUpdateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GroupUpdateRequest)
	return UpdateRequest{
		Group: NewGroup().
			SetActive(req.GetActive()).
			SetName(req.GetName()).
			SetUserID(int(req.GetUserID())).
			SetPhysicalEnvironmentID(int(req.GetPhysicalEnvironmentID())),
	}, nil

}
func encodeUpdateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(UpdateResponse)
	return &pb.GroupResponse{
		ID:                    int32(resp.Group.GetID()),
		Name:                  resp.Group.GetName(),
		Active:                resp.Group.GetActive(),
		UserID:                int32(resp.Group.GetID()),
		PhysicalEnvironmentID: int32(resp.Group.GetPhysicalEnvironmentID()),
	}, nil
}

// GET

func decodeGetGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GroupGetRequest)
	return GetRequest{
		Group: NewGroup().
			SetID(int(req.GetID())),
	}, nil

}

func encodeGetGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetResponse)
	return &pb.GroupResponse{
		ID:                    int32(resp.Group.GetID()),
		Name:                  resp.Group.GetName(),
		Active:                resp.Group.GetActive(),
		UserID:                int32(resp.Group.GetUserID()),
		PhysicalEnvironmentID: int32(resp.Group.GetPhysicalEnvironmentID()),
	}, nil
}

// List

func decodListGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	return ListRequest{}, nil

}
func encodListGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ListResponse)
	var listResponse []*pb.GroupResponse

	for _, group := range resp.Group {
		listResponse = append(listResponse, &pb.GroupResponse{
			ID:     int32(group.GetID()),
			Name:   group.GetName(),
			Active: group.GetActive(),
			UserID: int32(group.GetUserID()),
		})
	}

	return &pb.GroupListResponse{
		GroupResponse: listResponse,
	}, nil
}

func optionsGRPCServer(logger log.Logger) []gt.ServerOption {
	return []gt.ServerOption{
		gt.ServerErrorLogger(logger),
		gt.ServerErrorHandler(GrpcErrorHandler{}),
	}
}

type GrpcErrorHandler struct{}

func (ge GrpcErrorHandler) Handle(_ context.Context, err error) {
	if err == nil {
		panic("encodeError with nil error")
	}
	status.Error(codes.Canceled, err.Error())
}
