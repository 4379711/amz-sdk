package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateOptimizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOptimizationRule{}

// CreateOptimizationRule struct for CreateOptimizationRule
type CreateOptimizationRule struct {
	// Enum: \"CAMPAIGN\"  The type of entity passed.
	EntityType *string `json:"entityType,omitempty"`
	// Entity object identifier.
	EntityId   *string         `json:"entityId,omitempty"`
	Conditions []RuleCondition `json:"conditions,omitempty"`
}

// NewCreateOptimizationRule instantiates a new CreateOptimizationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOptimizationRule() *CreateOptimizationRule {
	this := CreateOptimizationRule{}
	return &this
}

// NewCreateOptimizationRuleWithDefaults instantiates a new CreateOptimizationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOptimizationRuleWithDefaults() *CreateOptimizationRule {
	this := CreateOptimizationRule{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *CreateOptimizationRule) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRule) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *CreateOptimizationRule) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *CreateOptimizationRule) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *CreateOptimizationRule) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRule) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *CreateOptimizationRule) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *CreateOptimizationRule) SetEntityId(v string) {
	o.EntityId = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *CreateOptimizationRule) GetConditions() []RuleCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []RuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOptimizationRule) GetConditionsOk() ([]RuleCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *CreateOptimizationRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []RuleCondition and assigns it to the Conditions field.
func (o *CreateOptimizationRule) SetConditions(v []RuleCondition) {
	o.Conditions = v
}

func (o CreateOptimizationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return toSerialize, nil
}

type NullableCreateOptimizationRule struct {
	value *CreateOptimizationRule
	isSet bool
}

func (v NullableCreateOptimizationRule) Get() *CreateOptimizationRule {
	return v.value
}

func (v *NullableCreateOptimizationRule) Set(val *CreateOptimizationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOptimizationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOptimizationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOptimizationRule(val *CreateOptimizationRule) *NullableCreateOptimizationRule {
	return &NullableCreateOptimizationRule{value: val, isSet: true}
}

func (v NullableCreateOptimizationRule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateOptimizationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
