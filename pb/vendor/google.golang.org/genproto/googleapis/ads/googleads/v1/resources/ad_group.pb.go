// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_group.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// An ad group.
type AdGroup struct {
	// The resource name of the ad group.
	// Ad group resource names have the form:
	//
	// `customers/{customer_id}/adGroups/{ad_group_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the ad group.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the ad group.
	//
	// This field is required and should not be empty when creating new ad
	// groups.
	//
	// It must contain fewer than 255 UTF-8 full-width characters.
	//
	// It must not contain any null (code point 0x0), NL line feed
	// (code point 0xA) or carriage return (code point 0xD) characters.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of the ad group.
	Status enums.AdGroupStatusEnum_AdGroupStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.AdGroupStatusEnum_AdGroupStatus" json:"status,omitempty"`
	// The type of the ad group.
	Type enums.AdGroupTypeEnum_AdGroupType `protobuf:"varint,12,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.AdGroupTypeEnum_AdGroupType" json:"type,omitempty"`
	// The ad rotation mode of the ad group.
	AdRotationMode enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode `protobuf:"varint,22,opt,name=ad_rotation_mode,json=adRotationMode,proto3,enum=google.ads.googleads.v1.enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode" json:"ad_rotation_mode,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,13,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The list of mappings used to substitute custom parameter tags in a
	// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,6,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The campaign to which the ad group belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,10,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The maximum CPC (cost-per-click) bid.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,14,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// The maximum CPM (cost-per-thousand viewable impressions) bid.
	CpmBidMicros *wrappers.Int64Value `protobuf:"bytes,15,opt,name=cpm_bid_micros,json=cpmBidMicros,proto3" json:"cpm_bid_micros,omitempty"`
	// The target CPA (cost-per-acquisition).
	TargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,27,opt,name=target_cpa_micros,json=targetCpaMicros,proto3" json:"target_cpa_micros,omitempty"`
	// The CPV (cost-per-view) bid.
	CpvBidMicros *wrappers.Int64Value `protobuf:"bytes,17,opt,name=cpv_bid_micros,json=cpvBidMicros,proto3" json:"cpv_bid_micros,omitempty"`
	// Average amount in micros that the advertiser is willing to pay for every
	// thousand times the ad is shown.
	TargetCpmMicros *wrappers.Int64Value `protobuf:"bytes,26,opt,name=target_cpm_micros,json=targetCpmMicros,proto3" json:"target_cpm_micros,omitempty"`
	// The target ROAS (return-on-ad-spend) override. If the ad group's campaign
	// bidding strategy is a standard Target ROAS strategy, then this field
	// overrides the target ROAS specified in the campaign's bidding strategy.
	// Otherwise, this value is ignored.
	TargetRoas *wrappers.DoubleValue `protobuf:"bytes,30,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	// The percent cpc bid amount, expressed as a fraction of the advertised price
	// for some good or service. The valid range for the fraction is [0,1) and the
	// value stored here is 1,000,000 * [fraction].
	PercentCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,20,opt,name=percent_cpc_bid_micros,json=percentCpcBidMicros,proto3" json:"percent_cpc_bid_micros,omitempty"`
	// Settings for the Display Campaign Optimizer, initially termed "Explorer".
	ExplorerAutoOptimizerSetting *common.ExplorerAutoOptimizerSetting `protobuf:"bytes,21,opt,name=explorer_auto_optimizer_setting,json=explorerAutoOptimizerSetting,proto3" json:"explorer_auto_optimizer_setting,omitempty"`
	// Allows advertisers to specify a targeting dimension on which to place
	// absolute bids. This is only applicable for campaigns that target only the
	// display network and not search.
	DisplayCustomBidDimension enums.TargetingDimensionEnum_TargetingDimension `protobuf:"varint,23,opt,name=display_custom_bid_dimension,json=displayCustomBidDimension,proto3,enum=google.ads.googleads.v1.enums.TargetingDimensionEnum_TargetingDimension" json:"display_custom_bid_dimension,omitempty"`
	// URL template for appending params to Final URL.
	FinalUrlSuffix *wrappers.StringValue `protobuf:"bytes,24,opt,name=final_url_suffix,json=finalUrlSuffix,proto3" json:"final_url_suffix,omitempty"`
	// Setting for targeting related features.
	TargetingSetting *common.TargetingSetting `protobuf:"bytes,25,opt,name=targeting_setting,json=targetingSetting,proto3" json:"targeting_setting,omitempty"`
	// The effective target CPA (cost-per-acquisition).
	// This field is read-only.
	EffectiveTargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,28,opt,name=effective_target_cpa_micros,json=effectiveTargetCpaMicros,proto3" json:"effective_target_cpa_micros,omitempty"`
	// Source of the effective target CPA.
	// This field is read-only.
	EffectiveTargetCpaSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,29,opt,name=effective_target_cpa_source,json=effectiveTargetCpaSource,proto3,enum=google.ads.googleads.v1.enums.BiddingSourceEnum_BiddingSource" json:"effective_target_cpa_source,omitempty"`
	// The effective target ROAS (return-on-ad-spend).
	// This field is read-only.
	EffectiveTargetRoas *wrappers.DoubleValue `protobuf:"bytes,31,opt,name=effective_target_roas,json=effectiveTargetRoas,proto3" json:"effective_target_roas,omitempty"`
	// Source of the effective target ROAS.
	// This field is read-only.
	EffectiveTargetRoasSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,32,opt,name=effective_target_roas_source,json=effectiveTargetRoasSource,proto3,enum=google.ads.googleads.v1.enums.BiddingSourceEnum_BiddingSource" json:"effective_target_roas_source,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                              `json:"-"`
	XXX_unrecognized          []byte                                `json:"-"`
	XXX_sizecache             int32                                 `json:"-"`
}

