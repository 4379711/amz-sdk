package sd_v1

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CreativeTypeInCreativeRequest The type of the creative. |Name|Description| |----|-----------| |IMAGE |The creative will display static assets (e.g. headline, brandLogo or custom image).| |VIDEO |The creative will display video assets. This type of creative must have video assets provided. Only supported when using productAds with ASIN or SKU.|
type CreativeTypeInCreativeRequest string

// List of CreativeTypeInCreativeRequest
const (
	CREATIVETYPEINCREATIVEREQUEST_IMAGE CreativeTypeInCreativeRequest = "IMAGE"
	CREATIVETYPEINCREATIVEREQUEST_VIDEO CreativeTypeInCreativeRequest = "VIDEO"
)

// All allowed values of CreativeTypeInCreativeRequest enum
var AllowedCreativeTypeInCreativeRequestEnumValues = []CreativeTypeInCreativeRequest{
	"IMAGE",
	"VIDEO",
}

func (v *CreativeTypeInCreativeRequest) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = CreativeTypeInCreativeRequest(value)
	return nil
}

// NewCreativeTypeInCreativeRequestFromValue returns a pointer to a valid CreativeTypeInCreativeRequest
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeTypeInCreativeRequestFromValue(v string) (*CreativeTypeInCreativeRequest, error) {
	ev := CreativeTypeInCreativeRequest(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeTypeInCreativeRequest: valid values are %v", v, AllowedCreativeTypeInCreativeRequestEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeTypeInCreativeRequest) IsValid() bool {
	for _, existing := range AllowedCreativeTypeInCreativeRequestEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreativeTypeInCreativeRequest value
func (v CreativeTypeInCreativeRequest) Ptr() *CreativeTypeInCreativeRequest {
	return &v
}

type NullableCreativeTypeInCreativeRequest struct {
	value *CreativeTypeInCreativeRequest
	isSet bool
}

func (v NullableCreativeTypeInCreativeRequest) Get() *CreativeTypeInCreativeRequest {
	return v.value
}

func (v *NullableCreativeTypeInCreativeRequest) Set(val *CreativeTypeInCreativeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeTypeInCreativeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeTypeInCreativeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeTypeInCreativeRequest(val *CreativeTypeInCreativeRequest) *NullableCreativeTypeInCreativeRequest {
	return &NullableCreativeTypeInCreativeRequest{value: val, isSet: true}
}

func (v NullableCreativeTypeInCreativeRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeTypeInCreativeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
