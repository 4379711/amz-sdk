package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRuleToEntityMappingSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRuleToEntityMappingSuccessResponseItem{}

// OptimizationRuleToEntityMappingSuccessResponseItem struct for OptimizationRuleToEntityMappingSuccessResponseItem
type OptimizationRuleToEntityMappingSuccessResponseItem struct {
	EntityType string `json:"entityType"`
	// The index of the entityId/optimizationId in the array from the request body.
	Index float32 `json:"index"`
	// Entity object identifier.
	EntityId string `json:"entityId"`
	// The identifier of the optimization rule.
	OptimizationRuleId string `json:"optimizationRuleId"`
}

type _OptimizationRuleToEntityMappingSuccessResponseItem OptimizationRuleToEntityMappingSuccessResponseItem

// NewOptimizationRuleToEntityMappingSuccessResponseItem instantiates a new OptimizationRuleToEntityMappingSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRuleToEntityMappingSuccessResponseItem(entityType string, index float32, entityId string, optimizationRuleId string) *OptimizationRuleToEntityMappingSuccessResponseItem {
	this := OptimizationRuleToEntityMappingSuccessResponseItem{}
	this.EntityType = entityType
	this.Index = index
	this.EntityId = entityId
	this.OptimizationRuleId = optimizationRuleId
	return &this
}

// NewOptimizationRuleToEntityMappingSuccessResponseItemWithDefaults instantiates a new OptimizationRuleToEntityMappingSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRuleToEntityMappingSuccessResponseItemWithDefaults() *OptimizationRuleToEntityMappingSuccessResponseItem {
	this := OptimizationRuleToEntityMappingSuccessResponseItem{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) SetEntityType(v string) {
	o.EntityType = v
}

// GetIndex returns the Index field value
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) SetIndex(v float32) {
	o.Index = v
}

// GetEntityId returns the EntityId field value
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) SetEntityId(v string) {
	o.EntityId = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) GetOptimizationRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRuleToEntityMappingSuccessResponseItem) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = v
}

func (o OptimizationRuleToEntityMappingSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityType"] = o.EntityType
	toSerialize["index"] = o.Index
	toSerialize["entityId"] = o.EntityId
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	return toSerialize, nil
}

type NullableOptimizationRuleToEntityMappingSuccessResponseItem struct {
	value *OptimizationRuleToEntityMappingSuccessResponseItem
	isSet bool
}

func (v NullableOptimizationRuleToEntityMappingSuccessResponseItem) Get() *OptimizationRuleToEntityMappingSuccessResponseItem {
	return v.value
}

func (v *NullableOptimizationRuleToEntityMappingSuccessResponseItem) Set(val *OptimizationRuleToEntityMappingSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRuleToEntityMappingSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRuleToEntityMappingSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRuleToEntityMappingSuccessResponseItem(val *OptimizationRuleToEntityMappingSuccessResponseItem) *NullableOptimizationRuleToEntityMappingSuccessResponseItem {
	return &NullableOptimizationRuleToEntityMappingSuccessResponseItem{value: val, isSet: true}
}

func (v NullableOptimizationRuleToEntityMappingSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRuleToEntityMappingSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
