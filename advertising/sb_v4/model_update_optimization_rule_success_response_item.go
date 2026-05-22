package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateOptimizationRuleSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOptimizationRuleSuccessResponseItem{}

// UpdateOptimizationRuleSuccessResponseItem struct for UpdateOptimizationRuleSuccessResponseItem
type UpdateOptimizationRuleSuccessResponseItem struct {
	OptimizationRule OptimizationRule `json:"optimizationRule"`
	// The index of the entityId in the array from the request body.
	Index float32 `json:"index"`
	// The identifier of the optimization rule.
	OptimizationRuleId string `json:"optimizationRuleId"`
}

type _UpdateOptimizationRuleSuccessResponseItem UpdateOptimizationRuleSuccessResponseItem

// NewUpdateOptimizationRuleSuccessResponseItem instantiates a new UpdateOptimizationRuleSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOptimizationRuleSuccessResponseItem(optimizationRule OptimizationRule, index float32, optimizationRuleId string) *UpdateOptimizationRuleSuccessResponseItem {
	this := UpdateOptimizationRuleSuccessResponseItem{}
	this.OptimizationRule = optimizationRule
	this.Index = index
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewUpdateOptimizationRuleSuccessResponseItemWithDefaults instantiates a new UpdateOptimizationRuleSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOptimizationRuleSuccessResponseItemWithDefaults() *UpdateOptimizationRuleSuccessResponseItem {
	this := UpdateOptimizationRuleSuccessResponseItem{}
	return &this
}

// GetOptimizationRule returns the OptimizationRule field value
func (o *UpdateOptimizationRuleSuccessResponseItem) GetOptimizationRule() OptimizationRule {
	if o == nil {
		var ret OptimizationRule
		return ret
	}

	return o.OptimizationRule
}

// GetOptimizationRuleOk returns a tuple with the OptimizationRule field value
// and a boolean to check if the value has been set.
func (o *UpdateOptimizationRuleSuccessResponseItem) GetOptimizationRuleOk() (*OptimizationRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRule, true
}

// SetOptimizationRule sets field value
func (o *UpdateOptimizationRuleSuccessResponseItem) SetOptimizationRule(v OptimizationRule) {
	o.OptimizationRule = v
}

// GetIndex returns the Index field value
func (o *UpdateOptimizationRuleSuccessResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *UpdateOptimizationRuleSuccessResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *UpdateOptimizationRuleSuccessResponseItem) SetIndex(v float32) {
	o.Index = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *UpdateOptimizationRuleSuccessResponseItem) GetOptimizationRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *UpdateOptimizationRuleSuccessResponseItem) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *UpdateOptimizationRuleSuccessResponseItem) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = v
}

func (o UpdateOptimizationRuleSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRule"] = o.OptimizationRule
	toSerialize["index"] = o.Index
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableUpdateOptimizationRuleSuccessResponseItem struct {
	value *UpdateOptimizationRuleSuccessResponseItem
	isSet bool
}

func (v NullableUpdateOptimizationRuleSuccessResponseItem) Get() *UpdateOptimizationRuleSuccessResponseItem {
	return v.value
}

func (v *NullableUpdateOptimizationRuleSuccessResponseItem) Set(val *UpdateOptimizationRuleSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOptimizationRuleSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOptimizationRuleSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOptimizationRuleSuccessResponseItem(val *UpdateOptimizationRuleSuccessResponseItem) *NullableUpdateOptimizationRuleSuccessResponseItem {
	return &NullableUpdateOptimizationRuleSuccessResponseItem{value: val, isSet: true}
}

func (v NullableUpdateOptimizationRuleSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateOptimizationRuleSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
