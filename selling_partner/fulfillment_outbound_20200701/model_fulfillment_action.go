package fulfillment_outbound_20200701

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// FulfillmentAction Specifies whether the fulfillment order should ship now or have an order hold put on it.
type FulfillmentAction string

// List of FulfillmentAction
const (
	FULFILLMENTACTION_SHIP FulfillmentAction = "Ship"
	FULFILLMENTACTION_HOLD FulfillmentAction = "Hold"
)

// All allowed values of FulfillmentAction enum
var AllowedFulfillmentActionEnumValues = []FulfillmentAction{
	FULFILLMENTACTION_SHIP,
	FULFILLMENTACTION_HOLD,
}

func (v *FulfillmentAction) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FulfillmentAction(value)
	for _, existing := range AllowedFulfillmentActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FulfillmentAction", value)
}

// NewFulfillmentActionFromValue returns a pointer to a valid FulfillmentAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFulfillmentActionFromValue(v string) (*FulfillmentAction, error) {
	ev := FulfillmentAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FulfillmentAction: valid values are %v", v, AllowedFulfillmentActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FulfillmentAction) IsValid() bool {
	for _, existing := range AllowedFulfillmentActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FulfillmentAction value
func (v FulfillmentAction) Ptr() *FulfillmentAction {
	return &v
}

type NullableFulfillmentAction struct {
	value *FulfillmentAction
	isSet bool
}

func (v NullableFulfillmentAction) Get() *FulfillmentAction {
	return v.value
}

func (v *NullableFulfillmentAction) Set(val *FulfillmentAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentAction(val *FulfillmentAction) *NullableFulfillmentAction {
	return &NullableFulfillmentAction{value: val, isSet: true}
}

func (v NullableFulfillmentAction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
