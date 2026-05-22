package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BaseLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseLocation{}

// BaseLocation struct for BaseLocation
type BaseLocation struct {
	State *string `json:"state,omitempty"`
}

// NewBaseLocation instantiates a new BaseLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseLocation() *BaseLocation {
	this := BaseLocation{}
	return &this
}

// NewBaseLocationWithDefaults instantiates a new BaseLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseLocationWithDefaults() *BaseLocation {
	this := BaseLocation{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BaseLocation) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseLocation) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BaseLocation) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BaseLocation) SetState(v string) {
	o.State = &v
}

func (o BaseLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableBaseLocation struct {
	value *BaseLocation
	isSet bool
}

func (v NullableBaseLocation) Get() *BaseLocation {
	return v.value
}

func (v *NullableBaseLocation) Set(val *BaseLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseLocation(val *BaseLocation) *NullableBaseLocation {
	return &NullableBaseLocation{value: val, isSet: true}
}

func (v NullableBaseLocation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBaseLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
