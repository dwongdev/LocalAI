// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: pkg/grpc/proto/llmserver.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LLMClient is the client API for LLM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LLMClient interface {
	Health(ctx context.Context, in *HealthMessage, opts ...grpc.CallOption) (*Reply, error)
	Predict(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*Reply, error)
	LoadModel(ctx context.Context, in *ModelOptions, opts ...grpc.CallOption) (*Result, error)
	PredictStream(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (LLM_PredictStreamClient, error)
	Embedding(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*EmbeddingResult, error)
}

type lLMClient struct {
	cc grpc.ClientConnInterface
}

func NewLLMClient(cc grpc.ClientConnInterface) LLMClient {
	return &lLMClient{cc}
}

func (c *lLMClient) Health(ctx context.Context, in *HealthMessage, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/llm.LLM/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lLMClient) Predict(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/llm.LLM/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lLMClient) LoadModel(ctx context.Context, in *ModelOptions, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/llm.LLM/LoadModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lLMClient) PredictStream(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (LLM_PredictStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &LLM_ServiceDesc.Streams[0], "/llm.LLM/PredictStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &lLMPredictStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LLM_PredictStreamClient interface {
	Recv() (*Reply, error)
	grpc.ClientStream
}

type lLMPredictStreamClient struct {
	grpc.ClientStream
}

func (x *lLMPredictStreamClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lLMClient) Embedding(ctx context.Context, in *PredictOptions, opts ...grpc.CallOption) (*EmbeddingResult, error) {
	out := new(EmbeddingResult)
	err := c.cc.Invoke(ctx, "/llm.LLM/Embedding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LLMServer is the server API for LLM service.
// All implementations must embed UnimplementedLLMServer
// for forward compatibility
type LLMServer interface {
	Health(context.Context, *HealthMessage) (*Reply, error)
	Predict(context.Context, *PredictOptions) (*Reply, error)
	LoadModel(context.Context, *ModelOptions) (*Result, error)
	PredictStream(*PredictOptions, LLM_PredictStreamServer) error
	Embedding(context.Context, *PredictOptions) (*EmbeddingResult, error)
	mustEmbedUnimplementedLLMServer()
}

// UnimplementedLLMServer must be embedded to have forward compatible implementations.
type UnimplementedLLMServer struct {
}

func (UnimplementedLLMServer) Health(context.Context, *HealthMessage) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedLLMServer) Predict(context.Context, *PredictOptions) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (UnimplementedLLMServer) LoadModel(context.Context, *ModelOptions) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadModel not implemented")
}
func (UnimplementedLLMServer) PredictStream(*PredictOptions, LLM_PredictStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PredictStream not implemented")
}
func (UnimplementedLLMServer) Embedding(context.Context, *PredictOptions) (*EmbeddingResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Embedding not implemented")
}
func (UnimplementedLLMServer) mustEmbedUnimplementedLLMServer() {}

// UnsafeLLMServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LLMServer will
// result in compilation errors.
type UnsafeLLMServer interface {
	mustEmbedUnimplementedLLMServer()
}

func RegisterLLMServer(s grpc.ServiceRegistrar, srv LLMServer) {
	s.RegisterService(&LLM_ServiceDesc, srv)
}

func _LLM_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llm.LLM/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServer).Health(ctx, req.(*HealthMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _LLM_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llm.LLM/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServer).Predict(ctx, req.(*PredictOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _LLM_LoadModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServer).LoadModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llm.LLM/LoadModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServer).LoadModel(ctx, req.(*ModelOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _LLM_PredictStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PredictOptions)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LLMServer).PredictStream(m, &lLMPredictStreamServer{stream})
}

type LLM_PredictStreamServer interface {
	Send(*Reply) error
	grpc.ServerStream
}

type lLMPredictStreamServer struct {
	grpc.ServerStream
}

func (x *lLMPredictStreamServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func _LLM_Embedding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LLMServer).Embedding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llm.LLM/Embedding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LLMServer).Embedding(ctx, req.(*PredictOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// LLM_ServiceDesc is the grpc.ServiceDesc for LLM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LLM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "llm.LLM",
	HandlerType: (*LLMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _LLM_Health_Handler,
		},
		{
			MethodName: "Predict",
			Handler:    _LLM_Predict_Handler,
		},
		{
			MethodName: "LoadModel",
			Handler:    _LLM_LoadModel_Handler,
		},
		{
			MethodName: "Embedding",
			Handler:    _LLM_Embedding_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PredictStream",
			Handler:       _LLM_PredictStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/grpc/proto/llmserver.proto",
}
