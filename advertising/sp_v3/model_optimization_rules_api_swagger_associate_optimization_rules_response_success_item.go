package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem{}

// OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem struct for OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem
type OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem struct {
	EntityType OptimizationRulesAPIEntityType `json:"entityType"`
	// Index of the successful association pair in the request.
	Index int32 `json:"index"`
	// The id that uniquely identifies the advertising entity that was associated to.
	EntityId string `json:"entityId"`
	// The id that uniquely identifies the optimization rule that was associated.
	OptimizationRuleId int32 `json:"optimizationRuleId"`
}

type _OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem

// NewOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem instantiates a new OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem(entityType OptimizationRulesAPIEntityType, index int32, entityId string, optimizationRuleId int32) *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem {
	this := OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem{}
	this.EntityType = entityType
	this.Index = index
	this.EntityId = entityId
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItemWithDefaults instantiates a new OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItemWithDefaults() *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem {
	this := OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) GetEntityType() OptimizationRulesAPIEntityType {
	if o == nil {
		var ret OptimizationRulesAPIEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) GetEntityTypeOk() (*OptimizationRulesAPIEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) SetEntityType(v OptimizationRulesAPIEntityType) {
	o.EntityType = v
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) SetIndex(v int32) {
	o.Index = v
}

// GetEntityId returns the EntityId field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) SetEntityId(v string) {
	o.EntityId = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) GetOptimizationRuleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) GetOptimizationRuleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) SetOptimizationRuleId(v int32) {
	o.OptimizationRuleId = v
}

func (o OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityType"] = o.EntityType
	toSerialize["index"] = o.Index
	toSerialize["entityId"] = o.EntityId
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem struct {
	value *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem
	isSet bool
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) Get() *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) Set(val *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem(val *OptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) *NullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem {
	return &NullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesResponseSuccessItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
