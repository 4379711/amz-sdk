package shipping_v2

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// NdrAction The type of NDR action shipper wants to take for a particular shipment.
type NdrAction string

// List of NdrAction
const (
	NDRACTION_RESCHEDULE NdrAction = "RESCHEDULE"
	NDRACTION_REATTEMPT  NdrAction = "REATTEMPT"
	NDRACTION_RTO        NdrAction = "RTO"
)

// All allowed values of NdrAction enum
var AllowedNdrActionEnumValues = []NdrAction{
	NDRACTION_RESCHEDULE,
	NDRACTION_REATTEMPT,
	NDRACTION_RTO,
}

func (v *NdrAction) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NdrAction(value)
	for _, existing := range AllowedNdrActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NdrAction", value)
}

// NewNdrActionFromValue returns a pointer to a valid NdrAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNdrActionFromValue(v string) (*NdrAction, error) {
	ev := NdrAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NdrAction: valid values are %v", v, AllowedNdrActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NdrAction) IsValid() bool {
	for _, existing := range AllowedNdrActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NdrAction value
func (v NdrAction) Ptr() *NdrAction {
	return &v
}

type NullableNdrAction struct {
	value *NdrAction
	isSet bool
}

func (v NullableNdrAction) Get() *NdrAction {
	return v.value
}

func (v *NullableNdrAction) Set(val *NdrAction) {
	v.value = val
	v.isSet = true
}

func (v NullableNdrAction) IsSet() bool {
	return v.isSet
}

func (v *NullableNdrAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNdrAction(val *NdrAction) *NullableNdrAction {
	return &NullableNdrAction{value: val, isSet: true}
}

func (v NullableNdrAction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableNdrAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
