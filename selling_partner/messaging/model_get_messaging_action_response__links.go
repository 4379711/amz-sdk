package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMessagingActionResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMessagingActionResponseLinks{}

// GetMessagingActionResponseLinks The links response that is associated with the messaging action.
type GetMessagingActionResponseLinks struct {
	Self   LinkObject `json:"self"`
	Schema LinkObject `json:"schema"`
}

type _GetMessagingActionResponseLinks GetMessagingActionResponseLinks

// NewGetMessagingActionResponseLinks instantiates a new GetMessagingActionResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessagingActionResponseLinks(self LinkObject, schema LinkObject) *GetMessagingActionResponseLinks {
	this := GetMessagingActionResponseLinks{}
	this.Self = self
	this.Schema = schema
	return &this
}

// NewGetMessagingActionResponseLinksWithDefaults instantiates a new GetMessagingActionResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessagingActionResponseLinksWithDefaults() *GetMessagingActionResponseLinks {
	this := GetMessagingActionResponseLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *GetMessagingActionResponseLinks) GetSelf() LinkObject {
	if o == nil {
		var ret LinkObject
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *GetMessagingActionResponseLinks) GetSelfOk() (*LinkObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *GetMessagingActionResponseLinks) SetSelf(v LinkObject) {
	o.Self = v
}

// GetSchema returns the Schema field value
func (o *GetMessagingActionResponseLinks) GetSchema() LinkObject {
	if o == nil {
		var ret LinkObject
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *GetMessagingActionResponseLinks) GetSchemaOk() (*LinkObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value
func (o *GetMessagingActionResponseLinks) SetSchema(v LinkObject) {
	o.Schema = v
}

func (o GetMessagingActionResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["schema"] = o.Schema
	return toSerialize, nil
}

type NullableGetMessagingActionResponseLinks struct {
	value *GetMessagingActionResponseLinks
	isSet bool
}

func (v NullableGetMessagingActionResponseLinks) Get() *GetMessagingActionResponseLinks {
	return v.value
}

func (v *NullableGetMessagingActionResponseLinks) Set(val *GetMessagingActionResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessagingActionResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessagingActionResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessagingActionResponseLinks(val *GetMessagingActionResponseLinks) *NullableGetMessagingActionResponseLinks {
	return &NullableGetMessagingActionResponseLinks{value: val, isSet: true}
}

func (v NullableGetMessagingActionResponseLinks) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMessagingActionResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
