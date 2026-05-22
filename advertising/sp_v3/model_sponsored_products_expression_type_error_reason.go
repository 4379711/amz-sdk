package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsExpressionTypeErrorReason the model 'SponsoredProductsExpressionTypeErrorReason'
type SponsoredProductsExpressionTypeErrorReason string

// List of SponsoredProductsExpressionTypeErrorReason
const (
	SPONSOREDPRODUCTSEXPRESSIONTYPEERRORREASON_UNSUPPORTED_EXPRESSION_TYPE SponsoredProductsExpressionTypeErrorReason = "UNSUPPORTED_EXPRESSION_TYPE"
)

// All allowed values of SponsoredProductsExpressionTypeErrorReason enum
var AllowedSponsoredProductsExpressionTypeErrorReasonEnumValues = []SponsoredProductsExpressionTypeErrorReason{
	"UNSUPPORTED_EXPRESSION_TYPE",
}

func (v *SponsoredProductsExpressionTypeErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsExpressionTypeErrorReason(value)
	for _, existing := range AllowedSponsoredProductsExpressionTypeErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsExpressionTypeErrorReason", value)
}

// NewSponsoredProductsExpressionTypeErrorReasonFromValue returns a pointer to a valid SponsoredProductsExpressionTypeErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsExpressionTypeErrorReasonFromValue(v string) (*SponsoredProductsExpressionTypeErrorReason, error) {
	ev := SponsoredProductsExpressionTypeErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsExpressionTypeErrorReason: valid values are %v", v, AllowedSponsoredProductsExpressionTypeErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsExpressionTypeErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsExpressionTypeErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsExpressionTypeErrorReason value
func (v SponsoredProductsExpressionTypeErrorReason) Ptr() *SponsoredProductsExpressionTypeErrorReason {
	return &v
}

type NullableSponsoredProductsExpressionTypeErrorReason struct {
	value *SponsoredProductsExpressionTypeErrorReason
	isSet bool
}

func (v NullableSponsoredProductsExpressionTypeErrorReason) Get() *SponsoredProductsExpressionTypeErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsExpressionTypeErrorReason) Set(val *SponsoredProductsExpressionTypeErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsExpressionTypeErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsExpressionTypeErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsExpressionTypeErrorReason(val *SponsoredProductsExpressionTypeErrorReason) *NullableSponsoredProductsExpressionTypeErrorReason {
	return &NullableSponsoredProductsExpressionTypeErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsExpressionTypeErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsExpressionTypeErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
