package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the Fee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Fee{}

// Fee Details of the fee.
type Fee struct {
	// The unique identifier assigned to the fee.
	FeeId   string  `json:"feeId"`
	FeeType FeeType `json:"feeType"`
	// A decimal number, such as an amount or FX rate.
	FeeRateValue float32  `json:"feeRateValue"`
	FeeAmount    Currency `json:"feeAmount"`
}

type _Fee Fee

// NewFee instantiates a new Fee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFee(feeId string, feeType FeeType, feeRateValue float32, feeAmount Currency) *Fee {
	this := Fee{}
	this.FeeId = feeId
	this.FeeType = feeType
	this.FeeRateValue = feeRateValue
	this.FeeAmount = feeAmount
	return &this
}

// NewFeeWithDefaults instantiates a new Fee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeWithDefaults() *Fee {
	this := Fee{}
	return &this
}

// GetFeeId returns the FeeId field value
func (o *Fee) GetFeeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeId
}

// GetFeeIdOk returns a tuple with the FeeId field value
// and a boolean to check if the value has been set.
func (o *Fee) GetFeeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeId, true
}

// SetFeeId sets field value
func (o *Fee) SetFeeId(v string) {
	o.FeeId = v
}

// GetFeeType returns the FeeType field value
func (o *Fee) GetFeeType() FeeType {
	if o == nil {
		var ret FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *Fee) GetFeeTypeOk() (*FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *Fee) SetFeeType(v FeeType) {
	o.FeeType = v
}

// GetFeeRateValue returns the FeeRateValue field value
func (o *Fee) GetFeeRateValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FeeRateValue
}

// GetFeeRateValueOk returns a tuple with the FeeRateValue field value
// and a boolean to check if the value has been set.
func (o *Fee) GetFeeRateValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeRateValue, true
}

// SetFeeRateValue sets field value
func (o *Fee) SetFeeRateValue(v float32) {
	o.FeeRateValue = v
}

// GetFeeAmount returns the FeeAmount field value
func (o *Fee) GetFeeAmount() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value
// and a boolean to check if the value has been set.
func (o *Fee) GetFeeAmountOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeAmount, true
}

// SetFeeAmount sets field value
func (o *Fee) SetFeeAmount(v Currency) {
	o.FeeAmount = v
}

func (o Fee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["feeId"] = o.FeeId
	toSerialize["feeType"] = o.FeeType
	toSerialize["feeRateValue"] = o.FeeRateValue
	toSerialize["feeAmount"] = o.FeeAmount
	return toSerialize, nil
}

type NullableFee struct {
	value *Fee
	isSet bool
}

func (v NullableFee) Get() *Fee {
	return v.value
}

func (v *NullableFee) Set(val *Fee) {
	v.value = val
	v.isSet = true
}

func (v NullableFee) IsSet() bool {
	return v.isSet
}

func (v *NullableFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFee(val *Fee) *NullableFee {
	return &NullableFee{value: val, isSet: true}
}

func (v NullableFee) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
