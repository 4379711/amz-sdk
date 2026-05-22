package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationsThrottlingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationsThrottlingError{}

// CreativeRecommendationsThrottlingError struct for CreativeRecommendationsThrottlingError
type CreativeRecommendationsThrottlingError struct {
	// Throttled error code.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error response.
	Details *string `json:"details,omitempty"`
}

// NewCreativeRecommendationsThrottlingError instantiates a new CreativeRecommendationsThrottlingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationsThrottlingError() *CreativeRecommendationsThrottlingError {
	this := CreativeRecommendationsThrottlingError{}
	return &this
}

// NewCreativeRecommendationsThrottlingErrorWithDefaults instantiates a new CreativeRecommendationsThrottlingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationsThrottlingErrorWithDefaults() *CreativeRecommendationsThrottlingError {
	this := CreativeRecommendationsThrottlingError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreativeRecommendationsThrottlingError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsThrottlingError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreativeRecommendationsThrottlingError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreativeRecommendationsThrottlingError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CreativeRecommendationsThrottlingError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsThrottlingError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CreativeRecommendationsThrottlingError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *CreativeRecommendationsThrottlingError) SetDetails(v string) {
	o.Details = &v
}

func (o CreativeRecommendationsThrottlingError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableCreativeRecommendationsThrottlingError struct {
	value *CreativeRecommendationsThrottlingError
	isSet bool
}

func (v NullableCreativeRecommendationsThrottlingError) Get() *CreativeRecommendationsThrottlingError {
	return v.value
}

func (v *NullableCreativeRecommendationsThrottlingError) Set(val *CreativeRecommendationsThrottlingError) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationsThrottlingError) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationsThrottlingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationsThrottlingError(val *CreativeRecommendationsThrottlingError) *NullableCreativeRecommendationsThrottlingError {
	return &NullableCreativeRecommendationsThrottlingError{value: val, isSet: true}
}

func (v NullableCreativeRecommendationsThrottlingError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationsThrottlingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
