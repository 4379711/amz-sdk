package product_eligibility

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// Action The action to be performed on the resource.
type Action string

// List of Action
const (
	ACTION_ADVERTISE_ASIN_LESS          Action = "advertiseAsinLess"
	ACTION_ASSOCIATE_HIGH_VALUE_ACTIONS Action = "associateHighValueActions"
	ACTION_VIEW_BOOK_METRICS            Action = "viewBookMetrics"
	ACTION_LAUNCH_CAMPAIGNS             Action = "launchCampaigns"
	ACTION_VIEW                         Action = "view"
)

// All allowed values of Action enum
var AllowedActionEnumValues = []Action{
	"advertiseAsinLess",
	"associateHighValueActions",
	"viewBookMetrics",
	"launchCampaigns",
	"view",
}

func (v *Action) UnmarshalJSON(src []byte) error {
	var value string
	err := sonic.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Action(value)
	for _, existing := range AllowedActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Action", value)
}

// NewActionFromValue returns a pointer to a valid Action
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionFromValue(v string) (*Action, error) {
	ev := Action(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Action: valid values are %v", v, AllowedActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Action) IsValid() bool {
	for _, existing := range AllowedActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Action value
func (v Action) Ptr() *Action {
	return &v
}

type NullableAction struct {
	value *Action
	isSet bool
}

func (v NullableAction) Get() *Action {
	return v.value
}

func (v *NullableAction) Set(val *Action) {
	v.value = val
	v.isSet = true
}

func (v NullableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAction(val *Action) *NullableAction {
	return &NullableAction{value: val, isSet: true}
}

func (v NullableAction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
