package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDProductTargetingRecommendationsSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDProductTargetingRecommendationsSuccess{}

// SDProductTargetingRecommendationsSuccess Recommendation results for contextual targeting.
type SDProductTargetingRecommendationsSuccess struct {
	// HTTP status code 200 indicating a successful response for product recomendations.
	Code *string `json:"code,omitempty"`
	// The theme name specified in the request.
	Name *string `json:"name,omitempty"`
	// A list of recommended products.
	Recommendations []SDProductRecommendationV32 `json:"recommendations,omitempty"`
}

// NewSDProductTargetingRecommendationsSuccess instantiates a new SDProductTargetingRecommendationsSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDProductTargetingRecommendationsSuccess() *SDProductTargetingRecommendationsSuccess {
	this := SDProductTargetingRecommendationsSuccess{}
	return &this
}

// NewSDProductTargetingRecommendationsSuccessWithDefaults instantiates a new SDProductTargetingRecommendationsSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDProductTargetingRecommendationsSuccessWithDefaults() *SDProductTargetingRecommendationsSuccess {
	this := SDProductTargetingRecommendationsSuccess{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDProductTargetingRecommendationsSuccess) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductTargetingRecommendationsSuccess) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDProductTargetingRecommendationsSuccess) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDProductTargetingRecommendationsSuccess) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SDProductTargetingRecommendationsSuccess) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductTargetingRecommendationsSuccess) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SDProductTargetingRecommendationsSuccess) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SDProductTargetingRecommendationsSuccess) SetName(v string) {
	o.Name = &v
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *SDProductTargetingRecommendationsSuccess) GetRecommendations() []SDProductRecommendationV32 {
	if o == nil || IsNil(o.Recommendations) {
		var ret []SDProductRecommendationV32
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductTargetingRecommendationsSuccess) GetRecommendationsOk() ([]SDProductRecommendationV32, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SDProductTargetingRecommendationsSuccess) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given []SDProductRecommendationV32 and assigns it to the Recommendations field.
func (o *SDProductTargetingRecommendationsSuccess) SetRecommendations(v []SDProductRecommendationV32) {
	o.Recommendations = v
}

func (o SDProductTargetingRecommendationsSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableSDProductTargetingRecommendationsSuccess struct {
	value *SDProductTargetingRecommendationsSuccess
	isSet bool
}

func (v NullableSDProductTargetingRecommendationsSuccess) Get() *SDProductTargetingRecommendationsSuccess {
	return v.value
}

func (v *NullableSDProductTargetingRecommendationsSuccess) Set(val *SDProductTargetingRecommendationsSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableSDProductTargetingRecommendationsSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableSDProductTargetingRecommendationsSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDProductTargetingRecommendationsSuccess(val *SDProductTargetingRecommendationsSuccess) *NullableSDProductTargetingRecommendationsSuccess {
	return &NullableSDProductTargetingRecommendationsSuccess{value: val, isSet: true}
}

func (v NullableSDProductTargetingRecommendationsSuccess) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDProductTargetingRecommendationsSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
