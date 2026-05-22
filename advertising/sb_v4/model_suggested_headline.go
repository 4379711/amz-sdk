package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SuggestedHeadline type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestedHeadline{}

// SuggestedHeadline Suggested Headline in response object.
type SuggestedHeadline struct {
	// Unique Id of suggested headline.
	HeadlineId *string `json:"headlineId,omitempty"`
	// String that contains suggested headline.
	Headline *string `json:"headline,omitempty"`
}

// NewSuggestedHeadline instantiates a new SuggestedHeadline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestedHeadline() *SuggestedHeadline {
	this := SuggestedHeadline{}
	return &this
}

// NewSuggestedHeadlineWithDefaults instantiates a new SuggestedHeadline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestedHeadlineWithDefaults() *SuggestedHeadline {
	this := SuggestedHeadline{}
	return &this
}

// GetHeadlineId returns the HeadlineId field value if set, zero value otherwise.
func (o *SuggestedHeadline) GetHeadlineId() string {
	if o == nil || IsNil(o.HeadlineId) {
		var ret string
		return ret
	}
	return *o.HeadlineId
}

// GetHeadlineIdOk returns a tuple with the HeadlineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedHeadline) GetHeadlineIdOk() (*string, bool) {
	if o == nil || IsNil(o.HeadlineId) {
		return nil, false
	}
	return o.HeadlineId, true
}

// HasHeadlineId returns a boolean if a field has been set.
func (o *SuggestedHeadline) HasHeadlineId() bool {
	if o != nil && !IsNil(o.HeadlineId) {
		return true
	}

	return false
}

// SetHeadlineId gets a reference to the given string and assigns it to the HeadlineId field.
func (o *SuggestedHeadline) SetHeadlineId(v string) {
	o.HeadlineId = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *SuggestedHeadline) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestedHeadline) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *SuggestedHeadline) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *SuggestedHeadline) SetHeadline(v string) {
	o.Headline = &v
}

func (o SuggestedHeadline) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeadlineId) {
		toSerialize["headlineId"] = o.HeadlineId
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	return toSerialize, nil
}

type NullableSuggestedHeadline struct {
	value *SuggestedHeadline
	isSet bool
}

func (v NullableSuggestedHeadline) Get() *SuggestedHeadline {
	return v.value
}

func (v *NullableSuggestedHeadline) Set(val *SuggestedHeadline) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestedHeadline) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestedHeadline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestedHeadline(val *SuggestedHeadline) *NullableSuggestedHeadline {
	return &NullableSuggestedHeadline{value: val, isSet: true}
}

func (v NullableSuggestedHeadline) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSuggestedHeadline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
