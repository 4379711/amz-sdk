package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRuleToEntityMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRuleToEntityMapping{}

// OptimizationRuleToEntityMapping struct for OptimizationRuleToEntityMapping
type OptimizationRuleToEntityMapping struct {
	// Enum: \"CAMPAIGN\"  The type of entity passed.
	EntityType string `json:"entityType"`
	// Entity object identifier.
	EntityId string `json:"entityId"`
	// The identifier of the optimization rule.
	OptimizationRuleId string `json:"optimizationRuleId"`
}

type _OptimizationRuleToEntityMapping OptimizationRuleToEntityMapping

// NewOptimizationRuleToEntityMapping instantiates a new OptimizationRuleToEntityMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRuleToEntityMapping(entityType string, entityId string, optimizationRuleId string) *OptimizationRuleToEntityMapping {
	this := OptimizationRuleToEntityMapping{}
	this.EntityType = entityType
	this.EntityId = entityId
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewOptimizationRuleToEntityMappingWithDefaults instantiates a new OptimizationRuleToEntityMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRuleToEntityMappingWithDefaults() *OptimizationRuleToEntityMapping {
	this := OptimizationRuleToEntityMapping{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *OptimizationRuleToEntityMapping) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *OptimizationRuleToEntityMapping) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *OptimizationRuleToEntityMapping) SetEntityType(v string) {
	o.EntityType = v
}

// GetEntityId returns the EntityId field value
func (o *OptimizationRuleToEntityMapping) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRuleToEntityMapping) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *OptimizationRuleToEntityMapping) SetEntityId(v string) {
	o.EntityId = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRuleToEntityMapping) GetOptimizationRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRuleToEntityMapping) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRuleToEntityMapping) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = v
}

func (o OptimizationRuleToEntityMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityType"] = o.EntityType
	toSerialize["entityId"] = o.EntityId
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableOptimizationRuleToEntityMapping struct {
	value *OptimizationRuleToEntityMapping
	isSet bool
}

func (v NullableOptimizationRuleToEntityMapping) Get() *OptimizationRuleToEntityMapping {
	return v.value
}

func (v *NullableOptimizationRuleToEntityMapping) Set(val *OptimizationRuleToEntityMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRuleToEntityMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRuleToEntityMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRuleToEntityMapping(val *OptimizationRuleToEntityMapping) *NullableOptimizationRuleToEntityMapping {
	return &NullableOptimizationRuleToEntityMapping{value: val, isSet: true}
}

func (v NullableOptimizationRuleToEntityMapping) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRuleToEntityMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
