package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardTextModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardTextModule{}

// StandardTextModule A standard headline and body text.
type StandardTextModule struct {
	Headline *TextComponent     `json:"headline,omitempty"`
	Body     ParagraphComponent `json:"body"`
}

type _StandardTextModule StandardTextModule

// NewStandardTextModule instantiates a new StandardTextModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardTextModule(body ParagraphComponent) *StandardTextModule {
	this := StandardTextModule{}
	this.Body = body
	return &this
}

// NewStandardTextModuleWithDefaults instantiates a new StandardTextModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardTextModuleWithDefaults() *StandardTextModule {
	this := StandardTextModule{}
	return &this
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *StandardTextModule) GetHeadline() TextComponent {
	if o == nil || IsNil(o.Headline) {
		var ret TextComponent
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardTextModule) GetHeadlineOk() (*TextComponent, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *StandardTextModule) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given TextComponent and assigns it to the Headline field.
func (o *StandardTextModule) SetHeadline(v TextComponent) {
	o.Headline = &v
}

// GetBody returns the Body field value
func (o *StandardTextModule) GetBody() ParagraphComponent {
	if o == nil {
		var ret ParagraphComponent
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *StandardTextModule) GetBodyOk() (*ParagraphComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *StandardTextModule) SetBody(v ParagraphComponent) {
	o.Body = v
}

func (o StandardTextModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

type NullableStandardTextModule struct {
	value *StandardTextModule
	isSet bool
}

func (v NullableStandardTextModule) Get() *StandardTextModule {
	return v.value
}

func (v *NullableStandardTextModule) Set(val *StandardTextModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardTextModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardTextModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardTextModule(val *StandardTextModule) *NullableStandardTextModule {
	return &NullableStandardTextModule{value: val, isSet: true}
}

func (v NullableStandardTextModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardTextModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
