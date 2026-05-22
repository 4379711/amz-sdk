package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the LabelSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelSpecification{}

// LabelSpecification The label specification info.
type LabelSpecification struct {
	// The format of the label. Enum of PNG only for now.
	LabelFormat string `json:"labelFormat"`
	// The label stock size specification in length and height. Enum of 4x6 only for now.
	LabelStockSize string `json:"labelStockSize"`
}

type _LabelSpecification LabelSpecification

// NewLabelSpecification instantiates a new LabelSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelSpecification(labelFormat string, labelStockSize string) *LabelSpecification {
	this := LabelSpecification{}
	this.LabelFormat = labelFormat
	this.LabelStockSize = labelStockSize
	return &this
}

// NewLabelSpecificationWithDefaults instantiates a new LabelSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelSpecificationWithDefaults() *LabelSpecification {
	this := LabelSpecification{}
	return &this
}

// GetLabelFormat returns the LabelFormat field value
func (o *LabelSpecification) GetLabelFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelFormat
}

// GetLabelFormatOk returns a tuple with the LabelFormat field value
// and a boolean to check if the value has been set.
func (o *LabelSpecification) GetLabelFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelFormat, true
}

// SetLabelFormat sets field value
func (o *LabelSpecification) SetLabelFormat(v string) {
	o.LabelFormat = v
}

// GetLabelStockSize returns the LabelStockSize field value
func (o *LabelSpecification) GetLabelStockSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelStockSize
}

// GetLabelStockSizeOk returns a tuple with the LabelStockSize field value
// and a boolean to check if the value has been set.
func (o *LabelSpecification) GetLabelStockSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelStockSize, true
}

// SetLabelStockSize sets field value
func (o *LabelSpecification) SetLabelStockSize(v string) {
	o.LabelStockSize = v
}

func (o LabelSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labelFormat"] = o.LabelFormat
	toSerialize["labelStockSize"] = o.LabelStockSize
	return toSerialize, nil
}

type NullableLabelSpecification struct {
	value *LabelSpecification
	isSet bool
}

func (v NullableLabelSpecification) Get() *LabelSpecification {
	return v.value
}

func (v *NullableLabelSpecification) Set(val *LabelSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelSpecification(val *LabelSpecification) *NullableLabelSpecification {
	return &NullableLabelSpecification{value: val, isSet: true}
}

func (v NullableLabelSpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabelSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
