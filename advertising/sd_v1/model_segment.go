package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// Segment Optional. A dimension used to further segment certain types of reports.  Note: matchedTarget reports only return targets that have generated at least one click. | Dimension | Report types | Tactics | Metrics | Description | |---------|------------------|-------------|-------------|------------| | matchedTarget | campaigns, adGroups, targets | T00020, T00030 | Existing metrics for each report type are accepted. |  Segments a report based on the ASIN of the product page where the ad appeared.|
type Segment string

// List of Segment
const (
	SEGMENT_MATCHED_TARGET Segment = "matchedTarget"
)

// All allowed values of Segment enum
var AllowedSegmentEnumValues = []Segment{
	"matchedTarget",
}

func (v *Segment) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Segment(value)
	for _, existing := range AllowedSegmentEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Segment", value)
}

// NewSegmentFromValue returns a pointer to a valid Segment
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSegmentFromValue(v string) (*Segment, error) {
	ev := Segment(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Segment: valid values are %v", v, AllowedSegmentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Segment) IsValid() bool {
	for _, existing := range AllowedSegmentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Segment value
func (v Segment) Ptr() *Segment {
	return &v
}

type NullableSegment struct {
	value *Segment
	isSet bool
}

func (v NullableSegment) Get() *Segment {
	return v.value
}

func (v *NullableSegment) Set(val *Segment) {
	v.value = val
	v.isSet = true
}

func (v NullableSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegment(val *Segment) *NullableSegment {
	return &NullableSegment{value: val, isSet: true}
}

func (v NullableSegment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
