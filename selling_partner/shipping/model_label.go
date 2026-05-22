package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the Label type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Label{}

// Label The label details of the container.
type Label struct {
	// Contains binary image data encoded as a base-64 string.
	LabelStream        *string             `json:"labelStream,omitempty"`
	LabelSpecification *LabelSpecification `json:"labelSpecification,omitempty"`
}

// NewLabel instantiates a new Label object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabel() *Label {
	this := Label{}
	return &this
}

// NewLabelWithDefaults instantiates a new Label object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelWithDefaults() *Label {
	this := Label{}
	return &this
}

// GetLabelStream returns the LabelStream field value if set, zero value otherwise.
func (o *Label) GetLabelStream() string {
	if o == nil || IsNil(o.LabelStream) {
		var ret string
		return ret
	}
	return *o.LabelStream
}

// GetLabelStreamOk returns a tuple with the LabelStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Label) GetLabelStreamOk() (*string, bool) {
	if o == nil || IsNil(o.LabelStream) {
		return nil, false
	}
	return o.LabelStream, true
}

// HasLabelStream returns a boolean if a field has been set.
func (o *Label) HasLabelStream() bool {
	if o != nil && !IsNil(o.LabelStream) {
		return true
	}

	return false
}

// SetLabelStream gets a reference to the given string and assigns it to the LabelStream field.
func (o *Label) SetLabelStream(v string) {
	o.LabelStream = &v
}

// GetLabelSpecification returns the LabelSpecification field value if set, zero value otherwise.
func (o *Label) GetLabelSpecification() LabelSpecification {
	if o == nil || IsNil(o.LabelSpecification) {
		var ret LabelSpecification
		return ret
	}
	return *o.LabelSpecification
}

// GetLabelSpecificationOk returns a tuple with the LabelSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Label) GetLabelSpecificationOk() (*LabelSpecification, bool) {
	if o == nil || IsNil(o.LabelSpecification) {
		return nil, false
	}
	return o.LabelSpecification, true
}

// HasLabelSpecification returns a boolean if a field has been set.
func (o *Label) HasLabelSpecification() bool {
	if o != nil && !IsNil(o.LabelSpecification) {
		return true
	}

	return false
}

// SetLabelSpecification gets a reference to the given LabelSpecification and assigns it to the LabelSpecification field.
func (o *Label) SetLabelSpecification(v LabelSpecification) {
	o.LabelSpecification = &v
}

func (o Label) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelStream) {
		toSerialize["labelStream"] = o.LabelStream
	}
	if !IsNil(o.LabelSpecification) {
		toSerialize["labelSpecification"] = o.LabelSpecification
	}
	return toSerialize, nil
}

type NullableLabel struct {
	value *Label
	isSet bool
}

func (v NullableLabel) Get() *Label {
	return v.value
}

func (v *NullableLabel) Set(val *Label) {
	v.value = val
	v.isSet = true
}

func (v NullableLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabel(val *Label) *NullableLabel {
	return &NullableLabel{value: val, isSet: true}
}

func (v NullableLabel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
