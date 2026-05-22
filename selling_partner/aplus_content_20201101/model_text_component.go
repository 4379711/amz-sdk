package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the TextComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextComponent{}

// TextComponent Rich text content.
type TextComponent struct {
	// The actual plain text.
	Value string `json:"value"`
	// A set of content decorators.
	DecoratorSet []Decorator `json:"decoratorSet,omitempty"`
}

type _TextComponent TextComponent

// NewTextComponent instantiates a new TextComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextComponent(value string) *TextComponent {
	this := TextComponent{}
	this.Value = value
	return &this
}

// NewTextComponentWithDefaults instantiates a new TextComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextComponentWithDefaults() *TextComponent {
	this := TextComponent{}
	return &this
}

// GetValue returns the Value field value
func (o *TextComponent) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TextComponent) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TextComponent) SetValue(v string) {
	o.Value = v
}

// GetDecoratorSet returns the DecoratorSet field value if set, zero value otherwise.
func (o *TextComponent) GetDecoratorSet() []Decorator {
	if o == nil || IsNil(o.DecoratorSet) {
		var ret []Decorator
		return ret
	}
	return o.DecoratorSet
}

// GetDecoratorSetOk returns a tuple with the DecoratorSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextComponent) GetDecoratorSetOk() ([]Decorator, bool) {
	if o == nil || IsNil(o.DecoratorSet) {
		return nil, false
	}
	return o.DecoratorSet, true
}

// HasDecoratorSet returns a boolean if a field has been set.
func (o *TextComponent) HasDecoratorSet() bool {
	if o != nil && !IsNil(o.DecoratorSet) {
		return true
	}

	return false
}

// SetDecoratorSet gets a reference to the given []Decorator and assigns it to the DecoratorSet field.
func (o *TextComponent) SetDecoratorSet(v []Decorator) {
	o.DecoratorSet = v
}

func (o TextComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.DecoratorSet) {
		toSerialize["decoratorSet"] = o.DecoratorSet
	}
	return toSerialize, nil
}

type NullableTextComponent struct {
	value *TextComponent
	isSet bool
}

func (v NullableTextComponent) Get() *TextComponent {
	return v.value
}

func (v *NullableTextComponent) Set(val *TextComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableTextComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableTextComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextComponent(val *TextComponent) *NullableTextComponent {
	return &NullableTextComponent{value: val, isSet: true}
}

func (v NullableTextComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTextComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
