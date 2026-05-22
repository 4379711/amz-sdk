package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsExpressionType the model 'SponsoredProductsExpressionType'
type SponsoredProductsExpressionType string

// List of SponsoredProductsExpressionType
const (
	SPONSOREDPRODUCTSEXPRESSIONTYPE_AUTO   SponsoredProductsExpressionType = "AUTO"
	SPONSOREDPRODUCTSEXPRESSIONTYPE_MANUAL SponsoredProductsExpressionType = "MANUAL"
	SPONSOREDPRODUCTSEXPRESSIONTYPE_OTHER  SponsoredProductsExpressionType = "OTHER"
)

// All allowed values of SponsoredProductsExpressionType enum
var AllowedSponsoredProductsExpressionTypeEnumValues = []SponsoredProductsExpressionType{
	"AUTO",
	"MANUAL",
	"OTHER",
}

func (v *SponsoredProductsExpressionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsExpressionType(value)
	for _, existing := range AllowedSponsoredProductsExpressionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsExpressionType", value)
}

// NewSponsoredProductsExpressionTypeFromValue returns a pointer to a valid SponsoredProductsExpressionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsExpressionTypeFromValue(v string) (*SponsoredProductsExpressionType, error) {
	ev := SponsoredProductsExpressionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsExpressionType: valid values are %v", v, AllowedSponsoredProductsExpressionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsExpressionType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsExpressionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsExpressionType value
func (v SponsoredProductsExpressionType) Ptr() *SponsoredProductsExpressionType {
	return &v
}

type NullableSponsoredProductsExpressionType struct {
	value *SponsoredProductsExpressionType
	isSet bool
}

func (v NullableSponsoredProductsExpressionType) Get() *SponsoredProductsExpressionType {
	return v.value
}

func (v *NullableSponsoredProductsExpressionType) Set(val *SponsoredProductsExpressionType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsExpressionType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsExpressionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsExpressionType(val *SponsoredProductsExpressionType) *NullableSponsoredProductsExpressionType {
	return &NullableSponsoredProductsExpressionType{value: val, isSet: true}
}

func (v NullableSponsoredProductsExpressionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsExpressionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
