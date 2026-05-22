package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CountryWithTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryWithTargets{}

// CountryWithTargets struct for CountryWithTargets
type CountryWithTargets struct {
	// A list of targets that need to be ranked
	Targets []KeywordTarget `json:"targets,omitempty"`
}

// NewCountryWithTargets instantiates a new CountryWithTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryWithTargets() *CountryWithTargets {
	this := CountryWithTargets{}
	return &this
}

// NewCountryWithTargetsWithDefaults instantiates a new CountryWithTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryWithTargetsWithDefaults() *CountryWithTargets {
	this := CountryWithTargets{}
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *CountryWithTargets) GetTargets() []KeywordTarget {
	if o == nil || IsNil(o.Targets) {
		var ret []KeywordTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryWithTargets) GetTargetsOk() ([]KeywordTarget, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *CountryWithTargets) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []KeywordTarget and assigns it to the Targets field.
func (o *CountryWithTargets) SetTargets(v []KeywordTarget) {
	o.Targets = v
}

func (o CountryWithTargets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableCountryWithTargets struct {
	value *CountryWithTargets
	isSet bool
}

func (v NullableCountryWithTargets) Get() *CountryWithTargets {
	return v.value
}

func (v *NullableCountryWithTargets) Set(val *CountryWithTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryWithTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryWithTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryWithTargets(val *CountryWithTargets) *NullableCountryWithTargets {
	return &NullableCountryWithTargets{value: val, isSet: true}
}

func (v NullableCountryWithTargets) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCountryWithTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
