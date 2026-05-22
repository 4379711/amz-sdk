package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ChargeComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeComponent{}

// ChargeComponent The type and amount of a charge applied on a package.
type ChargeComponent struct {
	Amount *Currency `json:"amount,omitempty"`
	// The type of charge.
	ChargeType *string `json:"chargeType,omitempty"`
}

// NewChargeComponent instantiates a new ChargeComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeComponent() *ChargeComponent {
	this := ChargeComponent{}
	return &this
}

// NewChargeComponentWithDefaults instantiates a new ChargeComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeComponentWithDefaults() *ChargeComponent {
	this := ChargeComponent{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ChargeComponent) GetAmount() Currency {
	if o == nil || IsNil(o.Amount) {
		var ret Currency
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeComponent) GetAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ChargeComponent) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Currency and assigns it to the Amount field.
func (o *ChargeComponent) SetAmount(v Currency) {
	o.Amount = &v
}

// GetChargeType returns the ChargeType field value if set, zero value otherwise.
func (o *ChargeComponent) GetChargeType() string {
	if o == nil || IsNil(o.ChargeType) {
		var ret string
		return ret
	}
	return *o.ChargeType
}

// GetChargeTypeOk returns a tuple with the ChargeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeComponent) GetChargeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeType) {
		return nil, false
	}
	return o.ChargeType, true
}

// HasChargeType returns a boolean if a field has been set.
func (o *ChargeComponent) HasChargeType() bool {
	if o != nil && !IsNil(o.ChargeType) {
		return true
	}

	return false
}

// SetChargeType gets a reference to the given string and assigns it to the ChargeType field.
func (o *ChargeComponent) SetChargeType(v string) {
	o.ChargeType = &v
}

func (o ChargeComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.ChargeType) {
		toSerialize["chargeType"] = o.ChargeType
	}
	return toSerialize, nil
}

type NullableChargeComponent struct {
	value *ChargeComponent
	isSet bool
}

func (v NullableChargeComponent) Get() *ChargeComponent {
	return v.value
}

func (v *NullableChargeComponent) Set(val *ChargeComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeComponent(val *ChargeComponent) *NullableChargeComponent {
	return &NullableChargeComponent{value: val, isSet: true}
}

func (v NullableChargeComponent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableChargeComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
