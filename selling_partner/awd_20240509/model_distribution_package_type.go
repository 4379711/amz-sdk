package awd_20240509

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// DistributionPackageType Type of distribution packages.
type DistributionPackageType string

// List of DistributionPackageType
const (
	DISTRIBUTIONPACKAGETYPE_CASE   DistributionPackageType = "CASE"
	DISTRIBUTIONPACKAGETYPE_PALLET DistributionPackageType = "PALLET"
)

// All allowed values of DistributionPackageType enum
var AllowedDistributionPackageTypeEnumValues = []DistributionPackageType{
	DISTRIBUTIONPACKAGETYPE_CASE,
	DISTRIBUTIONPACKAGETYPE_PALLET,
}

func (v *DistributionPackageType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	// 未知枚举值保留原值而非报错:Amazon 会在不升级 API 版本的情况下新增枚举值,
	// 严格校验会让含新值的整页响应反序列化失败。需要校验时用 IsValid()。
	*v = DistributionPackageType(value)
	return nil
}

// NewDistributionPackageTypeFromValue returns a pointer to a valid DistributionPackageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDistributionPackageTypeFromValue(v string) (*DistributionPackageType, error) {
	ev := DistributionPackageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DistributionPackageType: valid values are %v", v, AllowedDistributionPackageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DistributionPackageType) IsValid() bool {
	for _, existing := range AllowedDistributionPackageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionPackageType value
func (v DistributionPackageType) Ptr() *DistributionPackageType {
	return &v
}

type NullableDistributionPackageType struct {
	value *DistributionPackageType
	isSet bool
}

func (v NullableDistributionPackageType) Get() *DistributionPackageType {
	return v.value
}

func (v *NullableDistributionPackageType) Set(val *DistributionPackageType) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionPackageType) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionPackageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionPackageType(val *DistributionPackageType) *NullableDistributionPackageType {
	return &NullableDistributionPackageType{value: val, isSet: true}
}

func (v NullableDistributionPackageType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDistributionPackageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
