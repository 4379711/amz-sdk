package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseCondition{}

// PurchaseCondition Minimum purchase requirements using quantity-based or amount-based conditions for basket building promotions. You can specify conditions within purchase requirements and multi-tier benefit configurations. You must specify exactly one of `quantityThreshold` or `amountThreshold`.
type PurchaseCondition struct {
	QuantityThreshold *QuantityThreshold `json:"quantityThreshold,omitempty"`
	AmountThreshold   *AmountThreshold   `json:"amountThreshold,omitempty"`
}

// NewPurchaseCondition instantiates a new PurchaseCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseCondition() *PurchaseCondition {
	this := PurchaseCondition{}
	return &this
}

// NewPurchaseConditionWithDefaults instantiates a new PurchaseCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseConditionWithDefaults() *PurchaseCondition {
	this := PurchaseCondition{}
	return &this
}

// GetQuantityThreshold returns the QuantityThreshold field value if set, zero value otherwise.
func (o *PurchaseCondition) GetQuantityThreshold() QuantityThreshold {
	if o == nil || IsNil(o.QuantityThreshold) {
		var ret QuantityThreshold
		return ret
	}
	return *o.QuantityThreshold
}

// GetQuantityThresholdOk returns a tuple with the QuantityThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseCondition) GetQuantityThresholdOk() (*QuantityThreshold, bool) {
	if o == nil || IsNil(o.QuantityThreshold) {
		return nil, false
	}
	return o.QuantityThreshold, true
}

// HasQuantityThreshold returns a boolean if a field has been set.
func (o *PurchaseCondition) HasQuantityThreshold() bool {
	if o != nil && !IsNil(o.QuantityThreshold) {
		return true
	}

	return false
}

// SetQuantityThreshold gets a reference to the given QuantityThreshold and assigns it to the QuantityThreshold field.
func (o *PurchaseCondition) SetQuantityThreshold(v QuantityThreshold) {
	o.QuantityThreshold = &v
}

// GetAmountThreshold returns the AmountThreshold field value if set, zero value otherwise.
func (o *PurchaseCondition) GetAmountThreshold() AmountThreshold {
	if o == nil || IsNil(o.AmountThreshold) {
		var ret AmountThreshold
		return ret
	}
	return *o.AmountThreshold
}

// GetAmountThresholdOk returns a tuple with the AmountThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseCondition) GetAmountThresholdOk() (*AmountThreshold, bool) {
	if o == nil || IsNil(o.AmountThreshold) {
		return nil, false
	}
	return o.AmountThreshold, true
}

// HasAmountThreshold returns a boolean if a field has been set.
func (o *PurchaseCondition) HasAmountThreshold() bool {
	if o != nil && !IsNil(o.AmountThreshold) {
		return true
	}

	return false
}

// SetAmountThreshold gets a reference to the given AmountThreshold and assigns it to the AmountThreshold field.
func (o *PurchaseCondition) SetAmountThreshold(v AmountThreshold) {
	o.AmountThreshold = &v
}

func (o PurchaseCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuantityThreshold) {
		toSerialize["quantityThreshold"] = o.QuantityThreshold
	}
	if !IsNil(o.AmountThreshold) {
		toSerialize["amountThreshold"] = o.AmountThreshold
	}
	return toSerialize, nil
}

type NullablePurchaseCondition struct {
	value *PurchaseCondition
	isSet bool
}

func (v NullablePurchaseCondition) Get() *PurchaseCondition {
	return v.value
}

func (v *NullablePurchaseCondition) Set(val *PurchaseCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseCondition(val *PurchaseCondition) *NullablePurchaseCondition {
	return &NullablePurchaseCondition{value: val, isSet: true}
}

func (v NullablePurchaseCondition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
