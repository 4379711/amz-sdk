package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMessagingActionsForOrderResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMessagingActionsForOrderResponseLinks{}

// GetMessagingActionsForOrderResponseLinks The links response that is associated with the specified `amazonOrderId`.
type GetMessagingActionsForOrderResponseLinks struct {
	Self LinkObject `json:"self"`
	// Eligible actions for the specified amazonOrderId.
	Actions []LinkObject `json:"actions"`
}

type _GetMessagingActionsForOrderResponseLinks GetMessagingActionsForOrderResponseLinks

// NewGetMessagingActionsForOrderResponseLinks instantiates a new GetMessagingActionsForOrderResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessagingActionsForOrderResponseLinks(self LinkObject, actions []LinkObject) *GetMessagingActionsForOrderResponseLinks {
	this := GetMessagingActionsForOrderResponseLinks{}
	this.Self = self
	this.Actions = actions
	return &this
}

// NewGetMessagingActionsForOrderResponseLinksWithDefaults instantiates a new GetMessagingActionsForOrderResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessagingActionsForOrderResponseLinksWithDefaults() *GetMessagingActionsForOrderResponseLinks {
	this := GetMessagingActionsForOrderResponseLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *GetMessagingActionsForOrderResponseLinks) GetSelf() LinkObject {
	if o == nil {
		var ret LinkObject
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *GetMessagingActionsForOrderResponseLinks) GetSelfOk() (*LinkObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *GetMessagingActionsForOrderResponseLinks) SetSelf(v LinkObject) {
	o.Self = v
}

// GetActions returns the Actions field value
func (o *GetMessagingActionsForOrderResponseLinks) GetActions() []LinkObject {
	if o == nil {
		var ret []LinkObject
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *GetMessagingActionsForOrderResponseLinks) GetActionsOk() ([]LinkObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *GetMessagingActionsForOrderResponseLinks) SetActions(v []LinkObject) {
	o.Actions = v
}

func (o GetMessagingActionsForOrderResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

type NullableGetMessagingActionsForOrderResponseLinks struct {
	value *GetMessagingActionsForOrderResponseLinks
	isSet bool
}

func (v NullableGetMessagingActionsForOrderResponseLinks) Get() *GetMessagingActionsForOrderResponseLinks {
	return v.value
}

func (v *NullableGetMessagingActionsForOrderResponseLinks) Set(val *GetMessagingActionsForOrderResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessagingActionsForOrderResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessagingActionsForOrderResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessagingActionsForOrderResponseLinks(val *GetMessagingActionsForOrderResponseLinks) *NullableGetMessagingActionsForOrderResponseLinks {
	return &NullableGetMessagingActionsForOrderResponseLinks{value: val, isSet: true}
}

func (v NullableGetMessagingActionsForOrderResponseLinks) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMessagingActionsForOrderResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
