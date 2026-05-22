package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioMutationError{}

// PortfolioMutationError struct for PortfolioMutationError
type PortfolioMutationError struct {
	// The type of the error
	ErrorType  string                         `json:"errorType"`
	ErrorValue PortfolioMutationErrorSelector `json:"errorValue"`
}

type _PortfolioMutationError PortfolioMutationError

// NewPortfolioMutationError instantiates a new PortfolioMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMutationError(errorType string, errorValue PortfolioMutationErrorSelector) *PortfolioMutationError {
	this := PortfolioMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewPortfolioMutationErrorWithDefaults instantiates a new PortfolioMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMutationErrorWithDefaults() *PortfolioMutationError {
	this := PortfolioMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *PortfolioMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *PortfolioMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *PortfolioMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *PortfolioMutationError) GetErrorValue() PortfolioMutationErrorSelector {
	if o == nil {
		var ret PortfolioMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *PortfolioMutationError) GetErrorValueOk() (*PortfolioMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *PortfolioMutationError) SetErrorValue(v PortfolioMutationErrorSelector) {
	o.ErrorValue = v
}

func (o PortfolioMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullablePortfolioMutationError struct {
	value *PortfolioMutationError
	isSet bool
}

func (v NullablePortfolioMutationError) Get() *PortfolioMutationError {
	return v.value
}

func (v *NullablePortfolioMutationError) Set(val *PortfolioMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMutationError(val *PortfolioMutationError) *NullablePortfolioMutationError {
	return &NullablePortfolioMutationError{value: val, isSet: true}
}

func (v NullablePortfolioMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
