package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationsEligibilityRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationsEligibilityRequestContent{}

// CreativeRecommendationsEligibilityRequestContent struct for CreativeRecommendationsEligibilityRequestContent
type CreativeRecommendationsEligibilityRequestContent struct {
	Source Source `json:"source"`
}

type _CreativeRecommendationsEligibilityRequestContent CreativeRecommendationsEligibilityRequestContent

// NewCreativeRecommendationsEligibilityRequestContent instantiates a new CreativeRecommendationsEligibilityRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationsEligibilityRequestContent(source Source) *CreativeRecommendationsEligibilityRequestContent {
	this := CreativeRecommendationsEligibilityRequestContent{}
	this.Source = source
	return &this
}

// NewCreativeRecommendationsEligibilityRequestContentWithDefaults instantiates a new CreativeRecommendationsEligibilityRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationsEligibilityRequestContentWithDefaults() *CreativeRecommendationsEligibilityRequestContent {
	this := CreativeRecommendationsEligibilityRequestContent{}
	return &this
}

// GetSource returns the Source field value
func (o *CreativeRecommendationsEligibilityRequestContent) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsEligibilityRequestContent) GetSourceOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CreativeRecommendationsEligibilityRequestContent) SetSource(v Source) {
	o.Source = v
}

func (o CreativeRecommendationsEligibilityRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableCreativeRecommendationsEligibilityRequestContent struct {
	value *CreativeRecommendationsEligibilityRequestContent
	isSet bool
}

func (v NullableCreativeRecommendationsEligibilityRequestContent) Get() *CreativeRecommendationsEligibilityRequestContent {
	return v.value
}

func (v *NullableCreativeRecommendationsEligibilityRequestContent) Set(val *CreativeRecommendationsEligibilityRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationsEligibilityRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationsEligibilityRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationsEligibilityRequestContent(val *CreativeRecommendationsEligibilityRequestContent) *NullableCreativeRecommendationsEligibilityRequestContent {
	return &NullableCreativeRecommendationsEligibilityRequestContent{value: val, isSet: true}
}

func (v NullableCreativeRecommendationsEligibilityRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationsEligibilityRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
