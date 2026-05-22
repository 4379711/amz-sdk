package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem{}

// OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem struct for OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem
type OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem struct {
	EntityType OptimizationRulesAPIEntityType `json:"entityType"`
	// The name of the advertising entity in the association.
	EntityName string `json:"entityName"`
	// The id that uniquely identifies the advertising entity in the association.
	EntityId string `json:"entityId"`
	// The id that uniquely identifies the optimization rule in the association.
	OptimizationRuleId int32 `json:"optimizationRuleId"`
}

type _OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem

// NewOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem instantiates a new OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem(entityType OptimizationRulesAPIEntityType, entityName string, entityId string, optimizationRuleId int32) *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem {
	this := OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem{}
	this.EntityType = entityType
	this.EntityName = entityName
	this.EntityId = entityId
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItemWithDefaults instantiates a new OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItemWithDefaults() *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem {
	this := OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) GetEntityType() OptimizationRulesAPIEntityType {
	if o == nil {
		var ret OptimizationRulesAPIEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) GetEntityTypeOk() (*OptimizationRulesAPIEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) SetEntityType(v OptimizationRulesAPIEntityType) {
	o.EntityType = v
}

// GetEntityName returns the EntityName field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) GetEntityName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) GetEntityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityName, true
}

// SetEntityName sets field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) SetEntityName(v string) {
	o.EntityName = v
}

// GetEntityId returns the EntityId field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) SetEntityId(v string) {
	o.EntityId = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) GetOptimizationRuleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) GetOptimizationRuleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) SetOptimizationRuleId(v int32) {
	o.OptimizationRuleId = v
}

func (o OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityType"] = o.EntityType
	toSerialize["entityName"] = o.EntityName
	toSerialize["entityId"] = o.EntityId
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem struct {
	value *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem
	isSet bool
}

func (v NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) Get() *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem {
	return v.value
}

func (v *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) Set(val *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem(val *OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem {
	return &NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
