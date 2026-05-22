package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignNegativeKeywordAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignNegativeKeywordAccessError{}

// SponsoredProductsDraftCampaignNegativeKeywordAccessError struct for SponsoredProductsDraftCampaignNegativeKeywordAccessError
type SponsoredProductsDraftCampaignNegativeKeywordAccessError struct {
	// The type of the error
	ErrorType  string                                                           `json:"errorType"`
	ErrorValue SponsoredProductsDraftCampaignNegativeKeywordAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftCampaignNegativeKeywordAccessError SponsoredProductsDraftCampaignNegativeKeywordAccessError

// NewSponsoredProductsDraftCampaignNegativeKeywordAccessError instantiates a new SponsoredProductsDraftCampaignNegativeKeywordAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignNegativeKeywordAccessError(errorType string, errorValue SponsoredProductsDraftCampaignNegativeKeywordAccessErrorSelector) *SponsoredProductsDraftCampaignNegativeKeywordAccessError {
	this := SponsoredProductsDraftCampaignNegativeKeywordAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftCampaignNegativeKeywordAccessErrorWithDefaults instantiates a new SponsoredProductsDraftCampaignNegativeKeywordAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignNegativeKeywordAccessErrorWithDefaults() *SponsoredProductsDraftCampaignNegativeKeywordAccessError {
	this := SponsoredProductsDraftCampaignNegativeKeywordAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessError) GetErrorValue() SponsoredProductsDraftCampaignNegativeKeywordAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftCampaignNegativeKeywordAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessError) GetErrorValueOk() (*SponsoredProductsDraftCampaignNegativeKeywordAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordAccessError) SetErrorValue(v SponsoredProductsDraftCampaignNegativeKeywordAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftCampaignNegativeKeywordAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignNegativeKeywordAccessError struct {
	value *SponsoredProductsDraftCampaignNegativeKeywordAccessError
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordAccessError) Get() *SponsoredProductsDraftCampaignNegativeKeywordAccessError {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordAccessError) Set(val *SponsoredProductsDraftCampaignNegativeKeywordAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignNegativeKeywordAccessError(val *SponsoredProductsDraftCampaignNegativeKeywordAccessError) *NullableSponsoredProductsDraftCampaignNegativeKeywordAccessError {
	return &NullableSponsoredProductsDraftCampaignNegativeKeywordAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
