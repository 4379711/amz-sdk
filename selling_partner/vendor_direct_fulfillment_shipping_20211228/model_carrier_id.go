package vendor_direct_fulfillment_shipping_20211228

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// CarrierId The unique carrier code for the carrier for whom container labels are requested.
type CarrierId string

// List of CarrierId
const (
	CARRIERID_SWA CarrierId = "SWA"
)

// All allowed values of CarrierId enum
var AllowedCarrierIdEnumValues = []CarrierId{
	CARRIERID_SWA,
}

func (v *CarrierId) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CarrierId(value)
	for _, existing := range AllowedCarrierIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CarrierId", value)
}

// NewCarrierIdFromValue returns a pointer to a valid CarrierId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCarrierIdFromValue(v string) (*CarrierId, error) {
	ev := CarrierId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CarrierId: valid values are %v", v, AllowedCarrierIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CarrierId) IsValid() bool {
	for _, existing := range AllowedCarrierIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CarrierId value
func (v CarrierId) Ptr() *CarrierId {
	return &v
}

type NullableCarrierId struct {
	value *CarrierId
	isSet bool
}

func (v NullableCarrierId) Get() *CarrierId {
	return v.value
}

func (v *NullableCarrierId) Set(val *CarrierId) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierId) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierId(val *CarrierId) *NullableCarrierId {
	return &NullableCarrierId{value: val, isSet: true}
}

func (v NullableCarrierId) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCarrierId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
