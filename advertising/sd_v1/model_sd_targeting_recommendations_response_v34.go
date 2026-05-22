package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsResponseV34 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsResponseV34{}

// SDTargetingRecommendationsResponseV34 Response to a request for targeting recommendations.
type SDTargetingRecommendationsResponseV34 struct {
	Recommendations *SDTargetingRecommendationsV34 `json:"recommendations,omitempty"`
}

// NewSDTargetingRecommendationsResponseV34 instantiates a new SDTargetingRecommendationsResponseV34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsResponseV34() *SDTargetingRecommendationsResponseV34 {
	this := SDTargetingRecommendationsResponseV34{}
	return &this
}

// NewSDTargetingRecommendationsResponseV34WithDefaults instantiates a new SDTargetingRecommendationsResponseV34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsResponseV34WithDefaults() *SDTargetingRecommendationsResponseV34 {
	this := SDTargetingRecommendationsResponseV34{}
	return &this
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsResponseV34) GetRecommendations() SDTargetingRecommendationsV34 {
	if o == nil || IsNil(o.Recommendations) {
		var ret SDTargetingRecommendationsV34
		return ret
	}
	return *o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsResponseV34) GetRecommendationsOk() (*SDTargetingRecommendationsV34, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsResponseV34) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given SDTargetingRecommendationsV34 and assigns it to the Recommendations field.
func (o *SDTargetingRecommendationsResponseV34) SetRecommendations(v SDTargetingRecommendationsV34) {
	o.Recommendations = &v
}

func (o SDTargetingRecommendationsResponseV34) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsResponseV34 struct {
	value *SDTargetingRecommendationsResponseV34
	isSet bool
}

func (v NullableSDTargetingRecommendationsResponseV34) Get() *SDTargetingRecommendationsResponseV34 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsResponseV34) Set(val *SDTargetingRecommendationsResponseV34) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsResponseV34) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsResponseV34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsResponseV34(val *SDTargetingRecommendationsResponseV34) *NullableSDTargetingRecommendationsResponseV34 {
	return &NullableSDTargetingRecommendationsResponseV34{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsResponseV34) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsResponseV34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
