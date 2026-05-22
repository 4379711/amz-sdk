package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the EntityFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityFilter{}

// EntityFilter Filter optimization rules by entityId and entityType
type EntityFilter struct {
	// Enum: \"CAMPAIGN\"  The type of entity passed.
	EntityType *string `json:"entityType,omitempty"`
	// Entity object identifier.
	EntityId *string `json:"entityId,omitempty"`
}

// NewEntityFilter instantiates a new EntityFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityFilter() *EntityFilter {
	this := EntityFilter{}
	return &this
}

// NewEntityFilterWithDefaults instantiates a new EntityFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityFilterWithDefaults() *EntityFilter {
	this := EntityFilter{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *EntityFilter) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityFilter) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *EntityFilter) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *EntityFilter) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *EntityFilter) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityFilter) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *EntityFilter) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *EntityFilter) SetEntityId(v string) {
	o.EntityId = &v
}

func (o EntityFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

type NullableEntityFilter struct {
	value *EntityFilter
	isSet bool
}

func (v NullableEntityFilter) Get() *EntityFilter {
	return v.value
}

func (v *NullableEntityFilter) Set(val *EntityFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityFilter(val *EntityFilter) *NullableEntityFilter {
	return &NullableEntityFilter{value: val, isSet: true}
}

func (v NullableEntityFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEntityFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
