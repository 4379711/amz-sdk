package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDraftProductAdMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftProductAdMutationError{}

// SponsoredProductsDraftProductAdMutationError struct for SponsoredProductsDraftProductAdMutationError
type SponsoredProductsDraftProductAdMutationError struct {
	// The type of the error
	ErrorType  string                                               `json:"errorType"`
	ErrorValue SponsoredProductsDraftProductAdMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftProductAdMutationError SponsoredProductsDraftProductAdMutationError

// NewSponsoredProductsDraftProductAdMutationError instantiates a new SponsoredProductsDraftProductAdMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftProductAdMutationError(errorType string, errorValue SponsoredProductsDraftProductAdMutationErrorSelector) *SponsoredProductsDraftProductAdMutationError {
	this := SponsoredProductsDraftProductAdMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftProductAdMutationErrorWithDefaults instantiates a new SponsoredProductsDraftProductAdMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftProductAdMutationErrorWithDefaults() *SponsoredProductsDraftProductAdMutationError {
	this := SponsoredProductsDraftProductAdMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftProductAdMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftProductAdMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftProductAdMutationError) GetErrorValue() SponsoredProductsDraftProductAdMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftProductAdMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdMutationError) GetErrorValueOk() (*SponsoredProductsDraftProductAdMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftProductAdMutationError) SetErrorValue(v SponsoredProductsDraftProductAdMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftProductAdMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftProductAdMutationError struct {
	value *SponsoredProductsDraftProductAdMutationError
	isSet bool
}

func (v NullableSponsoredProductsDraftProductAdMutationError) Get() *SponsoredProductsDraftProductAdMutationError {
	return v.value
}

func (v *NullableSponsoredProductsDraftProductAdMutationError) Set(val *SponsoredProductsDraftProductAdMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftProductAdMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftProductAdMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftProductAdMutationError(val *SponsoredProductsDraftProductAdMutationError) *NullableSponsoredProductsDraftProductAdMutationError {
	return &NullableSponsoredProductsDraftProductAdMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftProductAdMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftProductAdMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