func (m *AdGroup) Reset()         { *m = AdGroup{} }
func (m *AdGroup) String() string { return proto.CompactTextString(m) }
func (*AdGroup) ProtoMessage()    {}
func (*AdGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_57dce8114718594f, []int{0}
}

func (m *AdGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroup.Unmarshal(m, b)
}
func (m *AdGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroup.Marshal(b, m, deterministic)
}
func (m *AdGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroup.Merge(m, src)
}
func (m *AdGroup) XXX_Size() int {
	return xxx_messageInfo_AdGroup.Size(m)
}
func (m *AdGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroup.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroup proto.InternalMessageInfo

func (m *AdGroup) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroup) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AdGroup) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AdGroup) GetStatus() enums.AdGroupStatusEnum_AdGroupStatus {
	if m != nil {
		return m.Status
	}
	return enums.AdGroupStatusEnum_UNSPECIFIED
}

func (m *AdGroup) GetType() enums.AdGroupTypeEnum_AdGroupType {
	if m != nil {
		return m.Type
	}
	return enums.AdGroupTypeEnum_UNSPECIFIED
}

func (m *AdGroup) GetAdRotationMode() enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode {
	if m != nil {
		return m.AdRotationMode
	}
	return enums.AdGroupAdRotationModeEnum_UNSPECIFIED
}

func (m *AdGroup) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *AdGroup) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *AdGroup) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *AdGroup) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *AdGroup) GetCpmBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpmBidMicros
	}
	return nil
}

func (m *AdGroup) GetTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpaMicros
	}
	return nil
}

func (m *AdGroup) GetCpvBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpvBidMicros
	}
	return nil
}

func (m *AdGroup) GetTargetCpmMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpmMicros
	}
	return nil
}

func (m *AdGroup) GetTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.TargetRoas
	}
	return nil
}

func (m *AdGroup) GetPercentCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.PercentCpcBidMicros
	}
	return nil
}

func (m *AdGroup) GetExplorerAutoOptimizerSetting() *common.ExplorerAutoOptimizerSetting {
	if m != nil {
		return m.ExplorerAutoOptimizerSetting
	}
	return nil
}

func (m *AdGroup) GetDisplayCustomBidDimension() enums.TargetingDimensionEnum_TargetingDimension {
	if m != nil {
		return m.DisplayCustomBidDimension
	}
	return enums.TargetingDimensionEnum_UNSPECIFIED
}

func (m *AdGroup) GetFinalUrlSuffix() *wrappers.StringValue {
	if m != nil {
		return m.FinalUrlSuffix
	}
	return nil
}

func (m *AdGroup) GetTargetingSetting() *common.TargetingSetting {
	if m != nil {
		return m.TargetingSetting
	}
	return nil
}

func (m *AdGroup) GetEffectiveTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.EffectiveTargetCpaMicros
	}
	return nil
}

func (m *AdGroup) GetEffectiveTargetCpaSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectiveTargetCpaSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func (m *AdGroup) GetEffectiveTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.EffectiveTargetRoas
	}
	return nil
}

