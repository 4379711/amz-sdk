package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeKeywordAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeKeywordAccessError{}

// SponsoredProductsCampaignNegativeKeywordAccessError struct for SponsoredProductsCampaignNegativeKeywordAccessError
type SponsoredProductsCampaignNegativeKeywordAccessError struct {
	// The type of the error
	ErrorType  string                                                      `json:"errorType"`
	ErrorValue SponsoredProductsCampaignNegativeKeywordAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsCampaignNegativeKeywordAccessError SponsoredProductsCampaignNegativeKeywordAccessError

// NewSponsoredProductsCampaignNegativeKeywordAccessError instantiates a new SponsoredProductsCampaignNegativeKeywordAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeKeywordAccessError(errorType string, errorValue SponsoredProductsCampaignNegativeKeywordAccessErrorSelector) *SponsoredProductsCampaignNegativeKeywordAccessError {
	this := SponsoredProductsCampaignNegativeKeywordAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsCampaignNegativeKeywordAccessErrorWithDefaults instantiates a new SponsoredProductsCampaignNegativeKeywordAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeKeywordAccessErrorWithDefaults() *SponsoredProductsCampaignNegativeKeywordAccessError {
	this := SponsoredProductsCampaignNegativeKeywordAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsCampaignNegativeKeywordAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsCampaignNegativeKeywordAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsCampaignNegativeKeywordAccessError) GetErrorValue() SponsoredProductsCampaignNegativeKeywordAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsCampaignNegativeKeywordAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordAccessError) GetErrorValueOk() (*SponsoredProductsCampaignNegativeKeywordAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsCampaignNegativeKeywordAccessError) SetErrorValue(v SponsoredProductsCampaignNegativeKeywordAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsCampaignNegativeKeywordAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeKeywordAccessError struct {
	value *SponsoredProductsCampaignNegativeKeywordAccessError
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeKeywordAccessError) Get() *SponsoredProductsCampaignNegativeKeywordAccessError {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordAccessError) Set(val *SponsoredProductsCampaignNegativeKeywordAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeKeywordAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeKeywordAccessError(val *SponsoredProductsCampaignNegativeKeywordAccessError) *NullableSponsoredProductsCampaignNegativeKeywordAccessError {
	return &NullableSponsoredProductsCampaignNegativeKeywordAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeKeywordAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
