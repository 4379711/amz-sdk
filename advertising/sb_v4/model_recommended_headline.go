package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the RecommendedHeadline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendedHeadline{}

// RecommendedHeadline Recommended Headline in response object. Recommended headline will be locale specific, i.e. for an asin input in ES, Recommended headline will be in ES.
type RecommendedHeadline struct {
	// Unique Id of Recommended headline.
	HeadlineId *string `json:"headlineId,omitempty"`
	// String that contains Recommended headline.
	Headline *string `json:"headline,omitempty"`
}

// NewRecommendedHeadline instantiates a new RecommendedHeadline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendedHeadline() *RecommendedHeadline {
	this := RecommendedHeadline{}
	return &this
}

// NewRecommendedHeadlineWithDefaults instantiates a new RecommendedHeadline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendedHeadlineWithDefaults() *RecommendedHeadline {
	this := RecommendedHeadline{}
	return &this
}

// GetHeadlineId returns the HeadlineId field value if set, zero value otherwise.
func (o *RecommendedHeadline) GetHeadlineId() string {
	if o == nil || IsNil(o.HeadlineId) {
		var ret string
		return ret
	}
	return *o.HeadlineId
}

// GetHeadlineIdOk returns a tuple with the HeadlineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendedHeadline) GetHeadlineIdOk() (*string, bool) {
	if o == nil || IsNil(o.HeadlineId) {
		return nil, false
	}
	return o.HeadlineId, true
}

// HasHeadlineId returns a boolean if a field has been set.
func (o *RecommendedHeadline) HasHeadlineId() bool {
	if o != nil && !IsNil(o.HeadlineId) {
		return true
	}

	return false
}

// SetHeadlineId gets a reference to the given string and assigns it to the HeadlineId field.
func (o *RecommendedHeadline) SetHeadlineId(v string) {
	o.HeadlineId = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *RecommendedHeadline) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendedHeadline) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *RecommendedHeadline) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *RecommendedHeadline) SetHeadline(v string) {
	o.Headline = &v
}

func (o RecommendedHeadline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeadlineId) {
		toSerialize["headlineId"] = o.HeadlineId
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	return toSerialize, nil
}

type NullableRecommendedHeadline struct {
	value *RecommendedHeadline
	isSet bool
}

func (v NullableRecommendedHeadline) Get() *RecommendedHeadline {
	return v.value
}

func (v *NullableRecommendedHeadline) Set(val *RecommendedHeadline) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendedHeadline) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendedHeadline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendedHeadline(val *RecommendedHeadline) *NullableRecommendedHeadline {
	return &NullableRecommendedHeadline{value: val, isSet: true}
}

func (v NullableRecommendedHeadline) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRecommendedHeadline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
