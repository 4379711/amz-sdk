package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsResponseV32 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsResponseV32{}

// SDTargetingRecommendationsResponseV32 Response to a request for targeting recommendations.
type SDTargetingRecommendationsResponseV32 struct {
	Recommendations *SDTargetingRecommendationsV32 `json:"recommendations,omitempty"`
}

// NewSDTargetingRecommendationsResponseV32 instantiates a new SDTargetingRecommendationsResponseV32 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsResponseV32() *SDTargetingRecommendationsResponseV32 {
	this := SDTargetingRecommendationsResponseV32{}
	return &this
}

// NewSDTargetingRecommendationsResponseV32WithDefaults instantiates a new SDTargetingRecommendationsResponseV32 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsResponseV32WithDefaults() *SDTargetingRecommendationsResponseV32 {
	this := SDTargetingRecommendationsResponseV32{}
	return &this
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsResponseV32) GetRecommendations() SDTargetingRecommendationsV32 {
	if o == nil || IsNil(o.Recommendations) {
		var ret SDTargetingRecommendationsV32
		return ret
	}
	return *o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsResponseV32) GetRecommendationsOk() (*SDTargetingRecommendationsV32, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsResponseV32) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given SDTargetingRecommendationsV32 and assigns it to the Recommendations field.
func (o *SDTargetingRecommendationsResponseV32) SetRecommendations(v SDTargetingRecommendationsV32) {
	o.Recommendations = &v
}

func (o SDTargetingRecommendationsResponseV32) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsResponseV32 struct {
	value *SDTargetingRecommendationsResponseV32
	isSet bool
}

func (v NullableSDTargetingRecommendationsResponseV32) Get() *SDTargetingRecommendationsResponseV32 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsResponseV32) Set(val *SDTargetingRecommendationsResponseV32) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsResponseV32) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsResponseV32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsResponseV32(val *SDTargetingRecommendationsResponseV32) *NullableSDTargetingRecommendationsResponseV32 {
	return &NullableSDTargetingRecommendationsResponseV32{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsResponseV32) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsResponseV32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
