package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the RequiredRecommendations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequiredRecommendations{}

// RequiredRecommendations struct for RequiredRecommendations
type RequiredRecommendations struct {
	// Maximum number of recommendations groups that API should return for given type. (recommendations are not guaranteed).
	MaxRecommendationGroups *int32 `json:"maxRecommendationGroups,omitempty"`
	// Type of recommendations.
	Type string `json:"type"`
}

type _RequiredRecommendations RequiredRecommendations

// NewRequiredRecommendations instantiates a new RequiredRecommendations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequiredRecommendations(type_ string) *RequiredRecommendations {
	this := RequiredRecommendations{}
	this.Type = type_
	return &this
}

// NewRequiredRecommendationsWithDefaults instantiates a new RequiredRecommendations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequiredRecommendationsWithDefaults() *RequiredRecommendations {
	this := RequiredRecommendations{}
	return &this
}

// GetMaxRecommendationGroups returns the MaxRecommendationGroups field value if set, zero value otherwise.
func (o *RequiredRecommendations) GetMaxRecommendationGroups() int32 {
	if o == nil || IsNil(o.MaxRecommendationGroups) {
		var ret int32
		return ret
	}
	return *o.MaxRecommendationGroups
}

// GetMaxRecommendationGroupsOk returns a tuple with the MaxRecommendationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredRecommendations) GetMaxRecommendationGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRecommendationGroups) {
		return nil, false
	}
	return o.MaxRecommendationGroups, true
}

// HasMaxRecommendationGroups returns a boolean if a field has been set.
func (o *RequiredRecommendations) HasMaxRecommendationGroups() bool {
	if o != nil && !IsNil(o.MaxRecommendationGroups) {
		return true
	}

	return false
}

// SetMaxRecommendationGroups gets a reference to the given int32 and assigns it to the MaxRecommendationGroups field.
func (o *RequiredRecommendations) SetMaxRecommendationGroups(v int32) {
	o.MaxRecommendationGroups = &v
}

// GetType returns the Type field value
func (o *RequiredRecommendations) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequiredRecommendations) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequiredRecommendations) SetType(v string) {
	o.Type = v
}

func (o RequiredRecommendations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxRecommendationGroups) {
		toSerialize["maxRecommendationGroups"] = o.MaxRecommendationGroups
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableRequiredRecommendations struct {
	value *RequiredRecommendations
	isSet bool
}

func (v NullableRequiredRecommendations) Get() *RequiredRecommendations {
	return v.value
}

func (v *NullableRequiredRecommendations) Set(val *RequiredRecommendations) {
	v.value = val
	v.isSet = true
}

func (v NullableRequiredRecommendations) IsSet() bool {
	return v.isSet
}

func (v *NullableRequiredRecommendations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequiredRecommendations(val *RequiredRecommendations) *NullableRequiredRecommendations {
	return &NullableRequiredRecommendations{value: val, isSet: true}
}

func (v NullableRequiredRecommendations) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRequiredRecommendations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
