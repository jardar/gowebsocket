// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/ad_group_bid_modifier.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Represents an ad group bid modifier.
type AdGroupBidModifier struct {
	// The resource name of the ad group bid modifier.
	// Ad group bid modifier resource names have the form:
	//
	// `customers/{customer_id}/adGroupBidModifiers/{ad_group_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ad group to which this criterion belongs.
	AdGroup *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The ID of the criterion to bid modify.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bid when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. The range is 1.0 - 6.0 for PreferredContent.
	// Use 0 to opt out of a Device type.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// The base ad group from which this draft/trial adgroup bid modifier was
	// created. If ad_group is a base ad group then this field will be equal to
	// ad_group. If the ad group was created in the draft or trial and has no
	// corresponding base ad group, then this field will be null.
	// This field is readonly.
	BaseAdGroup *wrappers.StringValue `protobuf:"bytes,9,opt,name=base_ad_group,json=baseAdGroup,proto3" json:"base_ad_group,omitempty"`
	// Bid modifier source.
	BidModifierSource enums.BidModifierSourceEnum_BidModifierSource `protobuf:"varint,10,opt,name=bid_modifier_source,json=bidModifierSource,proto3,enum=google.ads.googleads.v2.enums.BidModifierSourceEnum_BidModifierSource" json:"bid_modifier_source,omitempty"`
	// The criterion of this ad group bid modifier.
	//
	// Types that are valid to be assigned to Criterion:
	//	*AdGroupBidModifier_HotelDateSelectionType
	//	*AdGroupBidModifier_HotelAdvanceBookingWindow
	//	*AdGroupBidModifier_HotelLengthOfStay
	//	*AdGroupBidModifier_HotelCheckInDay
	//	*AdGroupBidModifier_Device
	//	*AdGroupBidModifier_PreferredContent
	Criterion            isAdGroupBidModifier_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AdGroupBidModifier) Reset()         { *m = AdGroupBidModifier{} }
func (m *AdGroupBidModifier) String() string { return proto.CompactTextString(m) }
func (*AdGroupBidModifier) ProtoMessage()    {}
func (*AdGroupBidModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_77d645d3b51cc0a0, []int{0}
}

func (m *AdGroupBidModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupBidModifier.Unmarshal(m, b)
}
func (m *AdGroupBidModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupBidModifier.Marshal(b, m, deterministic)
}
func (m *AdGroupBidModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupBidModifier.Merge(m, src)
}
func (m *AdGroupBidModifier) XXX_Size() int {
	return xxx_messageInfo_AdGroupBidModifier.Size(m)
}
func (m *AdGroupBidModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupBidModifier.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupBidModifier proto.InternalMessageInfo

func (m *AdGroupBidModifier) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupBidModifier) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupBidModifier) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *AdGroupBidModifier) GetBidModifier() *wrappers.DoubleValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *AdGroupBidModifier) GetBaseAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.BaseAdGroup
	}
	return nil
}

func (m *AdGroupBidModifier) GetBidModifierSource() enums.BidModifierSourceEnum_BidModifierSource {
	if m != nil {
		return m.BidModifierSource
	}
	return enums.BidModifierSourceEnum_UNSPECIFIED
}

type isAdGroupBidModifier_Criterion interface {
	isAdGroupBidModifier_Criterion()
}

type AdGroupBidModifier_HotelDateSelectionType struct {
	HotelDateSelectionType *common.HotelDateSelectionTypeInfo `protobuf:"bytes,5,opt,name=hotel_date_selection_type,json=hotelDateSelectionType,proto3,oneof"`
}

type AdGroupBidModifier_HotelAdvanceBookingWindow struct {
	HotelAdvanceBookingWindow *common.HotelAdvanceBookingWindowInfo `protobuf:"bytes,6,opt,name=hotel_advance_booking_window,json=hotelAdvanceBookingWindow,proto3,oneof"`
}

type AdGroupBidModifier_HotelLengthOfStay struct {
	HotelLengthOfStay *common.HotelLengthOfStayInfo `protobuf:"bytes,7,opt,name=hotel_length_of_stay,json=hotelLengthOfStay,proto3,oneof"`
}

