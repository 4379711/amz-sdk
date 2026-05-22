package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBudgetRulesRecommendationEventException type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRulesRecommendationEventException{}

// GlobalBudgetRulesRecommendationEventException struct for GlobalBudgetRulesRecommendationEventException
type GlobalBudgetRulesRecommendationEventException struct {
	Errors []GlobalBudgetRulesRecommendationError `json:"errors,omitempty"`
}

// NewGlobalBudgetRulesRecommendationEventException instantiates a new GlobalBudgetRulesRecommendationEventException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRulesRecommendationEventException() *GlobalBudgetRulesRecommendationEventException {
	this := GlobalBudgetRulesRecommendationEventException{}
	return &this
}

// NewGlobalBudgetRulesRecommendationEventExceptionWithDefaults instantiates a new GlobalBudgetRulesRecommendationEventException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRulesRecommendationEventExceptionWithDefaults() *GlobalBudgetRulesRecommendationEventException {
	this := GlobalBudgetRulesRecommendationEventException{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalBudgetRulesRecommendationEventException) GetErrors() []GlobalBudgetRulesRecommendationError {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalBudgetRulesRecommendationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRulesRecommendationEventException) GetErrorsOk() ([]GlobalBudgetRulesRecommendationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalBudgetRulesRecommendationEventException) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalBudgetRulesRecommendationError and assigns it to the Errors field.
func (o *GlobalBudgetRulesRecommendationEventException) SetErrors(v []GlobalBudgetRulesRecommendationError) {
	o.Errors = v
}

func (o GlobalBudgetRulesRecommendationEventException) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalBudgetRulesRecommendationEventException struct {
	value *GlobalBudgetRulesRecommendationEventException
	isSet bool
}

func (v NullableGlobalBudgetRulesRecommendationEventException) Get() *GlobalBudgetRulesRecommendationEventException {
	return v.value
}

func (v *NullableGlobalBudgetRulesRecommendationEventException) Set(val *GlobalBudgetRulesRecommendationEventException) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRulesRecommendationEventException) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRulesRecommendationEventException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRulesRecommendationEventException(val *GlobalBudgetRulesRecommendationEventException) *NullableGlobalBudgetRulesRecommendationEventException {
	return &NullableGlobalBudgetRulesRecommendationEventException{value: val, isSet: true}
}

func (v NullableGlobalBudgetRulesRecommendationEventException) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRulesRecommendationEventException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
