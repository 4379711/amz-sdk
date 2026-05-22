package solicitations

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSolicitationActionResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSolicitationActionResponseLinks{}

// GetSolicitationActionResponseLinks struct for GetSolicitationActionResponseLinks
type GetSolicitationActionResponseLinks struct {
	Self   LinkObject `json:"self"`
	Schema LinkObject `json:"schema"`
}

type _GetSolicitationActionResponseLinks GetSolicitationActionResponseLinks

// NewGetSolicitationActionResponseLinks instantiates a new GetSolicitationActionResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSolicitationActionResponseLinks(self LinkObject, schema LinkObject) *GetSolicitationActionResponseLinks {
	this := GetSolicitationActionResponseLinks{}
	this.Self = self
	this.Schema = schema
	return &this
}

// NewGetSolicitationActionResponseLinksWithDefaults instantiates a new GetSolicitationActionResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSolicitationActionResponseLinksWithDefaults() *GetSolicitationActionResponseLinks {
	this := GetSolicitationActionResponseLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *GetSolicitationActionResponseLinks) GetSelf() LinkObject {
	if o == nil {
		var ret LinkObject
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionResponseLinks) GetSelfOk() (*LinkObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *GetSolicitationActionResponseLinks) SetSelf(v LinkObject) {
	o.Self = v
}

// GetSchema returns the Schema field value
func (o *GetSolicitationActionResponseLinks) GetSchema() LinkObject {
	if o == nil {
		var ret LinkObject
		return ret
	}

	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value
// and a boolean to check if the value has been set.
func (o *GetSolicitationActionResponseLinks) GetSchemaOk() (*LinkObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schema, true
}

// SetSchema sets field value
func (o *GetSolicitationActionResponseLinks) SetSchema(v LinkObject) {
	o.Schema = v
}

func (o GetSolicitationActionResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	toSerialize["schema"] = o.Schema
	return toSerialize, nil
}

type NullableGetSolicitationActionResponseLinks struct {
	value *GetSolicitationActionResponseLinks
	isSet bool
}

func (v NullableGetSolicitationActionResponseLinks) Get() *GetSolicitationActionResponseLinks {
	return v.value
}

func (v *NullableGetSolicitationActionResponseLinks) Set(val *GetSolicitationActionResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSolicitationActionResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSolicitationActionResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSolicitationActionResponseLinks(val *GetSolicitationActionResponseLinks) *NullableGetSolicitationActionResponseLinks {
	return &NullableGetSolicitationActionResponseLinks{value: val, isSet: true}
}

func (v NullableGetSolicitationActionResponseLinks) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSolicitationActionResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
