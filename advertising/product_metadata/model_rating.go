package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the Rating type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Rating{}

// Rating struct for Rating
type Rating struct {
	FullStarCount *int32   `json:"fullStarCount,omitempty"`
	HasHalfStar   *bool    `json:"hasHalfStar,omitempty"`
	DisplayString *string  `json:"displayString,omitempty"`
	Value         *float32 `json:"value,omitempty"`
}

// NewRating instantiates a new Rating object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRating() *Rating {
	this := Rating{}
	return &this
}

// NewRatingWithDefaults instantiates a new Rating object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatingWithDefaults() *Rating {
	this := Rating{}
	return &this
}

// GetFullStarCount returns the FullStarCount field value if set, zero value otherwise.
func (o *Rating) GetFullStarCount() int32 {
	if o == nil || IsNil(o.FullStarCount) {
		var ret int32
		return ret
	}
	return *o.FullStarCount
}

// GetFullStarCountOk returns a tuple with the FullStarCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rating) GetFullStarCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FullStarCount) {
		return nil, false
	}
	return o.FullStarCount, true
}

// HasFullStarCount returns a boolean if a field has been set.
func (o *Rating) HasFullStarCount() bool {
	if o != nil && !IsNil(o.FullStarCount) {
		return true
	}

	return false
}

// SetFullStarCount gets a reference to the given int32 and assigns it to the FullStarCount field.
func (o *Rating) SetFullStarCount(v int32) {
	o.FullStarCount = &v
}

// GetHasHalfStar returns the HasHalfStar field value if set, zero value otherwise.
func (o *Rating) GetHasHalfStar() bool {
	if o == nil || IsNil(o.HasHalfStar) {
		var ret bool
		return ret
	}
	return *o.HasHalfStar
}

// GetHasHalfStarOk returns a tuple with the HasHalfStar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rating) GetHasHalfStarOk() (*bool, bool) {
	if o == nil || IsNil(o.HasHalfStar) {
		return nil, false
	}
	return o.HasHalfStar, true
}

// HasHasHalfStar returns a boolean if a field has been set.
func (o *Rating) HasHasHalfStar() bool {
	if o != nil && !IsNil(o.HasHalfStar) {
		return true
	}

	return false
}

// SetHasHalfStar gets a reference to the given bool and assigns it to the HasHalfStar field.
func (o *Rating) SetHasHalfStar(v bool) {
	o.HasHalfStar = &v
}

// GetDisplayString returns the DisplayString field value if set, zero value otherwise.
func (o *Rating) GetDisplayString() string {
	if o == nil || IsNil(o.DisplayString) {
		var ret string
		return ret
	}
	return *o.DisplayString
}

// GetDisplayStringOk returns a tuple with the DisplayString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rating) GetDisplayStringOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayString) {
		return nil, false
	}
	return o.DisplayString, true
}

// HasDisplayString returns a boolean if a field has been set.
func (o *Rating) HasDisplayString() bool {
	if o != nil && !IsNil(o.DisplayString) {
		return true
	}

	return false
}

// SetDisplayString gets a reference to the given string and assigns it to the DisplayString field.
func (o *Rating) SetDisplayString(v string) {
	o.DisplayString = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Rating) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rating) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Rating) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *Rating) SetValue(v float32) {
	o.Value = &v
}

func (o Rating) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FullStarCount) {
		toSerialize["fullStarCount"] = o.FullStarCount
	}
	if !IsNil(o.HasHalfStar) {
		toSerialize["hasHalfStar"] = o.HasHalfStar
	}
	if !IsNil(o.DisplayString) {
		toSerialize["displayString"] = o.DisplayString
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRating struct {
	value *Rating
	isSet bool
}

func (v NullableRating) Get() *Rating {
	return v.value
}

func (v *NullableRating) Set(val *Rating) {
	v.value = val
	v.isSet = true
}

func (v NullableRating) IsSet() bool {
	return v.isSet
}

func (v *NullableRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRating(val *Rating) *NullableRating {
	return &NullableRating{value: val, isSet: true}
}

func (v NullableRating) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
