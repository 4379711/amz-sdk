package sb_v4

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// SBInsightsKeywordAlertType Keyword alert insights associated with the selected keyword targets and bids. LOW_KEYWORD_TRAFFIC is provided if the keyword has very low traffic and is available in all marketplaces. LOW_BID is provided if the selected bid is low compared to the historical bids for this keyword and is only available in the following marketplaces: US, CA, MX, BR, UK, DE, FR, ES, IT, IN, AE, NL, SE, JP, AU, SG.
type SBInsightsKeywordAlertType string

// List of SBInsightsKeywordAlertType
const (
	SBINSIGHTSKEYWORDALERTTYPE_KEYWORD_TRAFFIC SBInsightsKeywordAlertType = "LOW_KEYWORD_TRAFFIC"
	SBINSIGHTSKEYWORDALERTTYPE_BID             SBInsightsKeywordAlertType = "LOW_BID"
)

// All allowed values of SBInsightsKeywordAlertType enum
var AllowedSBInsightsKeywordAlertTypeEnumValues = []SBInsightsKeywordAlertType{
	"LOW_KEYWORD_TRAFFIC",
	"LOW_BID",
}

func (v *SBInsightsKeywordAlertType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SBInsightsKeywordAlertType(value)
	for _, existing := range AllowedSBInsightsKeywordAlertTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SBInsightsKeywordAlertType", value)
}

// NewSBInsightsKeywordAlertTypeFromValue returns a pointer to a valid SBInsightsKeywordAlertType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSBInsightsKeywordAlertTypeFromValue(v string) (*SBInsightsKeywordAlertType, error) {
	ev := SBInsightsKeywordAlertType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SBInsightsKeywordAlertType: valid values are %v", v, AllowedSBInsightsKeywordAlertTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SBInsightsKeywordAlertType) IsValid() bool {
	for _, existing := range AllowedSBInsightsKeywordAlertTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SBInsightsKeywordAlertType value
func (v SBInsightsKeywordAlertType) Ptr() *SBInsightsKeywordAlertType {
	return &v
}

type NullableSBInsightsKeywordAlertType struct {
	value *SBInsightsKeywordAlertType
	isSet bool
}

func (v NullableSBInsightsKeywordAlertType) Get() *SBInsightsKeywordAlertType {
	return v.value
}

func (v *NullableSBInsightsKeywordAlertType) Set(val *SBInsightsKeywordAlertType) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsKeywordAlertType) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsKeywordAlertType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsKeywordAlertType(val *SBInsightsKeywordAlertType) *NullableSBInsightsKeywordAlertType {
	return &NullableSBInsightsKeywordAlertType{value: val, isSet: true}
}

func (v NullableSBInsightsKeywordAlertType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsKeywordAlertType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
