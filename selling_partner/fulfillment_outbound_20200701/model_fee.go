package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the Fee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Fee{}

// Fee Fee type and cost.
type Fee struct {
	// The type of fee.
	Name   string `json:"name"`
	Amount Money  `json:"amount"`
}

type _Fee Fee

// NewFee instantiates a new Fee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFee(name string, amount Money) *Fee {
	this := Fee{}
	this.Name = name
	this.Amount = amount
	return &this
}

// NewFeeWithDefaults instantiates a new Fee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeWithDefaults() *Fee {
	this := Fee{}
	return &this
}

// GetName returns the Name field value
func (o *Fee) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Fee) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Fee) SetName(v string) {
	o.Name = v
}

// GetAmount returns the Amount field value
func (o *Fee) GetAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Fee) GetAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Fee) SetAmount(v Money) {
	o.Amount = v
}

func (o Fee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["amount"] = o.Amount
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
