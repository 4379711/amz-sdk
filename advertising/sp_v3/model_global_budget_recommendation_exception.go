package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBudgetRecommendationException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRecommendationException{}

// GlobalBudgetRecommendationException struct for GlobalBudgetRecommendationException
type GlobalBudgetRecommendationException struct {
	Errors []GlobalBudgetRecommendationError `json:"errors,omitempty"`
}

// NewGlobalBudgetRecommendationException instantiates a new GlobalBudgetRecommendationException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRecommendationException() *GlobalBudgetRecommendationException {
	this := GlobalBudgetRecommendationException{}
	return &this
}

// NewGlobalBudgetRecommendationExceptionWithDefaults instantiates a new GlobalBudgetRecommendationException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRecommendationExceptionWithDefaults() *GlobalBudgetRecommendationException {
	this := GlobalBudgetRecommendationException{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalBudgetRecommendationException) GetErrors() []GlobalBudgetRecommendationError {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalBudgetRecommendationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationException) GetErrorsOk() ([]GlobalBudgetRecommendationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalBudgetRecommendationException) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalBudgetRecommendationError and assigns it to the Errors field.
func (o *GlobalBudgetRecommendationException) SetErrors(v []GlobalBudgetRecommendationError) {
	o.Errors = v
}

func (o GlobalBudgetRecommendationException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalBudgetRecommendationException struct {
	value *GlobalBudgetRecommendationException
	isSet bool
}

func (v NullableGlobalBudgetRecommendationException) Get() *GlobalBudgetRecommendationException {
	return v.value
}

func (v *NullableGlobalBudgetRecommendationException) Set(val *GlobalBudgetRecommendationException) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRecommendationException) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRecommendationException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRecommendationException(val *GlobalBudgetRecommendationException) *NullableGlobalBudgetRecommendationException {
	return &NullableGlobalBudgetRecommendationException{value: val, isSet: true}
}

func (v NullableGlobalBudgetRecommendationException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRecommendationException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
