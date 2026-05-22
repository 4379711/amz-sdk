package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsResponseV35 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsResponseV35{}

// SDTargetingRecommendationsResponseV35 Response to a request for targeting recommendations.
type SDTargetingRecommendationsResponseV35 struct {
	Recommendations *SDTargetingRecommendationsV35 `json:"recommendations,omitempty"`
}

// NewSDTargetingRecommendationsResponseV35 instantiates a new SDTargetingRecommendationsResponseV35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsResponseV35() *SDTargetingRecommendationsResponseV35 {
	this := SDTargetingRecommendationsResponseV35{}
	return &this
}

// NewSDTargetingRecommendationsResponseV35WithDefaults instantiates a new SDTargetingRecommendationsResponseV35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsResponseV35WithDefaults() *SDTargetingRecommendationsResponseV35 {
	this := SDTargetingRecommendationsResponseV35{}
	return &this
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsResponseV35) GetRecommendations() SDTargetingRecommendationsV35 {
	if o == nil || IsNil(o.Recommendations) {
		var ret SDTargetingRecommendationsV35
		return ret
	}
	return *o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsResponseV35) GetRecommendationsOk() (*SDTargetingRecommendationsV35, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsResponseV35) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given SDTargetingRecommendationsV35 and assigns it to the Recommendations field.
func (o *SDTargetingRecommendationsResponseV35) SetRecommendations(v SDTargetingRecommendationsV35) {
	o.Recommendations = &v
}

func (o SDTargetingRecommendationsResponseV35) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsResponseV35 struct {
	value *SDTargetingRecommendationsResponseV35
	isSet bool
}

func (v NullableSDTargetingRecommendationsResponseV35) Get() *SDTargetingRecommendationsResponseV35 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsResponseV35) Set(val *SDTargetingRecommendationsResponseV35) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsResponseV35) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsResponseV35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsResponseV35(val *SDTargetingRecommendationsResponseV35) *NullableSDTargetingRecommendationsResponseV35 {
	return &NullableSDTargetingRecommendationsResponseV35{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsResponseV35) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsResponseV35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
