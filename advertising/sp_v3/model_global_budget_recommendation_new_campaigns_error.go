package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBudgetRecommendationNewCampaignsError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRecommendationNewCampaignsError{}

// GlobalBudgetRecommendationNewCampaignsError struct for GlobalBudgetRecommendationNewCampaignsError
type GlobalBudgetRecommendationNewCampaignsError struct {
	CountryCodes []string `json:"countryCodes,omitempty"`
	Code         string   `json:"code"`
	Message      *string  `json:"message,omitempty"`
}

type _GlobalBudgetRecommendationNewCampaignsError GlobalBudgetRecommendationNewCampaignsError

// NewGlobalBudgetRecommendationNewCampaignsError instantiates a new GlobalBudgetRecommendationNewCampaignsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRecommendationNewCampaignsError(code string) *GlobalBudgetRecommendationNewCampaignsError {
	this := GlobalBudgetRecommendationNewCampaignsError{}
	this.Code = code
	return &this
}

// NewGlobalBudgetRecommendationNewCampaignsErrorWithDefaults instantiates a new GlobalBudgetRecommendationNewCampaignsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRecommendationNewCampaignsErrorWithDefaults() *GlobalBudgetRecommendationNewCampaignsError {
	this := GlobalBudgetRecommendationNewCampaignsError{}
	return &this
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *GlobalBudgetRecommendationNewCampaignsError) GetCountryCodes() []string {
	if o == nil || IsNil(o.CountryCodes) {
		var ret []string
		return ret
	}
	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationNewCampaignsError) GetCountryCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *GlobalBudgetRecommendationNewCampaignsError) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given []string and assigns it to the CountryCodes field.
func (o *GlobalBudgetRecommendationNewCampaignsError) SetCountryCodes(v []string) {
	o.CountryCodes = v
}

// GetCode returns the Code field value
func (o *GlobalBudgetRecommendationNewCampaignsError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationNewCampaignsError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GlobalBudgetRecommendationNewCampaignsError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GlobalBudgetRecommendationNewCampaignsError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationNewCampaignsError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GlobalBudgetRecommendationNewCampaignsError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GlobalBudgetRecommendationNewCampaignsError) SetMessage(v string) {
	o.Message = &v
}

func (o GlobalBudgetRecommendationNewCampaignsError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCodes) {
		toSerialize["countryCodes"] = o.CountryCodes
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGlobalBudgetRecommendationNewCampaignsError struct {
	value *GlobalBudgetRecommendationNewCampaignsError
	isSet bool
}

func (v NullableGlobalBudgetRecommendationNewCampaignsError) Get() *GlobalBudgetRecommendationNewCampaignsError {
	return v.value
}

func (v *NullableGlobalBudgetRecommendationNewCampaignsError) Set(val *GlobalBudgetRecommendationNewCampaignsError) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRecommendationNewCampaignsError) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRecommendationNewCampaignsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRecommendationNewCampaignsError(val *GlobalBudgetRecommendationNewCampaignsError) *NullableGlobalBudgetRecommendationNewCampaignsError {
	return &NullableGlobalBudgetRecommendationNewCampaignsError{value: val, isSet: true}
}

func (v NullableGlobalBudgetRecommendationNewCampaignsError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRecommendationNewCampaignsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
