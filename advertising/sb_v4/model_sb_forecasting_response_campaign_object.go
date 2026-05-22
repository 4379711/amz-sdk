package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingResponseCampaignObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingResponseCampaignObject{}

// SBForecastingResponseCampaignObject struct for SBForecastingResponseCampaignObject
type SBForecastingResponseCampaignObject struct {
	Successes []SBForecastingSuccessObject `json:"successes,omitempty"`
	Errors    []SBForecastingErrorObject   `json:"errors,omitempty"`
}

// NewSBForecastingResponseCampaignObject instantiates a new SBForecastingResponseCampaignObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingResponseCampaignObject() *SBForecastingResponseCampaignObject {
	this := SBForecastingResponseCampaignObject{}
	return &this
}

// NewSBForecastingResponseCampaignObjectWithDefaults instantiates a new SBForecastingResponseCampaignObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingResponseCampaignObjectWithDefaults() *SBForecastingResponseCampaignObject {
	this := SBForecastingResponseCampaignObject{}
	return &this
}

// GetSuccesses returns the Successes field value if set, zero value otherwise.
func (o *SBForecastingResponseCampaignObject) GetSuccesses() []SBForecastingSuccessObject {
	if o == nil || IsNil(o.Successes) {
		var ret []SBForecastingSuccessObject
		return ret
	}
	return o.Successes
}

// GetSuccessesOk returns a tuple with the Successes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingResponseCampaignObject) GetSuccessesOk() ([]SBForecastingSuccessObject, bool) {
	if o == nil || IsNil(o.Successes) {
		return nil, false
	}
	return o.Successes, true
}

// HasSuccesses returns a boolean if a field has been set.
func (o *SBForecastingResponseCampaignObject) HasSuccesses() bool {
	if o != nil && !IsNil(o.Successes) {
		return true
	}

	return false
}

// SetSuccesses gets a reference to the given []SBForecastingSuccessObject and assigns it to the Successes field.
func (o *SBForecastingResponseCampaignObject) SetSuccesses(v []SBForecastingSuccessObject) {
	o.Successes = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SBForecastingResponseCampaignObject) GetErrors() []SBForecastingErrorObject {
	if o == nil || IsNil(o.Errors) {
		var ret []SBForecastingErrorObject
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingResponseCampaignObject) GetErrorsOk() ([]SBForecastingErrorObject, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SBForecastingResponseCampaignObject) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SBForecastingErrorObject and assigns it to the Errors field.
func (o *SBForecastingResponseCampaignObject) SetErrors(v []SBForecastingErrorObject) {
	o.Errors = v
}

func (o SBForecastingResponseCampaignObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Successes) {
		toSerialize["successes"] = o.Successes
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSBForecastingResponseCampaignObject struct {
	value *SBForecastingResponseCampaignObject
	isSet bool
}

func (v NullableSBForecastingResponseCampaignObject) Get() *SBForecastingResponseCampaignObject {
	return v.value
}

func (v *NullableSBForecastingResponseCampaignObject) Set(val *SBForecastingResponseCampaignObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingResponseCampaignObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingResponseCampaignObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingResponseCampaignObject(val *SBForecastingResponseCampaignObject) *NullableSBForecastingResponseCampaignObject {
	return &NullableSBForecastingResponseCampaignObject{value: val, isSet: true}
}

func (v NullableSBForecastingResponseCampaignObject) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingResponseCampaignObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
