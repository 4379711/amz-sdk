package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationByIdResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationByIdResponseContent{}

// CreativeRecommendationByIdResponseContent Creative Recommendation by Id Response.
type CreativeRecommendationByIdResponseContent struct {
	// Supported are PRODUCT_COLLECTION, STORE_SPOTLIGHT, VIDEO, BRAND_VIDEO. More could be added in future.
	CreativeType       *string                           `json:"creativeType,omitempty"`
	CreativeProperties *CreativeRecommendationProperties `json:"creativeProperties,omitempty"`
}

// NewCreativeRecommendationByIdResponseContent instantiates a new CreativeRecommendationByIdResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationByIdResponseContent() *CreativeRecommendationByIdResponseContent {
	this := CreativeRecommendationByIdResponseContent{}
	return &this
}

// NewCreativeRecommendationByIdResponseContentWithDefaults instantiates a new CreativeRecommendationByIdResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationByIdResponseContentWithDefaults() *CreativeRecommendationByIdResponseContent {
	this := CreativeRecommendationByIdResponseContent{}
	return &this
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise.
func (o *CreativeRecommendationByIdResponseContent) GetCreativeType() string {
	if o == nil || IsNil(o.CreativeType) {
		var ret string
		return ret
	}
	return *o.CreativeType
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationByIdResponseContent) GetCreativeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreativeType) {
		return nil, false
	}
	return o.CreativeType, true
}

// HasCreativeType returns a boolean if a field has been set.
func (o *CreativeRecommendationByIdResponseContent) HasCreativeType() bool {
	if o != nil && !IsNil(o.CreativeType) {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given string and assigns it to the CreativeType field.
func (o *CreativeRecommendationByIdResponseContent) SetCreativeType(v string) {
	o.CreativeType = &v
}

// GetCreativeProperties returns the CreativeProperties field value if set, zero value otherwise.
func (o *CreativeRecommendationByIdResponseContent) GetCreativeProperties() CreativeRecommendationProperties {
	if o == nil || IsNil(o.CreativeProperties) {
		var ret CreativeRecommendationProperties
		return ret
	}
	return *o.CreativeProperties
}

// GetCreativePropertiesOk returns a tuple with the CreativeProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationByIdResponseContent) GetCreativePropertiesOk() (*CreativeRecommendationProperties, bool) {
	if o == nil || IsNil(o.CreativeProperties) {
		return nil, false
	}
	return o.CreativeProperties, true
}

// HasCreativeProperties returns a boolean if a field has been set.
func (o *CreativeRecommendationByIdResponseContent) HasCreativeProperties() bool {
	if o != nil && !IsNil(o.CreativeProperties) {
		return true
	}

	return false
}

// SetCreativeProperties gets a reference to the given CreativeRecommendationProperties and assigns it to the CreativeProperties field.
func (o *CreativeRecommendationByIdResponseContent) SetCreativeProperties(v CreativeRecommendationProperties) {
	o.CreativeProperties = &v
}

func (o CreativeRecommendationByIdResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreativeType) {
		toSerialize["creativeType"] = o.CreativeType
	}
	if !IsNil(o.CreativeProperties) {
		toSerialize["creativeProperties"] = o.CreativeProperties
	}
	return toSerialize, nil
}

type NullableCreativeRecommendationByIdResponseContent struct {
	value *CreativeRecommendationByIdResponseContent
	isSet bool
}

func (v NullableCreativeRecommendationByIdResponseContent) Get() *CreativeRecommendationByIdResponseContent {
	return v.value
}

func (v *NullableCreativeRecommendationByIdResponseContent) Set(val *CreativeRecommendationByIdResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationByIdResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationByIdResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationByIdResponseContent(val *CreativeRecommendationByIdResponseContent) *NullableCreativeRecommendationByIdResponseContent {
	return &NullableCreativeRecommendationByIdResponseContent{value: val, isSet: true}
}

func (v NullableCreativeRecommendationByIdResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationByIdResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
