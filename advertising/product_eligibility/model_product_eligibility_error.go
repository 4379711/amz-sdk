package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductEligibilityError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductEligibilityError{}

// ProductEligibilityError The error response object.
type ProductEligibilityError struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewProductEligibilityError instantiates a new ProductEligibilityError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductEligibilityError() *ProductEligibilityError {
	this := ProductEligibilityError{}
	return &this
}

// NewProductEligibilityErrorWithDefaults instantiates a new ProductEligibilityError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductEligibilityErrorWithDefaults() *ProductEligibilityError {
	this := ProductEligibilityError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProductEligibilityError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductEligibilityError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProductEligibilityError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProductEligibilityError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ProductEligibilityError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductEligibilityError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ProductEligibilityError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *ProductEligibilityError) SetDetails(v string) {
	o.Details = &v
}

func (o ProductEligibilityError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableProductEligibilityError struct {
	value *ProductEligibilityError
	isSet bool
}

func (v NullableProductEligibilityError) Get() *ProductEligibilityError {
	return v.value
}

func (v *NullableProductEligibilityError) Set(val *ProductEligibilityError) {
	v.value = val
	v.isSet = true
}

func (v NullableProductEligibilityError) IsSet() bool {
	return v.isSet
}

func (v *NullableProductEligibilityError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductEligibilityError(val *ProductEligibilityError) *NullableProductEligibilityError {
	return &NullableProductEligibilityError{value: val, isSet: true}
}

func (v NullableProductEligibilityError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductEligibilityError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
