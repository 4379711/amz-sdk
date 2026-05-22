package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem{}

// OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem struct for OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem
type OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem struct {
	// Index of the optimization rule in the request.
	Index int32 `json:"index"`
	// The id that uniquely identified the optimization rule that succeeded in deletion.
	OptimizationRuleId string `json:"optimizationRuleId"`
}

type _OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem

// NewOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem instantiates a new OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem(index int32, optimizationRuleId string) *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem {
	this := OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem{}
	this.Index = index
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItemWithDefaults instantiates a new OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItemWithDefaults() *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem {
	this := OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) SetIndex(v int32) {
	o.Index = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) GetOptimizationRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = v
}

func (o OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem struct {
	value *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem
	isSet bool
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) Get() *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) Set(val *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem(val *OptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) *NullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem {
	return &NullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesResponseSuccessItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
