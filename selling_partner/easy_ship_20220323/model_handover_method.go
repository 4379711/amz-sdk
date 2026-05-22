package easy_ship_20220323

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// HandoverMethod Identifies the method by which a seller will hand a package over to Amazon Logistics.
type HandoverMethod string

// List of HandoverMethod
const (
	HANDOVERMETHOD_PICKUP  HandoverMethod = "Pickup"
	HANDOVERMETHOD_DROPOFF HandoverMethod = "Dropoff"
)

// All allowed values of HandoverMethod enum
var AllowedHandoverMethodEnumValues = []HandoverMethod{
	HANDOVERMETHOD_PICKUP,
	HANDOVERMETHOD_DROPOFF,
}

func (v *HandoverMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HandoverMethod(value)
	for _, existing := range AllowedHandoverMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HandoverMethod", value)
}

// NewHandoverMethodFromValue returns a pointer to a valid HandoverMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHandoverMethodFromValue(v string) (*HandoverMethod, error) {
	ev := HandoverMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HandoverMethod: valid values are %v", v, AllowedHandoverMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HandoverMethod) IsValid() bool {
	for _, existing := range AllowedHandoverMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HandoverMethod value
func (v HandoverMethod) Ptr() *HandoverMethod {
	return &v
}

type NullableHandoverMethod struct {
	value *HandoverMethod
	isSet bool
}

func (v NullableHandoverMethod) Get() *HandoverMethod {
	return v.value
}

func (v *NullableHandoverMethod) Set(val *HandoverMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableHandoverMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableHandoverMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandoverMethod(val *HandoverMethod) *NullableHandoverMethod {
	return &NullableHandoverMethod{value: val, isSet: true}
}

func (v NullableHandoverMethod) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableHandoverMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
