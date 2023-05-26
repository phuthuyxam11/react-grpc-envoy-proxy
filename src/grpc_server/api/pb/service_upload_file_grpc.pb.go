// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service_upload_file.proto

package pb

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

// HrTimeSheetServiceClient is the client API for HrTimeSheetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HrTimeSheetServiceClient interface {
	UploadFile(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (HrTimeSheetService_UploadFileClient, error)
}

type hrTimeSheetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHrTimeSheetServiceClient(cc grpc.ClientConnInterface) HrTimeSheetServiceClient {
	return &hrTimeSheetServiceClient{cc}
}

func (c *hrTimeSheetServiceClient) UploadFile(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (HrTimeSheetService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &HrTimeSheetService_ServiceDesc.Streams[0], "/pb.HrTimeSheetService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &hrTimeSheetServiceUploadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HrTimeSheetService_UploadFileClient interface {
	Recv() (*FileUploadResponse, error)
	grpc.ClientStream
}

type hrTimeSheetServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *hrTimeSheetServiceUploadFileClient) Recv() (*FileUploadResponse, error) {
	m := new(FileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HrTimeSheetServiceServer is the server API for HrTimeSheetService service.
// All implementations must embed UnimplementedHrTimeSheetServiceServer
// for forward compatibility
type HrTimeSheetServiceServer interface {
	UploadFile(*FileUploadRequest, HrTimeSheetService_UploadFileServer) error
	mustEmbedUnimplementedHrTimeSheetServiceServer()
}

// UnimplementedHrTimeSheetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHrTimeSheetServiceServer struct {
}

func (UnimplementedHrTimeSheetServiceServer) UploadFile(*FileUploadRequest, HrTimeSheetService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedHrTimeSheetServiceServer) mustEmbedUnimplementedHrTimeSheetServiceServer() {}

// UnsafeHrTimeSheetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HrTimeSheetServiceServer will
// result in compilation errors.
type UnsafeHrTimeSheetServiceServer interface {
	mustEmbedUnimplementedHrTimeSheetServiceServer()
}

func RegisterHrTimeSheetServiceServer(s grpc.ServiceRegistrar, srv HrTimeSheetServiceServer) {
	s.RegisterService(&HrTimeSheetService_ServiceDesc, srv)
}

func _HrTimeSheetService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileUploadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HrTimeSheetServiceServer).UploadFile(m, &hrTimeSheetServiceUploadFileServer{stream})
}

type HrTimeSheetService_UploadFileServer interface {
	Send(*FileUploadResponse) error
	grpc.ServerStream
}

type hrTimeSheetServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *hrTimeSheetServiceUploadFileServer) Send(m *FileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

// HrTimeSheetService_ServiceDesc is the grpc.ServiceDesc for HrTimeSheetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HrTimeSheetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HrTimeSheetService",
	HandlerType: (*HrTimeSheetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _HrTimeSheetService_UploadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service_upload_file.proto",
}
