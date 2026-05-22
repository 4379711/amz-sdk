package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIRuleAssociationDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIRuleAssociationDefinition{}

// OptimizationRulesAPIRuleAssociationDefinition struct for OptimizationRulesAPIRuleAssociationDefinition
type OptimizationRulesAPIRuleAssociationDefinition struct {
	EntityType OptimizationRulesAPIEntityType `json:"entityType"`
	// The identifier of an advertising entity. Its type is defined in the entityType field.
	EntityId string `json:"entityId"`
	// The rule identifier.
	OptimizationRuleId int32 `json:"optimizationRuleId"`
}

type _OptimizationRulesAPIRuleAssociationDefinition OptimizationRulesAPIRuleAssociationDefinition

// NewOptimizationRulesAPIRuleAssociationDefinition instantiates a new OptimizationRulesAPIRuleAssociationDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIRuleAssociationDefinition(entityType OptimizationRulesAPIEntityType, entityId string, optimizationRuleId int32) *OptimizationRulesAPIRuleAssociationDefinition {
	this := OptimizationRulesAPIRuleAssociationDefinition{}
	this.EntityType = entityType
	this.EntityId = entityId
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewOptimizationRulesAPIRuleAssociationDefinitionWithDefaults instantiates a new OptimizationRulesAPIRuleAssociationDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIRuleAssociationDefinitionWithDefaults() *OptimizationRulesAPIRuleAssociationDefinition {
	this := OptimizationRulesAPIRuleAssociationDefinition{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *OptimizationRulesAPIRuleAssociationDefinition) GetEntityType() OptimizationRulesAPIEntityType {
	if o == nil {
		var ret OptimizationRulesAPIEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleAssociationDefinition) GetEntityTypeOk() (*OptimizationRulesAPIEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *OptimizationRulesAPIRuleAssociationDefinition) SetEntityType(v OptimizationRulesAPIEntityType) {
	o.EntityType = v
}

// GetEntityId returns the EntityId field value
func (o *OptimizationRulesAPIRuleAssociationDefinition) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleAssociationDefinition) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *OptimizationRulesAPIRuleAssociationDefinition) SetEntityId(v string) {
	o.EntityId = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRulesAPIRuleAssociationDefinition) GetOptimizationRuleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIRuleAssociationDefinition) GetOptimizationRuleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRulesAPIRuleAssociationDefinition) SetOptimizationRuleId(v int32) {
	o.OptimizationRuleId = v
}

func (o OptimizationRulesAPIRuleAssociationDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityType"] = o.EntityType
	toSerialize["entityId"] = o.EntityId
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableOptimizationRulesAPIRuleAssociationDefinition struct {
	value *OptimizationRulesAPIRuleAssociationDefinition
	isSet bool
}

func (v NullableOptimizationRulesAPIRuleAssociationDefinition) Get() *OptimizationRulesAPIRuleAssociationDefinition {
	return v.value
}

func (v *NullableOptimizationRulesAPIRuleAssociationDefinition) Set(val *OptimizationRulesAPIRuleAssociationDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIRuleAssociationDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIRuleAssociationDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIRuleAssociationDefinition(val *OptimizationRulesAPIRuleAssociationDefinition) *NullableOptimizationRulesAPIRuleAssociationDefinition {
	return &NullableOptimizationRulesAPIRuleAssociationDefinition{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIRuleAssociationDefinition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIRuleAssociationDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
