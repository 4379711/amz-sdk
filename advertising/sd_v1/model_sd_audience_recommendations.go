package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDAudienceRecommendations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDAudienceRecommendations{}

// SDAudienceRecommendations struct for SDAudienceRecommendations
type SDAudienceRecommendations struct {
	// List of recommended audience targets, broken down by audience category
	Audiences []SDAudienceCategoryRecommendations `json:"audiences,omitempty"`
}

// NewSDAudienceRecommendations instantiates a new SDAudienceRecommendations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDAudienceRecommendations() *SDAudienceRecommendations {
	this := SDAudienceRecommendations{}
	return &this
}

// NewSDAudienceRecommendationsWithDefaults instantiates a new SDAudienceRecommendations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDAudienceRecommendationsWithDefaults() *SDAudienceRecommendations {
	this := SDAudienceRecommendations{}
	return &this
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *SDAudienceRecommendations) GetAudiences() []SDAudienceCategoryRecommendations {
	if o == nil || IsNil(o.Audiences) {
		var ret []SDAudienceCategoryRecommendations
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDAudienceRecommendations) GetAudiencesOk() ([]SDAudienceCategoryRecommendations, bool) {
	if o == nil || IsNil(o.Audiences) {
		return nil, false
	}
	return o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *SDAudienceRecommendations) HasAudiences() bool {
	if o != nil && !IsNil(o.Audiences) {
		return true
	}

	return false
}

// SetAudiences gets a reference to the given []SDAudienceCategoryRecommendations and assigns it to the Audiences field.
func (o *SDAudienceRecommendations) SetAudiences(v []SDAudienceCategoryRecommendations) {
	o.Audiences = v
}

func (o SDAudienceRecommendations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audiences) {
		toSerialize["audiences"] = o.Audiences
	}
	return toSerialize, nil
}

type NullableSDAudienceRecommendations struct {
	value *SDAudienceRecommendations
	isSet bool
}

func (v NullableSDAudienceRecommendations) Get() *SDAudienceRecommendations {
	return v.value
}

func (v *NullableSDAudienceRecommendations) Set(val *SDAudienceRecommendations) {
	v.value = val
	v.isSet = true
}

func (v NullableSDAudienceRecommendations) IsSet() bool {
	return v.isSet
}

func (v *NullableSDAudienceRecommendations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDAudienceRecommendations(val *SDAudienceRecommendations) *NullableSDAudienceRecommendations {
	return &NullableSDAudienceRecommendations{value: val, isSet: true}
}

func (v NullableSDAudienceRecommendations) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDAudienceRecommendations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
