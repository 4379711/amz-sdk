package aplus_content_20201101

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// AsinBadge A flag that provides additional information about an ASIN. This is contextual and can change depending on the request that generated it.
type AsinBadge string

// List of AsinBadge
const (
	ASINBADGE_BRAND_NOT_ELIGIBLE    AsinBadge = "BRAND_NOT_ELIGIBLE"
	ASINBADGE_CATALOG_NOT_FOUND     AsinBadge = "CATALOG_NOT_FOUND"
	ASINBADGE_CONTENT_NOT_PUBLISHED AsinBadge = "CONTENT_NOT_PUBLISHED"
	ASINBADGE_CONTENT_PUBLISHED     AsinBadge = "CONTENT_PUBLISHED"
)

// All allowed values of AsinBadge enum
var AllowedAsinBadgeEnumValues = []AsinBadge{
	ASINBADGE_BRAND_NOT_ELIGIBLE,
	ASINBADGE_CATALOG_NOT_FOUND,
	ASINBADGE_CONTENT_NOT_PUBLISHED,
	ASINBADGE_CONTENT_PUBLISHED,
}

func (v *AsinBadge) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = AsinBadge(value)
	return nil
}

// NewAsinBadgeFromValue returns a pointer to a valid AsinBadge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAsinBadgeFromValue(v string) (*AsinBadge, error) {
	ev := AsinBadge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AsinBadge: valid values are %v", v, AllowedAsinBadgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AsinBadge) IsValid() bool {
	for _, existing := range AllowedAsinBadgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AsinBadge value
func (v AsinBadge) Ptr() *AsinBadge {
	return &v
}

type NullableAsinBadge struct {
	value *AsinBadge
	isSet bool
}

func (v NullableAsinBadge) Get() *AsinBadge {
	return v.value
}

func (v *NullableAsinBadge) Set(val *AsinBadge) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinBadge) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinBadge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinBadge(val *AsinBadge) *NullableAsinBadge {
	return &NullableAsinBadge{value: val, isSet: true}
}

func (v NullableAsinBadge) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinBadge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
