package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignMutationError{}

// SponsoredProductsDraftCampaignMutationError struct for SponsoredProductsDraftCampaignMutationError
type SponsoredProductsDraftCampaignMutationError struct {
	// The type of the error
	ErrorType  string                                              `json:"errorType"`
	ErrorValue SponsoredProductsDraftCampaignMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftCampaignMutationError SponsoredProductsDraftCampaignMutationError

// NewSponsoredProductsDraftCampaignMutationError instantiates a new SponsoredProductsDraftCampaignMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignMutationError(errorType string, errorValue SponsoredProductsDraftCampaignMutationErrorSelector) *SponsoredProductsDraftCampaignMutationError {
	this := SponsoredProductsDraftCampaignMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftCampaignMutationErrorWithDefaults instantiates a new SponsoredProductsDraftCampaignMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignMutationErrorWithDefaults() *SponsoredProductsDraftCampaignMutationError {
	this := SponsoredProductsDraftCampaignMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftCampaignMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftCampaignMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftCampaignMutationError) GetErrorValue() SponsoredProductsDraftCampaignMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftCampaignMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationError) GetErrorValueOk() (*SponsoredProductsDraftCampaignMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftCampaignMutationError) SetErrorValue(v SponsoredProductsDraftCampaignMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftCampaignMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignMutationError struct {
	value *SponsoredProductsDraftCampaignMutationError
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignMutationError) Get() *SponsoredProductsDraftCampaignMutationError {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignMutationError) Set(val *SponsoredProductsDraftCampaignMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignMutationError(val *SponsoredProductsDraftCampaignMutationError) *NullableSponsoredProductsDraftCampaignMutationError {
	return &NullableSponsoredProductsDraftCampaignMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
