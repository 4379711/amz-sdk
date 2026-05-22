package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the LabelFormatOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelFormatOption{}

// LabelFormatOption The label format details and whether to include a packing slip.
type LabelFormatOption struct {
	// When true, include a packing slip with the label.
	IncludePackingSlipWithLabel *bool        `json:"IncludePackingSlipWithLabel,omitempty"`
	LabelFormat                 *LabelFormat `json:"LabelFormat,omitempty"`
}

// NewLabelFormatOption instantiates a new LabelFormatOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelFormatOption() *LabelFormatOption {
	this := LabelFormatOption{}
	return &this
}

// NewLabelFormatOptionWithDefaults instantiates a new LabelFormatOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelFormatOptionWithDefaults() *LabelFormatOption {
	this := LabelFormatOption{}
	return &this
}

// GetIncludePackingSlipWithLabel returns the IncludePackingSlipWithLabel field value if set, zero value otherwise.
func (o *LabelFormatOption) GetIncludePackingSlipWithLabel() bool {
	if o == nil || IsNil(o.IncludePackingSlipWithLabel) {
		var ret bool
		return ret
	}
	return *o.IncludePackingSlipWithLabel
}

// GetIncludePackingSlipWithLabelOk returns a tuple with the IncludePackingSlipWithLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelFormatOption) GetIncludePackingSlipWithLabelOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludePackingSlipWithLabel) {
		return nil, false
	}
	return o.IncludePackingSlipWithLabel, true
}

// HasIncludePackingSlipWithLabel returns a boolean if a field has been set.
func (o *LabelFormatOption) HasIncludePackingSlipWithLabel() bool {
	if o != nil && !IsNil(o.IncludePackingSlipWithLabel) {
		return true
	}

	return false
}

// SetIncludePackingSlipWithLabel gets a reference to the given bool and assigns it to the IncludePackingSlipWithLabel field.
func (o *LabelFormatOption) SetIncludePackingSlipWithLabel(v bool) {
	o.IncludePackingSlipWithLabel = &v
}

// GetLabelFormat returns the LabelFormat field value if set, zero value otherwise.
func (o *LabelFormatOption) GetLabelFormat() LabelFormat {
	if o == nil || IsNil(o.LabelFormat) {
		var ret LabelFormat
		return ret
	}
	return *o.LabelFormat
}

// GetLabelFormatOk returns a tuple with the LabelFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelFormatOption) GetLabelFormatOk() (*LabelFormat, bool) {
	if o == nil || IsNil(o.LabelFormat) {
		return nil, false
	}
	return o.LabelFormat, true
}

// HasLabelFormat returns a boolean if a field has been set.
func (o *LabelFormatOption) HasLabelFormat() bool {
	if o != nil && !IsNil(o.LabelFormat) {
		return true
	}

	return false
}

// SetLabelFormat gets a reference to the given LabelFormat and assigns it to the LabelFormat field.
func (o *LabelFormatOption) SetLabelFormat(v LabelFormat) {
	o.LabelFormat = &v
}

func (o LabelFormatOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludePackingSlipWithLabel) {
		toSerialize["IncludePackingSlipWithLabel"] = o.IncludePackingSlipWithLabel
	}
	if !IsNil(o.LabelFormat) {
		toSerialize["LabelFormat"] = o.LabelFormat
	}
	return toSerialize, nil
}

type NullableLabelFormatOption struct {
	value *LabelFormatOption
	isSet bool
}

func (v NullableLabelFormatOption) Get() *LabelFormatOption {
	return v.value
}

func (v *NullableLabelFormatOption) Set(val *LabelFormatOption) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelFormatOption) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelFormatOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelFormatOption(val *LabelFormatOption) *NullableLabelFormatOption {
	return &NullableLabelFormatOption{value: val, isSet: true}
}

func (v NullableLabelFormatOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabelFormatOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
