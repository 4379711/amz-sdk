package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsResponseV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsResponseV31{}

// SDTargetingRecommendationsResponseV31 Response to a request for targeting recommendations.
type SDTargetingRecommendationsResponseV31 struct {
	Recommendations *SDTargetingRecommendationsV31 `json:"recommendations,omitempty"`
}

// NewSDTargetingRecommendationsResponseV31 instantiates a new SDTargetingRecommendationsResponseV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsResponseV31() *SDTargetingRecommendationsResponseV31 {
	this := SDTargetingRecommendationsResponseV31{}
	return &this
}

// NewSDTargetingRecommendationsResponseV31WithDefaults instantiates a new SDTargetingRecommendationsResponseV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsResponseV31WithDefaults() *SDTargetingRecommendationsResponseV31 {
	this := SDTargetingRecommendationsResponseV31{}
	return &this
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsResponseV31) GetRecommendations() SDTargetingRecommendationsV31 {
	if o == nil || IsNil(o.Recommendations) {
		var ret SDTargetingRecommendationsV31
		return ret
	}
	return *o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsResponseV31) GetRecommendationsOk() (*SDTargetingRecommendationsV31, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsResponseV31) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given SDTargetingRecommendationsV31 and assigns it to the Recommendations field.
func (o *SDTargetingRecommendationsResponseV31) SetRecommendations(v SDTargetingRecommendationsV31) {
	o.Recommendations = &v
}

func (o SDTargetingRecommendationsResponseV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsResponseV31 struct {
	value *SDTargetingRecommendationsResponseV31
	isSet bool
}

func (v NullableSDTargetingRecommendationsResponseV31) Get() *SDTargetingRecommendationsResponseV31 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsResponseV31) Set(val *SDTargetingRecommendationsResponseV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsResponseV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsResponseV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsResponseV31(val *SDTargetingRecommendationsResponseV31) *NullableSDTargetingRecommendationsResponseV31 {
	return &NullableSDTargetingRecommendationsResponseV31{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsResponseV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsResponseV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
