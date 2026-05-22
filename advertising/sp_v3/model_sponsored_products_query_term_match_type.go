package sp_v3

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// SponsoredProductsQueryTermMatchType Match type for query filters. | Value |  Description | |-----------|------------| | `BROAD_MATCH` | Match if the queried value contains the filter value (substring matching). Note: If queryTermMatchType is set to BROAD_MATCH, only matches for the first query included will be returned. | | `EXACT_MATCH` | Match if the queried value is exactly equivalent to the filter value. |
type SponsoredProductsQueryTermMatchType string

// List of SponsoredProductsQueryTermMatchType
const (
	SPONSOREDPRODUCTSQUERYTERMMATCHTYPE_BROAD_MATCH SponsoredProductsQueryTermMatchType = "BROAD_MATCH"
	SPONSOREDPRODUCTSQUERYTERMMATCHTYPE_EXACT_MATCH SponsoredProductsQueryTermMatchType = "EXACT_MATCH"
)

// All allowed values of SponsoredProductsQueryTermMatchType enum
var AllowedSponsoredProductsQueryTermMatchTypeEnumValues = []SponsoredProductsQueryTermMatchType{
	"BROAD_MATCH",
	"EXACT_MATCH",
}

func (v *SponsoredProductsQueryTermMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SponsoredProductsQueryTermMatchType(value)
	for _, existing := range AllowedSponsoredProductsQueryTermMatchTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SponsoredProductsQueryTermMatchType", value)
}

// NewSponsoredProductsQueryTermMatchTypeFromValue returns a pointer to a valid SponsoredProductsQueryTermMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSponsoredProductsQueryTermMatchTypeFromValue(v string) (*SponsoredProductsQueryTermMatchType, error) {
	ev := SponsoredProductsQueryTermMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SponsoredProductsQueryTermMatchType: valid values are %v", v, AllowedSponsoredProductsQueryTermMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SponsoredProductsQueryTermMatchType) IsValid() bool {
	for _, existing := range AllowedSponsoredProductsQueryTermMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SponsoredProductsQueryTermMatchType value
func (v SponsoredProductsQueryTermMatchType) Ptr() *SponsoredProductsQueryTermMatchType {
	return &v
}

type NullableSponsoredProductsQueryTermMatchType struct {
	value *SponsoredProductsQueryTermMatchType
	isSet bool
}

func (v NullableSponsoredProductsQueryTermMatchType) Get() *SponsoredProductsQueryTermMatchType {
	return v.value
}

func (v *NullableSponsoredProductsQueryTermMatchType) Set(val *SponsoredProductsQueryTermMatchType) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsQueryTermMatchType) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsQueryTermMatchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsQueryTermMatchType(val *SponsoredProductsQueryTermMatchType) *NullableSponsoredProductsQueryTermMatchType {
	return &NullableSponsoredProductsQueryTermMatchType{value: val, isSet: true}
}

func (v NullableSponsoredProductsQueryTermMatchType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsQueryTermMatchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
