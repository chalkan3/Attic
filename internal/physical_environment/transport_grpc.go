package physical_environment

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

	pb.UnimplementedPhysicalEnvironmentServer
}

func NewGRPCServer(svc Service, logger log.Logger) pb.PhysicalEnvironmentServer {
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

func (s *gRPCServer) Create(ctx context.Context, req *pb.PhysicalEnvironmentRequest) (*pb.PhysicalEnvironmentResponse, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PhysicalEnvironmentResponse), nil
}

func (s *gRPCServer) Update(ctx context.Context, req *pb.PhysicalEnvironmentRequest) (*pb.PhysicalEnvironmentResponse, error) {
	_, resp, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PhysicalEnvironmentResponse), nil
}

func (s *gRPCServer) Get(ctx context.Context, req *pb.PhysicalEnvironmentGetRequest) (*pb.PhysicalEnvironmentResponse, error) {
	_, resp, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.PhysicalEnvironmentResponse), nil
}

func (s *gRPCServer) List(ctx context.Context, req *pb.PhysicalEnvironmentRequestEmpty) (*pb.PhysicalEnvironmentResponseList, error) {
	_, resp, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PhysicalEnvironmentResponseList), nil
}

// CREATE
func decodeCreateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PhysicalEnvironmentRequest)
	return CreateRequest{
		PhysicalEnvironment: NewPhysicalEnvironment().
			SetActive(req.GetActive()).
			SetName(req.GetName()).
			SetUserID(int(req.GetUserID())),
	}, nil

}

func encodeCreateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(CreateResponse)
	return &pb.PhysicalEnvironmentResponse{
		ID:     int64(resp.PhysicalEnvironment.GetID()),
		Name:   resp.PhysicalEnvironment.GetName(),
		Active: resp.PhysicalEnvironment.GetActive(),
		UserID: int64(resp.PhysicalEnvironment.GetUserID()),
	}, nil
}

// UPDATE
func decodeUpdateGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PhysicalEnvironmentRequest)
	return UpdateRequest{
		PhysicalEnvironment: NewPhysicalEnvironment().
			SetActive(req.GetActive()).
			SetName(req.GetName()).
			SetUserID(int(req.GetUserID())),
	}, nil

}
func encodeUpdateGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(UpdateResponse)
	return &pb.PhysicalEnvironmentResponse{
		ID:     int64(resp.PhysicalEnvironment.GetID()),
		Name:   resp.PhysicalEnvironment.GetName(),
		Active: resp.PhysicalEnvironment.GetActive(),
		UserID: int64(resp.PhysicalEnvironment.GetUserID()),
	}, nil
}

// GET

func decodeGetGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PhysicalEnvironmentGetRequest)
	return GetRequest{
		PhysicalEnvironment: NewPhysicalEnvironment().
			SetID(int(req.GetID())),
	}, nil

}

func encodeGetGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(GetResponse)
	return &pb.PhysicalEnvironmentResponse{
		ID:     int64(resp.PhysicalEnvironment.GetID()),
		Name:   resp.PhysicalEnvironment.GetName(),
		Active: resp.PhysicalEnvironment.GetActive(),
		UserID: int64(resp.PhysicalEnvironment.GetUserID()),
	}, nil
}

// List

func decodListGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
	return ListRequest{}, nil

}
func encodListGRPCResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(ListResponse)
	var listResponse []*pb.PhysicalEnvironmentResponse

	for _, pe := range resp.PhysicalEnvironment {
		listResponse = append(listResponse, &pb.PhysicalEnvironmentResponse{
			ID:     int64(pe.GetID()),
			Name:   pe.GetName(),
			Active: pe.GetActive(),
			UserID: int64(pe.GetUserID()),
		})
	}

	return &pb.PhysicalEnvironmentResponseList{
		PhysicalEnvironment: listResponse,
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
