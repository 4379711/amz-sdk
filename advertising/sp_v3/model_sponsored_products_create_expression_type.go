package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateExpressionType the model 'SponsoredProductsCreateExpressionType'
type SponsoredProductsCreateExpressionType string

// List of SponsoredProductsCreateExpressionType
const (
	SPONSOREDPRODUCTSCREATEEXPRESSIONTYPE_MANUAL SponsoredProductsCreateExpressionType = "MANUAL"
)

// All allowed values of SponsoredProductsCreateExpressionType enum
var AllowedSponsoredProductsCreateExpressionTypeEnumValues = []SponsoredProductsCreateExpressionType{
	"MANUAL",
}

func (v *SponsoredProductsCreateExpressionType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCreateExpressionType(value)
	for _, existing := range AllowedSponsoredProductsCreateExpressionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCreateExpressionType", value)
}

// NewSponsoredProductsCreateExpressionTypeFromValue returns a pointer to a valid SponsoredProductsCreateExpressionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateExpressionTypeFromValue(v string) (*SponsoredProductsCreateExpressionType, error) {
	ev := SponsoredProductsCreateExpressionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateExpressionType: valid values are %v", v, AllowedSponsoredProductsCreateExpressionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateExpressionType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateExpressionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateExpressionType value
func (v SponsoredProductsCreateExpressionType) Ptr() *SponsoredProductsCreateExpressionType {
	return &v
}

type NullableSponsoredProductsCreateExpressionType struct {
	value *SponsoredProductsCreateExpressionType
	isSet bool
}

func (v NullableSponsoredProductsCreateExpressionType) Get() *SponsoredProductsCreateExpressionType {
	return v.value
}

func (v *NullableSponsoredProductsCreateExpressionType) Set(val *SponsoredProductsCreateExpressionType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateExpressionType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateExpressionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateExpressionType(val *SponsoredProductsCreateExpressionType) *NullableSponsoredProductsCreateExpressionType {
	return &NullableSponsoredProductsCreateExpressionType{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateExpressionType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateExpressionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
