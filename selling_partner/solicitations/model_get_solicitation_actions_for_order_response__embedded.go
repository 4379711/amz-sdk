package solicitations

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSolicitationActionsForOrderResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolicitationActionsForOrderResponseEmbedded{}

// GetSolicitationActionsForOrderResponseEmbedded struct for GetSolicitationActionsForOrderResponseEmbedded
type GetSolicitationActionsForOrderResponseEmbedded struct {
	Actions []GetSolicitationActionResponse `json:"actions"`
}

type _GetSolicitationActionsForOrderResponseEmbedded GetSolicitationActionsForOrderResponseEmbedded

// NewGetSolicitationActionsForOrderResponseEmbedded instantiates a new GetSolicitationActionsForOrderResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolicitationActionsForOrderResponseEmbedded(actions []GetSolicitationActionResponse) *GetSolicitationActionsForOrderResponseEmbedded {
	this := GetSolicitationActionsForOrderResponseEmbedded{}
	this.Actions = actions
	return &this
}

// NewGetSolicitationActionsForOrderResponseEmbeddedWithDefaults instantiates a new GetSolicitationActionsForOrderResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolicitationActionsForOrderResponseEmbeddedWithDefaults() *GetSolicitationActionsForOrderResponseEmbedded {
	this := GetSolicitationActionsForOrderResponseEmbedded{}
	return &this
}

// GetActions returns the Actions field value
func (o *GetSolicitationActionsForOrderResponseEmbedded) GetActions() []GetSolicitationActionResponse {
	if o == nil {
		var ret []GetSolicitationActionResponse
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionsForOrderResponseEmbedded) GetActionsOk() ([]GetSolicitationActionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *GetSolicitationActionsForOrderResponseEmbedded) SetActions(v []GetSolicitationActionResponse) {
	o.Actions = v
}

func (o GetSolicitationActionsForOrderResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

type NullableGetSolicitationActionsForOrderResponseEmbedded struct {
	value *GetSolicitationActionsForOrderResponseEmbedded
	isSet bool
}

func (v NullableGetSolicitationActionsForOrderResponseEmbedded) Get() *GetSolicitationActionsForOrderResponseEmbedded {
	return v.value
}

func (v *NullableGetSolicitationActionsForOrderResponseEmbedded) Set(val *GetSolicitationActionsForOrderResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolicitationActionsForOrderResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolicitationActionsForOrderResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolicitationActionsForOrderResponseEmbedded(val *GetSolicitationActionsForOrderResponseEmbedded) *NullableGetSolicitationActionsForOrderResponseEmbedded {
	return &NullableGetSolicitationActionsForOrderResponseEmbedded{value: val, isSet: true}
}

func (v NullableGetSolicitationActionsForOrderResponseEmbedded) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSolicitationActionsForOrderResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
