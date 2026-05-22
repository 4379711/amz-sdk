package fulfillment_outbound_20200701

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// ReturnItemDisposition The condition of the return item when received by an Amazon fulfillment center.
type ReturnItemDisposition string

// List of ReturnItemDisposition
const (
	RETURNITEMDISPOSITION_SELLABLE          ReturnItemDisposition = "Sellable"
	RETURNITEMDISPOSITION_DEFECTIVE         ReturnItemDisposition = "Defective"
	RETURNITEMDISPOSITION_CUSTOMER_DAMAGED  ReturnItemDisposition = "CustomerDamaged"
	RETURNITEMDISPOSITION_CARRIER_DAMAGED   ReturnItemDisposition = "CarrierDamaged"
	RETURNITEMDISPOSITION_FULFILLER_DAMAGED ReturnItemDisposition = "FulfillerDamaged"
)

// All allowed values of ReturnItemDisposition enum
var AllowedReturnItemDispositionEnumValues = []ReturnItemDisposition{
	RETURNITEMDISPOSITION_SELLABLE,
	RETURNITEMDISPOSITION_DEFECTIVE,
	RETURNITEMDISPOSITION_CUSTOMER_DAMAGED,
	RETURNITEMDISPOSITION_CARRIER_DAMAGED,
	RETURNITEMDISPOSITION_FULFILLER_DAMAGED,
}

func (v *ReturnItemDisposition) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnItemDisposition(value)
	for _, existing := range AllowedReturnItemDispositionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnItemDisposition", value)
}

// NewReturnItemDispositionFromValue returns a pointer to a valid ReturnItemDisposition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReturnItemDispositionFromValue(v string) (*ReturnItemDisposition, error) {
	ev := ReturnItemDisposition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReturnItemDisposition: valid values are %v", v, AllowedReturnItemDispositionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReturnItemDisposition) IsValid() bool {
	for _, existing := range AllowedReturnItemDispositionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnItemDisposition value
func (v ReturnItemDisposition) Ptr() *ReturnItemDisposition {
	return &v
}

type NullableReturnItemDisposition struct {
	value *ReturnItemDisposition
	isSet bool
}

func (v NullableReturnItemDisposition) Get() *ReturnItemDisposition {
	return v.value
}

func (v *NullableReturnItemDisposition) Set(val *ReturnItemDisposition) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnItemDisposition) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnItemDisposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnItemDisposition(val *ReturnItemDisposition) *NullableReturnItemDisposition {
	return &NullableReturnItemDisposition{value: val, isSet: true}
}

func (v NullableReturnItemDisposition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReturnItemDisposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
