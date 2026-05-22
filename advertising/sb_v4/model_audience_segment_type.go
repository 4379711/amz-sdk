package sb_v4

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// AudienceSegmentType The audience segment type is required to specify the type of audience being used to apply bid adjustments. - SPONSORED_ADS_AMC - Common AMC Audience created without a reference to a specific Ad Program.
type AudienceSegmentType string

// List of AudienceSegmentType
const (
	AUDIENCESEGMENTTYPE_SPONSORED_ADS_AMC AudienceSegmentType = "SPONSORED_ADS_AMC"
)

// All allowed values of AudienceSegmentType enum
var AllowedAudienceSegmentTypeEnumValues = []AudienceSegmentType{
	"SPONSORED_ADS_AMC",
}

func (v *AudienceSegmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AudienceSegmentType(value)
	for _, existing := range AllowedAudienceSegmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AudienceSegmentType", value)
}

// NewAudienceSegmentTypeFromValue returns a pointer to a valid AudienceSegmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudienceSegmentTypeFromValue(v string) (*AudienceSegmentType, error) {
	ev := AudienceSegmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudienceSegmentType: valid values are %v", v, AllowedAudienceSegmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudienceSegmentType) IsValid() bool {
	for _, existing := range AllowedAudienceSegmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AudienceSegmentType value
func (v AudienceSegmentType) Ptr() *AudienceSegmentType {
	return &v
}

type NullableAudienceSegmentType struct {
	value *AudienceSegmentType
	isSet bool
}

func (v NullableAudienceSegmentType) Get() *AudienceSegmentType {
	return v.value
}

func (v *NullableAudienceSegmentType) Set(val *AudienceSegmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceSegmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceSegmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceSegmentType(val *AudienceSegmentType) *NullableAudienceSegmentType {
	return &NullableAudienceSegmentType{value: val, isSet: true}
}

func (v NullableAudienceSegmentType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAudienceSegmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
