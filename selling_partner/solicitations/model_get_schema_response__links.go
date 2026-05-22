package solicitations

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSchemaResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSchemaResponseLinks{}

// GetSchemaResponseLinks struct for GetSchemaResponseLinks
type GetSchemaResponseLinks struct {
	Self LinkObject `json:"self"`
}

type _GetSchemaResponseLinks GetSchemaResponseLinks

// NewGetSchemaResponseLinks instantiates a new GetSchemaResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSchemaResponseLinks(self LinkObject) *GetSchemaResponseLinks {
	this := GetSchemaResponseLinks{}
	this.Self = self
	return &this
}

// NewGetSchemaResponseLinksWithDefaults instantiates a new GetSchemaResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSchemaResponseLinksWithDefaults() *GetSchemaResponseLinks {
	this := GetSchemaResponseLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *GetSchemaResponseLinks) GetSelf() LinkObject {
	if o == nil {
		var ret LinkObject
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *GetSchemaResponseLinks) GetSelfOk() (*LinkObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *GetSchemaResponseLinks) SetSelf(v LinkObject) {
	o.Self = v
}

func (o GetSchemaResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	return toSerialize, nil
}

type NullableGetSchemaResponseLinks struct {
	value *GetSchemaResponseLinks
	isSet bool
}

func (v NullableGetSchemaResponseLinks) Get() *GetSchemaResponseLinks {
	return v.value
}

func (v *NullableGetSchemaResponseLinks) Set(val *GetSchemaResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSchemaResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSchemaResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSchemaResponseLinks(val *GetSchemaResponseLinks) *NullableGetSchemaResponseLinks {
	return &NullableGetSchemaResponseLinks{value: val, isSet: true}
}

func (v NullableGetSchemaResponseLinks) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSchemaResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