func (m *AdGroup) GetEffectiveTargetRoasSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectiveTargetRoasSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroup)(nil), "google.ads.googleads.v1.resources.AdGroup")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_group.proto", fileDescriptor_57dce8114718594f)
}

var fileDescriptor_57dce8114718594f = []byte{
	// 952 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcf, 0x6f, 0x23, 0x35,
	0x14, 0xc7, 0x95, 0xb4, 0x14, 0x70, 0xbb, 0x69, 0x77, 0x4a, 0x97, 0x69, 0x1b, 0x76, 0xb3, 0xa0,
	0x95, 0x2a, 0x21, 0x4d, 0x9a, 0x2e, 0x2c, 0x28, 0xb0, 0x48, 0x49, 0xbb, 0x14, 0x90, 0xb6, 0x44,
	0x93, 0x6e, 0x0e, 0xab, 0xa2, 0x91, 0x33, 0x76, 0x46, 0x16, 0x33, 0x63, 0xcb, 0xf6, 0x84, 0x16,
	0x09, 0x71, 0xe0, 0xc4, 0xbf, 0xc1, 0x91, 0x3f, 0x85, 0x3f, 0x85, 0x2b, 0x27, 0x6e, 0x68, 0xfc,
	0x63, 0x9a, 0xa4, 0x4d, 0x67, 0x90, 0xf6, 0xe6, 0xb1, 0xdf, 0xf7, 0xe3, 0xf7, 0x9e, 0xdf, 0x1b,
	0x1b, 0x1c, 0x46, 0x94, 0x46, 0x31, 0x6e, 0x43, 0x24, 0xda, 0x7a, 0x98, 0x8f, 0xa6, 0x9d, 0x36,
	0xc7, 0x82, 0x66, 0x3c, 0xc4, 0xa2, 0x0d, 0x51, 0x10, 0x71, 0x9a, 0x31, 0x8f, 0x71, 0x2a, 0xa9,
	0xf3, 0x58, 0x9b, 0x79, 0x10, 0x09, 0xaf, 0x50, 0x78, 0xd3, 0x8e, 0x57, 0x28, 0xf6, 0x3e, 0x5d,
	0x06, 0x0d, 0x69, 0x92, 0xd0, 0xb4, 0x1d, 0x66, 0x42, 0xd2, 0x24, 0x60, 0x90, 0xc3, 0x04, 0x4b,
	0xcc, 0x35, 0x79, 0xef, 0xa4, 0x44, 0x86, 0x2f, 0x59, 0x4c, 0x39, 0xe6, 0x01, 0xcc, 0x24, 0x0d,
	0x28, 0x93, 0x24, 0x21, 0x3f, 0x63, 0x1e, 0x08, 0x2c, 0x25, 0x49, 0x23, 0x43, 0x79, 0x56, 0x42,
	0x91, 0x90, 0x47, 0x38, 0xb7, 0x5f, 0xd0, 0x3d, 0x5f, 0xa6, 0xc3, 0x69, 0x96, 0x5c, 0x67, 0x21,
	0x80, 0x28, 0xe0, 0x54, 0x42, 0x49, 0x68, 0x1a, 0x24, 0x14, 0x61, 0x23, 0x7f, 0x5a, 0x51, 0x2e,
	0x24, 0x94, 0x99, 0x30, 0xa2, 0x4e, 0x45, 0x91, 0xbc, 0x62, 0x76, 0x9f, 0xa3, 0xbb, 0x25, 0x63,
	0x82, 0x90, 0x8a, 0x4d, 0x1d, 0x85, 0xd1, 0x7c, 0x76, 0xb7, 0xe6, 0x3a, 0x23, 0x88, 0x24, 0x38,
	0x15, 0x84, 0xa6, 0x46, 0xf8, 0xd0, 0x08, 0xd5, 0xd7, 0x38, 0x9b, 0xb4, 0x7f, 0xe2, 0x90, 0x31,
	0xcc, 0xad, 0xff, 0x4d, 0x0b, 0x66, 0xa4, 0x0d, 0xd3, 0xd4, 0xa4, 0xc5, 0xac, 0x7e, 0xf8, 0xcf,
	0x26, 0x78, 0xbb, 0x87, 0x4e, 0xf3, 0x08, 0x9c, 0x8f, 0xc0, 0x3d, 0x5b, 0x1f, 0x41, 0x0a, 0x13,
	0xec, 0xd6, 0x5a, 0xb5, 0x83, 0x77, 0xfd, 0x0d, 0x3b, 0x79, 0x06, 0x13, 0xec, 0x7c, 0x0c, 0xea,
	0x04, 0xb9, 0x2b, 0xad, 0xda, 0xc1, 0xfa, 0xd1, 0xbe, 0x29, 0x2e, 0xcf, 0xee, 0xed, 0x7d, 0x9b,
	0xca, 0x67, 0x9f, 0x8c, 0x60, 0x9c, 0x61, 0xbf, 0x4e, 0x90, 0x73, 0x08, 0x56, 0x15, 0x68, 0x55,
	0x99, 0x37, 0x6f, 0x98, 0x0f, 0x25, 0x27, 0x69, 0xa4, 0xed, 0x95, 0xa5, 0x33, 0x02, 0x6b, 0x3a,
	0xfb, 0xee, 0x5b, 0xad, 0xda, 0x41, 0xe3, 0xe8, 0x2b, 0x6f, 0x59, 0x29, 0xab, 0xbc, 0x78, 0xc6,
	0xf7, 0xa1, 0xd2, 0xbc, 0x48, 0xb3, 0x64, 0x7e, 0xc6, 0x37, 0x34, 0xe7, 0x0c, 0xac, 0xe6, 0x07,
	0xe4, 0x6e, 0x28, 0x6a, 0xb7, 0x1a, 0xf5, 0xfc, 0x8a, 0xe1, 0x59, 0x66, 0xfe, 0xed, 0x2b, 0x8e,
	0x73, 0x09, 0xb6, 0x16, 0x8b, 0xcc, 0x7d, 0xa0, 0xd8, 0x67, 0xd5, 0xd8, 0x3d, 0xe4, 0x1b, 0xf1,
	0x4b, 0x8a, 0xe6, 0x76, 0x99, 0x5f, 0xf1, 0x1b, 0x70, 0xee, 0xdb, 0x19, 0x80, 0x1d, 0xc9, 0x61,
	0xf8, 0x63, 0x5e, 0x0b, 0x19, 0x8f, 0x03, 0x89, 0x13, 0x16, 0x43, 0x89, 0xdd, 0x7b, 0x15, 0x92,
	0xbc, 0x6d, 0xa5, 0xaf, 0x78, 0x7c, 0x6e, 0x84, 0x4e, 0x08, 0x76, 0x72, 0xd0, 0x62, 0xc7, 0x0b,
	0x77, 0xad, 0xb5, 0x72, 0xb0, 0x7e, 0xd4, 0x5e, 0x1a, 0x90, 0xee, 0x56, 0xef, 0x58, 0x09, 0x07,
	0x56, 0xe7, 0x6f, 0x67, 0x3c, 0x5e, 0x98, 0x13, 0xce, 0xe7, 0xe0, 0x9d, 0x10, 0x26, 0x0c, 0x92,
	0x28, 0x75, 0x41, 0x05, 0x4f, 0x0b, 0x6b, 0xa7, 0x07, 0x1a, 0x21, 0x0b, 0x83, 0x31, 0x41, 0x41,
	0x42, 0x42, 0x4e, 0x85, 0xdb, 0x28, 0xaf, 0xbe, 0x8d, 0x90, 0x85, 0x7d, 0x82, 0x5e, 0x2a, 0x81,
	0x46, 0x24, 0xb3, 0x88, 0xcd, 0x4a, 0x88, 0xe4, 0x1a, 0x71, 0x0a, 0xee, 0xeb, 0x1e, 0x0c, 0x42,
	0x06, 0x2d, 0x65, 0xbf, 0x9c, 0xb2, 0xa9, 0x55, 0xc7, 0x0c, 0xce, 0xfa, 0x32, 0x9d, 0xf5, 0xe5,
	0x7e, 0x25, 0x5f, 0xa6, 0xb7, 0xfb, 0x92, 0x58, 0xca, 0xde, 0xff, 0xf0, 0x25, 0x31, 0xa0, 0xe7,
	0x60, 0xdd, 0x80, 0x38, 0x85, 0xc2, 0x7d, 0xb8, 0xe4, 0x5c, 0x4e, 0x68, 0x36, 0x8e, 0xb1, 0x66,
	0x00, 0x2d, 0xf0, 0x29, 0x14, 0xce, 0x00, 0x3c, 0x60, 0x98, 0x87, 0x38, 0xcd, 0x1d, 0x99, 0x3b,
	0xa1, 0xf7, 0xca, 0x9d, 0xd9, 0x36, 0xd2, 0xe3, 0xd9, 0x83, 0xfa, 0xad, 0x06, 0x1e, 0x95, 0x5c,
	0x21, 0xee, 0x8e, 0x62, 0x7f, 0x59, 0x56, 0x95, 0x2f, 0x0c, 0xa6, 0x97, 0x49, 0xfa, 0xbd, 0x85,
	0x0c, 0x35, 0xc3, 0x6f, 0xe2, 0x3b, 0x56, 0x9d, 0xdf, 0x6b, 0xa0, 0x89, 0x88, 0x60, 0x31, 0xbc,
	0xb2, 0x5d, 0x91, 0xc7, 0x56, 0xfc, 0x79, 0xdd, 0xf7, 0x55, 0xa7, 0x7f, 0x53, 0xd2, 0xe9, 0xe7,
	0xf6, 0x9f, 0x7d, 0x62, 0x85, 0xaa, 0xcd, 0x6f, 0x4e, 0xfb, 0xbb, 0x66, 0x37, 0xdd, 0x35, 0x7d,
	0x82, 0x8a, 0x25, 0xe7, 0x6b, 0xb0, 0x35, 0x21, 0x29, 0x8c, 0x55, 0xaf, 0x8b, 0x6c, 0x32, 0x21,
	0x97, 0xae, 0x5b, 0xa1, 0x7f, 0x1a, 0x4a, 0xf5, 0x8a, 0xc7, 0x43, 0xa5, 0x71, 0x7e, 0xb0, 0x35,
	0x33, 0x73, 0xab, 0xba, 0xbb, 0x0a, 0x74, 0x58, 0x96, 0xca, 0xc2, 0x63, 0x9b, 0xbe, 0x2d, 0xb9,
	0x30, 0xe3, 0xbc, 0x06, 0xfb, 0x78, 0x32, 0xc1, 0xa1, 0x24, 0x53, 0x1c, 0xdc, 0x6c, 0x94, 0x66,
	0x79, 0x3d, 0xb8, 0x85, 0xfe, 0x7c, 0xa1, 0x63, 0x7e, 0x59, 0xc2, 0xd6, 0xb7, 0x92, 0xfb, 0x41,
	0xa5, 0x8b, 0xa2, 0xaf, 0x2f, 0xdd, 0xa1, 0xd2, 0xa8, 0x73, 0x98, 0x9b, 0xb9, 0x6d, 0x7b, 0xbd,
	0x92, 0xff, 0x70, 0x6f, 0x6c, 0xaf, 0xda, 0xe5, 0x51, 0x85, 0x76, 0xd9, 0x5e, 0xc0, 0xaa, 0xbe,
	0xf9, 0x15, 0x34, 0x6f, 0x25, 0xda, 0x88, 0x5a, 0x6f, 0x24, 0xa2, 0xdd, 0x5b, 0xb6, 0xd6, 0x4b,
	0xfd, 0x7f, 0x6b, 0xe0, 0x49, 0x48, 0x13, 0xaf, 0xf4, 0x99, 0xd8, 0xdf, 0x30, 0x97, 0xd2, 0x20,
	0x0f, 0x6e, 0x50, 0x7b, 0xfd, 0x9d, 0x91, 0x44, 0x34, 0x86, 0x69, 0xe4, 0x51, 0x1e, 0xb5, 0x23,
	0x9c, 0xaa, 0xd0, 0xed, 0xb3, 0x85, 0x11, 0x71, 0xc7, 0x53, 0xf5, 0x8b, 0x62, 0xf4, 0x47, 0x7d,
	0xe5, 0xb4, 0xd7, 0xfb, 0xb3, 0xfe, 0xf8, 0x54, 0x23, 0x7b, 0x48, 0x78, 0x7a, 0x98, 0x8f, 0x46,
	0x1d, 0xcf, 0xb7, 0x96, 0x7f, 0x59, 0x9b, 0x8b, 0x1e, 0x12, 0x17, 0x85, 0xcd, 0xc5, 0xa8, 0x73,
	0x51, 0xd8, 0xfc, 0x5d, 0x7f, 0xa2, 0x17, 0xba, 0xdd, 0x1e, 0x12, 0xdd, 0x6e, 0x61, 0xd5, 0xed,
	0x8e, 0x3a, 0xdd, 0x6e, 0x61, 0x37, 0x5e, 0x53, 0xce, 0x3e, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x33, 0x43, 0x6e, 0x63, 0x56, 0x0b, 0x00, 0x00,
}
