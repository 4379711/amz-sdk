package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeTargetsMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeTargetsMutationError{}

// SponsoredProductsCampaignNegativeTargetsMutationError struct for SponsoredProductsCampaignNegativeTargetsMutationError
type SponsoredProductsCampaignNegativeTargetsMutationError struct {
	// The type of the error
	ErrorType  string                                                        `json:"errorType"`
	ErrorValue SponsoredProductsCampaignNegativeTargetsMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsCampaignNegativeTargetsMutationError SponsoredProductsCampaignNegativeTargetsMutationError

// NewSponsoredProductsCampaignNegativeTargetsMutationError instantiates a new SponsoredProductsCampaignNegativeTargetsMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeTargetsMutationError(errorType string, errorValue SponsoredProductsCampaignNegativeTargetsMutationErrorSelector) *SponsoredProductsCampaignNegativeTargetsMutationError {
	this := SponsoredProductsCampaignNegativeTargetsMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsCampaignNegativeTargetsMutationErrorWithDefaults instantiates a new SponsoredProductsCampaignNegativeTargetsMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeTargetsMutationErrorWithDefaults() *SponsoredProductsCampaignNegativeTargetsMutationError {
	this := SponsoredProductsCampaignNegativeTargetsMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsCampaignNegativeTargetsMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsCampaignNegativeTargetsMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsCampaignNegativeTargetsMutationError) GetErrorValue() SponsoredProductsCampaignNegativeTargetsMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsCampaignNegativeTargetsMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsMutationError) GetErrorValueOk() (*SponsoredProductsCampaignNegativeTargetsMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsCampaignNegativeTargetsMutationError) SetErrorValue(v SponsoredProductsCampaignNegativeTargetsMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsCampaignNegativeTargetsMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeTargetsMutationError struct {
	value *SponsoredProductsCampaignNegativeTargetsMutationError
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeTargetsMutationError) Get() *SponsoredProductsCampaignNegativeTargetsMutationError {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsMutationError) Set(val *SponsoredProductsCampaignNegativeTargetsMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeTargetsMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeTargetsMutationError(val *SponsoredProductsCampaignNegativeTargetsMutationError) *NullableSponsoredProductsCampaignNegativeTargetsMutationError {
	return &NullableSponsoredProductsCampaignNegativeTargetsMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeTargetsMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
