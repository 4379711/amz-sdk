package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetUsageError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetUsageError{}

// BudgetUsageError The Error Response Object.
type BudgetUsageError struct {
	// An enumerated error code for machine use.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewBudgetUsageError instantiates a new BudgetUsageError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUsageError() *BudgetUsageError {
	this := BudgetUsageError{}
	return &this
}

// NewBudgetUsageErrorWithDefaults instantiates a new BudgetUsageError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUsageErrorWithDefaults() *BudgetUsageError {
	this := BudgetUsageError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BudgetUsageError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BudgetUsageError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *BudgetUsageError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *BudgetUsageError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BudgetUsageError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *BudgetUsageError) SetDetails(v string) {
	o.Details = &v
}

func (o BudgetUsageError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableBudgetUsageError struct {
	value *BudgetUsageError
	isSet bool
}

func (v NullableBudgetUsageError) Get() *BudgetUsageError {
	return v.value
}

func (v *NullableBudgetUsageError) Set(val *BudgetUsageError) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUsageError) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUsageError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUsageError(val *BudgetUsageError) *NullableBudgetUsageError {
	return &NullableBudgetUsageError{value: val, isSet: true}
}

func (v NullableBudgetUsageError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetUsageError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
