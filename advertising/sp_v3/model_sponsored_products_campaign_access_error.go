package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignAccessError{}

// SponsoredProductsCampaignAccessError struct for SponsoredProductsCampaignAccessError
type SponsoredProductsCampaignAccessError struct {
	// The type of the error
	ErrorType  string                                       `json:"errorType"`
	ErrorValue SponsoredProductsCampaignAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsCampaignAccessError SponsoredProductsCampaignAccessError

// NewSponsoredProductsCampaignAccessError instantiates a new SponsoredProductsCampaignAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignAccessError(errorType string, errorValue SponsoredProductsCampaignAccessErrorSelector) *SponsoredProductsCampaignAccessError {
	this := SponsoredProductsCampaignAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsCampaignAccessErrorWithDefaults instantiates a new SponsoredProductsCampaignAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignAccessErrorWithDefaults() *SponsoredProductsCampaignAccessError {
	this := SponsoredProductsCampaignAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsCampaignAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsCampaignAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsCampaignAccessError) GetErrorValue() SponsoredProductsCampaignAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsCampaignAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignAccessError) GetErrorValueOk() (*SponsoredProductsCampaignAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsCampaignAccessError) SetErrorValue(v SponsoredProductsCampaignAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsCampaignAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignAccessError struct {
	value *SponsoredProductsCampaignAccessError
	isSet bool
}

func (v NullableSponsoredProductsCampaignAccessError) Get() *SponsoredProductsCampaignAccessError {
	return v.value
}

func (v *NullableSponsoredProductsCampaignAccessError) Set(val *SponsoredProductsCampaignAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignAccessError(val *SponsoredProductsCampaignAccessError) *NullableSponsoredProductsCampaignAccessError {
	return &NullableSponsoredProductsCampaignAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
