// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/logging/v2/logging.proto

package loggingpb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LoggingServiceV2Client is the client API for LoggingServiceV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggingServiceV2Client interface {
	// Deletes all the log entries in a log for the _Default Log Bucket. The log
	// reappears if it receives new entries. Log entries written shortly before
	// the delete operation might not be deleted. Entries received after the
	// delete operation with a timestamp before the operation will be deleted.
	DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Writes log entries to Logging. This API method is the
	// only way to send log entries to Logging. This method
	// is used, directly or indirectly, by the Logging agent
	// (fluentd) and all logging libraries configured to use Logging.
	// A single request may contain log entries for a maximum of 1000
	// different resources (projects, organizations, billing accounts or
	// folders)
	WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries that originated
	// from a project/folder/organization/billing account.  For ways to export log
	// entries, see [Exporting
	// Logs](https://cloud.google.com/logging/docs/export).
	ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error)
	// Lists the descriptors for monitored resource types used by Logging.
	ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error)
	// Lists the logs in projects, organizations, folders, or billing accounts.
	// Only logs that have entries are listed.
	ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error)
	// Streaming read of log entries as they are ingested. Until the stream is
	// terminated, it will continue reading logs.
	TailLogEntries(ctx context.Context, opts ...grpc.CallOption) (LoggingServiceV2_TailLogEntriesClient, error)
}

type loggingServiceV2Client struct {
	cc grpc.ClientConnInterface
}

func NewLoggingServiceV2Client(cc grpc.ClientConnInterface) LoggingServiceV2Client {
	return &loggingServiceV2Client{cc}
}

