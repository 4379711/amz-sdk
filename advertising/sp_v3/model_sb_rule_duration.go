package sp_v3

import "github.com/bytedance/sonic"

// checks if the SBRuleDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBRuleDuration{}

// SBRuleDuration struct for SBRuleDuration
type SBRuleDuration struct {
	DateRangeTypeRuleDuration DateRangeTypeRuleDuration `json:"dateRangeTypeRuleDuration"`
}

type _SBRuleDuration SBRuleDuration

// NewSBRuleDuration instantiates a new SBRuleDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBRuleDuration(dateRangeTypeRuleDuration DateRangeTypeRuleDuration) *SBRuleDuration {
	this := SBRuleDuration{}
	this.DateRangeTypeRuleDuration = dateRangeTypeRuleDuration
	return &this
}

// NewSBRuleDurationWithDefaults instantiates a new SBRuleDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBRuleDurationWithDefaults() *SBRuleDuration {
	this := SBRuleDuration{}
	return &this
}

// GetDateRangeTypeRuleDuration returns the DateRangeTypeRuleDuration field value
func (o *SBRuleDuration) GetDateRangeTypeRuleDuration() DateRangeTypeRuleDuration {
	if o == nil {
		var ret DateRangeTypeRuleDuration
		return ret
	}

	return o.DateRangeTypeRuleDuration
}

// GetDateRangeTypeRuleDurationOk returns a tuple with the DateRangeTypeRuleDuration field value
// and a boolean to check if the value has been set.
func (o *SBRuleDuration) GetDateRangeTypeRuleDurationOk() (*DateRangeTypeRuleDuration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateRangeTypeRuleDuration, true
}

// SetDateRangeTypeRuleDuration sets field value
func (o *SBRuleDuration) SetDateRangeTypeRuleDuration(v DateRangeTypeRuleDuration) {
	o.DateRangeTypeRuleDuration = v
}

func (o SBRuleDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dateRangeTypeRuleDuration"] = o.DateRangeTypeRuleDuration
	return toSerialize, nil
}

type NullableSBRuleDuration struct {
	value *SBRuleDuration
	isSet bool
}

func (v NullableSBRuleDuration) Get() *SBRuleDuration {
	return v.value
}

func (v *NullableSBRuleDuration) Set(val *SBRuleDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableSBRuleDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableSBRuleDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBRuleDuration(val *SBRuleDuration) *NullableSBRuleDuration {
	return &NullableSBRuleDuration{value: val, isSet: true}
}

func (v NullableSBRuleDuration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBRuleDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
