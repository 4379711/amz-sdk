package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsResponseV33 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsResponseV33{}

// SDTargetingRecommendationsResponseV33 Response to a request for targeting recommendations.
type SDTargetingRecommendationsResponseV33 struct {
	Recommendations *SDTargetingRecommendationsV33 `json:"recommendations,omitempty"`
}

// NewSDTargetingRecommendationsResponseV33 instantiates a new SDTargetingRecommendationsResponseV33 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsResponseV33() *SDTargetingRecommendationsResponseV33 {
	this := SDTargetingRecommendationsResponseV33{}
	return &this
}

// NewSDTargetingRecommendationsResponseV33WithDefaults instantiates a new SDTargetingRecommendationsResponseV33 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsResponseV33WithDefaults() *SDTargetingRecommendationsResponseV33 {
	this := SDTargetingRecommendationsResponseV33{}
	return &this
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsResponseV33) GetRecommendations() SDTargetingRecommendationsV33 {
	if o == nil || IsNil(o.Recommendations) {
		var ret SDTargetingRecommendationsV33
		return ret
	}
	return *o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsResponseV33) GetRecommendationsOk() (*SDTargetingRecommendationsV33, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsResponseV33) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given SDTargetingRecommendationsV33 and assigns it to the Recommendations field.
func (o *SDTargetingRecommendationsResponseV33) SetRecommendations(v SDTargetingRecommendationsV33) {
	o.Recommendations = &v
}

func (o SDTargetingRecommendationsResponseV33) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsResponseV33 struct {
	value *SDTargetingRecommendationsResponseV33
	isSet bool
}

func (v NullableSDTargetingRecommendationsResponseV33) Get() *SDTargetingRecommendationsResponseV33 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsResponseV33) Set(val *SDTargetingRecommendationsResponseV33) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsResponseV33) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsResponseV33) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsResponseV33(val *SDTargetingRecommendationsResponseV33) *NullableSDTargetingRecommendationsResponseV33 {
	return &NullableSDTargetingRecommendationsResponseV33{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsResponseV33) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsResponseV33) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
