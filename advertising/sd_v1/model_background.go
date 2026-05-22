package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the Background type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Background{}

// Background This field denotes background which are displayed on the ad. This field is optional and mutable.
type Background struct {
	// The standard HTML hex color codes of the background (e.g. '#3cb371').
	Color *string `json:"color,omitempty"`
}

// NewBackground instantiates a new Background object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackground() *Background {
	this := Background{}
	return &this
}

// NewBackgroundWithDefaults instantiates a new Background object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackgroundWithDefaults() *Background {
	this := Background{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Background) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Background) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Background) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Background) SetColor(v string) {
	o.Color = &v
}

func (o Background) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	return toSerialize, nil
}

type NullableBackground struct {
	value *Background
	isSet bool
}

func (v NullableBackground) Get() *Background {
	return v.value
}

func (v *NullableBackground) Set(val *Background) {
	v.value = val
	v.isSet = true
}

func (v NullableBackground) IsSet() bool {
	return v.isSet
}

func (v *NullableBackground) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackground(val *Background) *NullableBackground {
	return &NullableBackground{value: val, isSet: true}
}

func (v NullableBackground) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBackground) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
