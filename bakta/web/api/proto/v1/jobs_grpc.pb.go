// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BaktaJobsClient is the client API for BaktaJobs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaktaJobsClient interface {
	InitJob(ctx context.Context, in *InitJobRequest, opts ...grpc.CallOption) (*InitJobResponse, error)
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*Empty, error)
	JobsStatus(ctx context.Context, in *JobStatusRequestList, opts ...grpc.CallOption) (*JobStatusReponseList, error)
	JobResult(ctx context.Context, in *JobAuth, opts ...grpc.CallOption) (*JobResultResponse, error)
	Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	Delete(ctx context.Context, in *JobAuth, opts ...grpc.CallOption) (*Empty, error)
}

type baktaJobsClient struct {
	cc grpc.ClientConnInterface
}

func NewBaktaJobsClient(cc grpc.ClientConnInterface) BaktaJobsClient {
	return &baktaJobsClient{cc}
}

func (c *baktaJobsClient) InitJob(ctx context.Context, in *InitJobRequest, opts ...grpc.CallOption) (*InitJobResponse, error) {
	out := new(InitJobResponse)
	err := c.cc.Invoke(ctx, "/bakta.web.api.proto.v1.BaktaJobs/InitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baktaJobsClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/bakta.web.api.proto.v1.BaktaJobs/StartJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baktaJobsClient) JobsStatus(ctx context.Context, in *JobStatusRequestList, opts ...grpc.CallOption) (*JobStatusReponseList, error) {
	out := new(JobStatusReponseList)
	err := c.cc.Invoke(ctx, "/bakta.web.api.proto.v1.BaktaJobs/JobsStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baktaJobsClient) JobResult(ctx context.Context, in *JobAuth, opts ...grpc.CallOption) (*JobResultResponse, error) {
	out := new(JobResultResponse)
	err := c.cc.Invoke(ctx, "/bakta.web.api.proto.v1.BaktaJobs/JobResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baktaJobsClient) Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/bakta.web.api.proto.v1.BaktaJobs/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baktaJobsClient) Delete(ctx context.Context, in *JobAuth, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/bakta.web.api.proto.v1.BaktaJobs/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaktaJobsServer is the server API for BaktaJobs service.
// All implementations should embed UnimplementedBaktaJobsServer
// for forward compatibility
type BaktaJobsServer interface {
	InitJob(context.Context, *InitJobRequest) (*InitJobResponse, error)
	StartJob(context.Context, *StartJobRequest) (*Empty, error)
	JobsStatus(context.Context, *JobStatusRequestList) (*JobStatusReponseList, error)
	JobResult(context.Context, *JobAuth) (*JobResultResponse, error)
	Version(context.Context, *Empty) (*VersionResponse, error)
	Delete(context.Context, *JobAuth) (*Empty, error)
}

// UnimplementedBaktaJobsServer should be embedded to have forward compatible implementations.
type UnimplementedBaktaJobsServer struct {
}

func (UnimplementedBaktaJobsServer) InitJob(context.Context, *InitJobRequest) (*InitJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitJob not implemented")
}
func (UnimplementedBaktaJobsServer) StartJob(context.Context, *StartJobRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (UnimplementedBaktaJobsServer) JobsStatus(context.Context, *JobStatusRequestList) (*JobStatusReponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobsStatus not implemented")
}
func (UnimplementedBaktaJobsServer) JobResult(context.Context, *JobAuth) (*JobResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JobResult not implemented")
}
func (UnimplementedBaktaJobsServer) Version(context.Context, *Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedBaktaJobsServer) Delete(context.Context, *JobAuth) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeBaktaJobsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaktaJobsServer will
// result in compilation errors.
type UnsafeBaktaJobsServer interface {
	mustEmbedUnimplementedBaktaJobsServer()
}

func RegisterBaktaJobsServer(s *grpc.Server, srv BaktaJobsServer) {
	s.RegisterService(&_BaktaJobs_serviceDesc, srv)
}

func _BaktaJobs_InitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaktaJobsServer).InitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bakta.web.api.proto.v1.BaktaJobs/InitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaktaJobsServer).InitJob(ctx, req.(*InitJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaktaJobs_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaktaJobsServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bakta.web.api.proto.v1.BaktaJobs/StartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaktaJobsServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaktaJobs_JobsStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobStatusRequestList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaktaJobsServer).JobsStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bakta.web.api.proto.v1.BaktaJobs/JobsStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaktaJobsServer).JobsStatus(ctx, req.(*JobStatusRequestList))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaktaJobs_JobResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaktaJobsServer).JobResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bakta.web.api.proto.v1.BaktaJobs/JobResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaktaJobsServer).JobResult(ctx, req.(*JobAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaktaJobs_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaktaJobsServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bakta.web.api.proto.v1.BaktaJobs/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaktaJobsServer).Version(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaktaJobs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaktaJobsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bakta.web.api.proto.v1.BaktaJobs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaktaJobsServer).Delete(ctx, req.(*JobAuth))
	}
	return interceptor(ctx, in, info, handler)
}

var _BaktaJobs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bakta.web.api.proto.v1.BaktaJobs",
	HandlerType: (*BaktaJobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitJob",
			Handler:    _BaktaJobs_InitJob_Handler,
		},
		{
			MethodName: "StartJob",
			Handler:    _BaktaJobs_StartJob_Handler,
		},
		{
			MethodName: "JobsStatus",
			Handler:    _BaktaJobs_JobsStatus_Handler,
		},
		{
			MethodName: "JobResult",
			Handler:    _BaktaJobs_JobResult_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _BaktaJobs_Version_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BaktaJobs_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bakta/web/api/proto/v1/jobs.proto",
}