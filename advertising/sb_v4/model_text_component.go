package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the TextComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextComponent{}

// TextComponent Text component which needs to be pre moderated
type TextComponent struct {
	// Type of text component.
	ComponentType string `json:"componentType"`
	// Id of the component. The same will be returned as part of the response as well. This can be used to uniquely identify the component from the pre moderation response.
	Id string `json:"id"`
	// Text which needs to be moderated.
	Text string `json:"text"`
}

type _TextComponent TextComponent

// NewTextComponent instantiates a new TextComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextComponent(componentType string, id string, text string) *TextComponent {
	this := TextComponent{}
	this.ComponentType = componentType
	this.Id = id
	this.Text = text
	return &this
}

// NewTextComponentWithDefaults instantiates a new TextComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextComponentWithDefaults() *TextComponent {
	this := TextComponent{}
	return &this
}

// GetComponentType returns the ComponentType field value
func (o *TextComponent) GetComponentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value
// and a boolean to check if the value has been set.
func (o *TextComponent) GetComponentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentType, true
}

// SetComponentType sets field value
func (o *TextComponent) SetComponentType(v string) {
	o.ComponentType = v
}

// GetId returns the Id field value
func (o *TextComponent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TextComponent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TextComponent) SetId(v string) {
	o.Id = v
}

// GetText returns the Text field value
func (o *TextComponent) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *TextComponent) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *TextComponent) SetText(v string) {
	o.Text = v
}

func (o TextComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["componentType"] = o.ComponentType
	toSerialize["id"] = o.Id
	toSerialize["text"] = o.Text
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
