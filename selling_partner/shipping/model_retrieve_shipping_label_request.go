package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the RetrieveShippingLabelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetrieveShippingLabelRequest{}

// RetrieveShippingLabelRequest The request schema for the retrieveShippingLabel operation.
type RetrieveShippingLabelRequest struct {
	LabelSpecification LabelSpecification `json:"labelSpecification"`
}

type _RetrieveShippingLabelRequest RetrieveShippingLabelRequest

// NewRetrieveShippingLabelRequest instantiates a new RetrieveShippingLabelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveShippingLabelRequest(labelSpecification LabelSpecification) *RetrieveShippingLabelRequest {
	this := RetrieveShippingLabelRequest{}
	this.LabelSpecification = labelSpecification
	return &this
}

// NewRetrieveShippingLabelRequestWithDefaults instantiates a new RetrieveShippingLabelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveShippingLabelRequestWithDefaults() *RetrieveShippingLabelRequest {
	this := RetrieveShippingLabelRequest{}
	return &this
}

// GetLabelSpecification returns the LabelSpecification field value
func (o *RetrieveShippingLabelRequest) GetLabelSpecification() LabelSpecification {
	if o == nil {
		var ret LabelSpecification
		return ret
	}

	return o.LabelSpecification
}

// GetLabelSpecificationOk returns a tuple with the LabelSpecification field value
// and a boolean to check if the value has been set.
func (o *RetrieveShippingLabelRequest) GetLabelSpecificationOk() (*LabelSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelSpecification, true
}

// SetLabelSpecification sets field value
func (o *RetrieveShippingLabelRequest) SetLabelSpecification(v LabelSpecification) {
	o.LabelSpecification = v
}

func (o RetrieveShippingLabelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labelSpecification"] = o.LabelSpecification
	return toSerialize, nil
}

type NullableRetrieveShippingLabelRequest struct {
	value *RetrieveShippingLabelRequest
	isSet bool
}

func (v NullableRetrieveShippingLabelRequest) Get() *RetrieveShippingLabelRequest {
	return v.value
}

func (v *NullableRetrieveShippingLabelRequest) Set(val *RetrieveShippingLabelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveShippingLabelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveShippingLabelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveShippingLabelRequest(val *RetrieveShippingLabelRequest) *NullableRetrieveShippingLabelRequest {
	return &NullableRetrieveShippingLabelRequest{value: val, isSet: true}
}

func (v NullableRetrieveShippingLabelRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRetrieveShippingLabelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
