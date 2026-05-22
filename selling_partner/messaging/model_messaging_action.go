package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the MessagingAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessagingAction{}

// MessagingAction A simple object containing the name of the template.
type MessagingAction struct {
	// The name of the template.
	Name string `json:"name"`
}

type _MessagingAction MessagingAction

// NewMessagingAction instantiates a new MessagingAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagingAction(name string) *MessagingAction {
	this := MessagingAction{}
	this.Name = name
	return &this
}

// NewMessagingActionWithDefaults instantiates a new MessagingAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagingActionWithDefaults() *MessagingAction {
	this := MessagingAction{}
	return &this
}

// GetName returns the Name field value
func (o *MessagingAction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MessagingAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MessagingAction) SetName(v string) {
	o.Name = v
}

func (o MessagingAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableMessagingAction struct {
	value *MessagingAction
	isSet bool
}

func (v NullableMessagingAction) Get() *MessagingAction {
	return v.value
}

func (v *NullableMessagingAction) Set(val *MessagingAction) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagingAction) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagingAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagingAction(val *MessagingAction) *NullableMessagingAction {
	return &NullableMessagingAction{value: val, isSet: true}
}

func (v NullableMessagingAction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMessagingAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
