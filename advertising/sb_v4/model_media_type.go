package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// MediaType Media type for assets from Asset Library
type MediaType string

// List of MediaType
const (
	MEDIATYPE_JPEG MediaType = "image/jpeg"
	MEDIATYPE_PNG  MediaType = "image/png"
	MEDIATYPE_GIF  MediaType = "image/gif"
)

// All allowed values of MediaType enum
var AllowedMediaTypeEnumValues = []MediaType{
	"image/jpeg",
	"image/png",
	"image/gif",
}

func (v *MediaType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = MediaType(value)
	return nil
}

// NewMediaTypeFromValue returns a pointer to a valid MediaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaTypeFromValue(v string) (*MediaType, error) {
	ev := MediaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaType: valid values are %v", v, AllowedMediaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaType) IsValid() bool {
	for _, existing := range AllowedMediaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaType value
func (v MediaType) Ptr() *MediaType {
	return &v
}

type NullableMediaType struct {
	value *MediaType
	isSet bool
}

func (v NullableMediaType) Get() *MediaType {
	return v.value
}

func (v *NullableMediaType) Set(val *MediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaType(val *MediaType) *NullableMediaType {
	return &NullableMediaType{value: val, isSet: true}
}

func (v NullableMediaType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
