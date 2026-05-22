package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the TextRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextRecommendation{}

// TextRecommendation struct for TextRecommendation
type TextRecommendation struct {
	// Unique ID for generated recommendation.
	Id *string `json:"id,omitempty"`
	// Recommendation value.
	Value *string `json:"value,omitempty"`
}

// NewTextRecommendation instantiates a new TextRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextRecommendation() *TextRecommendation {
	this := TextRecommendation{}
	return &this
}

// NewTextRecommendationWithDefaults instantiates a new TextRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextRecommendationWithDefaults() *TextRecommendation {
	this := TextRecommendation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TextRecommendation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextRecommendation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TextRecommendation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TextRecommendation) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TextRecommendation) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextRecommendation) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TextRecommendation) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TextRecommendation) SetValue(v string) {
	o.Value = &v
}

func (o TextRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTextRecommendation struct {
	value *TextRecommendation
	isSet bool
}

func (v NullableTextRecommendation) Get() *TextRecommendation {
	return v.value
}

func (v *NullableTextRecommendation) Set(val *TextRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableTextRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableTextRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextRecommendation(val *TextRecommendation) *NullableTextRecommendation {
	return &NullableTextRecommendation{value: val, isSet: true}
}

func (v NullableTextRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTextRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
