package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignPromotionError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignPromotionError{}

// SponsoredProductsDraftCampaignPromotionError struct for SponsoredProductsDraftCampaignPromotionError
type SponsoredProductsDraftCampaignPromotionError struct {
	// The type of the error
	ErrorType  string                                               `json:"errorType"`
	ErrorValue SponsoredProductsDraftCampaignPromotionErrorSelector `json:"errorValue"`
}

type _SponsoredProductsDraftCampaignPromotionError SponsoredProductsDraftCampaignPromotionError

// NewSponsoredProductsDraftCampaignPromotionError instantiates a new SponsoredProductsDraftCampaignPromotionError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignPromotionError(errorType string, errorValue SponsoredProductsDraftCampaignPromotionErrorSelector) *SponsoredProductsDraftCampaignPromotionError {
	this := SponsoredProductsDraftCampaignPromotionError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsDraftCampaignPromotionErrorWithDefaults instantiates a new SponsoredProductsDraftCampaignPromotionError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignPromotionErrorWithDefaults() *SponsoredProductsDraftCampaignPromotionError {
	this := SponsoredProductsDraftCampaignPromotionError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsDraftCampaignPromotionError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignPromotionError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsDraftCampaignPromotionError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsDraftCampaignPromotionError) GetErrorValue() SponsoredProductsDraftCampaignPromotionErrorSelector {
	if o == nil {
		var ret SponsoredProductsDraftCampaignPromotionErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignPromotionError) GetErrorValueOk() (*SponsoredProductsDraftCampaignPromotionErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsDraftCampaignPromotionError) SetErrorValue(v SponsoredProductsDraftCampaignPromotionErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsDraftCampaignPromotionError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignPromotionError struct {
	value *SponsoredProductsDraftCampaignPromotionError
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignPromotionError) Get() *SponsoredProductsDraftCampaignPromotionError {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignPromotionError) Set(val *SponsoredProductsDraftCampaignPromotionError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignPromotionError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignPromotionError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignPromotionError(val *SponsoredProductsDraftCampaignPromotionError) *NullableSponsoredProductsDraftCampaignPromotionError {
	return &NullableSponsoredProductsDraftCampaignPromotionError{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignPromotionError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignPromotionError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
