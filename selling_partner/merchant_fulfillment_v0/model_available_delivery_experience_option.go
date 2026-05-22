package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AvailableDeliveryExperienceOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableDeliveryExperienceOption{}

// AvailableDeliveryExperienceOption The available delivery confirmation options, and the fee charged, if any.
type AvailableDeliveryExperienceOption struct {
	DeliveryExperienceOption DeliveryExperienceOption `json:"DeliveryExperienceOption"`
	Charge                   CurrencyAmount           `json:"Charge"`
}

type _AvailableDeliveryExperienceOption AvailableDeliveryExperienceOption

// NewAvailableDeliveryExperienceOption instantiates a new AvailableDeliveryExperienceOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableDeliveryExperienceOption(deliveryExperienceOption DeliveryExperienceOption, charge CurrencyAmount) *AvailableDeliveryExperienceOption {
	this := AvailableDeliveryExperienceOption{}
	this.DeliveryExperienceOption = deliveryExperienceOption
	this.Charge = charge
	return &this
}

// NewAvailableDeliveryExperienceOptionWithDefaults instantiates a new AvailableDeliveryExperienceOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableDeliveryExperienceOptionWithDefaults() *AvailableDeliveryExperienceOption {
	this := AvailableDeliveryExperienceOption{}
	return &this
}

// GetDeliveryExperienceOption returns the DeliveryExperienceOption field value
func (o *AvailableDeliveryExperienceOption) GetDeliveryExperienceOption() DeliveryExperienceOption {
	if o == nil {
		var ret DeliveryExperienceOption
		return ret
	}

	return o.DeliveryExperienceOption
}

// GetDeliveryExperienceOptionOk returns a tuple with the DeliveryExperienceOption field value
// and a boolean to check if the value has been set.
func (o *AvailableDeliveryExperienceOption) GetDeliveryExperienceOptionOk() (*DeliveryExperienceOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryExperienceOption, true
}

// SetDeliveryExperienceOption sets field value
func (o *AvailableDeliveryExperienceOption) SetDeliveryExperienceOption(v DeliveryExperienceOption) {
	o.DeliveryExperienceOption = v
}

// GetCharge returns the Charge field value
func (o *AvailableDeliveryExperienceOption) GetCharge() CurrencyAmount {
	if o == nil {
		var ret CurrencyAmount
		return ret
	}

	return o.Charge
}

// GetChargeOk returns a tuple with the Charge field value
// and a boolean to check if the value has been set.
func (o *AvailableDeliveryExperienceOption) GetChargeOk() (*CurrencyAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Charge, true
}

// SetCharge sets field value
func (o *AvailableDeliveryExperienceOption) SetCharge(v CurrencyAmount) {
	o.Charge = v
}

func (o AvailableDeliveryExperienceOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DeliveryExperienceOption"] = o.DeliveryExperienceOption
	toSerialize["Charge"] = o.Charge
	return toSerialize, nil
}

type NullableAvailableDeliveryExperienceOption struct {
	value *AvailableDeliveryExperienceOption
	isSet bool
}

func (v NullableAvailableDeliveryExperienceOption) Get() *AvailableDeliveryExperienceOption {
	return v.value
}

func (v *NullableAvailableDeliveryExperienceOption) Set(val *AvailableDeliveryExperienceOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableDeliveryExperienceOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableDeliveryExperienceOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableDeliveryExperienceOption(val *AvailableDeliveryExperienceOption) *NullableAvailableDeliveryExperienceOption {
	return &NullableAvailableDeliveryExperienceOption{value: val, isSet: true}
}

func (v NullableAvailableDeliveryExperienceOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAvailableDeliveryExperienceOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
