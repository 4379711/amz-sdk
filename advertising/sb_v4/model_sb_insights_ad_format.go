package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SBInsightsAdFormat Type of Ad format.
type SBInsightsAdFormat string

// List of SBInsightsAdFormat
const (
	SBINSIGHTSADFORMAT_PRODUCT_COLLECTION SBInsightsAdFormat = "PRODUCT_COLLECTION"
	SBINSIGHTSADFORMAT_STORE_SPOTLIGHT    SBInsightsAdFormat = "STORE_SPOTLIGHT"
	SBINSIGHTSADFORMAT_VIDEO              SBInsightsAdFormat = "VIDEO"
	SBINSIGHTSADFORMAT_BRAND_VIDEO        SBInsightsAdFormat = "BRAND_VIDEO"
)

// All allowed values of SBInsightsAdFormat enum
var AllowedSBInsightsAdFormatEnumValues = []SBInsightsAdFormat{
	"PRODUCT_COLLECTION",
	"STORE_SPOTLIGHT",
	"VIDEO",
	"BRAND_VIDEO",
}

func (v *SBInsightsAdFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SBInsightsAdFormat(value)
	for _, existing := range AllowedSBInsightsAdFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SBInsightsAdFormat", value)
}

// NewSBInsightsAdFormatFromValue returns a pointer to a valid SBInsightsAdFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSBInsightsAdFormatFromValue(v string) (*SBInsightsAdFormat, error) {
	ev := SBInsightsAdFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SBInsightsAdFormat: valid values are %v", v, AllowedSBInsightsAdFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SBInsightsAdFormat) IsValid() bool {
	for _, existing := range AllowedSBInsightsAdFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBInsightsAdFormat value
func (v SBInsightsAdFormat) Ptr() *SBInsightsAdFormat {
	return &v
}

type NullableSBInsightsAdFormat struct {
	value *SBInsightsAdFormat
	isSet bool
}

func (v NullableSBInsightsAdFormat) Get() *SBInsightsAdFormat {
	return v.value
}

func (v *NullableSBInsightsAdFormat) Set(val *SBInsightsAdFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsAdFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsAdFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsAdFormat(val *SBInsightsAdFormat) *NullableSBInsightsAdFormat {
	return &NullableSBInsightsAdFormat{value: val, isSet: true}
}

func (v NullableSBInsightsAdFormat) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsAdFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
