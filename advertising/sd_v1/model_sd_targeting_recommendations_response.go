package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsResponse{}

// SDTargetingRecommendationsResponse Response to a request for targeting recommendations.
type SDTargetingRecommendationsResponse struct {
	Recommendations *SDTargetingRecommendations `json:"recommendations,omitempty"`
}

// NewSDTargetingRecommendationsResponse instantiates a new SDTargetingRecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsResponse() *SDTargetingRecommendationsResponse {
	this := SDTargetingRecommendationsResponse{}
	return &this
}

// NewSDTargetingRecommendationsResponseWithDefaults instantiates a new SDTargetingRecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsResponseWithDefaults() *SDTargetingRecommendationsResponse {
	this := SDTargetingRecommendationsResponse{}
	return &this
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsResponse) GetRecommendations() SDTargetingRecommendations {
	if o == nil || IsNil(o.Recommendations) {
		var ret SDTargetingRecommendations
		return ret
	}
	return *o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsResponse) GetRecommendationsOk() (*SDTargetingRecommendations, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsResponse) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given SDTargetingRecommendations and assigns it to the Recommendations field.
func (o *SDTargetingRecommendationsResponse) SetRecommendations(v SDTargetingRecommendations) {
	o.Recommendations = &v
}

func (o SDTargetingRecommendationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsResponse struct {
	value *SDTargetingRecommendationsResponse
	isSet bool
}

func (v NullableSDTargetingRecommendationsResponse) Get() *SDTargetingRecommendationsResponse {
	return v.value
}

func (v *NullableSDTargetingRecommendationsResponse) Set(val *SDTargetingRecommendationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsResponse(val *SDTargetingRecommendationsResponse) *NullableSDTargetingRecommendationsResponse {
	return &NullableSDTargetingRecommendationsResponse{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
