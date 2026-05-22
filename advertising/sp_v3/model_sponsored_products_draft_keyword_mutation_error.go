package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDraftKeywordMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftKeywordMutationError{}

// SponsoredProductsDraftKeywordMutationError struct for SponsoredProductsDraftKeywordMutationError
type SponsoredProductsDraftKeywordMutationError struct {
	// The type of the error
	ErrorType  string                                             `json:"errorType"`
	ErrorValue SponsoredProductsDraftKeywordMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftKeywordMutationError SponsoredProductsDraftKeywordMutationError

// NewSponsoredProductsDraftKeywordMutationError instantiates a new SponsoredProductsDraftKeywordMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftKeywordMutationError(errorType string, errorValue SponsoredProductsDraftKeywordMutationErrorSelector) *SponsoredProductsDraftKeywordMutationError {
	this := SponsoredProductsDraftKeywordMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftKeywordMutationErrorWithDefaults instantiates a new SponsoredProductsDraftKeywordMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftKeywordMutationErrorWithDefaults() *SponsoredProductsDraftKeywordMutationError {
	this := SponsoredProductsDraftKeywordMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftKeywordMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftKeywordMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftKeywordMutationError) GetErrorValue() SponsoredProductsDraftKeywordMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftKeywordMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordMutationError) GetErrorValueOk() (*SponsoredProductsDraftKeywordMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftKeywordMutationError) SetErrorValue(v SponsoredProductsDraftKeywordMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftKeywordMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftKeywordMutationError struct {
	value *SponsoredProductsDraftKeywordMutationError
	isSet bool
}

func (v NullableSponsoredProductsDraftKeywordMutationError) Get() *SponsoredProductsDraftKeywordMutationError {
	return v.value
}

func (v *NullableSponsoredProductsDraftKeywordMutationError) Set(val *SponsoredProductsDraftKeywordMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftKeywordMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftKeywordMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftKeywordMutationError(val *SponsoredProductsDraftKeywordMutationError) *NullableSponsoredProductsDraftKeywordMutationError {
	return &NullableSponsoredProductsDraftKeywordMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftKeywordMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftKeywordMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
