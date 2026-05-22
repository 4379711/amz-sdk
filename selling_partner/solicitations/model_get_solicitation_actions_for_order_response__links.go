package solicitations

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSolicitationActionsForOrderResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolicitationActionsForOrderResponseLinks{}

// GetSolicitationActionsForOrderResponseLinks struct for GetSolicitationActionsForOrderResponseLinks
type GetSolicitationActionsForOrderResponseLinks struct {
	Self LinkObject `json:"self"`
	// Eligible actions for the specified amazonOrderId.
	Actions []LinkObject `json:"actions"`
}

type _GetSolicitationActionsForOrderResponseLinks GetSolicitationActionsForOrderResponseLinks

// NewGetSolicitationActionsForOrderResponseLinks instantiates a new GetSolicitationActionsForOrderResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolicitationActionsForOrderResponseLinks(self LinkObject, actions []LinkObject) *GetSolicitationActionsForOrderResponseLinks {
	this := GetSolicitationActionsForOrderResponseLinks{}
	this.Self = self
	this.Actions = actions
	return &this
}

// NewGetSolicitationActionsForOrderResponseLinksWithDefaults instantiates a new GetSolicitationActionsForOrderResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolicitationActionsForOrderResponseLinksWithDefaults() *GetSolicitationActionsForOrderResponseLinks {
	this := GetSolicitationActionsForOrderResponseLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *GetSolicitationActionsForOrderResponseLinks) GetSelf() LinkObject {
	if o == nil {
		var ret LinkObject
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionsForOrderResponseLinks) GetSelfOk() (*LinkObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *GetSolicitationActionsForOrderResponseLinks) SetSelf(v LinkObject) {
	o.Self = v
}

// GetActions returns the Actions field value
func (o *GetSolicitationActionsForOrderResponseLinks) GetActions() []LinkObject {
	if o == nil {
		var ret []LinkObject
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionsForOrderResponseLinks) GetActionsOk() ([]LinkObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *GetSolicitationActionsForOrderResponseLinks) SetActions(v []LinkObject) {
	o.Actions = v
}

func (o GetSolicitationActionsForOrderResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

type NullableGetSolicitationActionsForOrderResponseLinks struct {
	value *GetSolicitationActionsForOrderResponseLinks
	isSet bool
}

func (v NullableGetSolicitationActionsForOrderResponseLinks) Get() *GetSolicitationActionsForOrderResponseLinks {
	return v.value
}

func (v *NullableGetSolicitationActionsForOrderResponseLinks) Set(val *GetSolicitationActionsForOrderResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolicitationActionsForOrderResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolicitationActionsForOrderResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolicitationActionsForOrderResponseLinks(val *GetSolicitationActionsForOrderResponseLinks) *NullableGetSolicitationActionsForOrderResponseLinks {
	return &NullableGetSolicitationActionsForOrderResponseLinks{value: val, isSet: true}
}

func (v NullableGetSolicitationActionsForOrderResponseLinks) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSolicitationActionsForOrderResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
