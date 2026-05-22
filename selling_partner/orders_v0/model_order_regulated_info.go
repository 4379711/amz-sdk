package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderRegulatedInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderRegulatedInfo{}

// OrderRegulatedInfo The order's regulated information along with its verification status.
type OrderRegulatedInfo struct {
	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId        string               `json:"AmazonOrderId"`
	RegulatedInformation RegulatedInformation `json:"RegulatedInformation"`
	// When true, the order requires attaching a dosage information label when shipped.
	RequiresDosageLabel              bool                             `json:"RequiresDosageLabel"`
	RegulatedOrderVerificationStatus RegulatedOrderVerificationStatus `json:"RegulatedOrderVerificationStatus"`
}

type _OrderRegulatedInfo OrderRegulatedInfo

// NewOrderRegulatedInfo instantiates a new OrderRegulatedInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderRegulatedInfo(amazonOrderId string, regulatedInformation RegulatedInformation, requiresDosageLabel bool, regulatedOrderVerificationStatus RegulatedOrderVerificationStatus) *OrderRegulatedInfo {
	this := OrderRegulatedInfo{}
	this.AmazonOrderId = amazonOrderId
	this.RegulatedInformation = regulatedInformation
	this.RequiresDosageLabel = requiresDosageLabel
	this.RegulatedOrderVerificationStatus = regulatedOrderVerificationStatus
	return &this
}

// NewOrderRegulatedInfoWithDefaults instantiates a new OrderRegulatedInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderRegulatedInfoWithDefaults() *OrderRegulatedInfo {
	this := OrderRegulatedInfo{}
	return &this
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *OrderRegulatedInfo) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderRegulatedInfo) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *OrderRegulatedInfo) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

// GetRegulatedInformation returns the RegulatedInformation field value
func (o *OrderRegulatedInfo) GetRegulatedInformation() RegulatedInformation {
	if o == nil {
		var ret RegulatedInformation
		return ret
	}

	return o.RegulatedInformation
}

// GetRegulatedInformationOk returns a tuple with the RegulatedInformation field value
// and a boolean to check if the value has been set.
func (o *OrderRegulatedInfo) GetRegulatedInformationOk() (*RegulatedInformation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegulatedInformation, true
}

// SetRegulatedInformation sets field value
func (o *OrderRegulatedInfo) SetRegulatedInformation(v RegulatedInformation) {
	o.RegulatedInformation = v
}

// GetRequiresDosageLabel returns the RequiresDosageLabel field value
func (o *OrderRegulatedInfo) GetRequiresDosageLabel() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequiresDosageLabel
}

// GetRequiresDosageLabelOk returns a tuple with the RequiresDosageLabel field value
// and a boolean to check if the value has been set.
func (o *OrderRegulatedInfo) GetRequiresDosageLabelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiresDosageLabel, true
}

// SetRequiresDosageLabel sets field value
func (o *OrderRegulatedInfo) SetRequiresDosageLabel(v bool) {
	o.RequiresDosageLabel = v
}

// GetRegulatedOrderVerificationStatus returns the RegulatedOrderVerificationStatus field value
func (o *OrderRegulatedInfo) GetRegulatedOrderVerificationStatus() RegulatedOrderVerificationStatus {
	if o == nil {
		var ret RegulatedOrderVerificationStatus
		return ret
	}

	return o.RegulatedOrderVerificationStatus
}

// GetRegulatedOrderVerificationStatusOk returns a tuple with the RegulatedOrderVerificationStatus field value
// and a boolean to check if the value has been set.
func (o *OrderRegulatedInfo) GetRegulatedOrderVerificationStatusOk() (*RegulatedOrderVerificationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegulatedOrderVerificationStatus, true
}

// SetRegulatedOrderVerificationStatus sets field value
func (o *OrderRegulatedInfo) SetRegulatedOrderVerificationStatus(v RegulatedOrderVerificationStatus) {
	o.RegulatedOrderVerificationStatus = v
}

func (o OrderRegulatedInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AmazonOrderId"] = o.AmazonOrderId
	toSerialize["RegulatedInformation"] = o.RegulatedInformation
	toSerialize["RequiresDosageLabel"] = o.RequiresDosageLabel
	toSerialize["RegulatedOrderVerificationStatus"] = o.RegulatedOrderVerificationStatus
	return toSerialize, nil
}

type NullableOrderRegulatedInfo struct {
	value *OrderRegulatedInfo
	isSet bool
}

func (v NullableOrderRegulatedInfo) Get() *OrderRegulatedInfo {
	return v.value
}

func (v *NullableOrderRegulatedInfo) Set(val *OrderRegulatedInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderRegulatedInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderRegulatedInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderRegulatedInfo(val *OrderRegulatedInfo) *NullableOrderRegulatedInfo {
	return &NullableOrderRegulatedInfo{value: val, isSet: true}
}

func (v NullableOrderRegulatedInfo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderRegulatedInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
