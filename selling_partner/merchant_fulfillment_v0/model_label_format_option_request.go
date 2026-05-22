package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the LabelFormatOptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelFormatOptionRequest{}

// LabelFormatOptionRequest Whether to include a packing slip.
type LabelFormatOptionRequest struct {
	// When true, include a packing slip with the label.
	IncludePackingSlipWithLabel *bool `json:"IncludePackingSlipWithLabel,omitempty"`
}

// NewLabelFormatOptionRequest instantiates a new LabelFormatOptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelFormatOptionRequest() *LabelFormatOptionRequest {
	this := LabelFormatOptionRequest{}
	return &this
}

// NewLabelFormatOptionRequestWithDefaults instantiates a new LabelFormatOptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelFormatOptionRequestWithDefaults() *LabelFormatOptionRequest {
	this := LabelFormatOptionRequest{}
	return &this
}

// GetIncludePackingSlipWithLabel returns the IncludePackingSlipWithLabel field value if set, zero value otherwise.
func (o *LabelFormatOptionRequest) GetIncludePackingSlipWithLabel() bool {
	if o == nil || IsNil(o.IncludePackingSlipWithLabel) {
		var ret bool
		return ret
	}
	return *o.IncludePackingSlipWithLabel
}

// GetIncludePackingSlipWithLabelOk returns a tuple with the IncludePackingSlipWithLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelFormatOptionRequest) GetIncludePackingSlipWithLabelOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludePackingSlipWithLabel) {
		return nil, false
	}
	return o.IncludePackingSlipWithLabel, true
}

// HasIncludePackingSlipWithLabel returns a boolean if a field has been set.
func (o *LabelFormatOptionRequest) HasIncludePackingSlipWithLabel() bool {
	if o != nil && !IsNil(o.IncludePackingSlipWithLabel) {
		return true
	}

	return false
}

// SetIncludePackingSlipWithLabel gets a reference to the given bool and assigns it to the IncludePackingSlipWithLabel field.
func (o *LabelFormatOptionRequest) SetIncludePackingSlipWithLabel(v bool) {
	o.IncludePackingSlipWithLabel = &v
}

func (o LabelFormatOptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludePackingSlipWithLabel) {
		toSerialize["IncludePackingSlipWithLabel"] = o.IncludePackingSlipWithLabel
	}
	return toSerialize, nil
}

type NullableLabelFormatOptionRequest struct {
	value *LabelFormatOptionRequest
	isSet bool
}

func (v NullableLabelFormatOptionRequest) Get() *LabelFormatOptionRequest {
	return v.value
}

func (v *NullableLabelFormatOptionRequest) Set(val *LabelFormatOptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelFormatOptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelFormatOptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelFormatOptionRequest(val *LabelFormatOptionRequest) *NullableLabelFormatOptionRequest {
	return &NullableLabelFormatOptionRequest{value: val, isSet: true}
}

func (v NullableLabelFormatOptionRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabelFormatOptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
