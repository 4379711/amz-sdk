package solicitations

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSolicitationActionResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolicitationActionResponseEmbedded{}

// GetSolicitationActionResponseEmbedded struct for GetSolicitationActionResponseEmbedded
type GetSolicitationActionResponseEmbedded struct {
	Schema *GetSchemaResponse `json:"schema,omitempty"`
}

// NewGetSolicitationActionResponseEmbedded instantiates a new GetSolicitationActionResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolicitationActionResponseEmbedded() *GetSolicitationActionResponseEmbedded {
	this := GetSolicitationActionResponseEmbedded{}
	return &this
}

// NewGetSolicitationActionResponseEmbeddedWithDefaults instantiates a new GetSolicitationActionResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolicitationActionResponseEmbeddedWithDefaults() *GetSolicitationActionResponseEmbedded {
	this := GetSolicitationActionResponseEmbedded{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *GetSolicitationActionResponseEmbedded) GetSchema() GetSchemaResponse {
	if o == nil || IsNil(o.Schema) {
		var ret GetSchemaResponse
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionResponseEmbedded) GetSchemaOk() (*GetSchemaResponse, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *GetSolicitationActionResponseEmbedded) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given GetSchemaResponse and assigns it to the Schema field.
func (o *GetSolicitationActionResponseEmbedded) SetSchema(v GetSchemaResponse) {
	o.Schema = &v
}

func (o GetSolicitationActionResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	return toSerialize, nil
}

type NullableGetSolicitationActionResponseEmbedded struct {
	value *GetSolicitationActionResponseEmbedded
	isSet bool
}

func (v NullableGetSolicitationActionResponseEmbedded) Get() *GetSolicitationActionResponseEmbedded {
	return v.value
}

func (v *NullableGetSolicitationActionResponseEmbedded) Set(val *GetSolicitationActionResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolicitationActionResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolicitationActionResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolicitationActionResponseEmbedded(val *GetSolicitationActionResponseEmbedded) *NullableGetSolicitationActionResponseEmbedded {
	return &NullableGetSolicitationActionResponseEmbedded{value: val, isSet: true}
}

func (v NullableGetSolicitationActionResponseEmbedded) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSolicitationActionResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
