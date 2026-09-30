package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// AcceptHeader Clients request a specific version of a resource using the Accept request-header field set to the value field of the desired content-type.
type AcceptHeader string

// List of AcceptHeader
const (
	ACCEPTHEADER_SB_AD_CREATIVE_RESOURCE_V4JSON                   AcceptHeader = "application/vnd.sbAdCreativeResource.v4+json"
	ACCEPTHEADER_SB_CREATIVE_IMAGE_RECOMMENDATION_RESOURCE_V4JSON AcceptHeader = "application/vnd.sbCreativeImageRecommendationResource.v4+json"
	ACCEPTHEADER_SB_CREATIVE_RECOMMENDATION_RESOURCE_V4JSON       AcceptHeader = "application/vnd.sbCreativeRecommendationResource.v4+json"
)

// All allowed values of AcceptHeader enum
var AllowedAcceptHeaderEnumValues = []AcceptHeader{
	"application/vnd.sbAdCreativeResource.v4+json",
	"application/vnd.sbCreativeImageRecommendationResource.v4+json",
	"application/vnd.sbCreativeRecommendationResource.v4+json",
}

func (v *AcceptHeader) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = AcceptHeader(value)
	return nil
}

// NewAcceptHeaderFromValue returns a pointer to a valid AcceptHeader
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAcceptHeaderFromValue(v string) (*AcceptHeader, error) {
	ev := AcceptHeader(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AcceptHeader: valid values are %v", v, AllowedAcceptHeaderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AcceptHeader) IsValid() bool {
	for _, existing := range AllowedAcceptHeaderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AcceptHeader value
func (v AcceptHeader) Ptr() *AcceptHeader {
	return &v
}

type NullableAcceptHeader struct {
	value *AcceptHeader
	isSet bool
}

func (v NullableAcceptHeader) Get() *AcceptHeader {
	return v.value
}

func (v *NullableAcceptHeader) Set(val *AcceptHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptHeader(val *AcceptHeader) *NullableAcceptHeader {
	return &NullableAcceptHeader{value: val, isSet: true}
}

func (v NullableAcceptHeader) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAcceptHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
