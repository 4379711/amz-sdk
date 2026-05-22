package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMessagingActionsForOrderResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMessagingActionsForOrderResponseEmbedded{}

// GetMessagingActionsForOrderResponseEmbedded The messaging actions response that is associated with the specified `amazonOrderId`.
type GetMessagingActionsForOrderResponseEmbedded struct {
	Actions []GetMessagingActionResponse `json:"actions"`
}

type _GetMessagingActionsForOrderResponseEmbedded GetMessagingActionsForOrderResponseEmbedded

// NewGetMessagingActionsForOrderResponseEmbedded instantiates a new GetMessagingActionsForOrderResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessagingActionsForOrderResponseEmbedded(actions []GetMessagingActionResponse) *GetMessagingActionsForOrderResponseEmbedded {
	this := GetMessagingActionsForOrderResponseEmbedded{}
	this.Actions = actions
	return &this
}

// NewGetMessagingActionsForOrderResponseEmbeddedWithDefaults instantiates a new GetMessagingActionsForOrderResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessagingActionsForOrderResponseEmbeddedWithDefaults() *GetMessagingActionsForOrderResponseEmbedded {
	this := GetMessagingActionsForOrderResponseEmbedded{}
	return &this
}

// GetActions returns the Actions field value
func (o *GetMessagingActionsForOrderResponseEmbedded) GetActions() []GetMessagingActionResponse {
	if o == nil {
		var ret []GetMessagingActionResponse
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *GetMessagingActionsForOrderResponseEmbedded) GetActionsOk() ([]GetMessagingActionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *GetMessagingActionsForOrderResponseEmbedded) SetActions(v []GetMessagingActionResponse) {
	o.Actions = v
}

func (o GetMessagingActionsForOrderResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

type NullableGetMessagingActionsForOrderResponseEmbedded struct {
	value *GetMessagingActionsForOrderResponseEmbedded
	isSet bool
}

func (v NullableGetMessagingActionsForOrderResponseEmbedded) Get() *GetMessagingActionsForOrderResponseEmbedded {
	return v.value
}

func (v *NullableGetMessagingActionsForOrderResponseEmbedded) Set(val *GetMessagingActionsForOrderResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessagingActionsForOrderResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessagingActionsForOrderResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessagingActionsForOrderResponseEmbedded(val *GetMessagingActionsForOrderResponseEmbedded) *NullableGetMessagingActionsForOrderResponseEmbedded {
	return &NullableGetMessagingActionsForOrderResponseEmbedded{value: val, isSet: true}
}

func (v NullableGetMessagingActionsForOrderResponseEmbedded) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMessagingActionsForOrderResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
