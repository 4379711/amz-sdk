package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseLabelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseLabelsRequest{}

// PurchaseLabelsRequest The request schema for the purchaseLabels operation.
type PurchaseLabelsRequest struct {
	// An identifier for the rating.
	RateId             string             `json:"rateId"`
	LabelSpecification LabelSpecification `json:"labelSpecification"`
}

type _PurchaseLabelsRequest PurchaseLabelsRequest

// NewPurchaseLabelsRequest instantiates a new PurchaseLabelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseLabelsRequest(rateId string, labelSpecification LabelSpecification) *PurchaseLabelsRequest {
	this := PurchaseLabelsRequest{}
	this.RateId = rateId
	this.LabelSpecification = labelSpecification
	return &this
}

// NewPurchaseLabelsRequestWithDefaults instantiates a new PurchaseLabelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseLabelsRequestWithDefaults() *PurchaseLabelsRequest {
	this := PurchaseLabelsRequest{}
	return &this
}

// GetRateId returns the RateId field value
func (o *PurchaseLabelsRequest) GetRateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RateId
}

// GetRateIdOk returns a tuple with the RateId field value
// and a boolean to check if the value has been set.
func (o *PurchaseLabelsRequest) GetRateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RateId, true
}

// SetRateId sets field value
func (o *PurchaseLabelsRequest) SetRateId(v string) {
	o.RateId = v
}

// GetLabelSpecification returns the LabelSpecification field value
func (o *PurchaseLabelsRequest) GetLabelSpecification() LabelSpecification {
	if o == nil {
		var ret LabelSpecification
		return ret
	}

	return o.LabelSpecification
}

// GetLabelSpecificationOk returns a tuple with the LabelSpecification field value
// and a boolean to check if the value has been set.
func (o *PurchaseLabelsRequest) GetLabelSpecificationOk() (*LabelSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelSpecification, true
}

// SetLabelSpecification sets field value
func (o *PurchaseLabelsRequest) SetLabelSpecification(v LabelSpecification) {
	o.LabelSpecification = v
}

func (o PurchaseLabelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rateId"] = o.RateId
	toSerialize["labelSpecification"] = o.LabelSpecification
	return toSerialize, nil
}

type NullablePurchaseLabelsRequest struct {
	value *PurchaseLabelsRequest
	isSet bool
}

func (v NullablePurchaseLabelsRequest) Get() *PurchaseLabelsRequest {
	return v.value
}

func (v *NullablePurchaseLabelsRequest) Set(val *PurchaseLabelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseLabelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseLabelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseLabelsRequest(val *PurchaseLabelsRequest) *NullablePurchaseLabelsRequest {
	return &NullablePurchaseLabelsRequest{value: val, isSet: true}
}

func (v NullablePurchaseLabelsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseLabelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
