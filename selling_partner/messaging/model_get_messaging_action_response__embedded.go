package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMessagingActionResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMessagingActionResponseEmbedded{}

// GetMessagingActionResponseEmbedded The embedded response associated with the messaging action.
type GetMessagingActionResponseEmbedded struct {
	Schema *GetSchemaResponse `json:"schema,omitempty"`
}

// NewGetMessagingActionResponseEmbedded instantiates a new GetMessagingActionResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessagingActionResponseEmbedded() *GetMessagingActionResponseEmbedded {
	this := GetMessagingActionResponseEmbedded{}
	return &this
}

// NewGetMessagingActionResponseEmbeddedWithDefaults instantiates a new GetMessagingActionResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessagingActionResponseEmbeddedWithDefaults() *GetMessagingActionResponseEmbedded {
	this := GetMessagingActionResponseEmbedded{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *GetMessagingActionResponseEmbedded) GetSchema() GetSchemaResponse {
	if o == nil || IsNil(o.Schema) {
		var ret GetSchemaResponse
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessagingActionResponseEmbedded) GetSchemaOk() (*GetSchemaResponse, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *GetMessagingActionResponseEmbedded) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given GetSchemaResponse and assigns it to the Schema field.
func (o *GetMessagingActionResponseEmbedded) SetSchema(v GetSchemaResponse) {
	o.Schema = &v
}

func (o GetMessagingActionResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	return toSerialize, nil
}

type NullableGetMessagingActionResponseEmbedded struct {
	value *GetMessagingActionResponseEmbedded
	isSet bool
}

func (v NullableGetMessagingActionResponseEmbedded) Get() *GetMessagingActionResponseEmbedded {
	return v.value
}

func (v *NullableGetMessagingActionResponseEmbedded) Set(val *GetMessagingActionResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessagingActionResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessagingActionResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessagingActionResponseEmbedded(val *GetMessagingActionResponseEmbedded) *NullableGetMessagingActionResponseEmbedded {
	return &NullableGetMessagingActionResponseEmbedded{value: val, isSet: true}
}

func (v NullableGetMessagingActionResponseEmbedded) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMessagingActionResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
