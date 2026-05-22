package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPTORBudgetRecommendationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPTORBudgetRecommendationError{}

// SPTORBudgetRecommendationError struct for SPTORBudgetRecommendationError
type SPTORBudgetRecommendationError struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewSPTORBudgetRecommendationError instantiates a new SPTORBudgetRecommendationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPTORBudgetRecommendationError() *SPTORBudgetRecommendationError {
	this := SPTORBudgetRecommendationError{}
	return &this
}

// NewSPTORBudgetRecommendationErrorWithDefaults instantiates a new SPTORBudgetRecommendationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPTORBudgetRecommendationErrorWithDefaults() *SPTORBudgetRecommendationError {
	this := SPTORBudgetRecommendationError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SPTORBudgetRecommendationError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPTORBudgetRecommendationError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SPTORBudgetRecommendationError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SPTORBudgetRecommendationError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SPTORBudgetRecommendationError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPTORBudgetRecommendationError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SPTORBudgetRecommendationError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SPTORBudgetRecommendationError) SetDetails(v string) {
	o.Details = &v
}

func (o SPTORBudgetRecommendationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSPTORBudgetRecommendationError struct {
	value *SPTORBudgetRecommendationError
	isSet bool
}

func (v NullableSPTORBudgetRecommendationError) Get() *SPTORBudgetRecommendationError {
	return v.value
}

func (v *NullableSPTORBudgetRecommendationError) Set(val *SPTORBudgetRecommendationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSPTORBudgetRecommendationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSPTORBudgetRecommendationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPTORBudgetRecommendationError(val *SPTORBudgetRecommendationError) *NullableSPTORBudgetRecommendationError {
	return &NullableSPTORBudgetRecommendationError{value: val, isSet: true}
}

func (v NullableSPTORBudgetRecommendationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPTORBudgetRecommendationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
