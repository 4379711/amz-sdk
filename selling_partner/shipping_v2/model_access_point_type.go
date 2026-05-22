package shipping_v2

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// AccessPointType The type of access point, like counter (HELIX), lockers, etc.
type AccessPointType string

// List of AccessPointType
const (
	ACCESSPOINTTYPE_HELIX         AccessPointType = "HELIX"
	ACCESSPOINTTYPE_CAMPUS_LOCKER AccessPointType = "CAMPUS_LOCKER"
	ACCESSPOINTTYPE_OMNI_LOCKER   AccessPointType = "OMNI_LOCKER"
	ACCESSPOINTTYPE_ODIN_LOCKER   AccessPointType = "ODIN_LOCKER"
	ACCESSPOINTTYPE_DOBBY_LOCKER  AccessPointType = "DOBBY_LOCKER"
	ACCESSPOINTTYPE_CORE_LOCKER   AccessPointType = "CORE_LOCKER"
	ACCESSPOINTTYPE__3_P          AccessPointType = "3P"
	ACCESSPOINTTYPE_CAMPUS_ROOM   AccessPointType = "CAMPUS_ROOM"
)

// All allowed values of AccessPointType enum
var AllowedAccessPointTypeEnumValues = []AccessPointType{
	ACCESSPOINTTYPE_HELIX,
	ACCESSPOINTTYPE_CAMPUS_LOCKER,
	ACCESSPOINTTYPE_OMNI_LOCKER,
	ACCESSPOINTTYPE_ODIN_LOCKER,
	ACCESSPOINTTYPE_DOBBY_LOCKER,
	ACCESSPOINTTYPE_CORE_LOCKER,
	ACCESSPOINTTYPE__3_P,
	ACCESSPOINTTYPE_CAMPUS_ROOM,
}

func (v *AccessPointType) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessPointType(value)
	for _, existing := range AllowedAccessPointTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessPointType", value)
}

// NewAccessPointTypeFromValue returns a pointer to a valid AccessPointType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessPointTypeFromValue(v string) (*AccessPointType, error) {
	ev := AccessPointType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessPointType: valid values are %v", v, AllowedAccessPointTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessPointType) IsValid() bool {
	for _, existing := range AllowedAccessPointTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessPointType value
func (v AccessPointType) Ptr() *AccessPointType {
	return &v
}

type NullableAccessPointType struct {
	value *AccessPointType
	isSet bool
}

func (v NullableAccessPointType) Get() *AccessPointType {
	return v.value
}

func (v *NullableAccessPointType) Set(val *AccessPointType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPointType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPointType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPointType(val *AccessPointType) *NullableAccessPointType {
	return &NullableAccessPointType{value: val, isSet: true}
}

func (v NullableAccessPointType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccessPointType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
