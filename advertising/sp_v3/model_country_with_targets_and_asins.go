package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CountryWithTargetsAndAsins type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryWithTargetsAndAsins{}

// CountryWithTargetsAndAsins struct for CountryWithTargetsAndAsins
type CountryWithTargetsAndAsins struct {
	// An array list of Asins
	Asins []string `json:"asins,omitempty"`
	// A list of targets that need to be ranked
	Targets []KeywordTarget `json:"targets,omitempty"`
}

// NewCountryWithTargetsAndAsins instantiates a new CountryWithTargetsAndAsins object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryWithTargetsAndAsins() *CountryWithTargetsAndAsins {
	this := CountryWithTargetsAndAsins{}
	return &this
}

// NewCountryWithTargetsAndAsinsWithDefaults instantiates a new CountryWithTargetsAndAsins object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryWithTargetsAndAsinsWithDefaults() *CountryWithTargetsAndAsins {
	this := CountryWithTargetsAndAsins{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *CountryWithTargetsAndAsins) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryWithTargetsAndAsins) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *CountryWithTargetsAndAsins) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *CountryWithTargetsAndAsins) SetAsins(v []string) {
	o.Asins = v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *CountryWithTargetsAndAsins) GetTargets() []KeywordTarget {
	if o == nil || IsNil(o.Targets) {
		var ret []KeywordTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryWithTargetsAndAsins) GetTargetsOk() ([]KeywordTarget, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *CountryWithTargetsAndAsins) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []KeywordTarget and assigns it to the Targets field.
func (o *CountryWithTargetsAndAsins) SetTargets(v []KeywordTarget) {
	o.Targets = v
}

func (o CountryWithTargetsAndAsins) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableCountryWithTargetsAndAsins struct {
	value *CountryWithTargetsAndAsins
	isSet bool
}

func (v NullableCountryWithTargetsAndAsins) Get() *CountryWithTargetsAndAsins {
	return v.value
}

func (v *NullableCountryWithTargetsAndAsins) Set(val *CountryWithTargetsAndAsins) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryWithTargetsAndAsins) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryWithTargetsAndAsins) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryWithTargetsAndAsins(val *CountryWithTargetsAndAsins) *NullableCountryWithTargetsAndAsins {
	return &NullableCountryWithTargetsAndAsins{value: val, isSet: true}
}

func (v NullableCountryWithTargetsAndAsins) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCountryWithTargetsAndAsins) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
