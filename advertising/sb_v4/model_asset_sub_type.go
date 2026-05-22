package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// AssetSubType Asset subtype from Asset Library which you are looking to get recommendations for. Asset Library documentation can be found here: https://advertising.amazon.com/API/docs/en-us/creative-asset-library
type AssetSubType string

// List of AssetSubType
const (
	ASSETSUBTYPE_CUSTOM_IMAGE  AssetSubType = "CUSTOM_IMAGE"
	ASSETSUBTYPE_LOGO          AssetSubType = "LOGO"
	ASSETSUBTYPE_PRODUCT_IMAGE AssetSubType = "PRODUCT_IMAGE"
	ASSETSUBTYPE_AUTHOR_IMAGE  AssetSubType = "AUTHOR_IMAGE"
)

// All allowed values of AssetSubType enum
var AllowedAssetSubTypeEnumValues = []AssetSubType{
	"CUSTOM_IMAGE",
	"LOGO",
	"PRODUCT_IMAGE",
	"AUTHOR_IMAGE",
}

func (v *AssetSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssetSubType(value)
	for _, existing := range AllowedAssetSubTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssetSubType", value)
}

// NewAssetSubTypeFromValue returns a pointer to a valid AssetSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetSubTypeFromValue(v string) (*AssetSubType, error) {
	ev := AssetSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetSubType: valid values are %v", v, AllowedAssetSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetSubType) IsValid() bool {
	for _, existing := range AllowedAssetSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetSubType value
func (v AssetSubType) Ptr() *AssetSubType {
	return &v
}

type NullableAssetSubType struct {
	value *AssetSubType
	isSet bool
}

func (v NullableAssetSubType) Get() *AssetSubType {
	return v.value
}

func (v *NullableAssetSubType) Set(val *AssetSubType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSubType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSubType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSubType(val *AssetSubType) *NullableAssetSubType {
	return &NullableAssetSubType{value: val, isSet: true}
}

func (v NullableAssetSubType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssetSubType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
