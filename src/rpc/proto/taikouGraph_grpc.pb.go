// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: taikouGraph.proto

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

// TaiKouClient is the client API for TaiKou service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaiKouClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
	GetSingleVertex(ctx context.Context, in *GetSingleVertexReq, opts ...grpc.CallOption) (*GetSingleVertexResp, error)
	GetSubGraph(ctx context.Context, in *GetSubGraphReq, opts ...grpc.CallOption) (*GetSubGraphResp, error)
	GetEdgeList(ctx context.Context, in *GetEdgeListReq, opts ...grpc.CallOption) (*GetEdgeListResp, error)
	CreateKG(ctx context.Context, in *CreateKGReq, opts ...grpc.CallOption) (*CreateKGResp, error)
	PutVertexList(ctx context.Context, in *PutVertexListReq, opts ...grpc.CallOption) (*PutVertexListResp, error)
	PutEdgeList(ctx context.Context, in *PutEdgeListReq, opts ...grpc.CallOption) (*PutEdgeListResp, error)
}

type taiKouClient struct {
	cc grpc.ClientConnInterface
}

func NewTaiKouClient(cc grpc.ClientConnInterface) TaiKouClient {
	return &taiKouClient{cc}
}

func (c *taiKouClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	out := new(PingResp)
	err := c.cc.Invoke(ctx, "/TaiKou/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taiKouClient) GetSingleVertex(ctx context.Context, in *GetSingleVertexReq, opts ...grpc.CallOption) (*GetSingleVertexResp, error) {
	out := new(GetSingleVertexResp)
	err := c.cc.Invoke(ctx, "/TaiKou/GetSingleVertex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taiKouClient) GetSubGraph(ctx context.Context, in *GetSubGraphReq, opts ...grpc.CallOption) (*GetSubGraphResp, error) {
	out := new(GetSubGraphResp)
	err := c.cc.Invoke(ctx, "/TaiKou/GetSubGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taiKouClient) GetEdgeList(ctx context.Context, in *GetEdgeListReq, opts ...grpc.CallOption) (*GetEdgeListResp, error) {
	out := new(GetEdgeListResp)
	err := c.cc.Invoke(ctx, "/TaiKou/GetEdgeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taiKouClient) CreateKG(ctx context.Context, in *CreateKGReq, opts ...grpc.CallOption) (*CreateKGResp, error) {
	out := new(CreateKGResp)
	err := c.cc.Invoke(ctx, "/TaiKou/CreateKG", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taiKouClient) PutVertexList(ctx context.Context, in *PutVertexListReq, opts ...grpc.CallOption) (*PutVertexListResp, error) {
	out := new(PutVertexListResp)
	err := c.cc.Invoke(ctx, "/TaiKou/PutVertexList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taiKouClient) PutEdgeList(ctx context.Context, in *PutEdgeListReq, opts ...grpc.CallOption) (*PutEdgeListResp, error) {
	out := new(PutEdgeListResp)
	err := c.cc.Invoke(ctx, "/TaiKou/PutEdgeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaiKouServer is the server API for TaiKou service.
// All implementations must embed UnimplementedTaiKouServer
// for forward compatibility
type TaiKouServer interface {
	Ping(context.Context, *PingReq) (*PingResp, error)
	GetSingleVertex(context.Context, *GetSingleVertexReq) (*GetSingleVertexResp, error)
	GetSubGraph(context.Context, *GetSubGraphReq) (*GetSubGraphResp, error)
	GetEdgeList(context.Context, *GetEdgeListReq) (*GetEdgeListResp, error)
	CreateKG(context.Context, *CreateKGReq) (*CreateKGResp, error)
	PutVertexList(context.Context, *PutVertexListReq) (*PutVertexListResp, error)
	PutEdgeList(context.Context, *PutEdgeListReq) (*PutEdgeListResp, error)
	mustEmbedUnimplementedTaiKouServer()
}

// UnimplementedTaiKouServer must be embedded to have forward compatible implementations.
type UnimplementedTaiKouServer struct {
}

func (UnimplementedTaiKouServer) Ping(context.Context, *PingReq) (*PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTaiKouServer) GetSingleVertex(context.Context, *GetSingleVertexReq) (*GetSingleVertexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleVertex not implemented")
}
func (UnimplementedTaiKouServer) GetSubGraph(context.Context, *GetSubGraphReq) (*GetSubGraphResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubGraph not implemented")
}
func (UnimplementedTaiKouServer) GetEdgeList(context.Context, *GetEdgeListReq) (*GetEdgeListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEdgeList not implemented")
}
func (UnimplementedTaiKouServer) CreateKG(context.Context, *CreateKGReq) (*CreateKGResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKG not implemented")
}
func (UnimplementedTaiKouServer) PutVertexList(context.Context, *PutVertexListReq) (*PutVertexListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutVertexList not implemented")
}
func (UnimplementedTaiKouServer) PutEdgeList(context.Context, *PutEdgeListReq) (*PutEdgeListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutEdgeList not implemented")
}
func (UnimplementedTaiKouServer) mustEmbedUnimplementedTaiKouServer() {}

// UnsafeTaiKouServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaiKouServer will
// result in compilation errors.
type UnsafeTaiKouServer interface {
	mustEmbedUnimplementedTaiKouServer()
}

func RegisterTaiKouServer(s grpc.ServiceRegistrar, srv TaiKouServer) {
	s.RegisterService(&TaiKou_ServiceDesc, srv)
}

func _TaiKou_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiKouServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TaiKou/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiKouServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaiKou_GetSingleVertex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSingleVertexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiKouServer).GetSingleVertex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TaiKou/GetSingleVertex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiKouServer).GetSingleVertex(ctx, req.(*GetSingleVertexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaiKou_GetSubGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubGraphReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiKouServer).GetSubGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TaiKou/GetSubGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiKouServer).GetSubGraph(ctx, req.(*GetSubGraphReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaiKou_GetEdgeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEdgeListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiKouServer).GetEdgeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TaiKou/GetEdgeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiKouServer).GetEdgeList(ctx, req.(*GetEdgeListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaiKou_CreateKG_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKGReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiKouServer).CreateKG(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TaiKou/CreateKG",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiKouServer).CreateKG(ctx, req.(*CreateKGReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaiKou_PutVertexList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutVertexListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiKouServer).PutVertexList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TaiKou/PutVertexList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiKouServer).PutVertexList(ctx, req.(*PutVertexListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaiKou_PutEdgeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutEdgeListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaiKouServer).PutEdgeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TaiKou/PutEdgeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaiKouServer).PutEdgeList(ctx, req.(*PutEdgeListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TaiKou_ServiceDesc is the grpc.ServiceDesc for TaiKou service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaiKou_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TaiKou",
	HandlerType: (*TaiKouServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _TaiKou_Ping_Handler,
		},
		{
			MethodName: "GetSingleVertex",
			Handler:    _TaiKou_GetSingleVertex_Handler,
		},
		{
			MethodName: "GetSubGraph",
			Handler:    _TaiKou_GetSubGraph_Handler,
		},
		{
			MethodName: "GetEdgeList",
			Handler:    _TaiKou_GetEdgeList_Handler,
		},
		{
			MethodName: "CreateKG",
			Handler:    _TaiKou_CreateKG_Handler,
		},
		{
			MethodName: "PutVertexList",
			Handler:    _TaiKou_PutVertexList_Handler,
		},
		{
			MethodName: "PutEdgeList",
			Handler:    _TaiKou_PutEdgeList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taikouGraph.proto",
}
