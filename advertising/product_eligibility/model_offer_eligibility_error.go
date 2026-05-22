package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the OfferEligibilityError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferEligibilityError{}

// OfferEligibilityError The HTTP status code of the response
type OfferEligibilityError struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewOfferEligibilityError instantiates a new OfferEligibilityError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferEligibilityError() *OfferEligibilityError {
	this := OfferEligibilityError{}
	return &this
}

// NewOfferEligibilityErrorWithDefaults instantiates a new OfferEligibilityError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferEligibilityErrorWithDefaults() *OfferEligibilityError {
	this := OfferEligibilityError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OfferEligibilityError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferEligibilityError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OfferEligibilityError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OfferEligibilityError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OfferEligibilityError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferEligibilityError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OfferEligibilityError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *OfferEligibilityError) SetDetails(v string) {
	o.Details = &v
}

func (o OfferEligibilityError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableOfferEligibilityError struct {
	value *OfferEligibilityError
	isSet bool
}

func (v NullableOfferEligibilityError) Get() *OfferEligibilityError {
	return v.value
}

func (v *NullableOfferEligibilityError) Set(val *OfferEligibilityError) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferEligibilityError) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferEligibilityError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferEligibilityError(val *OfferEligibilityError) *NullableOfferEligibilityError {
	return &NullableOfferEligibilityError{value: val, isSet: true}
}

func (v NullableOfferEligibilityError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOfferEligibilityError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
