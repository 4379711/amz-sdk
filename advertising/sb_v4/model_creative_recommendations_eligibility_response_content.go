package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationsEligibilityResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationsEligibilityResponseContent{}

// CreativeRecommendationsEligibilityResponseContent struct for CreativeRecommendationsEligibilityResponseContent
type CreativeRecommendationsEligibilityResponseContent struct {
	// Returns false if there is no creative recommendation possible with the given landing page.
	IsEligible *bool `json:"isEligible,omitempty"`
	// Supported are PRODUCT_COLLECTION, STORE_SPOTLIGHT, VIDEO, BRAND_VIDEO. More could be added in future.
	CreativeTypes []string `json:"creativeTypes,omitempty"`
}

// NewCreativeRecommendationsEligibilityResponseContent instantiates a new CreativeRecommendationsEligibilityResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationsEligibilityResponseContent() *CreativeRecommendationsEligibilityResponseContent {
	this := CreativeRecommendationsEligibilityResponseContent{}
	return &this
}

// NewCreativeRecommendationsEligibilityResponseContentWithDefaults instantiates a new CreativeRecommendationsEligibilityResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationsEligibilityResponseContentWithDefaults() *CreativeRecommendationsEligibilityResponseContent {
	this := CreativeRecommendationsEligibilityResponseContent{}
	return &this
}

// GetIsEligible returns the IsEligible field value if set, zero value otherwise.
func (o *CreativeRecommendationsEligibilityResponseContent) GetIsEligible() bool {
	if o == nil || IsNil(o.IsEligible) {
		var ret bool
		return ret
	}
	return *o.IsEligible
}

// GetIsEligibleOk returns a tuple with the IsEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsEligibilityResponseContent) GetIsEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEligible) {
		return nil, false
	}
	return o.IsEligible, true
}

// HasIsEligible returns a boolean if a field has been set.
func (o *CreativeRecommendationsEligibilityResponseContent) HasIsEligible() bool {
	if o != nil && !IsNil(o.IsEligible) {
		return true
	}

	return false
}

// SetIsEligible gets a reference to the given bool and assigns it to the IsEligible field.
func (o *CreativeRecommendationsEligibilityResponseContent) SetIsEligible(v bool) {
	o.IsEligible = &v
}

// GetCreativeTypes returns the CreativeTypes field value if set, zero value otherwise.
func (o *CreativeRecommendationsEligibilityResponseContent) GetCreativeTypes() []string {
	if o == nil || IsNil(o.CreativeTypes) {
		var ret []string
		return ret
	}
	return o.CreativeTypes
}

// GetCreativeTypesOk returns a tuple with the CreativeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsEligibilityResponseContent) GetCreativeTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.CreativeTypes) {
		return nil, false
	}
	return o.CreativeTypes, true
}

// HasCreativeTypes returns a boolean if a field has been set.
func (o *CreativeRecommendationsEligibilityResponseContent) HasCreativeTypes() bool {
	if o != nil && !IsNil(o.CreativeTypes) {
		return true
	}

	return false
}

// SetCreativeTypes gets a reference to the given []string and assigns it to the CreativeTypes field.
func (o *CreativeRecommendationsEligibilityResponseContent) SetCreativeTypes(v []string) {
	o.CreativeTypes = v
}

func (o CreativeRecommendationsEligibilityResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsEligible) {
		toSerialize["isEligible"] = o.IsEligible
	}
	if !IsNil(o.CreativeTypes) {
		toSerialize["creativeTypes"] = o.CreativeTypes
	}
	return toSerialize, nil
}

type NullableCreativeRecommendationsEligibilityResponseContent struct {
	value *CreativeRecommendationsEligibilityResponseContent
	isSet bool
}

func (v NullableCreativeRecommendationsEligibilityResponseContent) Get() *CreativeRecommendationsEligibilityResponseContent {
	return v.value
}

func (v *NullableCreativeRecommendationsEligibilityResponseContent) Set(val *CreativeRecommendationsEligibilityResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationsEligibilityResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationsEligibilityResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationsEligibilityResponseContent(val *CreativeRecommendationsEligibilityResponseContent) *NullableCreativeRecommendationsEligibilityResponseContent {
	return &NullableCreativeRecommendationsEligibilityResponseContent{value: val, isSet: true}
}

func (v NullableCreativeRecommendationsEligibilityResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationsEligibilityResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
