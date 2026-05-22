package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDAudienceCategoryRecommendations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDAudienceCategoryRecommendations{}

// SDAudienceCategoryRecommendations List of recommended standard Amazon audience targets of a specific audience category
type SDAudienceCategoryRecommendations struct {
	Category *SDAudienceCategory `json:"category,omitempty"`
	// List of recommended standard Amazon audience targets
	Audiences []SDAudienceRecommendation `json:"audiences,omitempty"`
}

// NewSDAudienceCategoryRecommendations instantiates a new SDAudienceCategoryRecommendations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDAudienceCategoryRecommendations() *SDAudienceCategoryRecommendations {
	this := SDAudienceCategoryRecommendations{}
	return &this
}

// NewSDAudienceCategoryRecommendationsWithDefaults instantiates a new SDAudienceCategoryRecommendations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDAudienceCategoryRecommendationsWithDefaults() *SDAudienceCategoryRecommendations {
	this := SDAudienceCategoryRecommendations{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SDAudienceCategoryRecommendations) GetCategory() SDAudienceCategory {
	if o == nil || IsNil(o.Category) {
		var ret SDAudienceCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDAudienceCategoryRecommendations) GetCategoryOk() (*SDAudienceCategory, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SDAudienceCategoryRecommendations) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given SDAudienceCategory and assigns it to the Category field.
func (o *SDAudienceCategoryRecommendations) SetCategory(v SDAudienceCategory) {
	o.Category = &v
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *SDAudienceCategoryRecommendations) GetAudiences() []SDAudienceRecommendation {
	if o == nil || IsNil(o.Audiences) {
		var ret []SDAudienceRecommendation
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDAudienceCategoryRecommendations) GetAudiencesOk() ([]SDAudienceRecommendation, bool) {
	if o == nil || IsNil(o.Audiences) {
		return nil, false
	}
	return o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *SDAudienceCategoryRecommendations) HasAudiences() bool {
	if o != nil && !IsNil(o.Audiences) {
		return true
	}

	return false
}

// SetAudiences gets a reference to the given []SDAudienceRecommendation and assigns it to the Audiences field.
func (o *SDAudienceCategoryRecommendations) SetAudiences(v []SDAudienceRecommendation) {
	o.Audiences = v
}

func (o SDAudienceCategoryRecommendations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Audiences) {
		toSerialize["audiences"] = o.Audiences
	}
	return toSerialize, nil
}

type NullableSDAudienceCategoryRecommendations struct {
	value *SDAudienceCategoryRecommendations
	isSet bool
}

func (v NullableSDAudienceCategoryRecommendations) Get() *SDAudienceCategoryRecommendations {
	return v.value
}

func (v *NullableSDAudienceCategoryRecommendations) Set(val *SDAudienceCategoryRecommendations) {
	v.value = val
	v.isSet = true
}

func (v NullableSDAudienceCategoryRecommendations) IsSet() bool {
	return v.isSet
}

func (v *NullableSDAudienceCategoryRecommendations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDAudienceCategoryRecommendations(val *SDAudienceCategoryRecommendations) *NullableSDAudienceCategoryRecommendations {
	return &NullableSDAudienceCategoryRecommendations{value: val, isSet: true}
}

func (v NullableSDAudienceCategoryRecommendations) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDAudienceCategoryRecommendations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
