package seller_wallet_20240301

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// RecurringFrequency The frequency at which the transaction is repeated.
type RecurringFrequency string

// List of RecurringFrequency
const (
	RECURRINGFREQUENCY_BIWEEKLY RecurringFrequency = "BIWEEKLY"
	RECURRINGFREQUENCY_DAILY    RecurringFrequency = "DAILY"
	RECURRINGFREQUENCY_MONTHLY  RecurringFrequency = "MONTHLY"
	RECURRINGFREQUENCY_WEEKLY   RecurringFrequency = "WEEKLY"
)

// All allowed values of RecurringFrequency enum
var AllowedRecurringFrequencyEnumValues = []RecurringFrequency{
	RECURRINGFREQUENCY_BIWEEKLY,
	RECURRINGFREQUENCY_DAILY,
	RECURRINGFREQUENCY_MONTHLY,
	RECURRINGFREQUENCY_WEEKLY,
}

func (v *RecurringFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecurringFrequency(value)
	for _, existing := range AllowedRecurringFrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecurringFrequency", value)
}

// NewRecurringFrequencyFromValue returns a pointer to a valid RecurringFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecurringFrequencyFromValue(v string) (*RecurringFrequency, error) {
	ev := RecurringFrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecurringFrequency: valid values are %v", v, AllowedRecurringFrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecurringFrequency) IsValid() bool {
	for _, existing := range AllowedRecurringFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecurringFrequency value
func (v RecurringFrequency) Ptr() *RecurringFrequency {
	return &v
}

type NullableRecurringFrequency struct {
	value *RecurringFrequency
	isSet bool
}

func (v NullableRecurringFrequency) Get() *RecurringFrequency {
	return v.value
}

func (v *NullableRecurringFrequency) Set(val *RecurringFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringFrequency(val *RecurringFrequency) *NullableRecurringFrequency {
	return &NullableRecurringFrequency{value: val, isSet: true}
}

func (v NullableRecurringFrequency) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRecurringFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