func (c *loggingServiceV2Client) DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mockgcp.logging.v2.LoggingServiceV2/DeleteLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error) {
	out := new(WriteLogEntriesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.logging.v2.LoggingServiceV2/WriteLogEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error) {
	out := new(ListLogEntriesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.logging.v2.LoggingServiceV2/ListLogEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
	out := new(ListMonitoredResourceDescriptorsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error) {
	out := new(ListLogsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.logging.v2.LoggingServiceV2/ListLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) TailLogEntries(ctx context.Context, opts ...grpc.CallOption) (LoggingServiceV2_TailLogEntriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoggingServiceV2_ServiceDesc.Streams[0], "/mockgcp.logging.v2.LoggingServiceV2/TailLogEntries", opts...)
	if err != nil {
		return nil, err
	}
	x := &loggingServiceV2TailLogEntriesClient{stream}
	return x, nil
}

type LoggingServiceV2_TailLogEntriesClient interface {
	Send(*TailLogEntriesRequest) error
	Recv() (*TailLogEntriesResponse, error)
	grpc.ClientStream
}

type loggingServiceV2TailLogEntriesClient struct {
	grpc.ClientStream
}

func (x *loggingServiceV2TailLogEntriesClient) Send(m *TailLogEntriesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loggingServiceV2TailLogEntriesClient) Recv() (*TailLogEntriesResponse, error) {
	m := new(TailLogEntriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoggingServiceV2Server is the server API for LoggingServiceV2 service.
// All implementations must embed UnimplementedLoggingServiceV2Server
// for forward compatibility
type LoggingServiceV2Server interface {
	// Deletes all the log entries in a log for the _Default Log Bucket. The log
	// reappears if it receives new entries. Log entries written shortly before
	// the delete operation might not be deleted. Entries received after the
	// delete operation with a timestamp before the operation will be deleted.
	DeleteLog(context.Context, *DeleteLogRequest) (*empty.Empty, error)
	// Writes log entries to Logging. This API method is the
	// only way to send log entries to Logging. This method
	// is used, directly or indirectly, by the Logging agent
	// (fluentd) and all logging libraries configured to use Logging.
	// A single request may contain log entries for a maximum of 1000
	// different resources (projects, organizations, billing accounts or
	// folders)
	WriteLogEntries(context.Context, *WriteLogEntriesRequest) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries that originated
	// from a project/folder/organization/billing account.  For ways to export log
	// entries, see [Exporting
	// Logs](https://cloud.google.com/logging/docs/export).
	ListLogEntries(context.Context, *ListLogEntriesRequest) (*ListLogEntriesResponse, error)
	// Lists the descriptors for monitored resource types used by Logging.
	ListMonitoredResourceDescriptors(context.Context, *ListMonitoredResourceDescriptorsRequest) (*ListMonitoredResourceDescriptorsResponse, error)
	// Lists the logs in projects, organizations, folders, or billing accounts.
	// Only logs that have entries are listed.
	ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error)
	// Streaming read of log entries as they are ingested. Until the stream is
	// terminated, it will continue reading logs.
	TailLogEntries(LoggingServiceV2_TailLogEntriesServer) error
	mustEmbedUnimplementedLoggingServiceV2Server()
}

// UnimplementedLoggingServiceV2Server must be embedded to have forward compatible implementations.
type UnimplementedLoggingServiceV2Server struct {
}

func (UnimplementedLoggingServiceV2Server) DeleteLog(context.Context, *DeleteLogRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLog not implemented")
}
func (UnimplementedLoggingServiceV2Server) WriteLogEntries(context.Context, *WriteLogEntriesRequest) (*WriteLogEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteLogEntries not implemented")
}
func (UnimplementedLoggingServiceV2Server) ListLogEntries(context.Context, *ListLogEntriesRequest) (*ListLogEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLogEntries not implemented")
}
func (UnimplementedLoggingServiceV2Server) ListMonitoredResourceDescriptors(context.Context, *ListMonitoredResourceDescriptorsRequest) (*ListMonitoredResourceDescriptorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMonitoredResourceDescriptors not implemented")
}
func (UnimplementedLoggingServiceV2Server) ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLogs not implemented")
}
func (UnimplementedLoggingServiceV2Server) TailLogEntries(LoggingServiceV2_TailLogEntriesServer) error {
	return status.Errorf(codes.Unimplemented, "method TailLogEntries not implemented")
}
func (UnimplementedLoggingServiceV2Server) mustEmbedUnimplementedLoggingServiceV2Server() {}

// UnsafeLoggingServiceV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggingServiceV2Server will
// result in compilation errors.
type UnsafeLoggingServiceV2Server interface {
	mustEmbedUnimplementedLoggingServiceV2Server()
}

func RegisterLoggingServiceV2Server(s grpc.ServiceRegistrar, srv LoggingServiceV2Server) {
	s.RegisterService(&LoggingServiceV2_ServiceDesc, srv)
}

func _LoggingServiceV2_DeleteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.logging.v2.LoggingServiceV2/DeleteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, req.(*DeleteLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_WriteLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.logging.v2.LoggingServiceV2/WriteLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, req.(*WriteLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.logging.v2.LoggingServiceV2/ListLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, req.(*ListLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMonitoredResourceDescriptorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, req.(*ListMonitoredResourceDescriptorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.logging.v2.LoggingServiceV2/ListLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListLogs(ctx, req.(*ListLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_TailLogEntries_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoggingServiceV2Server).TailLogEntries(&loggingServiceV2TailLogEntriesServer{stream})
}

type LoggingServiceV2_TailLogEntriesServer interface {
	Send(*TailLogEntriesResponse) error
	Recv() (*TailLogEntriesRequest, error)
	grpc.ServerStream
}

type loggingServiceV2TailLogEntriesServer struct {
	grpc.ServerStream
}

func (x *loggingServiceV2TailLogEntriesServer) Send(m *TailLogEntriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loggingServiceV2TailLogEntriesServer) Recv() (*TailLogEntriesRequest, error) {
	m := new(TailLogEntriesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoggingServiceV2_ServiceDesc is the grpc.ServiceDesc for LoggingServiceV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggingServiceV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.logging.v2.LoggingServiceV2",
	HandlerType: (*LoggingServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteLog",
			Handler:    _LoggingServiceV2_DeleteLog_Handler,
		},
		{
			MethodName: "WriteLogEntries",
			Handler:    _LoggingServiceV2_WriteLogEntries_Handler,
		},
		{
			MethodName: "ListLogEntries",
			Handler:    _LoggingServiceV2_ListLogEntries_Handler,
		},
		{
			MethodName: "ListMonitoredResourceDescriptors",
			Handler:    _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler,
		},
		{
			MethodName: "ListLogs",
			Handler:    _LoggingServiceV2_ListLogs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TailLogEntries",
			Handler:       _LoggingServiceV2_TailLogEntries_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mockgcp/logging/v2/logging.proto",
}