type AdGroupBidModifier_HotelCheckInDay struct {
	HotelCheckInDay *common.HotelCheckInDayInfo `protobuf:"bytes,8,opt,name=hotel_check_in_day,json=hotelCheckInDay,proto3,oneof"`
}

type AdGroupBidModifier_Device struct {
	Device *common.DeviceInfo `protobuf:"bytes,11,opt,name=device,proto3,oneof"`
}

type AdGroupBidModifier_PreferredContent struct {
	PreferredContent *common.PreferredContentInfo `protobuf:"bytes,12,opt,name=preferred_content,json=preferredContent,proto3,oneof"`
}

func (*AdGroupBidModifier_HotelDateSelectionType) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelAdvanceBookingWindow) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelLengthOfStay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_HotelCheckInDay) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_Device) isAdGroupBidModifier_Criterion() {}

func (*AdGroupBidModifier_PreferredContent) isAdGroupBidModifier_Criterion() {}

func (m *AdGroupBidModifier) GetCriterion() isAdGroupBidModifier_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelDateSelectionType() *common.HotelDateSelectionTypeInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelDateSelectionType); ok {
		return x.HotelDateSelectionType
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelAdvanceBookingWindow() *common.HotelAdvanceBookingWindowInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelAdvanceBookingWindow); ok {
		return x.HotelAdvanceBookingWindow
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelLengthOfStay() *common.HotelLengthOfStayInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelLengthOfStay); ok {
		return x.HotelLengthOfStay
	}
	return nil
}

func (m *AdGroupBidModifier) GetHotelCheckInDay() *common.HotelCheckInDayInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_HotelCheckInDay); ok {
		return x.HotelCheckInDay
	}
	return nil
}

func (m *AdGroupBidModifier) GetDevice() *common.DeviceInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_Device); ok {
		return x.Device
	}
	return nil
}

func (m *AdGroupBidModifier) GetPreferredContent() *common.PreferredContentInfo {
	if x, ok := m.GetCriterion().(*AdGroupBidModifier_PreferredContent); ok {
		return x.PreferredContent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupBidModifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupBidModifier_HotelDateSelectionType)(nil),
		(*AdGroupBidModifier_HotelAdvanceBookingWindow)(nil),
		(*AdGroupBidModifier_HotelLengthOfStay)(nil),
		(*AdGroupBidModifier_HotelCheckInDay)(nil),
		(*AdGroupBidModifier_Device)(nil),
		(*AdGroupBidModifier_PreferredContent)(nil),
	}
}

func init() {
	proto.RegisterType((*AdGroupBidModifier)(nil), "google.ads.googleads.v2.resources.AdGroupBidModifier")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/ad_group_bid_modifier.proto", fileDescriptor_77d645d3b51cc0a0)
}

