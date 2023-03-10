// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: proto/sysmon.proto

package sysmonpb

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
	SysMon_GetSnapshot_FullMethodName = "/sysmon.SysMon/GetSnapshot"
)

// SysMonClient is the client API for SysMon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysMonClient interface {
	GetSnapshot(ctx context.Context, in *Request, opts ...grpc.CallOption) (SysMon_GetSnapshotClient, error)
}

type sysMonClient struct {
	cc grpc.ClientConnInterface
}

func NewSysMonClient(cc grpc.ClientConnInterface) SysMonClient {
	return &sysMonClient{cc}
}

func (c *sysMonClient) GetSnapshot(ctx context.Context, in *Request, opts ...grpc.CallOption) (SysMon_GetSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &SysMon_ServiceDesc.Streams[0], SysMon_GetSnapshot_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sysMonGetSnapshotClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SysMon_GetSnapshotClient interface {
	Recv() (*Snapshot, error)
	grpc.ClientStream
}

type sysMonGetSnapshotClient struct {
	grpc.ClientStream
}

func (x *sysMonGetSnapshotClient) Recv() (*Snapshot, error) {
	m := new(Snapshot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SysMonServer is the server API for SysMon service.
// All implementations must embed UnimplementedSysMonServer
// for forward compatibility
type SysMonServer interface {
	GetSnapshot(*Request, SysMon_GetSnapshotServer) error
	mustEmbedUnimplementedSysMonServer()
}

// UnimplementedSysMonServer must be embedded to have forward compatible implementations.
type UnimplementedSysMonServer struct {
}

func (UnimplementedSysMonServer) GetSnapshot(*Request, SysMon_GetSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSnapshot not implemented")
}
func (UnimplementedSysMonServer) mustEmbedUnimplementedSysMonServer() {}

// UnsafeSysMonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysMonServer will
// result in compilation errors.
type UnsafeSysMonServer interface {
	mustEmbedUnimplementedSysMonServer()
}

func RegisterSysMonServer(s grpc.ServiceRegistrar, srv SysMonServer) {
	s.RegisterService(&SysMon_ServiceDesc, srv)
}

func _SysMon_GetSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SysMonServer).GetSnapshot(m, &sysMonGetSnapshotServer{stream})
}

type SysMon_GetSnapshotServer interface {
	Send(*Snapshot) error
	grpc.ServerStream
}

type sysMonGetSnapshotServer struct {
	grpc.ServerStream
}

func (x *sysMonGetSnapshotServer) Send(m *Snapshot) error {
	return x.ServerStream.SendMsg(m)
}

// SysMon_ServiceDesc is the grpc.ServiceDesc for SysMon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysMon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sysmon.SysMon",
	HandlerType: (*SysMonServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSnapshot",
			Handler:       _SysMon_GetSnapshot_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/sysmon.proto",
}
