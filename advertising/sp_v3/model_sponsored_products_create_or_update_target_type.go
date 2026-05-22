package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsCreateOrUpdateTargetType the model 'SponsoredProductsCreateOrUpdateTargetType'
type SponsoredProductsCreateOrUpdateTargetType string

// List of SponsoredProductsCreateOrUpdateTargetType
const (
	SPONSOREDPRODUCTSCREATEORUPDATETARGETTYPE_KEYWORD          SponsoredProductsCreateOrUpdateTargetType = "KEYWORD"
	SPONSOREDPRODUCTSCREATEORUPDATETARGETTYPE_PRODUCT          SponsoredProductsCreateOrUpdateTargetType = "PRODUCT"
	SPONSOREDPRODUCTSCREATEORUPDATETARGETTYPE_PRODUCT_CATEGORY SponsoredProductsCreateOrUpdateTargetType = "PRODUCT_CATEGORY"
)

// All allowed values of SponsoredProductsCreateOrUpdateTargetType enum
var AllowedSponsoredProductsCreateOrUpdateTargetTypeEnumValues = []SponsoredProductsCreateOrUpdateTargetType{
	"KEYWORD",
	"PRODUCT",
	"PRODUCT_CATEGORY",
}

func (v *SponsoredProductsCreateOrUpdateTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsCreateOrUpdateTargetType(value)
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateTargetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsCreateOrUpdateTargetType", value)
}

// NewSponsoredProductsCreateOrUpdateTargetTypeFromValue returns a pointer to a valid SponsoredProductsCreateOrUpdateTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsCreateOrUpdateTargetTypeFromValue(v string) (*SponsoredProductsCreateOrUpdateTargetType, error) {
	ev := SponsoredProductsCreateOrUpdateTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsCreateOrUpdateTargetType: valid values are %v", v, AllowedSponsoredProductsCreateOrUpdateTargetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsCreateOrUpdateTargetType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsCreateOrUpdateTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsCreateOrUpdateTargetType value
func (v SponsoredProductsCreateOrUpdateTargetType) Ptr() *SponsoredProductsCreateOrUpdateTargetType {
	return &v
}

type NullableSponsoredProductsCreateOrUpdateTargetType struct {
	value *SponsoredProductsCreateOrUpdateTargetType
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateTargetType) Get() *SponsoredProductsCreateOrUpdateTargetType {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateTargetType) Set(val *SponsoredProductsCreateOrUpdateTargetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateTargetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateTargetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateTargetType(val *SponsoredProductsCreateOrUpdateTargetType) *NullableSponsoredProductsCreateOrUpdateTargetType {
	return &NullableSponsoredProductsCreateOrUpdateTargetType{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateTargetType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateTargetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
