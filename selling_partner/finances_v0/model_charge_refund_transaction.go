package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ChargeRefundTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeRefundTransaction{}

// ChargeRefundTransaction The charge refund transaction.
type ChargeRefundTransaction struct {
	ChargeAmount *Currency `json:"ChargeAmount,omitempty"`
	// The type of charge.
	ChargeType *string `json:"ChargeType,omitempty"`
}

// NewChargeRefundTransaction instantiates a new ChargeRefundTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeRefundTransaction() *ChargeRefundTransaction {
	this := ChargeRefundTransaction{}
	return &this
}

// NewChargeRefundTransactionWithDefaults instantiates a new ChargeRefundTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeRefundTransactionWithDefaults() *ChargeRefundTransaction {
	this := ChargeRefundTransaction{}
	return &this
}

// GetChargeAmount returns the ChargeAmount field value if set, zero value otherwise.
func (o *ChargeRefundTransaction) GetChargeAmount() Currency {
	if o == nil || IsNil(o.ChargeAmount) {
		var ret Currency
		return ret
	}
	return *o.ChargeAmount
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRefundTransaction) GetChargeAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.ChargeAmount) {
		return nil, false
	}
	return o.ChargeAmount, true
}

// HasChargeAmount returns a boolean if a field has been set.
func (o *ChargeRefundTransaction) HasChargeAmount() bool {
	if o != nil && !IsNil(o.ChargeAmount) {
		return true
	}

	return false
}

// SetChargeAmount gets a reference to the given Currency and assigns it to the ChargeAmount field.
func (o *ChargeRefundTransaction) SetChargeAmount(v Currency) {
	o.ChargeAmount = &v
}

// GetChargeType returns the ChargeType field value if set, zero value otherwise.
func (o *ChargeRefundTransaction) GetChargeType() string {
	if o == nil || IsNil(o.ChargeType) {
		var ret string
		return ret
	}
	return *o.ChargeType
}

// GetChargeTypeOk returns a tuple with the ChargeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRefundTransaction) GetChargeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeType) {
		return nil, false
	}
	return o.ChargeType, true
}

// HasChargeType returns a boolean if a field has been set.
func (o *ChargeRefundTransaction) HasChargeType() bool {
	if o != nil && !IsNil(o.ChargeType) {
		return true
	}

	return false
}

// SetChargeType gets a reference to the given string and assigns it to the ChargeType field.
func (o *ChargeRefundTransaction) SetChargeType(v string) {
	o.ChargeType = &v
}

func (o ChargeRefundTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChargeAmount) {
		toSerialize["ChargeAmount"] = o.ChargeAmount
	}
	if !IsNil(o.ChargeType) {
		toSerialize["ChargeType"] = o.ChargeType
	}
	return toSerialize, nil
}

type NullableChargeRefundTransaction struct {
	value *ChargeRefundTransaction
	isSet bool
}

func (v NullableChargeRefundTransaction) Get() *ChargeRefundTransaction {
	return v.value
}

func (v *NullableChargeRefundTransaction) Set(val *ChargeRefundTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeRefundTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeRefundTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeRefundTransaction(val *ChargeRefundTransaction) *NullableChargeRefundTransaction {
	return &NullableChargeRefundTransaction{value: val, isSet: true}
}

func (v NullableChargeRefundTransaction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableChargeRefundTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
