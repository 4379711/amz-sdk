package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignAccessError{}

// SponsoredProductsDraftCampaignAccessError struct for SponsoredProductsDraftCampaignAccessError
type SponsoredProductsDraftCampaignAccessError struct {
	// The type of the error
	ErrorType  string                                            `json:"errorType"`
	ErrorValue SponsoredProductsDraftCampaignAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftCampaignAccessError SponsoredProductsDraftCampaignAccessError

// NewSponsoredProductsDraftCampaignAccessError instantiates a new SponsoredProductsDraftCampaignAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignAccessError(errorType string, errorValue SponsoredProductsDraftCampaignAccessErrorSelector) *SponsoredProductsDraftCampaignAccessError {
	this := SponsoredProductsDraftCampaignAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftCampaignAccessErrorWithDefaults instantiates a new SponsoredProductsDraftCampaignAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignAccessErrorWithDefaults() *SponsoredProductsDraftCampaignAccessError {
	this := SponsoredProductsDraftCampaignAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftCampaignAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftCampaignAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftCampaignAccessError) GetErrorValue() SponsoredProductsDraftCampaignAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftCampaignAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignAccessError) GetErrorValueOk() (*SponsoredProductsDraftCampaignAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftCampaignAccessError) SetErrorValue(v SponsoredProductsDraftCampaignAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftCampaignAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignAccessError struct {
	value *SponsoredProductsDraftCampaignAccessError
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignAccessError) Get() *SponsoredProductsDraftCampaignAccessError {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignAccessError) Set(val *SponsoredProductsDraftCampaignAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignAccessError(val *SponsoredProductsDraftCampaignAccessError) *NullableSponsoredProductsDraftCampaignAccessError {
	return &NullableSponsoredProductsDraftCampaignAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
