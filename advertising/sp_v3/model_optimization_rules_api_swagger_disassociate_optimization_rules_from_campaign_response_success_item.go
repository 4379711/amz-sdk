package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem{}

// OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem struct for OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem
type OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem struct {
	// Index of the optimization rule in the request.
	Index int32 `json:"index"`
	// The id that uniquely identifies the optimization rule that was disassociated.
	OptimizationRuleId string `json:"optimizationRuleId"`
}

type _OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem

// NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem(index int32, optimizationRuleId string) *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem {
	this := OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem{}
	this.Index = index
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItemWithDefaults instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItemWithDefaults() *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem {
	this := OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) SetIndex(v int32) {
	o.Index = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) GetOptimizationRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = v
}

func (o OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem struct {
	value *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem
	isSet bool
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) Get() *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) Set(val *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem(val *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem {
	return &NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseSuccessItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
