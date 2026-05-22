package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Impressions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Impressions{}

// Impressions Impressions benchmark.
type Impressions struct {
	// lower bound.
	Lower *int32 `json:"lower,omitempty"`
	// upper bound.
	Upper *int32 `json:"upper,omitempty"`
}

// NewImpressions instantiates a new Impressions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImpressions() *Impressions {
	this := Impressions{}
	return &this
}

// NewImpressionsWithDefaults instantiates a new Impressions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImpressionsWithDefaults() *Impressions {
	this := Impressions{}
	return &this
}

// GetLower returns the Lower field value if set, zero value otherwise.
func (o *Impressions) GetLower() int32 {
	if o == nil || IsNil(o.Lower) {
		var ret int32
		return ret
	}
	return *o.Lower
}

// GetLowerOk returns a tuple with the Lower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impressions) GetLowerOk() (*int32, bool) {
	if o == nil || IsNil(o.Lower) {
		return nil, false
	}
	return o.Lower, true
}

// HasLower returns a boolean if a field has been set.
func (o *Impressions) HasLower() bool {
	if o != nil && !IsNil(o.Lower) {
		return true
	}

	return false
}

// SetLower gets a reference to the given int32 and assigns it to the Lower field.
func (o *Impressions) SetLower(v int32) {
	o.Lower = &v
}

// GetUpper returns the Upper field value if set, zero value otherwise.
func (o *Impressions) GetUpper() int32 {
	if o == nil || IsNil(o.Upper) {
		var ret int32
		return ret
	}
	return *o.Upper
}

// GetUpperOk returns a tuple with the Upper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Impressions) GetUpperOk() (*int32, bool) {
	if o == nil || IsNil(o.Upper) {
		return nil, false
	}
	return o.Upper, true
}

// HasUpper returns a boolean if a field has been set.
func (o *Impressions) HasUpper() bool {
	if o != nil && !IsNil(o.Upper) {
		return true
	}

	return false
}

// SetUpper gets a reference to the given int32 and assigns it to the Upper field.
func (o *Impressions) SetUpper(v int32) {
	o.Upper = &v
}

func (o Impressions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lower) {
		toSerialize["lower"] = o.Lower
	}
	if !IsNil(o.Upper) {
		toSerialize["upper"] = o.Upper
	}
	return toSerialize, nil
}

type NullableImpressions struct {
	value *Impressions
	isSet bool
}

func (v NullableImpressions) Get() *Impressions {
	return v.value
}

func (v *NullableImpressions) Set(val *Impressions) {
	v.value = val
	v.isSet = true
}

func (v NullableImpressions) IsSet() bool {
	return v.isSet
}

func (v *NullableImpressions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImpressions(val *Impressions) *NullableImpressions {
	return &NullableImpressions{value: val, isSet: true}
}

func (v NullableImpressions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImpressions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
