package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem{}

// OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem struct for OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem
type OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem struct {
	EntityType OptimizationRulesAPIEntityType `json:"entityType"`
	// Index of the successful disassociation pair in the request.
	Index int32 `json:"index"`
	// The id that uniquely identifies the advertising entity that was disassociated from.
	EntityId string `json:"entityId"`
	// The id that uniquely identifies the optimization rule that was disassociated.
	OptimizationRuleId int32 `json:"optimizationRuleId"`
}

type _OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem

// NewOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem(entityType OptimizationRulesAPIEntityType, index int32, entityId string, optimizationRuleId int32) *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem {
	this := OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem{}
	this.EntityType = entityType
	this.Index = index
	this.EntityId = entityId
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItemWithDefaults instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItemWithDefaults() *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem {
	this := OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) GetEntityType() OptimizationRulesAPIEntityType {
	if o == nil {
		var ret OptimizationRulesAPIEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) GetEntityTypeOk() (*OptimizationRulesAPIEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) SetEntityType(v OptimizationRulesAPIEntityType) {
	o.EntityType = v
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) SetIndex(v int32) {
	o.Index = v
}

// GetEntityId returns the EntityId field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) SetEntityId(v string) {
	o.EntityId = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) GetOptimizationRuleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) GetOptimizationRuleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) SetOptimizationRuleId(v int32) {
	o.OptimizationRuleId = v
}

func (o OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityType"] = o.EntityType
	toSerialize["index"] = o.Index
	toSerialize["entityId"] = o.EntityId
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem struct {
	value *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem
	isSet bool
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) Get() *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) Set(val *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem(val *OptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem {
	return &NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseSuccessItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
