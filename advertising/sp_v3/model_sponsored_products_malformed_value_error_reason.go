package sp_v3

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SponsoredProductsMalformedValueErrorReason the model 'SponsoredProductsMalformedValueErrorReason'
type SponsoredProductsMalformedValueErrorReason string

// List of SponsoredProductsMalformedValueErrorReason
const (
	SPONSOREDPRODUCTSMALFORMEDVALUEERRORREASON_FORBIDDEN_CHARS                SponsoredProductsMalformedValueErrorReason = "FORBIDDEN_CHARS"
	SPONSOREDPRODUCTSMALFORMEDVALUEERRORREASON_PATTERN_NOT_MATCHED            SponsoredProductsMalformedValueErrorReason = "PATTERN_NOT_MATCHED"
	SPONSOREDPRODUCTSMALFORMEDVALUEERRORREASON_TOO_LONG                       SponsoredProductsMalformedValueErrorReason = "TOO_LONG"
	SPONSOREDPRODUCTSMALFORMEDVALUEERRORREASON_TOO_SHORT                      SponsoredProductsMalformedValueErrorReason = "TOO_SHORT"
	SPONSOREDPRODUCTSMALFORMEDVALUEERRORREASON_LEADING_OR_TRAILING_WHITESPACE SponsoredProductsMalformedValueErrorReason = "LEADING_OR_TRAILING_WHITESPACE"
	SPONSOREDPRODUCTSMALFORMEDVALUEERRORREASON_BLANK                          SponsoredProductsMalformedValueErrorReason = "BLANK"
)

// All allowed values of SponsoredProductsMalformedValueErrorReason enum
var AllowedSponsoredProductsMalformedValueErrorReasonEnumValues = []SponsoredProductsMalformedValueErrorReason{
	"FORBIDDEN_CHARS",
	"PATTERN_NOT_MATCHED",
	"TOO_LONG",
	"TOO_SHORT",
	"LEADING_OR_TRAILING_WHITESPACE",
	"BLANK",
}

func (v *SponsoredProductsMalformedValueErrorReason) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsMalformedValueErrorReason(value)
	for _, existing := range AllowedSponsoredProductsMalformedValueErrorReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsMalformedValueErrorReason", value)
}

// NewSponsoredProductsMalformedValueErrorReasonFromValue returns a pointer to a valid SponsoredProductsMalformedValueErrorReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsMalformedValueErrorReasonFromValue(v string) (*SponsoredProductsMalformedValueErrorReason, error) {
	ev := SponsoredProductsMalformedValueErrorReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsMalformedValueErrorReason: valid values are %v", v, AllowedSponsoredProductsMalformedValueErrorReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsMalformedValueErrorReason) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsMalformedValueErrorReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsMalformedValueErrorReason value
func (v SponsoredProductsMalformedValueErrorReason) Ptr() *SponsoredProductsMalformedValueErrorReason {
	return &v
}

type NullableSponsoredProductsMalformedValueErrorReason struct {
	value *SponsoredProductsMalformedValueErrorReason
	isSet bool
}

func (v NullableSponsoredProductsMalformedValueErrorReason) Get() *SponsoredProductsMalformedValueErrorReason {
	return v.value
}

func (v *NullableSponsoredProductsMalformedValueErrorReason) Set(val *SponsoredProductsMalformedValueErrorReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMalformedValueErrorReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMalformedValueErrorReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMalformedValueErrorReason(val *SponsoredProductsMalformedValueErrorReason) *NullableSponsoredProductsMalformedValueErrorReason {
	return &NullableSponsoredProductsMalformedValueErrorReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsMalformedValueErrorReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMalformedValueErrorReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
