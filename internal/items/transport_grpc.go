package items

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

	pb.UnimplementedItemServer
}

func NewGRPCServer(svc Service, logger log.Logger) pb.ItemServer {
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

func (s *gRPCServer) Create(ctx context.Context, req *pb.ItemCreateRequest) (*pb.ItemResponse, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ItemResponse), nil
}

func (s *gRPCServer) Update(ctx context.Context, req *pb.ItemUpdateRequest) (*pb.ItemResponse, error) {
	_, resp, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ItemResponse), nil
}

func (s *gRPCServer) Get(ctx context.Context, req *pb.ItemGetRequest) (*pb.ItemResponse, error) {
	_, resp, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.ItemResponse), nil
}

func (s *gRPCServer) List(ctx context.Context, req *pb.ItemListRequest) (*pb.ItemListResponse, error) {
	_, resp, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ItemListResponse), nil
}

// CREATE
func decodeCreateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ItemCreateRequest)
	return CreateRequest{
		Item: NewItem().
			SetActive(req.GetActive()).
			SetName(req.GetName()).
			SetUserID(int(req.GetUserID())).
			SetPhysicalEnvironmentID(int(req.GetPhysicalEnvironmentID())),
	}, nil

}

func encodeCreateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(CreateResponse)
	return &pb.ItemResponse{
		ID:     int32(resp.Item.GetID()),
		Name:   resp.Item.GetName(),
		Active: resp.Item.GetActive(),
		UserID: int32(resp.Item.GetUserID()),
	}, nil
}

// UPDATE
func decodeUpdateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ItemUpdateRequest)
	return UpdateRequest{
		Item: NewItem().
			SetActive(req.GetActive()).
			SetName(req.GetName()).
			SetUserID(int(req.GetUserID())).
			SetPhysicalEnvironmentID(int(req.GetPhysicalEnvironmentID())),
	}, nil

}
func encodeUpdateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(UpdateResponse)
	return &pb.ItemResponse{
		ID:                    int32(resp.Item.GetID()),
		Name:                  resp.Item.GetName(),
		Active:                resp.Item.GetActive(),
		UserID:                int32(resp.Item.GetID()),
		PhysicalEnvironmentID: int32(resp.Item.GetPhysicalEnvironmentID()),
	}, nil
}

// GET

func decodeGetGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ItemGetRequest)
	return GetRequest{
		Item: NewItem().
			SetID(int(req.GetID())),
	}, nil

}

func encodeGetGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetResponse)
	return &pb.ItemResponse{
		ID:                    int32(resp.Item.GetID()),
		Name:                  resp.Item.GetName(),
		Active:                resp.Item.GetActive(),
		UserID:                int32(resp.Item.GetUserID()),
		PhysicalEnvironmentID: int32(resp.Item.GetPhysicalEnvironmentID()),
	}, nil
}

// List

func decodListGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	return ListRequest{}, nil

}
func encodListGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ListResponse)
	var listResponse []*pb.ItemResponse

	for _, group := range resp.Item {
		listResponse = append(listResponse, &pb.ItemResponse{
			ID:     int32(group.GetID()),
			Name:   group.GetName(),
			Active: group.GetActive(),
			UserID: int32(group.GetUserID()),
		})
	}

	return &pb.ItemListResponse{
		ItemResponse: listResponse,
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
