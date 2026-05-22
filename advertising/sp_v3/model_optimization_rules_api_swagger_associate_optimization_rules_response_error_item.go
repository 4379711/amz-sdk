package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem{}

// OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem struct for OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem
type OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem struct {
	// Index of the failed association pair in the request.
	Index  int32                                       `json:"index"`
	Errors []OptimizationRulesAPIOptimizationRuleError `json:"errors"`
}

type _OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem

// NewOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem instantiates a new OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem(index int32, errors []OptimizationRulesAPIOptimizationRuleError) *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem {
	this := OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem{}
	this.Index = index
	this.Errors = errors
	return &this
}

// NewOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItemWithDefaults instantiates a new OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItemWithDefaults() *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem {
	this := OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) GetErrors() []OptimizationRulesAPIOptimizationRuleError {
	if o == nil {
		var ret []OptimizationRulesAPIOptimizationRuleError
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) GetErrorsOk() ([]OptimizationRulesAPIOptimizationRuleError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) SetErrors(v []OptimizationRulesAPIOptimizationRuleError) {
	o.Errors = v
}

func (o OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem struct {
	value *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem
	isSet bool
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) Get() *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) Set(val *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem(val *OptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) *NullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem {
	return &NullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesResponseErrorItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
