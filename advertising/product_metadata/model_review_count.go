package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the ReviewCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewCount{}

// ReviewCount struct for ReviewCount
type ReviewCount struct {
	DisplayString *string `json:"displayString,omitempty"`
	Value         *int32  `json:"value,omitempty"`
}

// NewReviewCount instantiates a new ReviewCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewCount() *ReviewCount {
	this := ReviewCount{}
	return &this
}

// NewReviewCountWithDefaults instantiates a new ReviewCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewCountWithDefaults() *ReviewCount {
	this := ReviewCount{}
	return &this
}

// GetDisplayString returns the DisplayString field value if set, zero value otherwise.
func (o *ReviewCount) GetDisplayString() string {
	if o == nil || IsNil(o.DisplayString) {
		var ret string
		return ret
	}
	return *o.DisplayString
}

// GetDisplayStringOk returns a tuple with the DisplayString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewCount) GetDisplayStringOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayString) {
		return nil, false
	}
	return o.DisplayString, true
}

// HasDisplayString returns a boolean if a field has been set.
func (o *ReviewCount) HasDisplayString() bool {
	if o != nil && !IsNil(o.DisplayString) {
		return true
	}

	return false
}

// SetDisplayString gets a reference to the given string and assigns it to the DisplayString field.
func (o *ReviewCount) SetDisplayString(v string) {
	o.DisplayString = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ReviewCount) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewCount) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ReviewCount) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ReviewCount) SetValue(v int32) {
	o.Value = &v
}

func (o ReviewCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayString) {
		toSerialize["displayString"] = o.DisplayString
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableReviewCount struct {
	value *ReviewCount
	isSet bool
}

func (v NullableReviewCount) Get() *ReviewCount {
	return v.value
}

func (v *NullableReviewCount) Set(val *ReviewCount) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewCount) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewCount(val *ReviewCount) *NullableReviewCount {
	return &NullableReviewCount{value: val, isSet: true}
}

func (v NullableReviewCount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReviewCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
