// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: proto/notif.proto

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

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationClient interface {
	Notify(ctx context.Context, opts ...grpc.CallOption) (Notification_NotifyClient, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) Notify(ctx context.Context, opts ...grpc.CallOption) (Notification_NotifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &Notification_ServiceDesc.Streams[0], "/proto.Notification/Notify", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationNotifyClient{stream}
	return x, nil
}

type Notification_NotifyClient interface {
	Send(*Comment) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type notificationNotifyClient struct {
	grpc.ClientStream
}

func (x *notificationNotifyClient) Send(m *Comment) error {
	return x.ClientStream.SendMsg(m)
}

func (x *notificationNotifyClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationServer is the server API for Notification service.
// All implementations must embed UnimplementedNotificationServer
// for forward compatibility
type NotificationServer interface {
	Notify(Notification_NotifyServer) error
	mustEmbedUnimplementedNotificationServer()
}

// UnimplementedNotificationServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (UnimplementedNotificationServer) Notify(Notification_NotifyServer) error {
	return status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedNotificationServer) mustEmbedUnimplementedNotificationServer() {}

// UnsafeNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServer will
// result in compilation errors.
type UnsafeNotificationServer interface {
	mustEmbedUnimplementedNotificationServer()
}

func RegisterNotificationServer(s grpc.ServiceRegistrar, srv NotificationServer) {
	s.RegisterService(&Notification_ServiceDesc, srv)
}

func _Notification_Notify_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NotificationServer).Notify(&notificationNotifyServer{stream})
}

type Notification_NotifyServer interface {
	Send(*Response) error
	Recv() (*Comment, error)
	grpc.ServerStream
}

type notificationNotifyServer struct {
	grpc.ServerStream
}

func (x *notificationNotifyServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *notificationNotifyServer) Recv() (*Comment, error) {
	m := new(Comment)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Notification_ServiceDesc is the grpc.ServiceDesc for Notification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Notify",
			Handler:       _Notification_Notify_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/notif.proto",
}
