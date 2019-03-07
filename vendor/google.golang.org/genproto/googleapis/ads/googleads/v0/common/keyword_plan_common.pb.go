// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/common/keyword_plan_common.proto

package common // import "google.golang.org/genproto/googleapis/ads/googleads/v0/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Historical metrics.
type KeywordPlanHistoricalMetrics struct {
	// Average monthly searches for the past 12 months.
	AvgMonthlySearches *wrappers.Int64Value `protobuf:"bytes,1,opt,name=avg_monthly_searches,json=avgMonthlySearches,proto3" json:"avg_monthly_searches,omitempty"`
	// The competition level for the query.
	Competition          enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel `protobuf:"varint,2,opt,name=competition,proto3,enum=google.ads.googleads.v0.enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel" json:"competition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                          `json:"-"`
	XXX_unrecognized     []byte                                                            `json:"-"`
	XXX_sizecache        int32                                                             `json:"-"`
}

func (m *KeywordPlanHistoricalMetrics) Reset()         { *m = KeywordPlanHistoricalMetrics{} }
func (m *KeywordPlanHistoricalMetrics) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanHistoricalMetrics) ProtoMessage()    {}
func (*KeywordPlanHistoricalMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_common_37a8ebf6dcecb987, []int{0}
}
func (m *KeywordPlanHistoricalMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanHistoricalMetrics.Unmarshal(m, b)
}
func (m *KeywordPlanHistoricalMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanHistoricalMetrics.Marshal(b, m, deterministic)
}
func (dst *KeywordPlanHistoricalMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanHistoricalMetrics.Merge(dst, src)
}
func (m *KeywordPlanHistoricalMetrics) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanHistoricalMetrics.Size(m)
}
func (m *KeywordPlanHistoricalMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanHistoricalMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanHistoricalMetrics proto.InternalMessageInfo

func (m *KeywordPlanHistoricalMetrics) GetAvgMonthlySearches() *wrappers.Int64Value {
	if m != nil {
		return m.AvgMonthlySearches
	}
	return nil
}

func (m *KeywordPlanHistoricalMetrics) GetCompetition() enums.KeywordPlanCompetitionLevelEnum_KeywordPlanCompetitionLevel {
	if m != nil {
		return m.Competition
	}
	return enums.KeywordPlanCompetitionLevelEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*KeywordPlanHistoricalMetrics)(nil), "google.ads.googleads.v0.common.KeywordPlanHistoricalMetrics")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/common/keyword_plan_common.proto", fileDescriptor_keyword_plan_common_37a8ebf6dcecb987)
}

var fileDescriptor_keyword_plan_common_37a8ebf6dcecb987 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcf, 0x4a, 0xe3, 0x40,
	0x18, 0x27, 0x59, 0xd8, 0x43, 0x0a, 0x7b, 0x08, 0xcb, 0x52, 0xba, 0x52, 0x4a, 0x4f, 0x3d, 0x4d,
	0x42, 0x15, 0x91, 0xf1, 0x94, 0xaa, 0x54, 0xd1, 0x42, 0xa9, 0x90, 0x43, 0x09, 0x84, 0x69, 0x32,
	0x4e, 0x83, 0x93, 0x99, 0x30, 0x93, 0xa4, 0x14, 0x7c, 0x1a, 0x8f, 0x3e, 0x8a, 0x8f, 0xa2, 0x27,
	0xdf, 0x40, 0x92, 0x2f, 0xad, 0x15, 0x69, 0x4f, 0xf9, 0x32, 0xf3, 0xfb, 0xfb, 0x25, 0xd6, 0x19,
	0x93, 0x92, 0x71, 0xea, 0x90, 0x58, 0x3b, 0x30, 0x56, 0x53, 0xe9, 0x3a, 0x91, 0x4c, 0x53, 0x29,
	0x9c, 0x47, 0xba, 0x5e, 0x49, 0x15, 0x87, 0x19, 0x27, 0x22, 0x84, 0x33, 0x94, 0x29, 0x99, 0x4b,
	0xbb, 0x0b, 0x70, 0x44, 0x62, 0x8d, 0xb6, 0x4c, 0x54, 0xba, 0x08, 0x50, 0x9d, 0xd1, 0x3e, 0x65,
	0x2a, 0x8a, 0x54, 0xff, 0x10, 0xce, 0x68, 0x9e, 0xe4, 0x89, 0x14, 0x21, 0xa7, 0x25, 0xe5, 0xe0,
	0xd1, 0x69, 0x3c, 0x9c, 0xfa, 0x6d, 0x51, 0x3c, 0x38, 0x2b, 0x45, 0xb2, 0x8c, 0x2a, 0x0d, 0xf7,
	0xfd, 0x77, 0xc3, 0x3a, 0xba, 0x05, 0xa1, 0x29, 0x27, 0xe2, 0x3a, 0xd1, 0xb9, 0x54, 0x49, 0x44,
	0xf8, 0x84, 0xe6, 0x2a, 0x89, 0xb4, 0x3d, 0xb1, 0xfe, 0x92, 0x92, 0x85, 0xa9, 0x14, 0xf9, 0x92,
	0xaf, 0x43, 0x4d, 0x89, 0x8a, 0x96, 0x54, 0xb7, 0x8d, 0x9e, 0x31, 0x68, 0x0d, 0xff, 0x37, 0xc1,
	0xd1, 0x46, 0x1f, 0xdd, 0x88, 0xfc, 0xf4, 0xc4, 0x27, 0xbc, 0xa0, 0x33, 0x9b, 0x94, 0x6c, 0x02,
	0xbc, 0xfb, 0x86, 0x66, 0x3f, 0x59, 0xad, 0x9d, 0xa8, 0x6d, 0xb3, 0x67, 0x0c, 0xfe, 0x0c, 0xe7,
	0x68, 0xdf, 0x26, 0xea, 0xa6, 0x68, 0x27, 0xe0, 0xc5, 0x17, 0xf9, 0xae, 0xaa, 0x79, 0x25, 0x8a,
	0xf4, 0xd0, 0xfd, 0x6c, 0xd7, 0x6e, 0xf4, 0x61, 0x58, 0xfd, 0x48, 0xa6, 0xe8, 0xf0, 0xe2, 0x47,
	0xff, 0xbe, 0x0b, 0xa6, 0x52, 0x4c, 0xab, 0x7a, 0x53, 0x63, 0x7e, 0xd9, 0x30, 0x99, 0xe4, 0x44,
	0x30, 0x24, 0x15, 0x73, 0x18, 0x15, 0x75, 0xf9, 0xcd, 0x27, 0xca, 0x12, 0xbd, 0xef, 0x5f, 0x38,
	0x87, 0xc7, 0xb3, 0xf9, 0x6b, 0xec, 0x79, 0x2f, 0x66, 0x77, 0x0c, 0x62, 0x5e, 0xac, 0x11, 0x8c,
	0xd5, 0xe4, 0xbb, 0x08, 0x3c, 0x5f, 0x37, 0x80, 0xc0, 0x8b, 0x75, 0xb0, 0x05, 0x04, 0xbe, 0x1b,
	0x00, 0xe0, 0xcd, 0xec, 0xc3, 0x29, 0xc6, 0x5e, 0xac, 0x31, 0xde, 0x42, 0x30, 0xf6, 0x5d, 0x8c,
	0x01, 0xb4, 0xf8, 0x5d, 0xa7, 0x3b, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xa2, 0xd0, 0x96,
	0xa8, 0x02, 0x00, 0x00,
}