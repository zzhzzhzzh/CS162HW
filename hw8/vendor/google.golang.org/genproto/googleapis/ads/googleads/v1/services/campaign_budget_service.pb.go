// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_budget_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for
// [CampaignBudgetService.GetCampaignBudget][google.ads.googleads.v1.services.CampaignBudgetService.GetCampaignBudget].
type GetCampaignBudgetRequest struct {
	// The resource name of the campaign budget to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignBudgetRequest) Reset()         { *m = GetCampaignBudgetRequest{} }
func (m *GetCampaignBudgetRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignBudgetRequest) ProtoMessage()    {}
func (*GetCampaignBudgetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d7b628a151e6ac, []int{0}
}

func (m *GetCampaignBudgetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignBudgetRequest.Unmarshal(m, b)
}
func (m *GetCampaignBudgetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignBudgetRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignBudgetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignBudgetRequest.Merge(m, src)
}
func (m *GetCampaignBudgetRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignBudgetRequest.Size(m)
}
func (m *GetCampaignBudgetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignBudgetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignBudgetRequest proto.InternalMessageInfo

func (m *GetCampaignBudgetRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CampaignBudgetService.MutateCampaignBudgets][google.ads.googleads.v1.services.CampaignBudgetService.MutateCampaignBudgets].
type MutateCampaignBudgetsRequest struct {
	// The ID of the customer whose campaign budgets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual campaign budgets.
	Operations []*CampaignBudgetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignBudgetsRequest) Reset()         { *m = MutateCampaignBudgetsRequest{} }
func (m *MutateCampaignBudgetsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignBudgetsRequest) ProtoMessage()    {}
func (*MutateCampaignBudgetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d7b628a151e6ac, []int{1}
}

func (m *MutateCampaignBudgetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignBudgetsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignBudgetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignBudgetsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignBudgetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignBudgetsRequest.Merge(m, src)
}
func (m *MutateCampaignBudgetsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignBudgetsRequest.Size(m)
}
func (m *MutateCampaignBudgetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignBudgetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignBudgetsRequest proto.InternalMessageInfo

func (m *MutateCampaignBudgetsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignBudgetsRequest) GetOperations() []*CampaignBudgetOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignBudgetsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignBudgetsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on a campaign budget.
type CampaignBudgetOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignBudgetOperation_Create
	//	*CampaignBudgetOperation_Update
	//	*CampaignBudgetOperation_Remove
	Operation            isCampaignBudgetOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *CampaignBudgetOperation) Reset()         { *m = CampaignBudgetOperation{} }
func (m *CampaignBudgetOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignBudgetOperation) ProtoMessage()    {}
func (*CampaignBudgetOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d7b628a151e6ac, []int{2}
}

func (m *CampaignBudgetOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignBudgetOperation.Unmarshal(m, b)
}
func (m *CampaignBudgetOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignBudgetOperation.Marshal(b, m, deterministic)
}
func (m *CampaignBudgetOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignBudgetOperation.Merge(m, src)
}
func (m *CampaignBudgetOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignBudgetOperation.Size(m)
}
func (m *CampaignBudgetOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignBudgetOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignBudgetOperation proto.InternalMessageInfo

func (m *CampaignBudgetOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignBudgetOperation_Operation interface {
	isCampaignBudgetOperation_Operation()
}

type CampaignBudgetOperation_Create struct {
	Create *resources.CampaignBudget `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignBudgetOperation_Update struct {
	Update *resources.CampaignBudget `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignBudgetOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignBudgetOperation_Create) isCampaignBudgetOperation_Operation() {}

func (*CampaignBudgetOperation_Update) isCampaignBudgetOperation_Operation() {}

func (*CampaignBudgetOperation_Remove) isCampaignBudgetOperation_Operation() {}

func (m *CampaignBudgetOperation) GetOperation() isCampaignBudgetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignBudgetOperation) GetCreate() *resources.CampaignBudget {
	if x, ok := m.GetOperation().(*CampaignBudgetOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignBudgetOperation) GetUpdate() *resources.CampaignBudget {
	if x, ok := m.GetOperation().(*CampaignBudgetOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignBudgetOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignBudgetOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignBudgetOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignBudgetOperation_Create)(nil),
		(*CampaignBudgetOperation_Update)(nil),
		(*CampaignBudgetOperation_Remove)(nil),
	}
}

// Response message for campaign budget mutate.
type MutateCampaignBudgetsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignBudgetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MutateCampaignBudgetsResponse) Reset()         { *m = MutateCampaignBudgetsResponse{} }
func (m *MutateCampaignBudgetsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignBudgetsResponse) ProtoMessage()    {}
func (*MutateCampaignBudgetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d7b628a151e6ac, []int{3}
}

func (m *MutateCampaignBudgetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignBudgetsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignBudgetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignBudgetsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignBudgetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignBudgetsResponse.Merge(m, src)
}
func (m *MutateCampaignBudgetsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignBudgetsResponse.Size(m)
}
func (m *MutateCampaignBudgetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignBudgetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignBudgetsResponse proto.InternalMessageInfo

func (m *MutateCampaignBudgetsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignBudgetsResponse) GetResults() []*MutateCampaignBudgetResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign budget mutate.
type MutateCampaignBudgetResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignBudgetResult) Reset()         { *m = MutateCampaignBudgetResult{} }
func (m *MutateCampaignBudgetResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignBudgetResult) ProtoMessage()    {}
func (*MutateCampaignBudgetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_58d7b628a151e6ac, []int{4}
}

func (m *MutateCampaignBudgetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignBudgetResult.Unmarshal(m, b)
}
func (m *MutateCampaignBudgetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignBudgetResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignBudgetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignBudgetResult.Merge(m, src)
}
func (m *MutateCampaignBudgetResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignBudgetResult.Size(m)
}
func (m *MutateCampaignBudgetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignBudgetResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignBudgetResult proto.InternalMessageInfo

func (m *MutateCampaignBudgetResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignBudgetRequest)(nil), "google.ads.googleads.v1.services.GetCampaignBudgetRequest")
	proto.RegisterType((*MutateCampaignBudgetsRequest)(nil), "google.ads.googleads.v1.services.MutateCampaignBudgetsRequest")
	proto.RegisterType((*CampaignBudgetOperation)(nil), "google.ads.googleads.v1.services.CampaignBudgetOperation")
	proto.RegisterType((*MutateCampaignBudgetsResponse)(nil), "google.ads.googleads.v1.services.MutateCampaignBudgetsResponse")
	proto.RegisterType((*MutateCampaignBudgetResult)(nil), "google.ads.googleads.v1.services.MutateCampaignBudgetResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_budget_service.proto", fileDescriptor_58d7b628a151e6ac)
}

var fileDescriptor_58d7b628a151e6ac = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xc1, 0x6b, 0xd4, 0x4e,
	0x14, 0xfe, 0x25, 0xfb, 0xa3, 0xda, 0xd9, 0xaa, 0x38, 0x52, 0xba, 0x84, 0xaa, 0x4b, 0x2c, 0x58,
	0xf6, 0x90, 0x90, 0xad, 0x20, 0x4d, 0x6d, 0xcb, 0xae, 0xd8, 0x56, 0xa4, 0xb6, 0xa4, 0xb0, 0xa0,
	0x2c, 0x84, 0x69, 0x32, 0x0d, 0xa1, 0x49, 0x26, 0xce, 0x4c, 0x56, 0x4a, 0xe9, 0x41, 0xff, 0x05,
	0x0f, 0xde, 0x3d, 0x7a, 0xf5, 0x2c, 0xde, 0xbd, 0x7a, 0xf2, 0xee, 0x41, 0xfc, 0x2b, 0x24, 0x99,
	0xcc, 0xda, 0x5d, 0x37, 0xac, 0xd6, 0xdb, 0xcb, 0x9b, 0xef, 0x7d, 0xef, 0x7d, 0xf3, 0xde, 0x9b,
	0x80, 0x8d, 0x80, 0x90, 0x20, 0xc2, 0x26, 0xf2, 0x99, 0x29, 0xcc, 0xdc, 0x1a, 0x58, 0x26, 0xc3,
	0x74, 0x10, 0x7a, 0x98, 0x99, 0x1e, 0x8a, 0x53, 0x14, 0x06, 0x89, 0x7b, 0x98, 0xf9, 0x01, 0xe6,
	0x6e, 0x79, 0x60, 0xa4, 0x94, 0x70, 0x02, 0x9b, 0x22, 0xc8, 0x40, 0x3e, 0x33, 0x86, 0xf1, 0xc6,
	0xc0, 0x32, 0x64, 0xbc, 0x76, 0xbf, 0x2a, 0x03, 0xc5, 0x8c, 0x64, 0x74, 0x42, 0x0a, 0x41, 0xad,
	0x2d, 0xca, 0xc0, 0x34, 0x34, 0x51, 0x92, 0x10, 0x8e, 0x78, 0x48, 0x12, 0x56, 0x9e, 0x96, 0x89,
	0xcd, 0xe2, 0xeb, 0x30, 0x3b, 0x32, 0x8f, 0x42, 0x1c, 0xf9, 0x6e, 0x8c, 0xd8, 0x71, 0x89, 0xb8,
	0x35, 0x8e, 0x78, 0x49, 0x51, 0x9a, 0x62, 0x2a, 0x19, 0x16, 0xca, 0x73, 0x9a, 0x7a, 0x26, 0xe3,
	0x88, 0x67, 0xe5, 0x81, 0xbe, 0x09, 0x1a, 0xdb, 0x98, 0x3f, 0x2c, 0x8b, 0xea, 0x16, 0x35, 0x39,
	0xf8, 0x45, 0x86, 0x19, 0x87, 0x77, 0xc0, 0x15, 0x59, 0xb7, 0x9b, 0xa0, 0x18, 0x37, 0x94, 0xa6,
	0xb2, 0x3c, 0xeb, 0xcc, 0x49, 0xe7, 0x53, 0x14, 0x63, 0xfd, 0xbb, 0x02, 0x16, 0x77, 0x33, 0x8e,
	0x38, 0x1e, 0x25, 0x61, 0x92, 0xe5, 0x36, 0xa8, 0x7b, 0x19, 0xe3, 0x24, 0xc6, 0xd4, 0x0d, 0xfd,
	0x92, 0x03, 0x48, 0xd7, 0x63, 0x1f, 0x3e, 0x03, 0x80, 0xa4, 0x98, 0x0a, 0xc5, 0x0d, 0xb5, 0x59,
	0x5b, 0xae, 0xb7, 0x57, 0x8d, 0x69, 0x77, 0x6d, 0x8c, 0xa6, 0xdb, 0x93, 0x0c, 0xce, 0x39, 0x32,
	0x78, 0x17, 0x5c, 0x4b, 0x11, 0xe5, 0x21, 0x8a, 0xdc, 0x23, 0x14, 0x46, 0x19, 0xc5, 0x8d, 0x5a,
	0x53, 0x59, 0xbe, 0xec, 0x5c, 0x2d, 0xdd, 0x5b, 0xc2, 0x9b, 0x4b, 0x1d, 0xa0, 0x28, 0xf4, 0x11,
	0xc7, 0x2e, 0x49, 0xa2, 0x93, 0xc6, 0xff, 0x05, 0x6c, 0x4e, 0x3a, 0xf7, 0x92, 0xe8, 0x44, 0x7f,
	0xab, 0x82, 0x85, 0x8a, 0xac, 0x70, 0x0d, 0xd4, 0xb3, 0xb4, 0x08, 0xcf, 0xbb, 0x52, 0x84, 0xd7,
	0xdb, 0x9a, 0x54, 0x21, 0xdb, 0x62, 0x6c, 0xe5, 0x8d, 0xdb, 0x45, 0xec, 0xd8, 0x01, 0x02, 0x9e,
	0xdb, 0xf0, 0x09, 0x98, 0xf1, 0x28, 0x46, 0x5c, 0xdc, 0x70, 0xbd, 0x6d, 0x55, 0xaa, 0x1f, 0xce,
	0xd1, 0x98, 0xfc, 0x9d, 0xff, 0x9c, 0x92, 0x22, 0x27, 0x13, 0xd4, 0x0d, 0xf5, 0x1f, 0xc8, 0x04,
	0x05, 0x6c, 0x80, 0x19, 0x8a, 0x63, 0x32, 0x10, 0xf7, 0x36, 0x9b, 0x9f, 0x88, 0xef, 0x6e, 0x1d,
	0xcc, 0x0e, 0x2f, 0x5a, 0xff, 0xa4, 0x80, 0x9b, 0x15, 0x43, 0xc0, 0x52, 0x92, 0x30, 0x0c, 0xb7,
	0xc0, 0xfc, 0x58, 0x27, 0x5c, 0x4c, 0x29, 0xa1, 0x05, 0x6f, 0xbd, 0x0d, 0x65, 0x91, 0x34, 0xf5,
	0x8c, 0x83, 0x62, 0x40, 0x9d, 0x1b, 0xa3, 0x3d, 0x7a, 0x94, 0xc3, 0x61, 0x0f, 0x5c, 0xa2, 0x98,
	0x65, 0x11, 0x97, 0x93, 0xf2, 0x60, 0xfa, 0xa4, 0x4c, 0xaa, 0xcc, 0x29, 0x48, 0x1c, 0x49, 0xa6,
	0x77, 0x80, 0x56, 0x0d, 0xfb, 0xa3, 0x4d, 0x68, 0x7f, 0xa8, 0x81, 0xf9, 0xd1, 0xe8, 0x03, 0x51,
	0x01, 0xfc, 0xa8, 0x80, 0xeb, 0xbf, 0x6d, 0x19, 0xb4, 0xa7, 0x57, 0x5e, 0xb5, 0x9a, 0xda, 0xdf,
	0x37, 0x55, 0x5f, 0x7d, 0xfd, 0xe5, 0xdb, 0x1b, 0x75, 0x05, 0x5a, 0xf9, 0x7b, 0x74, 0x3a, 0x22,
	0x67, 0x5d, 0x6e, 0x23, 0x33, 0x5b, 0xc3, 0x07, 0xaa, 0xec, 0xa0, 0xd9, 0x3a, 0x83, 0x5f, 0x15,
	0x30, 0x3f, 0xb1, 0xbd, 0x70, 0xe3, 0x62, 0xb7, 0x2f, 0x1f, 0x07, 0x6d, 0xf3, 0xc2, 0xf1, 0x62,
	0xae, 0xf4, 0xcd, 0x42, 0xd5, 0xaa, 0x7e, 0x2f, 0x57, 0xf5, 0x4b, 0xc6, 0xe9, 0xb9, 0x27, 0x67,
	0xbd, 0x75, 0x36, 0x2e, 0xca, 0x8e, 0x0b, 0x52, 0x5b, 0x69, 0x75, 0x5f, 0xa9, 0x60, 0xc9, 0x23,
	0xf1, 0xd4, 0x3a, 0xba, 0xda, 0xc4, 0xde, 0xee, 0xe7, 0x9b, 0xbd, 0xaf, 0x3c, 0xdf, 0x29, 0xe3,
	0x03, 0x12, 0xa1, 0x24, 0x30, 0x08, 0x0d, 0xcc, 0x00, 0x27, 0xc5, 0xde, 0xcb, 0x3f, 0x41, 0x1a,
	0xb2, 0xea, 0x5f, 0xcf, 0x9a, 0x34, 0xde, 0xa9, 0xb5, 0xed, 0x4e, 0xe7, 0xbd, 0xda, 0xdc, 0x16,
	0x84, 0x1d, 0x9f, 0x19, 0xc2, 0xcc, 0xad, 0x9e, 0x65, 0x94, 0x89, 0xd9, 0x67, 0x09, 0xe9, 0x77,
	0x7c, 0xd6, 0x1f, 0x42, 0xfa, 0x3d, 0xab, 0x2f, 0x21, 0x3f, 0xd4, 0x25, 0xe1, 0xb7, 0xed, 0x8e,
	0xcf, 0x6c, 0x7b, 0x08, 0xb2, 0xed, 0x9e, 0x65, 0xdb, 0x12, 0x76, 0x38, 0x53, 0xd4, 0xb9, 0xf2,
	0x33, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x22, 0xee, 0x20, 0x21, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignBudgetServiceClient is the client API for CampaignBudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignBudgetServiceClient interface {
	// Returns the requested Campaign Budget in full detail.
	GetCampaignBudget(ctx context.Context, in *GetCampaignBudgetRequest, opts ...grpc.CallOption) (*resources.CampaignBudget, error)
	// Creates, updates, or removes campaign budgets. Operation statuses are
	// returned.
	MutateCampaignBudgets(ctx context.Context, in *MutateCampaignBudgetsRequest, opts ...grpc.CallOption) (*MutateCampaignBudgetsResponse, error)
}

type campaignBudgetServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignBudgetServiceClient(cc *grpc.ClientConn) CampaignBudgetServiceClient {
	return &campaignBudgetServiceClient{cc}
}

func (c *campaignBudgetServiceClient) GetCampaignBudget(ctx context.Context, in *GetCampaignBudgetRequest, opts ...grpc.CallOption) (*resources.CampaignBudget, error) {
	out := new(resources.CampaignBudget)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignBudgetService/GetCampaignBudget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignBudgetServiceClient) MutateCampaignBudgets(ctx context.Context, in *MutateCampaignBudgetsRequest, opts ...grpc.CallOption) (*MutateCampaignBudgetsResponse, error) {
	out := new(MutateCampaignBudgetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignBudgetService/MutateCampaignBudgets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignBudgetServiceServer is the server API for CampaignBudgetService service.
type CampaignBudgetServiceServer interface {
	// Returns the requested Campaign Budget in full detail.
	GetCampaignBudget(context.Context, *GetCampaignBudgetRequest) (*resources.CampaignBudget, error)
	// Creates, updates, or removes campaign budgets. Operation statuses are
	// returned.
	MutateCampaignBudgets(context.Context, *MutateCampaignBudgetsRequest) (*MutateCampaignBudgetsResponse, error)
}

// UnimplementedCampaignBudgetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignBudgetServiceServer struct {
}

func (*UnimplementedCampaignBudgetServiceServer) GetCampaignBudget(ctx context.Context, req *GetCampaignBudgetRequest) (*resources.CampaignBudget, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignBudget not implemented")
}
func (*UnimplementedCampaignBudgetServiceServer) MutateCampaignBudgets(ctx context.Context, req *MutateCampaignBudgetsRequest) (*MutateCampaignBudgetsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignBudgets not implemented")
}

func RegisterCampaignBudgetServiceServer(s *grpc.Server, srv CampaignBudgetServiceServer) {
	s.RegisterService(&_CampaignBudgetService_serviceDesc, srv)
}

func _CampaignBudgetService_GetCampaignBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBudgetServiceServer).GetCampaignBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignBudgetService/GetCampaignBudget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBudgetServiceServer).GetCampaignBudget(ctx, req.(*GetCampaignBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignBudgetService_MutateCampaignBudgets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignBudgetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBudgetServiceServer).MutateCampaignBudgets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignBudgetService/MutateCampaignBudgets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBudgetServiceServer).MutateCampaignBudgets(ctx, req.(*MutateCampaignBudgetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignBudgetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignBudgetService",
	HandlerType: (*CampaignBudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignBudget",
			Handler:    _CampaignBudgetService_GetCampaignBudget_Handler,
		},
		{
			MethodName: "MutateCampaignBudgets",
			Handler:    _CampaignBudgetService_MutateCampaignBudgets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_budget_service.proto",
}
