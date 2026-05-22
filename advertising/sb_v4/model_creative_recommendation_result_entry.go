package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationResultEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationResultEntry{}

// CreativeRecommendationResultEntry Creative Recommendation Result.
type CreativeRecommendationResultEntry struct {
	// Supported are PRODUCT_COLLECTION, STORE_SPOTLIGHT, VIDEO, BRAND_VIDEO. More could be added in future.
	CreativeType       *string                           `json:"creativeType,omitempty"`
	CreativeProperties *CreativeRecommendationProperties `json:"creativeProperties,omitempty"`
}

// NewCreativeRecommendationResultEntry instantiates a new CreativeRecommendationResultEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationResultEntry() *CreativeRecommendationResultEntry {
	this := CreativeRecommendationResultEntry{}
	return &this
}

// NewCreativeRecommendationResultEntryWithDefaults instantiates a new CreativeRecommendationResultEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationResultEntryWithDefaults() *CreativeRecommendationResultEntry {
	this := CreativeRecommendationResultEntry{}
	return &this
}

// GetCreativeType returns the CreativeType field value if set, zero value otherwise.
func (o *CreativeRecommendationResultEntry) GetCreativeType() string {
	if o == nil || IsNil(o.CreativeType) {
		var ret string
		return ret
	}
	return *o.CreativeType
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationResultEntry) GetCreativeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CreativeType) {
		return nil, false
	}
	return o.CreativeType, true
}

// HasCreativeType returns a boolean if a field has been set.
func (o *CreativeRecommendationResultEntry) HasCreativeType() bool {
	if o != nil && !IsNil(o.CreativeType) {
		return true
	}

	return false
}

// SetCreativeType gets a reference to the given string and assigns it to the CreativeType field.
func (o *CreativeRecommendationResultEntry) SetCreativeType(v string) {
	o.CreativeType = &v
}

// GetCreativeProperties returns the CreativeProperties field value if set, zero value otherwise.
func (o *CreativeRecommendationResultEntry) GetCreativeProperties() CreativeRecommendationProperties {
	if o == nil || IsNil(o.CreativeProperties) {
		var ret CreativeRecommendationProperties
		return ret
	}
	return *o.CreativeProperties
}

// GetCreativePropertiesOk returns a tuple with the CreativeProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationResultEntry) GetCreativePropertiesOk() (*CreativeRecommendationProperties, bool) {
	if o == nil || IsNil(o.CreativeProperties) {
		return nil, false
	}
	return o.CreativeProperties, true
}

// HasCreativeProperties returns a boolean if a field has been set.
func (o *CreativeRecommendationResultEntry) HasCreativeProperties() bool {
	if o != nil && !IsNil(o.CreativeProperties) {
		return true
	}

	return false
}

// SetCreativeProperties gets a reference to the given CreativeRecommendationProperties and assigns it to the CreativeProperties field.
func (o *CreativeRecommendationResultEntry) SetCreativeProperties(v CreativeRecommendationProperties) {
	o.CreativeProperties = &v
}

func (o CreativeRecommendationResultEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreativeType) {
		toSerialize["creativeType"] = o.CreativeType
	}
	if !IsNil(o.CreativeProperties) {
		toSerialize["creativeProperties"] = o.CreativeProperties
	}
	return toSerialize, nil
}

type NullableCreativeRecommendationResultEntry struct {
	value *CreativeRecommendationResultEntry
	isSet bool
}

func (v NullableCreativeRecommendationResultEntry) Get() *CreativeRecommendationResultEntry {
	return v.value
}

func (v *NullableCreativeRecommendationResultEntry) Set(val *CreativeRecommendationResultEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationResultEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationResultEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationResultEntry(val *CreativeRecommendationResultEntry) *NullableCreativeRecommendationResultEntry {
	return &NullableCreativeRecommendationResultEntry{value: val, isSet: true}
}

func (v NullableCreativeRecommendationResultEntry) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationResultEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
