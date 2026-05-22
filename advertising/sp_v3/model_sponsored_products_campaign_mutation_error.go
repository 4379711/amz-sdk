package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCampaignMutationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignMutationError{}

// SponsoredProductsCampaignMutationError struct for SponsoredProductsCampaignMutationError
type SponsoredProductsCampaignMutationError struct {
	// The type of the error
	ErrorType  string                                         `json:"errorType"`
	ErrorValue SponsoredProductsCampaignMutationErrorSelector `json:"errorValue"`
}

type _SponsoredProductsCampaignMutationError SponsoredProductsCampaignMutationError

// NewSponsoredProductsCampaignMutationError instantiates a new SponsoredProductsCampaignMutationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignMutationError(errorType string, errorValue SponsoredProductsCampaignMutationErrorSelector) *SponsoredProductsCampaignMutationError {
	this := SponsoredProductsCampaignMutationError{}
	this.ErrorType = errorType
	this.ErrorValue = errorValue
	return &this
}

// NewSponsoredProductsCampaignMutationErrorWithDefaults instantiates a new SponsoredProductsCampaignMutationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignMutationErrorWithDefaults() *SponsoredProductsCampaignMutationError {
	this := SponsoredProductsCampaignMutationError{}
	return &this
}

// GetErrorType returns the ErrorType field value
func (o *SponsoredProductsCampaignMutationError) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationError) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value
func (o *SponsoredProductsCampaignMutationError) SetErrorType(v string) {
	o.ErrorType = v
}

// GetErrorValue returns the ErrorValue field value
func (o *SponsoredProductsCampaignMutationError) GetErrorValue() SponsoredProductsCampaignMutationErrorSelector {
	if o == nil {
		var ret SponsoredProductsCampaignMutationErrorSelector
		return ret
	}

	return o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationError) GetErrorValueOk() (*SponsoredProductsCampaignMutationErrorSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorValue, true
}

// SetErrorValue sets field value
func (o *SponsoredProductsCampaignMutationError) SetErrorValue(v SponsoredProductsCampaignMutationErrorSelector) {
	o.ErrorValue = v
}

func (o SponsoredProductsCampaignMutationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["errorValue"] = o.ErrorValue
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignMutationError struct {
	value *SponsoredProductsCampaignMutationError
	isSet bool
}

func (v NullableSponsoredProductsCampaignMutationError) Get() *SponsoredProductsCampaignMutationError {
	return v.value
}

func (v *NullableSponsoredProductsCampaignMutationError) Set(val *SponsoredProductsCampaignMutationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignMutationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignMutationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignMutationError(val *SponsoredProductsCampaignMutationError) *NullableSponsoredProductsCampaignMutationError {
	return &NullableSponsoredProductsCampaignMutationError{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignMutationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignMutationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