var fileDescriptor_77d645d3b51cc0a0 = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0x49, 0xcf, 0xe9, 0xcf, 0x26, 0x3d, 0xe7, 0xd4, 0x20, 0x30, 0xa5, 0x42, 0x2d, 0xa8,
	0x52, 0x85, 0x84, 0x2d, 0xa5, 0x85, 0x4a, 0x46, 0x05, 0x92, 0x06, 0xda, 0x20, 0x7e, 0xaa, 0x04,
	0x05, 0x09, 0x45, 0x5a, 0xad, 0xbd, 0x13, 0x67, 0xd5, 0x78, 0xd7, 0xb2, 0x37, 0x89, 0x72, 0xc7,
	0x05, 0x4f, 0xc2, 0x25, 0x37, 0xbc, 0x07, 0x8f, 0xc2, 0x53, 0x20, 0xef, 0xae, 0xad, 0xd0, 0x36,
	0x34, 0x77, 0xe3, 0x99, 0xef, 0x67, 0x66, 0xd6, 0xbb, 0xe8, 0x28, 0x14, 0x22, 0x1c, 0x82, 0x4b,
	0x68, 0xea, 0xea, 0x30, 0x8b, 0xc6, 0x35, 0x37, 0x81, 0x54, 0x8c, 0x92, 0x00, 0x52, 0x97, 0x50,
	0x1c, 0x26, 0x62, 0x14, 0x63, 0x9f, 0x51, 0x1c, 0x09, 0xca, 0xfa, 0x0c, 0x12, 0x27, 0x4e, 0x84,
	0x14, 0xd6, 0x8e, 0xe6, 0x38, 0x84, 0xa6, 0x4e, 0x41, 0x77, 0xc6, 0x35, 0xa7, 0xa0, 0x6f, 0x3e,
	0x9a, 0xe7, 0x10, 0x88, 0x28, 0x12, 0xdc, 0x0d, 0x12, 0x26, 0x21, 0x61, 0x44, 0x2b, 0x6e, 0x1e,
	0xce, 0x83, 0x03, 0x1f, 0x45, 0xa9, 0x3b, 0xdb, 0x03, 0xd6, 0x16, 0x86, 0x78, 0xcf, 0x10, 0xd5,
	0x97, 0x3f, 0xea, 0xbb, 0x93, 0x84, 0xc4, 0x31, 0x24, 0xa9, 0xa9, 0x6f, 0xe5, 0xc2, 0x31, 0x73,
	0x09, 0xe7, 0x42, 0x12, 0xc9, 0x04, 0x37, 0xd5, 0xfb, 0xdf, 0x57, 0x91, 0x55, 0xa7, 0x27, 0xd9,
	0x9c, 0x0d, 0x46, 0xdf, 0x1a, 0x07, 0xeb, 0x01, 0x5a, 0xcf, 0x27, 0xc1, 0x9c, 0x44, 0x60, 0x97,
	0xb6, 0x4b, 0x7b, 0x6b, 0xed, 0x6a, 0x9e, 0x7c, 0x47, 0x22, 0xb0, 0x0e, 0xd1, 0x6a, 0xbe, 0x23,
	0xbb, 0xbc, 0x5d, 0xda, 0xab, 0xd4, 0xb6, 0xcc, 0x32, 0x9c, 0xbc, 0x19, 0xa7, 0x23, 0x13, 0xc6,
	0xc3, 0x2e, 0x19, 0x8e, 0xa0, 0xbd, 0x42, 0xb4, 0x91, 0xf5, 0x0c, 0x55, 0xcd, 0xf4, 0x82, 0x63,
	0x46, 0xed, 0x25, 0x45, 0xbe, 0x7b, 0x89, 0xdc, 0xe2, 0xf2, 0xc9, 0x81, 0xe6, 0x56, 0x0a, 0x42,
	0x8b, 0x5a, 0xcf, 0x51, 0x75, 0x76, 0x1f, 0xf6, 0xdf, 0x73, 0xcc, 0x9b, 0x62, 0xe4, 0x0f, 0xc1,
	0x08, 0xf8, 0x33, 0xe3, 0xbd, 0x40, 0xeb, 0x3e, 0x49, 0x01, 0x17, 0xed, 0xaf, 0x2d, 0xd0, 0x7e,
	0x25, 0xa3, 0x98, 0x5d, 0x59, 0x63, 0x74, 0xe3, 0x8a, 0x23, 0xb1, 0xd1, 0x76, 0x69, 0xef, 0xdf,
	0xda, 0x2b, 0x67, 0xde, 0xef, 0xa1, 0x0e, 0xd3, 0x99, 0xd9, 0x74, 0x47, 0xf1, 0x5e, 0xf2, 0x51,
	0x74, 0x39, 0xdb, 0xde, 0xf0, 0x2f, 0xa6, 0xac, 0x09, 0xba, 0x33, 0x10, 0x12, 0x86, 0x98, 0x12,
	0x09, 0x38, 0x85, 0x21, 0x04, 0xd9, 0x71, 0x62, 0x39, 0x8d, 0xc1, 0xfe, 0x47, 0x4d, 0xe1, 0xcd,
	0x75, 0xd7, 0x7f, 0x9e, 0x73, 0x9a, 0x09, 0x34, 0x89, 0x84, 0x4e, 0x4e, 0xff, 0x30, 0x8d, 0xa1,
	0xc5, 0xfb, 0xe2, 0xf4, 0xaf, 0xf6, 0xad, 0xc1, 0x95, 0x55, 0xeb, 0x73, 0x09, 0x6d, 0x69, 0x67,
	0x42, 0xc7, 0x84, 0x07, 0x80, 0x7d, 0x21, 0xce, 0x19, 0x0f, 0xf1, 0x84, 0x71, 0x2a, 0x26, 0xf6,
	0xb2, 0x32, 0x3f, 0x5a, 0xc8, 0xbc, 0xae, 0x25, 0x1a, 0x5a, 0xe1, 0xa3, 0x12, 0x30, 0xfe, 0x7a,
	0xbc, 0xab, 0x00, 0xd6, 0x00, 0xdd, 0xd4, 0x1d, 0x0c, 0x81, 0x87, 0x72, 0x80, 0x45, 0x1f, 0xa7,
	0x92, 0x4c, 0xed, 0x15, 0xe5, 0xfc, 0x78, 0x21, 0xe7, 0x37, 0x8a, 0xfa, 0xbe, 0xdf, 0x91, 0x64,
	0x6a, 0x1c, 0x37, 0x06, 0x17, 0x0b, 0x96, 0x8f, 0x2c, 0xed, 0x14, 0x0c, 0x20, 0x38, 0xc7, 0x8c,
	0x63, 0x4a, 0xa6, 0xf6, 0xaa, 0xf2, 0xd9, 0x5f, 0xc8, 0xe7, 0x38, 0x23, 0xb6, 0x78, 0xb3, 0x70,
	0xf9, 0x6f, 0xf0, 0x7b, 0xda, 0x6a, 0xa2, 0x65, 0x0a, 0x63, 0x16, 0x80, 0x5d, 0x51, 0xba, 0x0f,
	0xaf, 0xd3, 0x6d, 0x2a, 0xb4, 0x91, 0x33, 0x5c, 0x2b, 0x40, 0x1b, 0x71, 0x02, 0x7d, 0x48, 0x12,
	0xa0, 0x38, 0x10, 0x5c, 0x02, 0x97, 0x76, 0x55, 0x09, 0x1e, 0x5c, 0x27, 0x78, 0x96, 0x13, 0x8f,
	0x35, 0xcf, 0x48, 0xff, 0x1f, 0x5f, 0xc8, 0x37, 0x2a, 0x68, 0xad, 0xb8, 0x7e, 0x8d, 0x2f, 0x65,
	0xb4, 0x1b, 0x88, 0xc8, 0xb9, 0xf6, 0x05, 0x6c, 0xdc, 0xbe, 0xfc, 0xb0, 0x9c, 0x65, 0x17, 0xeb,
	0xac, 0xf4, 0xe9, 0xb5, 0x61, 0x87, 0x62, 0x48, 0x78, 0xe8, 0x88, 0x24, 0x74, 0x43, 0xe0, 0xea,
	0xda, 0xe5, 0xaf, 0x5f, 0xcc, 0xd2, 0x3f, 0xbc, 0xce, 0x4f, 0x8b, 0xe8, 0x6b, 0x79, 0xe9, 0xa4,
	0x5e, 0xff, 0x56, 0xde, 0x39, 0xd1, 0x92, 0x75, 0x9a, 0x3a, 0x3a, 0xcc, 0xa2, 0x6e, 0xcd, 0x69,
	0xe7, 0xc8, 0x1f, 0x39, 0xa6, 0x57, 0xa7, 0x69, 0xaf, 0xc0, 0xf4, 0xba, 0xb5, 0x5e, 0x81, 0xf9,
	0x59, 0xde, 0xd5, 0x05, 0xcf, 0xab, 0xd3, 0xd4, 0xf3, 0x0a, 0x94, 0xe7, 0x75, 0x6b, 0x9e, 0x57,
	0xe0, 0xfc, 0x65, 0xd5, 0xec, 0xfe, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x51, 0x83, 0x53,
	0x49, 0x06, 0x00, 0x00,
}
