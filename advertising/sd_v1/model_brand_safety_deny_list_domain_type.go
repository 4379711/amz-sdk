package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// BrandSafetyDenyListDomainType The domain type.
type BrandSafetyDenyListDomainType string

// List of BrandSafetyDenyListDomainType
const (
	BRANDSAFETYDENYLISTDOMAINTYPE_WEBSITE BrandSafetyDenyListDomainType = "WEBSITE"
	BRANDSAFETYDENYLISTDOMAINTYPE_APP     BrandSafetyDenyListDomainType = "APP"
)

// All allowed values of BrandSafetyDenyListDomainType enum
var AllowedBrandSafetyDenyListDomainTypeEnumValues = []BrandSafetyDenyListDomainType{
	"WEBSITE",
	"APP",
}

func (v *BrandSafetyDenyListDomainType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BrandSafetyDenyListDomainType(value)
	for _, existing := range AllowedBrandSafetyDenyListDomainTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BrandSafetyDenyListDomainType", value)
}

// NewBrandSafetyDenyListDomainTypeFromValue returns a pointer to a valid BrandSafetyDenyListDomainType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandSafetyDenyListDomainTypeFromValue(v string) (*BrandSafetyDenyListDomainType, error) {
	ev := BrandSafetyDenyListDomainType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandSafetyDenyListDomainType: valid values are %v", v, AllowedBrandSafetyDenyListDomainTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandSafetyDenyListDomainType) IsValid() bool {
	for _, existing := range AllowedBrandSafetyDenyListDomainTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BrandSafetyDenyListDomainType value
func (v BrandSafetyDenyListDomainType) Ptr() *BrandSafetyDenyListDomainType {
	return &v
}

type NullableBrandSafetyDenyListDomainType struct {
	value *BrandSafetyDenyListDomainType
	isSet bool
}

func (v NullableBrandSafetyDenyListDomainType) Get() *BrandSafetyDenyListDomainType {
	return v.value
}

func (v *NullableBrandSafetyDenyListDomainType) Set(val *BrandSafetyDenyListDomainType) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyDenyListDomainType) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyDenyListDomainType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyDenyListDomainType(val *BrandSafetyDenyListDomainType) *NullableBrandSafetyDenyListDomainType {
	return &NullableBrandSafetyDenyListDomainType{value: val, isSet: true}
}

func (v NullableBrandSafetyDenyListDomainType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyDenyListDomainType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
