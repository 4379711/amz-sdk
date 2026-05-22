package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBudgetRecommendationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRecommendationError{}

// GlobalBudgetRecommendationError struct for GlobalBudgetRecommendationError
type GlobalBudgetRecommendationError struct {
	CountryCodes []string `json:"countryCodes,omitempty"`
	Code         string   `json:"code"`
	Message      *string  `json:"message,omitempty"`
}

type _GlobalBudgetRecommendationError GlobalBudgetRecommendationError

// NewGlobalBudgetRecommendationError instantiates a new GlobalBudgetRecommendationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRecommendationError(code string) *GlobalBudgetRecommendationError {
	this := GlobalBudgetRecommendationError{}
	this.Code = code
	return &this
}

// NewGlobalBudgetRecommendationErrorWithDefaults instantiates a new GlobalBudgetRecommendationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRecommendationErrorWithDefaults() *GlobalBudgetRecommendationError {
	this := GlobalBudgetRecommendationError{}
	return &this
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *GlobalBudgetRecommendationError) GetCountryCodes() []string {
	if o == nil || IsNil(o.CountryCodes) {
		var ret []string
		return ret
	}
	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationError) GetCountryCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *GlobalBudgetRecommendationError) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given []string and assigns it to the CountryCodes field.
func (o *GlobalBudgetRecommendationError) SetCountryCodes(v []string) {
	o.CountryCodes = v
}

// GetCode returns the Code field value
func (o *GlobalBudgetRecommendationError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GlobalBudgetRecommendationError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GlobalBudgetRecommendationError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GlobalBudgetRecommendationError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GlobalBudgetRecommendationError) SetMessage(v string) {
	o.Message = &v
}

func (o GlobalBudgetRecommendationError) ToMap() (map[string]interface{}, error) {
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

type NullableGlobalBudgetRecommendationError struct {
	value *GlobalBudgetRecommendationError
	isSet bool
}

func (v NullableGlobalBudgetRecommendationError) Get() *GlobalBudgetRecommendationError {
	return v.value
}

func (v *NullableGlobalBudgetRecommendationError) Set(val *GlobalBudgetRecommendationError) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRecommendationError) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRecommendationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRecommendationError(val *GlobalBudgetRecommendationError) *NullableGlobalBudgetRecommendationError {
	return &NullableGlobalBudgetRecommendationError{value: val, isSet: true}
}

func (v NullableGlobalBudgetRecommendationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRecommendationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
