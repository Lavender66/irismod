// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: irismod/htlc/tx.proto

package htlc

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

const (
	Msg_CreateHTLC_FullMethodName   = "/irismod.htlc.Msg/CreateHTLC"
	Msg_ClaimHTLC_FullMethodName    = "/irismod.htlc.Msg/ClaimHTLC"
	Msg_UpdateParams_FullMethodName = "/irismod.htlc.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateHTLC defines a method for creating a HTLC
	CreateHTLC(ctx context.Context, in *MsgCreateHTLC, opts ...grpc.CallOption) (*MsgCreateHTLCResponse, error)
	// ClaimHTLC defines a method for claiming a HTLC
	ClaimHTLC(ctx context.Context, in *MsgClaimHTLC, opts ...grpc.CallOption) (*MsgClaimHTLCResponse, error)
	// UpdateParams defines a governance operation for updating the x/htlc
	// module parameters. The authority is defined in the keeper.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateHTLC(ctx context.Context, in *MsgCreateHTLC, opts ...grpc.CallOption) (*MsgCreateHTLCResponse, error) {
	out := new(MsgCreateHTLCResponse)
	err := c.cc.Invoke(ctx, Msg_CreateHTLC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimHTLC(ctx context.Context, in *MsgClaimHTLC, opts ...grpc.CallOption) (*MsgClaimHTLCResponse, error) {
	out := new(MsgClaimHTLCResponse)
	err := c.cc.Invoke(ctx, Msg_ClaimHTLC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateHTLC defines a method for creating a HTLC
	CreateHTLC(context.Context, *MsgCreateHTLC) (*MsgCreateHTLCResponse, error)
	// ClaimHTLC defines a method for claiming a HTLC
	ClaimHTLC(context.Context, *MsgClaimHTLC) (*MsgClaimHTLCResponse, error)
	// UpdateParams defines a governance operation for updating the x/htlc
	// module parameters. The authority is defined in the keeper.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateHTLC(context.Context, *MsgCreateHTLC) (*MsgCreateHTLCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTLC not implemented")
}
func (UnimplementedMsgServer) ClaimHTLC(context.Context, *MsgClaimHTLC) (*MsgClaimHTLCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimHTLC not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateHTLC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateHTLC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateHTLC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateHTLC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateHTLC(ctx, req.(*MsgCreateHTLC))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimHTLC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimHTLC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimHTLC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ClaimHTLC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimHTLC(ctx, req.(*MsgClaimHTLC))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.htlc.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHTLC",
			Handler:    _Msg_CreateHTLC_Handler,
		},
		{
			MethodName: "ClaimHTLC",
			Handler:    _Msg_ClaimHTLC_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/htlc/tx.proto",
}