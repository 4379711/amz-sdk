package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetableAsinCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetableAsinCounts{}

// TargetableAsinCounts Response object to get number of targetable asins for refinements provided by the user
type TargetableAsinCounts struct {
	AsinCounts *IntegerRange `json:"asinCounts,omitempty"`
}

// NewTargetableAsinCounts instantiates a new TargetableAsinCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetableAsinCounts() *TargetableAsinCounts {
	this := TargetableAsinCounts{}
	return &this
}

// NewTargetableAsinCountsWithDefaults instantiates a new TargetableAsinCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetableAsinCountsWithDefaults() *TargetableAsinCounts {
	this := TargetableAsinCounts{}
	return &this
}

// GetAsinCounts returns the AsinCounts field value if set, zero value otherwise.
func (o *TargetableAsinCounts) GetAsinCounts() IntegerRange {
	if o == nil || IsNil(o.AsinCounts) {
		var ret IntegerRange
		return ret
	}
	return *o.AsinCounts
}

// GetAsinCountsOk returns a tuple with the AsinCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetableAsinCounts) GetAsinCountsOk() (*IntegerRange, bool) {
	if o == nil || IsNil(o.AsinCounts) {
		return nil, false
	}
	return o.AsinCounts, true
}

// HasAsinCounts returns a boolean if a field has been set.
func (o *TargetableAsinCounts) HasAsinCounts() bool {
	if o != nil && !IsNil(o.AsinCounts) {
		return true
	}

	return false
}

// SetAsinCounts gets a reference to the given IntegerRange and assigns it to the AsinCounts field.
func (o *TargetableAsinCounts) SetAsinCounts(v IntegerRange) {
	o.AsinCounts = &v
}

func (o TargetableAsinCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AsinCounts) {
		toSerialize["asinCounts"] = o.AsinCounts
	}
	return toSerialize, nil
}

type NullableTargetableAsinCounts struct {
	value *TargetableAsinCounts
	isSet bool
}

func (v NullableTargetableAsinCounts) Get() *TargetableAsinCounts {
	return v.value
}

func (v *NullableTargetableAsinCounts) Set(val *TargetableAsinCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetableAsinCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetableAsinCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetableAsinCounts(val *TargetableAsinCounts) *NullableTargetableAsinCounts {
	return &NullableTargetableAsinCounts{value: val, isSet: true}
}

func (v NullableTargetableAsinCounts) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetableAsinCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
