package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingRecommendationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingRecommendationsResponse{}

// TargetingRecommendationsResponse Response to a request for targeting recommendations.
type TargetingRecommendationsResponse struct {
	Recommendations *TargetingRecommendations `json:"recommendations,omitempty"`
}

// NewTargetingRecommendationsResponse instantiates a new TargetingRecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingRecommendationsResponse() *TargetingRecommendationsResponse {
	this := TargetingRecommendationsResponse{}
	return &this
}

// NewTargetingRecommendationsResponseWithDefaults instantiates a new TargetingRecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingRecommendationsResponseWithDefaults() *TargetingRecommendationsResponse {
	this := TargetingRecommendationsResponse{}
	return &this
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *TargetingRecommendationsResponse) GetRecommendations() TargetingRecommendations {
	if o == nil || IsNil(o.Recommendations) {
		var ret TargetingRecommendations
		return ret
	}
	return *o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingRecommendationsResponse) GetRecommendationsOk() (*TargetingRecommendations, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *TargetingRecommendationsResponse) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given TargetingRecommendations and assigns it to the Recommendations field.
func (o *TargetingRecommendationsResponse) SetRecommendations(v TargetingRecommendations) {
	o.Recommendations = &v
}

func (o TargetingRecommendationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableTargetingRecommendationsResponse struct {
	value *TargetingRecommendationsResponse
	isSet bool
}

func (v NullableTargetingRecommendationsResponse) Get() *TargetingRecommendationsResponse {
	return v.value
}

func (v *NullableTargetingRecommendationsResponse) Set(val *TargetingRecommendationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingRecommendationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingRecommendationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingRecommendationsResponse(val *TargetingRecommendationsResponse) *NullableTargetingRecommendationsResponse {
	return &NullableTargetingRecommendationsResponse{value: val, isSet: true}
}

func (v NullableTargetingRecommendationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingRecommendationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
