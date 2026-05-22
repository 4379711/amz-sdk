package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsInternalErrorErrorCode the model 'SponsoredProductsInternalErrorErrorCode'
type SponsoredProductsInternalErrorErrorCode string

// List of SponsoredProductsInternalErrorErrorCode
const (
	SPONSOREDPRODUCTSINTERNALERRORERRORCODE_INTERNAL_ERROR SponsoredProductsInternalErrorErrorCode = "INTERNAL_ERROR"
)

// All allowed values of SponsoredProductsInternalErrorErrorCode enum
var AllowedSponsoredProductsInternalErrorErrorCodeEnumValues = []SponsoredProductsInternalErrorErrorCode{
	"INTERNAL_ERROR",
}

func (v *SponsoredProductsInternalErrorErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsInternalErrorErrorCode(value)
	for _, existing := range AllowedSponsoredProductsInternalErrorErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsInternalErrorErrorCode", value)
}

// NewSponsoredProductsInternalErrorErrorCodeFromValue returns a pointer to a valid SponsoredProductsInternalErrorErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsInternalErrorErrorCodeFromValue(v string) (*SponsoredProductsInternalErrorErrorCode, error) {
	ev := SponsoredProductsInternalErrorErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsInternalErrorErrorCode: valid values are %v", v, AllowedSponsoredProductsInternalErrorErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsInternalErrorErrorCode) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsInternalErrorErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsInternalErrorErrorCode value
func (v SponsoredProductsInternalErrorErrorCode) Ptr() *SponsoredProductsInternalErrorErrorCode {
	return &v
}

type NullableSponsoredProductsInternalErrorErrorCode struct {
	value *SponsoredProductsInternalErrorErrorCode
	isSet bool
}

func (v NullableSponsoredProductsInternalErrorErrorCode) Get() *SponsoredProductsInternalErrorErrorCode {
	return v.value
}

func (v *NullableSponsoredProductsInternalErrorErrorCode) Set(val *SponsoredProductsInternalErrorErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsInternalErrorErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsInternalErrorErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsInternalErrorErrorCode(val *SponsoredProductsInternalErrorErrorCode) *NullableSponsoredProductsInternalErrorErrorCode {
	return &NullableSponsoredProductsInternalErrorErrorCode{value: val, isSet: true}
}

func (v NullableSponsoredProductsInternalErrorErrorCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsInternalErrorErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
