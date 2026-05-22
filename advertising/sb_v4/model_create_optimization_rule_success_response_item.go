package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateOptimizationRuleSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOptimizationRuleSuccessResponseItem{}

// CreateOptimizationRuleSuccessResponseItem struct for CreateOptimizationRuleSuccessResponseItem
type CreateOptimizationRuleSuccessResponseItem struct {
	OptimizationRule OptimizationRule `json:"optimizationRule"`
	EntityType       string           `json:"entityType"`
	// The index of the entityId in the array from the request body.
	Index float32 `json:"index"`
	// Entity object identifier.
	EntityId string `json:"entityId"`
	// The identifier of the optimization rule.
	OptimizationRuleId string `json:"optimizationRuleId"`
}

type _CreateOptimizationRuleSuccessResponseItem CreateOptimizationRuleSuccessResponseItem

// NewCreateOptimizationRuleSuccessResponseItem instantiates a new CreateOptimizationRuleSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOptimizationRuleSuccessResponseItem(optimizationRule OptimizationRule, entityType string, index float32, entityId string, optimizationRuleId string) *CreateOptimizationRuleSuccessResponseItem {
	this := CreateOptimizationRuleSuccessResponseItem{}
	this.OptimizationRule = optimizationRule
	this.EntityType = entityType
	this.Index = index
	this.EntityId = entityId
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewCreateOptimizationRuleSuccessResponseItemWithDefaults instantiates a new CreateOptimizationRuleSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOptimizationRuleSuccessResponseItemWithDefaults() *CreateOptimizationRuleSuccessResponseItem {
	this := CreateOptimizationRuleSuccessResponseItem{}
	return &this
}

// GetOptimizationRule returns the OptimizationRule field value
func (o *CreateOptimizationRuleSuccessResponseItem) GetOptimizationRule() OptimizationRule {
	if o == nil {
		var ret OptimizationRule
		return ret
	}

	return o.OptimizationRule
}

// GetOptimizationRuleOk returns a tuple with the OptimizationRule field value
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRuleSuccessResponseItem) GetOptimizationRuleOk() (*OptimizationRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRule, true
}

// SetOptimizationRule sets field value
func (o *CreateOptimizationRuleSuccessResponseItem) SetOptimizationRule(v OptimizationRule) {
	o.OptimizationRule = v
}

// GetEntityType returns the EntityType field value
func (o *CreateOptimizationRuleSuccessResponseItem) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRuleSuccessResponseItem) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *CreateOptimizationRuleSuccessResponseItem) SetEntityType(v string) {
	o.EntityType = v
}

// GetIndex returns the Index field value
func (o *CreateOptimizationRuleSuccessResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRuleSuccessResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *CreateOptimizationRuleSuccessResponseItem) SetIndex(v float32) {
	o.Index = v
}

// GetEntityId returns the EntityId field value
func (o *CreateOptimizationRuleSuccessResponseItem) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRuleSuccessResponseItem) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *CreateOptimizationRuleSuccessResponseItem) SetEntityId(v string) {
	o.EntityId = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *CreateOptimizationRuleSuccessResponseItem) GetOptimizationRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRuleSuccessResponseItem) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *CreateOptimizationRuleSuccessResponseItem) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = v
}

func (o CreateOptimizationRuleSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["optimizationRule"] = o.OptimizationRule
	toSerialize["entityType"] = o.EntityType
	toSerialize["index"] = o.Index
	toSerialize["entityId"] = o.EntityId
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableCreateOptimizationRuleSuccessResponseItem struct {
	value *CreateOptimizationRuleSuccessResponseItem
	isSet bool
}

func (v NullableCreateOptimizationRuleSuccessResponseItem) Get() *CreateOptimizationRuleSuccessResponseItem {
	return v.value
}

func (v *NullableCreateOptimizationRuleSuccessResponseItem) Set(val *CreateOptimizationRuleSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOptimizationRuleSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOptimizationRuleSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOptimizationRuleSuccessResponseItem(val *CreateOptimizationRuleSuccessResponseItem) *NullableCreateOptimizationRuleSuccessResponseItem {
	return &NullableCreateOptimizationRuleSuccessResponseItem{value: val, isSet: true}
}

func (v NullableCreateOptimizationRuleSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateOptimizationRuleSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
