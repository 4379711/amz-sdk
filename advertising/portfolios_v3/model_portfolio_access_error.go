package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the PortfolioAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioAccessError{}

// PortfolioAccessError struct for PortfolioAccessError
type PortfolioAccessError struct {
	// The type of the error
	ErrorType  string                       `json:"errorType"`
	ErrorValue PortfolioAccessErrorSelector `json:"errorValue"`
}

type _PortfolioAccessError PortfolioAccessError

// NewPortfolioAccessError instantiates a new PortfolioAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioAccessError(errorType string, errorValue PortfolioAccessErrorSelector) *PortfolioAccessError {
	this := PortfolioAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewPortfolioAccessErrorWithDefaults instantiates a new PortfolioAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioAccessErrorWithDefaults() *PortfolioAccessError {
	this := PortfolioAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *PortfolioAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *PortfolioAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *PortfolioAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *PortfolioAccessError) GetErrorValue() PortfolioAccessErrorSelector {
	if o == nil {
		var ret PortfolioAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *PortfolioAccessError) GetErrorValueOk() (*PortfolioAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *PortfolioAccessError) SetErrorValue(v PortfolioAccessErrorSelector) {
	o.ErrorValue = v
}

func (o PortfolioAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullablePortfolioAccessError struct {
	value *PortfolioAccessError
	isSet bool
}

func (v NullablePortfolioAccessError) Get() *PortfolioAccessError {
	return v.value
}

func (v *NullablePortfolioAccessError) Set(val *PortfolioAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioAccessError(val *PortfolioAccessError) *NullablePortfolioAccessError {
	return &NullablePortfolioAccessError{value: val, isSet: true}
}

func (v NullablePortfolioAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePortfolioAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
