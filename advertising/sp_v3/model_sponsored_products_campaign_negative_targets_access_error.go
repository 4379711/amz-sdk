package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeTargetsAccessError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeTargetsAccessError{}

// SponsoredProductsCampaignNegativeTargetsAccessError struct for SponsoredProductsCampaignNegativeTargetsAccessError
type SponsoredProductsCampaignNegativeTargetsAccessError struct {
	// The type of the error
	ErrorType  string                                                      `json:"errorType"`
	ErrorValue SponsoredProductsCampaignNegativeTargetsAccessErrorSelector `json:"errorValue"`
}

type _SponsoredProductsCampaignNegativeTargetsAccessError SponsoredProductsCampaignNegativeTargetsAccessError

// NewSponsoredProductsCampaignNegativeTargetsAccessError instantiates a new SponsoredProductsCampaignNegativeTargetsAccessError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeTargetsAccessError(errorType string, errorValue SponsoredProductsCampaignNegativeTargetsAccessErrorSelector) *SponsoredProductsCampaignNegativeTargetsAccessError {
	this := SponsoredProductsCampaignNegativeTargetsAccessError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsCampaignNegativeTargetsAccessErrorWithDefaults instantiates a new SponsoredProductsCampaignNegativeTargetsAccessError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeTargetsAccessErrorWithDefaults() *SponsoredProductsCampaignNegativeTargetsAccessError {
	this := SponsoredProductsCampaignNegativeTargetsAccessError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsCampaignNegativeTargetsAccessError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsAccessError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsCampaignNegativeTargetsAccessError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsCampaignNegativeTargetsAccessError) GetErrorValue() SponsoredProductsCampaignNegativeTargetsAccessErrorSelector {
	if o == nil {
		var ret SponsoredProductsCampaignNegativeTargetsAccessErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetsAccessError) GetErrorValueOk() (*SponsoredProductsCampaignNegativeTargetsAccessErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsCampaignNegativeTargetsAccessError) SetErrorValue(v SponsoredProductsCampaignNegativeTargetsAccessErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsCampaignNegativeTargetsAccessError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeTargetsAccessError struct {
	value *SponsoredProductsCampaignNegativeTargetsAccessError
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeTargetsAccessError) Get() *SponsoredProductsCampaignNegativeTargetsAccessError {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsAccessError) Set(val *SponsoredProductsCampaignNegativeTargetsAccessError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeTargetsAccessError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsAccessError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeTargetsAccessError(val *SponsoredProductsCampaignNegativeTargetsAccessError) *NullableSponsoredProductsCampaignNegativeTargetsAccessError {
	return &NullableSponsoredProductsCampaignNegativeTargetsAccessError{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeTargetsAccessError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeTargetsAccessError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
