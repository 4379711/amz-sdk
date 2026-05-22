package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the RetrieveShippingLabelResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetrieveShippingLabelResult{}

// RetrieveShippingLabelResult The payload schema for the retrieveShippingLabel operation.
type RetrieveShippingLabelResult struct {
	// Contains binary image data encoded as a base-64 string.
	LabelStream        string             `json:"labelStream"`
	LabelSpecification LabelSpecification `json:"labelSpecification"`
}

type _RetrieveShippingLabelResult RetrieveShippingLabelResult

// NewRetrieveShippingLabelResult instantiates a new RetrieveShippingLabelResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveShippingLabelResult(labelStream string, labelSpecification LabelSpecification) *RetrieveShippingLabelResult {
	this := RetrieveShippingLabelResult{}
	this.LabelStream = labelStream
	this.LabelSpecification = labelSpecification
	return &this
}

// NewRetrieveShippingLabelResultWithDefaults instantiates a new RetrieveShippingLabelResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveShippingLabelResultWithDefaults() *RetrieveShippingLabelResult {
	this := RetrieveShippingLabelResult{}
	return &this
}

// GetLabelStream returns the LabelStream field value
func (o *RetrieveShippingLabelResult) GetLabelStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelStream
}

// GetLabelStreamOk returns a tuple with the LabelStream field value
// and a boolean to check if the value has been set.
func (o *RetrieveShippingLabelResult) GetLabelStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelStream, true
}

// SetLabelStream sets field value
func (o *RetrieveShippingLabelResult) SetLabelStream(v string) {
	o.LabelStream = v
}

// GetLabelSpecification returns the LabelSpecification field value
func (o *RetrieveShippingLabelResult) GetLabelSpecification() LabelSpecification {
	if o == nil {
		var ret LabelSpecification
		return ret
	}

	return o.LabelSpecification
}

// GetLabelSpecificationOk returns a tuple with the LabelSpecification field value
// and a boolean to check if the value has been set.
func (o *RetrieveShippingLabelResult) GetLabelSpecificationOk() (*LabelSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelSpecification, true
}

// SetLabelSpecification sets field value
func (o *RetrieveShippingLabelResult) SetLabelSpecification(v LabelSpecification) {
	o.LabelSpecification = v
}

func (o RetrieveShippingLabelResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labelStream"] = o.LabelStream
	toSerialize["labelSpecification"] = o.LabelSpecification
	return toSerialize, nil
}

type NullableRetrieveShippingLabelResult struct {
	value *RetrieveShippingLabelResult
	isSet bool
}

func (v NullableRetrieveShippingLabelResult) Get() *RetrieveShippingLabelResult {
	return v.value
}

func (v *NullableRetrieveShippingLabelResult) Set(val *RetrieveShippingLabelResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveShippingLabelResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveShippingLabelResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveShippingLabelResult(val *RetrieveShippingLabelResult) *NullableRetrieveShippingLabelResult {
	return &NullableRetrieveShippingLabelResult{value: val, isSet: true}
}

func (v NullableRetrieveShippingLabelResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRetrieveShippingLabelResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
