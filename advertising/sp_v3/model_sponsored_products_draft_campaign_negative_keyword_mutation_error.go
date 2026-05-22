package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignNegativeKeywordMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignNegativeKeywordMutationError{}

// SponsoredProductsDraftCampaignNegativeKeywordMutationError struct for SponsoredProductsDraftCampaignNegativeKeywordMutationError
type SponsoredProductsDraftCampaignNegativeKeywordMutationError struct {
	// The type of the error
	ErrorType  string                                                             `json:"errorType"`
	ErrorValue SponsoredProductsDraftCampaignNegativeKeywordMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftCampaignNegativeKeywordMutationError SponsoredProductsDraftCampaignNegativeKeywordMutationError

// NewSponsoredProductsDraftCampaignNegativeKeywordMutationError instantiates a new SponsoredProductsDraftCampaignNegativeKeywordMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignNegativeKeywordMutationError(errorType string, errorValue SponsoredProductsDraftCampaignNegativeKeywordMutationErrorSelector) *SponsoredProductsDraftCampaignNegativeKeywordMutationError {
	this := SponsoredProductsDraftCampaignNegativeKeywordMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftCampaignNegativeKeywordMutationErrorWithDefaults instantiates a new SponsoredProductsDraftCampaignNegativeKeywordMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignNegativeKeywordMutationErrorWithDefaults() *SponsoredProductsDraftCampaignNegativeKeywordMutationError {
	this := SponsoredProductsDraftCampaignNegativeKeywordMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationError) GetErrorValue() SponsoredProductsDraftCampaignNegativeKeywordMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftCampaignNegativeKeywordMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationError) GetErrorValueOk() (*SponsoredProductsDraftCampaignNegativeKeywordMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordMutationError) SetErrorValue(v SponsoredProductsDraftCampaignNegativeKeywordMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftCampaignNegativeKeywordMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignNegativeKeywordMutationError struct {
	value *SponsoredProductsDraftCampaignNegativeKeywordMutationError
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordMutationError) Get() *SponsoredProductsDraftCampaignNegativeKeywordMutationError {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordMutationError) Set(val *SponsoredProductsDraftCampaignNegativeKeywordMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignNegativeKeywordMutationError(val *SponsoredProductsDraftCampaignNegativeKeywordMutationError) *NullableSponsoredProductsDraftCampaignNegativeKeywordMutationError {
	return &NullableSponsoredProductsDraftCampaignNegativeKeywordMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
