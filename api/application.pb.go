// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application.proto

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

type IntegrationKind int32

const (
	IntegrationKind_HTTP IntegrationKind = 0
)

var IntegrationKind_name = map[int32]string{
	0: "HTTP",
}
var IntegrationKind_value = map[string]int32{
	"HTTP": 0,
}

func (x IntegrationKind) String() string {
	return proto.EnumName(IntegrationKind_name, int32(x))
}
func (IntegrationKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type CreateApplicationRequest struct {
	// Name of the application (must be unique).
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// ID of the organization to which the application belongs.
	OrganizationID int64 `protobuf:"varint,14,opt,name=organizationID" json:"organizationID,omitempty"`
	// ID of the service profile.
	ServiceProfileID string `protobuf:"bytes,15,opt,name=serviceProfileID" json:"serviceProfileID,omitempty"`
}

func (m *CreateApplicationRequest) Reset()                    { *m = CreateApplicationRequest{} }
func (m *CreateApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateApplicationRequest) ProtoMessage()               {}
func (*CreateApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateApplicationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateApplicationRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateApplicationRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

func (m *CreateApplicationRequest) GetServiceProfileID() string {
	if m != nil {
		return m.ServiceProfileID
	}
	return ""
}

type CreateApplicationResponse struct {
	// ID of the application that was created.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateApplicationResponse) Reset()                    { *m = CreateApplicationResponse{} }
func (m *CreateApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateApplicationResponse) ProtoMessage()               {}
func (*CreateApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CreateApplicationResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetApplicationRequest struct {
	// Name of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetApplicationRequest) Reset()                    { *m = GetApplicationRequest{} }
func (m *GetApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetApplicationRequest) ProtoMessage()               {}
func (*GetApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetApplicationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetApplicationResponse struct {
	// ID of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the application.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// ID of the organization to which the application belongs.
	OrganizationID int64 `protobuf:"varint,14,opt,name=organizationID" json:"organizationID,omitempty"`
	// ID of the service profile.
	ServiceProfileID string `protobuf:"bytes,15,opt,name=serviceProfileID" json:"serviceProfileID,omitempty"`
}

func (m *GetApplicationResponse) Reset()                    { *m = GetApplicationResponse{} }
func (m *GetApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*GetApplicationResponse) ProtoMessage()               {}
func (*GetApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetApplicationResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetApplicationResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetApplicationResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetApplicationResponse) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

func (m *GetApplicationResponse) GetServiceProfileID() string {
	if m != nil {
		return m.ServiceProfileID
	}
	return ""
}

type UpdateApplicationRequest struct {
	// ID of the application to update.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the application (must be unique).
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Description of the application.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// ID of the organization to which the application belongs.
	OrganizationID int64 `protobuf:"varint,14,opt,name=organizationID" json:"organizationID,omitempty"`
	// ID of the service profile.
	ServiceProfileID string `protobuf:"bytes,15,opt,name=serviceProfileID" json:"serviceProfileID,omitempty"`
}

func (m *UpdateApplicationRequest) Reset()                    { *m = UpdateApplicationRequest{} }
func (m *UpdateApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateApplicationRequest) ProtoMessage()               {}
func (*UpdateApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *UpdateApplicationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateApplicationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateApplicationRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateApplicationRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

func (m *UpdateApplicationRequest) GetServiceProfileID() string {
	if m != nil {
		return m.ServiceProfileID
	}
	return ""
}

type UpdateApplicationResponse struct {
}

func (m *UpdateApplicationResponse) Reset()                    { *m = UpdateApplicationResponse{} }
func (m *UpdateApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateApplicationResponse) ProtoMessage()               {}
func (*UpdateApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type DeleteApplicationRequest struct {
	// ID of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteApplicationRequest) Reset()                    { *m = DeleteApplicationRequest{} }
func (m *DeleteApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteApplicationRequest) ProtoMessage()               {}
func (*DeleteApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *DeleteApplicationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteApplicationResponse struct {
}

func (m *DeleteApplicationResponse) Reset()                    { *m = DeleteApplicationResponse{} }
func (m *DeleteApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteApplicationResponse) ProtoMessage()               {}
func (*DeleteApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type ListApplicationRequest struct {
	// Max number of applications to return in the result-test.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	// ID of the organization to filter on.
	OrganizationID int64 `protobuf:"varint,3,opt,name=organizationID" json:"organizationID,omitempty"`
}

func (m *ListApplicationRequest) Reset()                    { *m = ListApplicationRequest{} }
func (m *ListApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*ListApplicationRequest) ProtoMessage()               {}
func (*ListApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ListApplicationRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListApplicationRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListApplicationRequest) GetOrganizationID() int64 {
	if m != nil {
		return m.OrganizationID
	}
	return 0
}

type ListApplicationResponse struct {
	// Total number of applications available within the result-set.
	TotalCount int64 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// Applications within this result-set.
	Result []*GetApplicationResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListApplicationResponse) Reset()                    { *m = ListApplicationResponse{} }
func (m *ListApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*ListApplicationResponse) ProtoMessage()               {}
func (*ListApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ListApplicationResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListApplicationResponse) GetResult() []*GetApplicationResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

type HTTPIntegrationHeader struct {
	// Key
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Value
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *HTTPIntegrationHeader) Reset()                    { *m = HTTPIntegrationHeader{} }
func (m *HTTPIntegrationHeader) String() string            { return proto.CompactTextString(m) }
func (*HTTPIntegrationHeader) ProtoMessage()               {}
func (*HTTPIntegrationHeader) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *HTTPIntegrationHeader) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HTTPIntegrationHeader) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HTTPIntegration struct {
	// The id of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The headers to use when making HTTP callbacks.
	Headers []*HTTPIntegrationHeader `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty"`
	// The URL to call for uplink data.
	DataUpURL string `protobuf:"bytes,3,opt,name=dataUpURL" json:"dataUpURL,omitempty"`
	// The URL to call for join notifications.
	JoinNotificationURL string `protobuf:"bytes,4,opt,name=joinNotificationURL" json:"joinNotificationURL,omitempty"`
	// The URL to call for ACK notifications (for confirmed downlink data).
	AckNotificationURL string `protobuf:"bytes,5,opt,name=ackNotificationURL" json:"ackNotificationURL,omitempty"`
	// The URL to call for error notifications.
	ErrorNotificationURL string `protobuf:"bytes,6,opt,name=errorNotificationURL" json:"errorNotificationURL,omitempty"`
}

func (m *HTTPIntegration) Reset()                    { *m = HTTPIntegration{} }
func (m *HTTPIntegration) String() string            { return proto.CompactTextString(m) }
func (*HTTPIntegration) ProtoMessage()               {}
func (*HTTPIntegration) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *HTTPIntegration) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HTTPIntegration) GetHeaders() []*HTTPIntegrationHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPIntegration) GetDataUpURL() string {
	if m != nil {
		return m.DataUpURL
	}
	return ""
}

func (m *HTTPIntegration) GetJoinNotificationURL() string {
	if m != nil {
		return m.JoinNotificationURL
	}
	return ""
}

func (m *HTTPIntegration) GetAckNotificationURL() string {
	if m != nil {
		return m.AckNotificationURL
	}
	return ""
}

func (m *HTTPIntegration) GetErrorNotificationURL() string {
	if m != nil {
		return m.ErrorNotificationURL
	}
	return ""
}

type GetHTTPIntegrationRequest struct {
	// The id of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetHTTPIntegrationRequest) Reset()                    { *m = GetHTTPIntegrationRequest{} }
func (m *GetHTTPIntegrationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetHTTPIntegrationRequest) ProtoMessage()               {}
func (*GetHTTPIntegrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *GetHTTPIntegrationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteIntegrationRequest struct {
	// The id of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteIntegrationRequest) Reset()                    { *m = DeleteIntegrationRequest{} }
func (m *DeleteIntegrationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteIntegrationRequest) ProtoMessage()               {}
func (*DeleteIntegrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *DeleteIntegrationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListIntegrationRequest struct {
	// The id of the application.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ListIntegrationRequest) Reset()                    { *m = ListIntegrationRequest{} }
func (m *ListIntegrationRequest) String() string            { return proto.CompactTextString(m) }
func (*ListIntegrationRequest) ProtoMessage()               {}
func (*ListIntegrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *ListIntegrationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListIntegrationResponse struct {
	// The integration kinds associated with the application.
	Kinds []IntegrationKind `protobuf:"varint,1,rep,packed,name=kinds,enum=api.IntegrationKind" json:"kinds,omitempty"`
}

func (m *ListIntegrationResponse) Reset()                    { *m = ListIntegrationResponse{} }
func (m *ListIntegrationResponse) String() string            { return proto.CompactTextString(m) }
func (*ListIntegrationResponse) ProtoMessage()               {}
func (*ListIntegrationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *ListIntegrationResponse) GetKinds() []IntegrationKind {
	if m != nil {
		return m.Kinds
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateApplicationRequest)(nil), "api.CreateApplicationRequest")
	proto.RegisterType((*CreateApplicationResponse)(nil), "api.CreateApplicationResponse")
	proto.RegisterType((*GetApplicationRequest)(nil), "api.GetApplicationRequest")
	proto.RegisterType((*GetApplicationResponse)(nil), "api.GetApplicationResponse")
	proto.RegisterType((*UpdateApplicationRequest)(nil), "api.UpdateApplicationRequest")
	proto.RegisterType((*UpdateApplicationResponse)(nil), "api.UpdateApplicationResponse")
	proto.RegisterType((*DeleteApplicationRequest)(nil), "api.DeleteApplicationRequest")
	proto.RegisterType((*DeleteApplicationResponse)(nil), "api.DeleteApplicationResponse")
	proto.RegisterType((*ListApplicationRequest)(nil), "api.ListApplicationRequest")
	proto.RegisterType((*ListApplicationResponse)(nil), "api.ListApplicationResponse")
	proto.RegisterType((*EmptyResponse)(nil), "api.EmptyResponse")
	proto.RegisterType((*HTTPIntegrationHeader)(nil), "api.HTTPIntegrationHeader")
	proto.RegisterType((*HTTPIntegration)(nil), "api.HTTPIntegration")
	proto.RegisterType((*GetHTTPIntegrationRequest)(nil), "api.GetHTTPIntegrationRequest")
	proto.RegisterType((*DeleteIntegrationRequest)(nil), "api.DeleteIntegrationRequest")
	proto.RegisterType((*ListIntegrationRequest)(nil), "api.ListIntegrationRequest")
	proto.RegisterType((*ListIntegrationResponse)(nil), "api.ListIntegrationResponse")
	proto.RegisterEnum("api.IntegrationKind", IntegrationKind_name, IntegrationKind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Application service

type ApplicationClient interface {
	// Create creates the given application.
	Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error)
	// Get returns the requested application.
	Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
	// Update updates the given application.
	Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*UpdateApplicationResponse, error)
	// Delete deletes the given application.
	Delete(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*DeleteApplicationResponse, error)
	// List lists the available applications.
	List(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ListApplicationResponse, error)
	// CreateHTTPIntegration creates an HTTP application-integration.
	CreateHTTPIntegration(ctx context.Context, in *HTTPIntegration, opts ...grpc.CallOption) (*EmptyResponse, error)
	// GetHTTPIntegration returns the HTTP application-itegration.
	GetHTTPIntegration(ctx context.Context, in *GetHTTPIntegrationRequest, opts ...grpc.CallOption) (*HTTPIntegration, error)
	// UpdateHTTPIntegration updates the HTTP application-integration.
	UpdateHTTPIntegration(ctx context.Context, in *HTTPIntegration, opts ...grpc.CallOption) (*EmptyResponse, error)
	// DeleteIntegration deletes the application-integration of the given type.
	DeleteHTTPIntegration(ctx context.Context, in *DeleteIntegrationRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// ListIntegrations lists all configured integrations.
	ListIntegrations(ctx context.Context, in *ListIntegrationRequest, opts ...grpc.CallOption) (*ListIntegrationResponse, error)
}

type applicationClient struct {
	cc *grpc.ClientConn
}

func NewApplicationClient(cc *grpc.ClientConn) ApplicationClient {
	return &applicationClient{cc}
}

func (c *applicationClient) Create(ctx context.Context, in *CreateApplicationRequest, opts ...grpc.CallOption) (*CreateApplicationResponse, error) {
	out := new(CreateApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Get(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	out := new(GetApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Update(ctx context.Context, in *UpdateApplicationRequest, opts ...grpc.CallOption) (*UpdateApplicationResponse, error) {
	out := new(UpdateApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Delete(ctx context.Context, in *DeleteApplicationRequest, opts ...grpc.CallOption) (*DeleteApplicationResponse, error) {
	out := new(DeleteApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) List(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ListApplicationResponse, error) {
	out := new(ListApplicationResponse)
	err := grpc.Invoke(ctx, "/api.Application/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) CreateHTTPIntegration(ctx context.Context, in *HTTPIntegration, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/api.Application/CreateHTTPIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) GetHTTPIntegration(ctx context.Context, in *GetHTTPIntegrationRequest, opts ...grpc.CallOption) (*HTTPIntegration, error) {
	out := new(HTTPIntegration)
	err := grpc.Invoke(ctx, "/api.Application/GetHTTPIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) UpdateHTTPIntegration(ctx context.Context, in *HTTPIntegration, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/api.Application/UpdateHTTPIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) DeleteHTTPIntegration(ctx context.Context, in *DeleteIntegrationRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/api.Application/DeleteHTTPIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) ListIntegrations(ctx context.Context, in *ListIntegrationRequest, opts ...grpc.CallOption) (*ListIntegrationResponse, error) {
	out := new(ListIntegrationResponse)
	err := grpc.Invoke(ctx, "/api.Application/ListIntegrations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Application service

type ApplicationServer interface {
	// Create creates the given application.
	Create(context.Context, *CreateApplicationRequest) (*CreateApplicationResponse, error)
	// Get returns the requested application.
	Get(context.Context, *GetApplicationRequest) (*GetApplicationResponse, error)
	// Update updates the given application.
	Update(context.Context, *UpdateApplicationRequest) (*UpdateApplicationResponse, error)
	// Delete deletes the given application.
	Delete(context.Context, *DeleteApplicationRequest) (*DeleteApplicationResponse, error)
	// List lists the available applications.
	List(context.Context, *ListApplicationRequest) (*ListApplicationResponse, error)
	// CreateHTTPIntegration creates an HTTP application-integration.
	CreateHTTPIntegration(context.Context, *HTTPIntegration) (*EmptyResponse, error)
	// GetHTTPIntegration returns the HTTP application-itegration.
	GetHTTPIntegration(context.Context, *GetHTTPIntegrationRequest) (*HTTPIntegration, error)
	// UpdateHTTPIntegration updates the HTTP application-integration.
	UpdateHTTPIntegration(context.Context, *HTTPIntegration) (*EmptyResponse, error)
	// DeleteIntegration deletes the application-integration of the given type.
	DeleteHTTPIntegration(context.Context, *DeleteIntegrationRequest) (*EmptyResponse, error)
	// ListIntegrations lists all configured integrations.
	ListIntegrations(context.Context, *ListIntegrationRequest) (*ListIntegrationResponse, error)
}

func RegisterApplicationServer(s *grpc.Server, srv ApplicationServer) {
	s.RegisterService(&_Application_serviceDesc, srv)
}

func _Application_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Create(ctx, req.(*CreateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Get(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Update(ctx, req.(*UpdateApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Delete(ctx, req.(*DeleteApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).List(ctx, req.(*ListApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_CreateHTTPIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPIntegration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).CreateHTTPIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/CreateHTTPIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).CreateHTTPIntegration(ctx, req.(*HTTPIntegration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_GetHTTPIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHTTPIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).GetHTTPIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/GetHTTPIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).GetHTTPIntegration(ctx, req.(*GetHTTPIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_UpdateHTTPIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPIntegration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).UpdateHTTPIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/UpdateHTTPIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).UpdateHTTPIntegration(ctx, req.(*HTTPIntegration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_DeleteHTTPIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).DeleteHTTPIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/DeleteHTTPIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).DeleteHTTPIntegration(ctx, req.(*DeleteIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_ListIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).ListIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Application/ListIntegrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).ListIntegrations(ctx, req.(*ListIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Application_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Application",
	HandlerType: (*ApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Application_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Application_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Application_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Application_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Application_List_Handler,
		},
		{
			MethodName: "CreateHTTPIntegration",
			Handler:    _Application_CreateHTTPIntegration_Handler,
		},
		{
			MethodName: "GetHTTPIntegration",
			Handler:    _Application_GetHTTPIntegration_Handler,
		},
		{
			MethodName: "UpdateHTTPIntegration",
			Handler:    _Application_UpdateHTTPIntegration_Handler,
		},
		{
			MethodName: "DeleteHTTPIntegration",
			Handler:    _Application_DeleteHTTPIntegration_Handler,
		},
		{
			MethodName: "ListIntegrations",
			Handler:    _Application_ListIntegrations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application.proto",
}

func init() { proto.RegisterFile("application.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 800 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0x39, 0x4e, 0xf3, 0x6b, 0xa7, 0x22, 0x49, 0xb7, 0x49, 0x70, 0x9d, 0x10, 0x05, 0x23,
	0x20, 0x72, 0x45, 0x52, 0xa5, 0x9c, 0xb8, 0x20, 0xd4, 0x56, 0x6d, 0x45, 0x85, 0x2a, 0xab, 0xbd,
	0x21, 0x24, 0x13, 0x6f, 0xd2, 0x6d, 0x5d, 0xaf, 0xb1, 0x37, 0x95, 0x5a, 0xe0, 0xc2, 0x95, 0x23,
	0xaf, 0xc0, 0x13, 0x20, 0xde, 0x84, 0x57, 0xe0, 0xce, 0x2b, 0x20, 0xef, 0x6e, 0x9a, 0xd4, 0x59,
	0xa3, 0x20, 0x24, 0xc4, 0x2d, 0xbb, 0xf3, 0xcd, 0x9f, 0x6f, 0xe6, 0xdb, 0x71, 0x60, 0xc5, 0x0d,
	0x43, 0x9f, 0xf4, 0x5d, 0x46, 0x68, 0xd0, 0x09, 0x23, 0xca, 0x28, 0xd2, 0xdd, 0x90, 0x98, 0x8d,
	0x21, 0xa5, 0x43, 0x1f, 0x77, 0xdd, 0x90, 0x74, 0xdd, 0x20, 0xa0, 0x8c, 0x23, 0x62, 0x01, 0xb1,
	0x3e, 0x6b, 0x60, 0x6c, 0x45, 0xd8, 0x65, 0xf8, 0xd9, 0xc4, 0xdd, 0xc1, 0x6f, 0x46, 0x38, 0x66,
	0x08, 0x41, 0x3e, 0x70, 0xcf, 0xb1, 0xa1, 0xb5, 0xb4, 0xf6, 0x92, 0xc3, 0x7f, 0xa3, 0x16, 0x2c,
	0x7b, 0x38, 0xee, 0x47, 0x24, 0x4c, 0x90, 0x46, 0x8e, 0x9b, 0xa6, 0xaf, 0xd0, 0x03, 0x28, 0xd2,
	0x68, 0xe8, 0x06, 0xe4, 0x8a, 0x07, 0xdb, 0xdf, 0x36, 0x8a, 0x2d, 0xad, 0xad, 0x3b, 0xa9, 0x5b,
	0x64, 0x43, 0x39, 0xc6, 0xd1, 0x05, 0xe9, 0xe3, 0xc3, 0x88, 0x0e, 0x88, 0x8f, 0xf7, 0xb7, 0x8d,
	0x12, 0x0f, 0x37, 0x73, 0x6f, 0xad, 0xc3, 0x9a, 0xa2, 0xca, 0x38, 0xa4, 0x41, 0x8c, 0x51, 0x11,
	0x72, 0xc4, 0xe3, 0x45, 0xea, 0x4e, 0x8e, 0x78, 0xd6, 0x43, 0xa8, 0xee, 0x62, 0xa6, 0xe0, 0x93,
	0x06, 0x7e, 0xd1, 0xa0, 0x96, 0x46, 0xaa, 0x63, 0x5e, 0xb7, 0x22, 0x97, 0xdd, 0x0a, 0xfd, 0xef,
	0xb4, 0xe2, 0xab, 0x06, 0xc6, 0x71, 0xe8, 0xa9, 0x27, 0xf6, 0xef, 0x96, 0x5d, 0x87, 0x35, 0x45,
	0xd5, 0xa2, 0xdb, 0x96, 0x0d, 0xc6, 0x36, 0xf6, 0xf1, 0x3c, 0x94, 0x92, 0x40, 0x0a, 0xac, 0x0c,
	0x14, 0x40, 0xed, 0x80, 0xc4, 0xaa, 0xd9, 0x57, 0x60, 0xc1, 0x27, 0xe7, 0x84, 0xc9, 0x48, 0xe2,
	0x80, 0x6a, 0x50, 0xa0, 0x83, 0x41, 0x8c, 0x19, 0xef, 0x90, 0xee, 0xc8, 0x93, 0xa2, 0x03, 0xba,
	0xaa, 0x03, 0x56, 0x00, 0xb7, 0x67, 0xf2, 0x49, 0x05, 0x35, 0x01, 0x18, 0x65, 0xae, 0xbf, 0x45,
	0x47, 0xc1, 0x38, 0xeb, 0xd4, 0x0d, 0xda, 0x84, 0x42, 0x84, 0xe3, 0x91, 0x9f, 0xa4, 0xd6, 0xdb,
	0xcb, 0xbd, 0x7a, 0xc7, 0x0d, 0x49, 0x47, 0x2d, 0x47, 0x47, 0x42, 0xad, 0x12, 0xdc, 0xda, 0x39,
	0x0f, 0xd9, 0xe5, 0x35, 0xe1, 0xa7, 0x50, 0xdd, 0x3b, 0x3a, 0x3a, 0xdc, 0x0f, 0x18, 0x1e, 0x46,
	0xdc, 0x67, 0x0f, 0xbb, 0x1e, 0x8e, 0x50, 0x19, 0xf4, 0x33, 0x7c, 0x29, 0x9f, 0x6e, 0xf2, 0x33,
	0xe9, 0xc0, 0x85, 0xeb, 0x8f, 0xc6, 0x62, 0x10, 0x07, 0xeb, 0x63, 0x0e, 0x4a, 0xa9, 0x08, 0x33,
	0x2a, 0x7a, 0x0c, 0xff, 0x9f, 0xf0, 0xa8, 0xb1, 0xac, 0xd5, 0xe4, 0xb5, 0x2a, 0x13, 0x3b, 0x63,
	0x28, 0x6a, 0xc0, 0x92, 0xe7, 0x32, 0xf7, 0x38, 0x3c, 0x76, 0x0e, 0xa4, 0xca, 0x26, 0x17, 0x68,
	0x03, 0x56, 0x4f, 0x29, 0x09, 0x5e, 0x50, 0x46, 0x06, 0x92, 0x6d, 0x82, 0xcb, 0x73, 0x9c, 0xca,
	0x84, 0x3a, 0x80, 0xdc, 0xfe, 0x59, 0xda, 0x61, 0x81, 0x3b, 0x28, 0x2c, 0xa8, 0x07, 0x15, 0x1c,
	0x45, 0x34, 0x4a, 0x7b, 0x14, 0xb8, 0x87, 0xd2, 0x96, 0xec, 0x99, 0x5d, 0xcc, 0x52, 0xc4, 0xb2,
	0x94, 0x78, 0xad, 0xda, 0x39, 0xb0, 0x6d, 0x21, 0xcc, 0x39, 0x90, 0x3b, 0x42, 0x52, 0x37, 0x90,
	0x52, 0x52, 0x36, 0x2c, 0x9c, 0x91, 0xc0, 0x8b, 0x0d, 0xad, 0xa5, 0xb7, 0x8b, 0xbd, 0x0a, 0x9f,
	0xc2, 0x14, 0xf0, 0x39, 0x09, 0x3c, 0x47, 0x40, 0xec, 0x3a, 0x94, 0x52, 0x16, 0xb4, 0x08, 0xf9,
	0x84, 0x59, 0xf9, 0xbf, 0xde, 0x8f, 0x45, 0x58, 0x9e, 0x92, 0x19, 0xc2, 0x50, 0x10, 0xeb, 0x15,
	0xdd, 0xe1, 0x31, 0xb3, 0xbe, 0x08, 0x66, 0x33, 0xcb, 0x2c, 0xe5, 0xd8, 0xf8, 0xf0, 0xed, 0xfb,
	0xa7, 0x5c, 0xcd, 0x5a, 0x11, 0x9f, 0x9b, 0x09, 0x22, 0x7e, 0xa2, 0xd9, 0xe8, 0x15, 0xe8, 0xbb,
	0x98, 0x21, 0x53, 0xa9, 0x74, 0x91, 0xe0, 0x57, 0xaf, 0xc0, 0x6a, 0xf2, 0xe8, 0x06, 0xaa, 0xcd,
	0x44, 0xef, 0xbe, 0x25, 0xde, 0x7b, 0x74, 0x0a, 0x05, 0xb1, 0x63, 0x24, 0x8d, 0xac, 0x35, 0x29,
	0x69, 0x64, 0xef, 0xa3, 0xbb, 0x3c, 0x51, 0xdd, 0xcc, 0x48, 0x94, 0x70, 0x19, 0x42, 0x41, 0x0c,
	0x5f, 0xe6, 0xca, 0xda, 0x5f, 0x32, 0x57, 0xf6, 0xca, 0x92, 0xa4, 0xec, 0x2c, 0x52, 0x2f, 0x21,
	0x9f, 0xe8, 0x01, 0x89, 0xce, 0xa8, 0xb7, 0x9b, 0xd9, 0x50, 0x1b, 0x65, 0x8a, 0x35, 0x9e, 0x62,
	0x15, 0xcd, 0x4e, 0x05, 0x5d, 0x40, 0x55, 0x4c, 0x33, 0xbd, 0x03, 0x2a, 0xaa, 0x27, 0x6e, 0x22,
	0x7e, 0x7b, 0x73, 0x05, 0x6d, 0xf2, 0xe8, 0x8f, 0xac, 0xb6, 0x9a, 0x40, 0x97, 0x4c, 0xfc, 0xe3,
	0xee, 0x09, 0x63, 0x61, 0xd2, 0xbe, 0x77, 0x80, 0x66, 0x1f, 0x1a, 0x6a, 0x8e, 0xa7, 0xaf, 0x7e,
	0x81, 0xa6, 0xb2, 0x28, 0x6b, 0x83, 0x17, 0x60, 0xa3, 0xb9, 0x0b, 0x48, 0x58, 0x8b, 0xe1, 0xff,
	0x31, 0x6b, 0xf3, 0x37, 0x59, 0x57, 0x85, 0x10, 0xd2, 0x79, 0xa7, 0x35, 0xa4, 0xe0, 0xad, 0x2a,
	0x40, 0xb2, 0xb6, 0xe7, 0x67, 0x7d, 0x05, 0xe5, 0xd4, 0x66, 0x89, 0xa7, 0x54, 0xa5, 0x48, 0xdb,
	0x50, 0x1b, 0x65, 0x01, 0xeb, 0xbc, 0x80, 0xfb, 0xe8, 0xde, 0x1c, 0x05, 0xbc, 0x2e, 0xf0, 0xbf,
	0x9b, 0x9b, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xd5, 0xfe, 0xea, 0xa6, 0x0a, 0x00, 0x00,
}
