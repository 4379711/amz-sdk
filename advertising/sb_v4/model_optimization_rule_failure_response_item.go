package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRuleFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRuleFailureResponseItem{}

// OptimizationRuleFailureResponseItem struct for OptimizationRuleFailureResponseItem
type OptimizationRuleFailureResponseItem struct {
	// the index of the optimization rule id/entity Id in the array from the request body.
	Index float32 `json:"index"`
	// A list of validation errors
	Errors []OptimizationRulesError `json:"errors,omitempty"`
}

type _OptimizationRuleFailureResponseItem OptimizationRuleFailureResponseItem

// NewOptimizationRuleFailureResponseItem instantiates a new OptimizationRuleFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRuleFailureResponseItem(index float32) *OptimizationRuleFailureResponseItem {
	this := OptimizationRuleFailureResponseItem{}
	this.Index = index
	return &this
}

// NewOptimizationRuleFailureResponseItemWithDefaults instantiates a new OptimizationRuleFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRuleFailureResponseItemWithDefaults() *OptimizationRuleFailureResponseItem {
	this := OptimizationRuleFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *OptimizationRuleFailureResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRuleFailureResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRuleFailureResponseItem) SetIndex(v float32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *OptimizationRuleFailureResponseItem) GetErrors() []OptimizationRulesError {
	if o == nil || IsNil(o.Errors) {
		var ret []OptimizationRulesError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRuleFailureResponseItem) GetErrorsOk() ([]OptimizationRulesError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *OptimizationRuleFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []OptimizationRulesError and assigns it to the Errors field.
func (o *OptimizationRuleFailureResponseItem) SetErrors(v []OptimizationRulesError) {
	o.Errors = v
}

func (o OptimizationRuleFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableOptimizationRuleFailureResponseItem struct {
	value *OptimizationRuleFailureResponseItem
	isSet bool
}

func (v NullableOptimizationRuleFailureResponseItem) Get() *OptimizationRuleFailureResponseItem {
	return v.value
}

func (v *NullableOptimizationRuleFailureResponseItem) Set(val *OptimizationRuleFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRuleFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRuleFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRuleFailureResponseItem(val *OptimizationRuleFailureResponseItem) *NullableOptimizationRuleFailureResponseItem {
	return &NullableOptimizationRuleFailureResponseItem{value: val, isSet: true}
}

func (v NullableOptimizationRuleFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRuleFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
