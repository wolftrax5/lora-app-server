// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deviceProfile.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateDeviceProfileRequest struct {
	DeviceProfile *DeviceProfile `protobuf:"bytes,1,opt,name=deviceProfile" json:"deviceProfile,omitempty"`
	// Name of the device-profile.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Organization id of the device-profile.
	OrganizationID int64 `protobuf:"varint,3,opt,name=organizationID" json:"organizationID,omitempty"`
	// Network-server id of the device-profile.
	NetworkServerID int64 `protobuf:"varint,4,opt,name=networkServerID" json:"networkServerID,omitempty"`
}

func (m *CreateDeviceProfileRequest) Reset()                    { *m = CreateDeviceProfileRequest{} }
func (m *CreateDeviceProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDeviceProfileRequest) ProtoMessage()               {}
func (*CreateDeviceProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *CreateDeviceProfileRequest) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

func (m *CreateDeviceProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateDeviceProfileRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

func (m *CreateDeviceProfileRequest) GetNetworkServerID() int64 {
	if m != nil {
		return m.NetworkServerID
	}
	return 0
}

type CreateDeviceProfileResponse struct {
	// ID of the device-profile.
	DeviceProfileID string `protobuf:"bytes,1,opt,name=deviceProfileID" json:"deviceProfileID,omitempty"`
}

func (m *CreateDeviceProfileResponse) Reset()                    { *m = CreateDeviceProfileResponse{} }
func (m *CreateDeviceProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDeviceProfileResponse) ProtoMessage()               {}
func (*CreateDeviceProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *CreateDeviceProfileResponse) GetDeviceProfileID() string {
	if m != nil {
		return m.DeviceProfileID
	}
	return ""
}

type GetDeviceProfileRequest struct {
	// ID of the device-profile.
	DeviceProfileID string `protobuf:"bytes,1,opt,name=deviceProfileID" json:"deviceProfileID,omitempty"`
}

func (m *GetDeviceProfileRequest) Reset()                    { *m = GetDeviceProfileRequest{} }
func (m *GetDeviceProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDeviceProfileRequest) ProtoMessage()               {}
func (*GetDeviceProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *GetDeviceProfileRequest) GetDeviceProfileID() string {
	if m != nil {
		return m.DeviceProfileID
	}
	return ""
}

type GetDeviceProfileResponse struct {
	DeviceProfile *DeviceProfile `protobuf:"bytes,1,opt,name=deviceProfile" json:"deviceProfile,omitempty"`
	// Name of the device-profile.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Organization id of the device-profile.
	OrganizationID int64 `protobuf:"varint,3,opt,name=organizationID" json:"organizationID,omitempty"`
	// Network-server id of the device-profile.
	NetworkServerID int64 `protobuf:"varint,4,opt,name=networkServerID" json:"networkServerID,omitempty"`
	// Timestamp when the record was created.
	CreatedAt string `protobuf:"bytes,5,opt,name=createdAt" json:"createdAt,omitempty"`
	// Timestamp when the record was last updated.
	UpdatedAt string `protobuf:"bytes,6,opt,name=updatedAt" json:"updatedAt,omitempty"`
}

func (m *GetDeviceProfileResponse) Reset()                    { *m = GetDeviceProfileResponse{} }
func (m *GetDeviceProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*GetDeviceProfileResponse) ProtoMessage()               {}
func (*GetDeviceProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *GetDeviceProfileResponse) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

func (m *GetDeviceProfileResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetDeviceProfileResponse) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

func (m *GetDeviceProfileResponse) GetNetworkServerID() int64 {
	if m != nil {
		return m.NetworkServerID
	}
	return 0
}

func (m *GetDeviceProfileResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetDeviceProfileResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type UpdateDeviceProfileRequest struct {
	DeviceProfile *DeviceProfile `protobuf:"bytes,1,opt,name=deviceProfile" json:"deviceProfile,omitempty"`
	// Name of the device-profile.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *UpdateDeviceProfileRequest) Reset()                    { *m = UpdateDeviceProfileRequest{} }
func (m *UpdateDeviceProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDeviceProfileRequest) ProtoMessage()               {}
func (*UpdateDeviceProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *UpdateDeviceProfileRequest) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

func (m *UpdateDeviceProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateDeviceProfileResponse struct {
}

func (m *UpdateDeviceProfileResponse) Reset()                    { *m = UpdateDeviceProfileResponse{} }
func (m *UpdateDeviceProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateDeviceProfileResponse) ProtoMessage()               {}
func (*UpdateDeviceProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

type DeleteDeviceProfileRequest struct {
	// ID of the device-profile.
	DeviceProfileID string `protobuf:"bytes,1,opt,name=deviceProfileID" json:"deviceProfileID,omitempty"`
}

func (m *DeleteDeviceProfileRequest) Reset()                    { *m = DeleteDeviceProfileRequest{} }
func (m *DeleteDeviceProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDeviceProfileRequest) ProtoMessage()               {}
func (*DeleteDeviceProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *DeleteDeviceProfileRequest) GetDeviceProfileID() string {
	if m != nil {
		return m.DeviceProfileID
	}
	return ""
}

type DeleteDeviceProfileResponse struct {
}

func (m *DeleteDeviceProfileResponse) Reset()                    { *m = DeleteDeviceProfileResponse{} }
func (m *DeleteDeviceProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDeviceProfileResponse) ProtoMessage()               {}
func (*DeleteDeviceProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

type ListDeviceProfileRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// Organization id to filter on.
	OrganizationID int64 `protobuf:"varint,3,opt,name=organizationID" json:"organizationID,omitempty"`
}

func (m *ListDeviceProfileRequest) Reset()                    { *m = ListDeviceProfileRequest{} }
func (m *ListDeviceProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDeviceProfileRequest) ProtoMessage()               {}
func (*ListDeviceProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *ListDeviceProfileRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

type DeviceProfileMeta struct {
	// ID of the device-profile.
	DeviceProfileID string `protobuf:"bytes,1,opt,name=deviceProfileID" json:"deviceProfileID,omitempty"`
	// Name of the device-profile.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Organization id of the device-profile.
	OrganizationID int64 `protobuf:"varint,3,opt,name=organizationID" json:"organizationID,omitempty"`
	// Network-server id of the device-profile.
	NetworkServerID int64 `protobuf:"varint,4,opt,name=networkServerID" json:"networkServerID,omitempty"`
	// Timestamp when the record was created.
	CreatedAt string `protobuf:"bytes,5,opt,name=createdAt" json:"createdAt,omitempty"`
	// Timestamp when the record was last updated.
	UpdatedAt string `protobuf:"bytes,6,opt,name=updatedAt" json:"updatedAt,omitempty"`
}

func (m *DeviceProfileMeta) Reset()                    { *m = DeviceProfileMeta{} }
func (m *DeviceProfileMeta) String() string            { return proto.CompactTextString(m) }
func (*DeviceProfileMeta) ProtoMessage()               {}
func (*DeviceProfileMeta) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *DeviceProfileMeta) GetDeviceProfileID() string {
	if m != nil {
		return m.DeviceProfileID
	}
	return ""
}

func (m *DeviceProfileMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceProfileMeta) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

func (m *DeviceProfileMeta) GetNetworkServerID() int64 {
	if m != nil {
		return m.NetworkServerID
	}
	return 0
}

func (m *DeviceProfileMeta) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *DeviceProfileMeta) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ListDeviceProfileResponse struct {
	// Total number of device-profiles.
	TotalCount int64                `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	Result     []*DeviceProfileMeta `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListDeviceProfileResponse) Reset()                    { *m = ListDeviceProfileResponse{} }
func (m *ListDeviceProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDeviceProfileResponse) ProtoMessage()               {}
func (*ListDeviceProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *ListDeviceProfileResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListDeviceProfileResponse) GetResult() []*DeviceProfileMeta {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDeviceProfileRequest)(nil), "api.CreateDeviceProfileRequest")
	proto.RegisterType((*CreateDeviceProfileResponse)(nil), "api.CreateDeviceProfileResponse")
	proto.RegisterType((*GetDeviceProfileRequest)(nil), "api.GetDeviceProfileRequest")
	proto.RegisterType((*GetDeviceProfileResponse)(nil), "api.GetDeviceProfileResponse")
	proto.RegisterType((*UpdateDeviceProfileRequest)(nil), "api.UpdateDeviceProfileRequest")
	proto.RegisterType((*UpdateDeviceProfileResponse)(nil), "api.UpdateDeviceProfileResponse")
	proto.RegisterType((*DeleteDeviceProfileRequest)(nil), "api.DeleteDeviceProfileRequest")
	proto.RegisterType((*DeleteDeviceProfileResponse)(nil), "api.DeleteDeviceProfileResponse")
	proto.RegisterType((*ListDeviceProfileRequest)(nil), "api.ListDeviceProfileRequest")
	proto.RegisterType((*DeviceProfileMeta)(nil), "api.DeviceProfileMeta")
	proto.RegisterType((*ListDeviceProfileResponse)(nil), "api.ListDeviceProfileResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeviceProfileService service

type DeviceProfileServiceClient interface {
	// Create creates the given device-profile.
	Create(ctx context.Context, in *CreateDeviceProfileRequest, opts ...grpc.CallOption) (*CreateDeviceProfileResponse, error)
	// Get returns the device-profile matching the given id.
	Get(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error)
	// Update updates the given device-profile.
	Update(ctx context.Context, in *UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*UpdateDeviceProfileResponse, error)
	// Delete deletes the device-profile matching the given id.
	Delete(ctx context.Context, in *DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*DeleteDeviceProfileResponse, error)
	// List lists the available device-profiles.
	List(ctx context.Context, in *ListDeviceProfileRequest, opts ...grpc.CallOption) (*ListDeviceProfileResponse, error)
}

type deviceProfileServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceProfileServiceClient(cc *grpc.ClientConn) DeviceProfileServiceClient {
	return &deviceProfileServiceClient{cc}
}

func (c *deviceProfileServiceClient) Create(ctx context.Context, in *CreateDeviceProfileRequest, opts ...grpc.CallOption) (*CreateDeviceProfileResponse, error) {
	out := new(CreateDeviceProfileResponse)
	err := grpc.Invoke(ctx, "/api.DeviceProfileService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Get(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error) {
	out := new(GetDeviceProfileResponse)
	err := grpc.Invoke(ctx, "/api.DeviceProfileService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Update(ctx context.Context, in *UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*UpdateDeviceProfileResponse, error) {
	out := new(UpdateDeviceProfileResponse)
	err := grpc.Invoke(ctx, "/api.DeviceProfileService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Delete(ctx context.Context, in *DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*DeleteDeviceProfileResponse, error) {
	out := new(DeleteDeviceProfileResponse)
	err := grpc.Invoke(ctx, "/api.DeviceProfileService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) List(ctx context.Context, in *ListDeviceProfileRequest, opts ...grpc.CallOption) (*ListDeviceProfileResponse, error) {
	out := new(ListDeviceProfileResponse)
	err := grpc.Invoke(ctx, "/api.DeviceProfileService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceProfileService service

type DeviceProfileServiceServer interface {
	// Create creates the given device-profile.
	Create(context.Context, *CreateDeviceProfileRequest) (*CreateDeviceProfileResponse, error)
	// Get returns the device-profile matching the given id.
	Get(context.Context, *GetDeviceProfileRequest) (*GetDeviceProfileResponse, error)
	// Update updates the given device-profile.
	Update(context.Context, *UpdateDeviceProfileRequest) (*UpdateDeviceProfileResponse, error)
	// Delete deletes the device-profile matching the given id.
	Delete(context.Context, *DeleteDeviceProfileRequest) (*DeleteDeviceProfileResponse, error)
	// List lists the available device-profiles.
	List(context.Context, *ListDeviceProfileRequest) (*ListDeviceProfileResponse, error)
}

func RegisterDeviceProfileServiceServer(s *grpc.Server, srv DeviceProfileServiceServer) {
	s.RegisterService(&_DeviceProfileService_serviceDesc, srv)
}

func _DeviceProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Create(ctx, req.(*CreateDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Get(ctx, req.(*GetDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Update(ctx, req.(*UpdateDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Delete(ctx, req.(*DeleteDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceProfileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).List(ctx, req.(*ListDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceProfileService",
	HandlerType: (*DeviceProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DeviceProfileService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DeviceProfileService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeviceProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeviceProfileService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceProfileService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deviceProfile.proto",
}

func init() { proto.RegisterFile("deviceProfile.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xeb, 0xd4, 0x52, 0xa7, 0xa2, 0x88, 0x25, 0x0a, 0xae, 0x93, 0xb4, 0x91, 0x0f, 0x95,
	0x55, 0x09, 0x47, 0x0a, 0x1c, 0x80, 0x0b, 0x42, 0xb1, 0x88, 0x22, 0x81, 0x84, 0x8c, 0xf8, 0x80,
	0x25, 0x99, 0x44, 0x4b, 0x5d, 0xaf, 0xb1, 0x37, 0x45, 0xa2, 0xe2, 0xc2, 0x99, 0x1b, 0xbf, 0xc3,
	0x5f, 0xf0, 0x01, 0x5c, 0xb8, 0xf0, 0x03, 0x9c, 0x91, 0x77, 0xb7, 0x50, 0xbb, 0xde, 0xca, 0x17,
	0x24, 0xb8, 0x65, 0x67, 0xc6, 0xf3, 0x66, 0xde, 0xcc, 0x9b, 0xc0, 0xed, 0x25, 0x9e, 0xb1, 0x05,
	0xbe, 0xc8, 0xf9, 0x8a, 0x25, 0x18, 0x66, 0x39, 0x17, 0x9c, 0xd8, 0x34, 0x63, 0xde, 0x60, 0xcd,
	0xf9, 0x3a, 0xc1, 0x31, 0xcd, 0xd8, 0x98, 0xa6, 0x29, 0x17, 0x54, 0x30, 0x9e, 0x16, 0x2a, 0xc4,
	0xdb, 0xcb, 0xd4, 0x17, 0xfa, 0xed, 0x7f, 0xb1, 0xc0, 0x9b, 0xe6, 0x48, 0x05, 0x46, 0x97, 0x13,
	0xc6, 0xf8, 0x76, 0x83, 0x85, 0x20, 0x0f, 0xe0, 0x46, 0x05, 0xc8, 0xb5, 0x46, 0x56, 0xb0, 0x3b,
	0x21, 0x21, 0xcd, 0x58, 0x58, 0xfd, 0xa2, 0x1a, 0x48, 0x08, 0x74, 0x52, 0x7a, 0x8a, 0xee, 0xd6,
	0xc8, 0x0a, 0x76, 0x62, 0xf9, 0x9b, 0x1c, 0xc1, 0x1e, 0xcf, 0xd7, 0x34, 0x65, 0xef, 0x65, 0x4d,
	0xf3, 0xc8, 0xb5, 0x47, 0x56, 0x60, 0xc7, 0x35, 0x2b, 0x09, 0xe0, 0x66, 0x8a, 0xe2, 0x1d, 0xcf,
	0x4f, 0x5e, 0x62, 0x7e, 0x86, 0xf9, 0x3c, 0x72, 0x3b, 0x32, 0xb0, 0x6e, 0xf6, 0x67, 0xd0, 0x6f,
	0xac, 0xbe, 0xc8, 0x78, 0x5a, 0x60, 0x99, 0xa8, 0x52, 0xd5, 0x3c, 0x92, 0x0d, 0xec, 0xc4, 0x75,
	0xb3, 0x3f, 0x85, 0x3b, 0x33, 0x14, 0x8d, 0x1c, 0xb4, 0x4f, 0xf2, 0xd3, 0x02, 0xf7, 0x6a, 0x16,
	0x5d, 0xcb, 0x3f, 0x4e, 0x25, 0x19, 0xc0, 0xce, 0x42, 0x52, 0xb9, 0x7c, 0x22, 0xdc, 0x6d, 0x09,
	0xf5, 0xc7, 0x50, 0x7a, 0x37, 0xd9, 0x52, 0x7b, 0x1d, 0xe5, 0xfd, 0x6d, 0xf0, 0xdf, 0x80, 0xf7,
	0x4a, 0x3e, 0xfe, 0xfe, 0x12, 0xf9, 0x43, 0xe8, 0x37, 0x62, 0x29, 0x9a, 0xfd, 0xa7, 0xe0, 0x45,
	0x98, 0xa0, 0xa1, 0x94, 0xf6, 0xb3, 0x1c, 0x42, 0xbf, 0x31, 0x8f, 0x86, 0xc9, 0xc0, 0x7d, 0xc6,
	0x8a, 0xe6, 0x85, 0xe9, 0xc2, 0x76, 0xc2, 0x4e, 0x99, 0x90, 0xa9, 0xed, 0x58, 0x3d, 0x48, 0x0f,
	0x1c, 0xbe, 0x5a, 0x15, 0x28, 0x64, 0x37, 0x76, 0xac, 0x5f, 0x6d, 0x27, 0xe9, 0x7f, 0xb3, 0xe0,
	0x56, 0x05, 0xee, 0x39, 0x0a, 0xda, 0xbe, 0xa1, 0xff, 0x60, 0x8b, 0x4e, 0x60, 0xbf, 0x81, 0x53,
	0x2d, 0x9f, 0x03, 0x00, 0xc1, 0x05, 0x4d, 0xa6, 0x7c, 0x93, 0x5e, 0x30, 0x7b, 0xc9, 0x42, 0x42,
	0x70, 0x72, 0x2c, 0x36, 0x49, 0x49, 0xaf, 0x1d, 0xec, 0x4e, 0x7a, 0x57, 0xb7, 0xab, 0x24, 0x2c,
	0xd6, 0x51, 0x93, 0x1f, 0x1d, 0xe8, 0x56, 0xbc, 0x65, 0x0b, 0x6c, 0x81, 0x24, 0x01, 0x47, 0x9d,
	0x14, 0x72, 0x28, 0x53, 0x98, 0xaf, 0xa3, 0x37, 0x32, 0x07, 0xe8, 0x35, 0x39, 0xfc, 0xf8, 0xf5,
	0xfb, 0xe7, 0xad, 0x7d, 0xbf, 0x2b, 0xcf, 0xb1, 0x1a, 0xc9, 0xdd, 0x8b, 0x13, 0xfc, 0xc8, 0x3a,
	0x26, 0x39, 0xd8, 0x33, 0x14, 0x64, 0x20, 0x33, 0x19, 0x2e, 0x90, 0x37, 0x34, 0x78, 0x35, 0x48,
	0x28, 0x41, 0x02, 0x72, 0xd4, 0x04, 0x32, 0x3e, 0xaf, 0x2d, 0xc2, 0x07, 0xf2, 0xc9, 0x02, 0x47,
	0x49, 0x48, 0xb7, 0x68, 0xd6, 0xae, 0x6e, 0xf1, 0x3a, 0xc1, 0x3d, 0x96, 0xe8, 0x0f, 0xbd, 0xfb,
	0x2d, 0xd0, 0xc3, 0x7a, 0x2d, 0x25, 0x05, 0xe7, 0xe0, 0x28, 0xa5, 0xe9, 0x6a, 0xcc, 0xf2, 0xd5,
	0xd5, 0x5c, 0xa7, 0x4b, 0xcd, 0xc5, 0x71, 0x5b, 0x2e, 0x16, 0xd0, 0x29, 0x77, 0x8e, 0x28, 0x8a,
	0x4d, 0x92, 0xf6, 0x0e, 0x4c, 0x6e, 0x0d, 0x3b, 0x90, 0xb0, 0x3d, 0xd2, 0x38, 0xe7, 0xd7, 0x8e,
	0xfc, 0xaf, 0xbd, 0xf7, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x9b, 0xb4, 0x5e, 0xb5, 0x07, 0x00,
	0x00,
}