package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardProductDescriptionModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardProductDescriptionModule{}

// StandardProductDescriptionModule Standard product description text.
type StandardProductDescriptionModule struct {
	Body ParagraphComponent `json:"body"`
}

type _StandardProductDescriptionModule StandardProductDescriptionModule

// NewStandardProductDescriptionModule instantiates a new StandardProductDescriptionModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardProductDescriptionModule(body ParagraphComponent) *StandardProductDescriptionModule {
	this := StandardProductDescriptionModule{}
	this.Body = body
	return &this
}

// NewStandardProductDescriptionModuleWithDefaults instantiates a new StandardProductDescriptionModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardProductDescriptionModuleWithDefaults() *StandardProductDescriptionModule {
	this := StandardProductDescriptionModule{}
	return &this
}

// GetBody returns the Body field value
func (o *StandardProductDescriptionModule) GetBody() ParagraphComponent {
	if o == nil {
		var ret ParagraphComponent
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *StandardProductDescriptionModule) GetBodyOk() (*ParagraphComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *StandardProductDescriptionModule) SetBody(v ParagraphComponent) {
	o.Body = v
}

func (o StandardProductDescriptionModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

type NullableStandardProductDescriptionModule struct {
	value *StandardProductDescriptionModule
	isSet bool
}

func (v NullableStandardProductDescriptionModule) Get() *StandardProductDescriptionModule {
	return v.value
}

func (v *NullableStandardProductDescriptionModule) Set(val *StandardProductDescriptionModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardProductDescriptionModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardProductDescriptionModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardProductDescriptionModule(val *StandardProductDescriptionModule) *NullableStandardProductDescriptionModule {
	return &NullableStandardProductDescriptionModule{value: val, isSet: true}
}

func (v NullableStandardProductDescriptionModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardProductDescriptionModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
