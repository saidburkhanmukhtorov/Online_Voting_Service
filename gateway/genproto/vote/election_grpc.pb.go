// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: protos/vote/election.proto

package vote

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ElectionService_Create_FullMethodName           = "/protos.ElectionService/Create"
	ElectionService_Update_FullMethodName           = "/protos.ElectionService/Update"
	ElectionService_Delete_FullMethodName           = "/protos.ElectionService/Delete"
	ElectionService_GetById_FullMethodName          = "/protos.ElectionService/GetById"
	ElectionService_GetAll_FullMethodName           = "/protos.ElectionService/GetAll"
	ElectionService_GetCandidateVoes_FullMethodName = "/protos.ElectionService/GetCandidateVoes"
)

// ElectionServiceClient is the client API for ElectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectionServiceClient interface {
	// Create a new election
	Create(ctx context.Context, in *ElectionCreate, opts ...grpc.CallOption) (*Election, error)
	// Update an existing election
	Update(ctx context.Context, in *ElectionUpdate, opts ...grpc.CallOption) (*Void, error)
	// Delete an existing election
	Delete(ctx context.Context, in *ElectionDelete, opts ...grpc.CallOption) (*Void, error)
	// Get an election by its ID
	GetById(ctx context.Context, in *ElectionById, opts ...grpc.CallOption) (*Election, error)
	// Get all elections
	GetAll(ctx context.Context, in *GetAllElectionReq, opts ...grpc.CallOption) (*GetAllElectionRes, error)
	// Get election result
	GetCandidateVoes(ctx context.Context, in *GetCandidateVotesReq, opts ...grpc.CallOption) (*GetCandidateVotesRes, error)
}

type electionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewElectionServiceClient(cc grpc.ClientConnInterface) ElectionServiceClient {
	return &electionServiceClient{cc}
}

func (c *electionServiceClient) Create(ctx context.Context, in *ElectionCreate, opts ...grpc.CallOption) (*Election, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Election)
	err := c.cc.Invoke(ctx, ElectionService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) Update(ctx context.Context, in *ElectionUpdate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ElectionService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) Delete(ctx context.Context, in *ElectionDelete, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ElectionService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) GetById(ctx context.Context, in *ElectionById, opts ...grpc.CallOption) (*Election, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Election)
	err := c.cc.Invoke(ctx, ElectionService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) GetAll(ctx context.Context, in *GetAllElectionReq, opts ...grpc.CallOption) (*GetAllElectionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllElectionRes)
	err := c.cc.Invoke(ctx, ElectionService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) GetCandidateVoes(ctx context.Context, in *GetCandidateVotesReq, opts ...grpc.CallOption) (*GetCandidateVotesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCandidateVotesRes)
	err := c.cc.Invoke(ctx, ElectionService_GetCandidateVoes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectionServiceServer is the server API for ElectionService service.
// All implementations must embed UnimplementedElectionServiceServer
// for forward compatibility
type ElectionServiceServer interface {
	// Create a new election
	Create(context.Context, *ElectionCreate) (*Election, error)
	// Update an existing election
	Update(context.Context, *ElectionUpdate) (*Void, error)
	// Delete an existing election
	Delete(context.Context, *ElectionDelete) (*Void, error)
	// Get an election by its ID
	GetById(context.Context, *ElectionById) (*Election, error)
	// Get all elections
	GetAll(context.Context, *GetAllElectionReq) (*GetAllElectionRes, error)
	// Get election result
	GetCandidateVoes(context.Context, *GetCandidateVotesReq) (*GetCandidateVotesRes, error)
	mustEmbedUnimplementedElectionServiceServer()
}

// UnimplementedElectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedElectionServiceServer struct {
}

func (UnimplementedElectionServiceServer) Create(context.Context, *ElectionCreate) (*Election, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedElectionServiceServer) Update(context.Context, *ElectionUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedElectionServiceServer) Delete(context.Context, *ElectionDelete) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedElectionServiceServer) GetById(context.Context, *ElectionById) (*Election, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedElectionServiceServer) GetAll(context.Context, *GetAllElectionReq) (*GetAllElectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedElectionServiceServer) GetCandidateVoes(context.Context, *GetCandidateVotesReq) (*GetCandidateVotesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCandidateVoes not implemented")
}
func (UnimplementedElectionServiceServer) mustEmbedUnimplementedElectionServiceServer() {}

// UnsafeElectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectionServiceServer will
// result in compilation errors.
type UnsafeElectionServiceServer interface {
	mustEmbedUnimplementedElectionServiceServer()
}

func RegisterElectionServiceServer(s grpc.ServiceRegistrar, srv ElectionServiceServer) {
	s.RegisterService(&ElectionService_ServiceDesc, srv)
}

func _ElectionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).Create(ctx, req.(*ElectionCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).Update(ctx, req.(*ElectionUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).Delete(ctx, req.(*ElectionDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).GetById(ctx, req.(*ElectionById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllElectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).GetAll(ctx, req.(*GetAllElectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_GetCandidateVoes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCandidateVotesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).GetCandidateVoes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectionService_GetCandidateVoes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).GetCandidateVoes(ctx, req.(*GetCandidateVotesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ElectionService_ServiceDesc is the grpc.ServiceDesc for ElectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ElectionService",
	HandlerType: (*ElectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ElectionService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ElectionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ElectionService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ElectionService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ElectionService_GetAll_Handler,
		},
		{
			MethodName: "GetCandidateVoes",
			Handler:    _ElectionService_GetCandidateVoes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/vote/election.proto",
}