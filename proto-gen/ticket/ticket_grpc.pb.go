// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: proto/ticket/ticket.proto

package ticket

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
	TrainTicketService_SubmitPurchase_FullMethodName    = "/ticketingapp.TrainTicketService/SubmitPurchase"
	TrainTicketService_AllocateSeat_FullMethodName      = "/ticketingapp.TrainTicketService/AllocateSeat"
	TrainTicketService_GetReceiptDetails_FullMethodName = "/ticketingapp.TrainTicketService/GetReceiptDetails"
	TrainTicketService_GetUsersBySection_FullMethodName = "/ticketingapp.TrainTicketService/GetUsersBySection"
	TrainTicketService_RemoveUser_FullMethodName        = "/ticketingapp.TrainTicketService/RemoveUser"
	TrainTicketService_ModifySeat_FullMethodName        = "/ticketingapp.TrainTicketService/ModifySeat"
	TrainTicketService_ApplyDiscount_FullMethodName     = "/ticketingapp.TrainTicketService/ApplyDiscount"
)

// TrainTicketServiceClient is the client API for TrainTicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrainTicketServiceClient interface {
	SubmitPurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Receipt, error)
	AllocateSeat(ctx context.Context, in *SeatAllocationRequest, opts ...grpc.CallOption) (*SeatAllocationResponse, error)
	GetReceiptDetails(ctx context.Context, in *ReceiptDetailsRequest, opts ...grpc.CallOption) (*Receipt, error)
	GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*UsersBySectionResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error)
	ApplyDiscount(ctx context.Context, in *DiscountRequest, opts ...grpc.CallOption) (*DiscountResponse, error)
}

type trainTicketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrainTicketServiceClient(cc grpc.ClientConnInterface) TrainTicketServiceClient {
	return &trainTicketServiceClient{cc}
}

func (c *trainTicketServiceClient) SubmitPurchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, TrainTicketService_SubmitPurchase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) AllocateSeat(ctx context.Context, in *SeatAllocationRequest, opts ...grpc.CallOption) (*SeatAllocationResponse, error) {
	out := new(SeatAllocationResponse)
	err := c.cc.Invoke(ctx, TrainTicketService_AllocateSeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) GetReceiptDetails(ctx context.Context, in *ReceiptDetailsRequest, opts ...grpc.CallOption) (*Receipt, error) {
	out := new(Receipt)
	err := c.cc.Invoke(ctx, TrainTicketService_GetReceiptDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) GetUsersBySection(ctx context.Context, in *SectionRequest, opts ...grpc.CallOption) (*UsersBySectionResponse, error) {
	out := new(UsersBySectionResponse)
	err := c.cc.Invoke(ctx, TrainTicketService_GetUsersBySection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, TrainTicketService_RemoveUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error) {
	out := new(ModifySeatResponse)
	err := c.cc.Invoke(ctx, TrainTicketService_ModifySeat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainTicketServiceClient) ApplyDiscount(ctx context.Context, in *DiscountRequest, opts ...grpc.CallOption) (*DiscountResponse, error) {
	out := new(DiscountResponse)
	err := c.cc.Invoke(ctx, TrainTicketService_ApplyDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainTicketServiceServer is the server API for TrainTicketService service.
// All implementations must embed UnimplementedTrainTicketServiceServer
// for forward compatibility
type TrainTicketServiceServer interface {
	SubmitPurchase(context.Context, *PurchaseRequest) (*Receipt, error)
	AllocateSeat(context.Context, *SeatAllocationRequest) (*SeatAllocationResponse, error)
	GetReceiptDetails(context.Context, *ReceiptDetailsRequest) (*Receipt, error)
	GetUsersBySection(context.Context, *SectionRequest) (*UsersBySectionResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error)
	ApplyDiscount(context.Context, *DiscountRequest) (*DiscountResponse, error)
	mustEmbedUnimplementedTrainTicketServiceServer()
}

// UnimplementedTrainTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrainTicketServiceServer struct {
}

func (UnimplementedTrainTicketServiceServer) SubmitPurchase(context.Context, *PurchaseRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPurchase not implemented")
}
func (UnimplementedTrainTicketServiceServer) AllocateSeat(context.Context, *SeatAllocationRequest) (*SeatAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateSeat not implemented")
}
func (UnimplementedTrainTicketServiceServer) GetReceiptDetails(context.Context, *ReceiptDetailsRequest) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceiptDetails not implemented")
}
func (UnimplementedTrainTicketServiceServer) GetUsersBySection(context.Context, *SectionRequest) (*UsersBySectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersBySection not implemented")
}
func (UnimplementedTrainTicketServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedTrainTicketServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedTrainTicketServiceServer) ApplyDiscount(context.Context, *DiscountRequest) (*DiscountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyDiscount not implemented")
}
func (UnimplementedTrainTicketServiceServer) mustEmbedUnimplementedTrainTicketServiceServer() {}

// UnsafeTrainTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrainTicketServiceServer will
// result in compilation errors.
type UnsafeTrainTicketServiceServer interface {
	mustEmbedUnimplementedTrainTicketServiceServer()
}

func RegisterTrainTicketServiceServer(s grpc.ServiceRegistrar, srv TrainTicketServiceServer) {
	s.RegisterService(&TrainTicketService_ServiceDesc, srv)
}

func _TrainTicketService_SubmitPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).SubmitPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_SubmitPurchase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).SubmitPurchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_AllocateSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeatAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).AllocateSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_AllocateSeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).AllocateSeat(ctx, req.(*SeatAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_GetReceiptDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiptDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).GetReceiptDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_GetReceiptDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).GetReceiptDetails(ctx, req.(*ReceiptDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_GetUsersBySection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).GetUsersBySection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_GetUsersBySection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).GetUsersBySection(ctx, req.(*SectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_RemoveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_ModifySeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainTicketService_ApplyDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainTicketServiceServer).ApplyDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrainTicketService_ApplyDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainTicketServiceServer).ApplyDiscount(ctx, req.(*DiscountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TrainTicketService_ServiceDesc is the grpc.ServiceDesc for TrainTicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrainTicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticketingapp.TrainTicketService",
	HandlerType: (*TrainTicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitPurchase",
			Handler:    _TrainTicketService_SubmitPurchase_Handler,
		},
		{
			MethodName: "AllocateSeat",
			Handler:    _TrainTicketService_AllocateSeat_Handler,
		},
		{
			MethodName: "GetReceiptDetails",
			Handler:    _TrainTicketService_GetReceiptDetails_Handler,
		},
		{
			MethodName: "GetUsersBySection",
			Handler:    _TrainTicketService_GetUsersBySection_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _TrainTicketService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _TrainTicketService_ModifySeat_Handler,
		},
		{
			MethodName: "ApplyDiscount",
			Handler:    _TrainTicketService_ApplyDiscount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ticket/ticket.proto",
}
