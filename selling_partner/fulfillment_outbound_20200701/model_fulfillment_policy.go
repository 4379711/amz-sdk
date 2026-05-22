package fulfillment_outbound_20200701

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// FulfillmentPolicy The `FulfillmentPolicy` value specified when you submitted the `createFulfillmentOrder` operation.
type FulfillmentPolicy string

// List of FulfillmentPolicy
const (
	FULFILLMENTPOLICY_FILL_OR_KILL       FulfillmentPolicy = "FillOrKill"
	FULFILLMENTPOLICY_FILL_ALL           FulfillmentPolicy = "FillAll"
	FULFILLMENTPOLICY_FILL_ALL_AVAILABLE FulfillmentPolicy = "FillAllAvailable"
)

// All allowed values of FulfillmentPolicy enum
var AllowedFulfillmentPolicyEnumValues = []FulfillmentPolicy{
	FULFILLMENTPOLICY_FILL_OR_KILL,
	FULFILLMENTPOLICY_FILL_ALL,
	FULFILLMENTPOLICY_FILL_ALL_AVAILABLE,
}

func (v *FulfillmentPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FulfillmentPolicy(value)
	for _, existing := range AllowedFulfillmentPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FulfillmentPolicy", value)
}

// NewFulfillmentPolicyFromValue returns a pointer to a valid FulfillmentPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFulfillmentPolicyFromValue(v string) (*FulfillmentPolicy, error) {
	ev := FulfillmentPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FulfillmentPolicy: valid values are %v", v, AllowedFulfillmentPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FulfillmentPolicy) IsValid() bool {
	for _, existing := range AllowedFulfillmentPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FulfillmentPolicy value
func (v FulfillmentPolicy) Ptr() *FulfillmentPolicy {
	return &v
}

type NullableFulfillmentPolicy struct {
	value *FulfillmentPolicy
	isSet bool
}

func (v NullableFulfillmentPolicy) Get() *FulfillmentPolicy {
	return v.value
}

func (v *NullableFulfillmentPolicy) Set(val *FulfillmentPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentPolicy(val *FulfillmentPolicy) *NullableFulfillmentPolicy {
	return &NullableFulfillmentPolicy{value: val, isSet: true}
}

func (v NullableFulfillmentPolicy) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
