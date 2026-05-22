package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeKeywordMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeKeywordMutationError{}

// SponsoredProductsCampaignNegativeKeywordMutationError struct for SponsoredProductsCampaignNegativeKeywordMutationError
type SponsoredProductsCampaignNegativeKeywordMutationError struct {
	// The type of the error
	ErrorType  string                                                        `json:"errorType"`
	ErrorValue SponsoredProductsCampaignNegativeKeywordMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsCampaignNegativeKeywordMutationError SponsoredProductsCampaignNegativeKeywordMutationError

// NewSponsoredProductsCampaignNegativeKeywordMutationError instantiates a new SponsoredProductsCampaignNegativeKeywordMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeKeywordMutationError(errorType string, errorValue SponsoredProductsCampaignNegativeKeywordMutationErrorSelector) *SponsoredProductsCampaignNegativeKeywordMutationError {
	this := SponsoredProductsCampaignNegativeKeywordMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsCampaignNegativeKeywordMutationErrorWithDefaults instantiates a new SponsoredProductsCampaignNegativeKeywordMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeKeywordMutationErrorWithDefaults() *SponsoredProductsCampaignNegativeKeywordMutationError {
	this := SponsoredProductsCampaignNegativeKeywordMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsCampaignNegativeKeywordMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsCampaignNegativeKeywordMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsCampaignNegativeKeywordMutationError) GetErrorValue() SponsoredProductsCampaignNegativeKeywordMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsCampaignNegativeKeywordMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordMutationError) GetErrorValueOk() (*SponsoredProductsCampaignNegativeKeywordMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsCampaignNegativeKeywordMutationError) SetErrorValue(v SponsoredProductsCampaignNegativeKeywordMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsCampaignNegativeKeywordMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeKeywordMutationError struct {
	value *SponsoredProductsCampaignNegativeKeywordMutationError
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeKeywordMutationError) Get() *SponsoredProductsCampaignNegativeKeywordMutationError {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordMutationError) Set(val *SponsoredProductsCampaignNegativeKeywordMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeKeywordMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeKeywordMutationError(val *SponsoredProductsCampaignNegativeKeywordMutationError) *NullableSponsoredProductsCampaignNegativeKeywordMutationError {
	return &NullableSponsoredProductsCampaignNegativeKeywordMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeKeywordMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
