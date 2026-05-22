package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AvailableCarrierWillPickUpOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableCarrierWillPickUpOption{}

// AvailableCarrierWillPickUpOption Indicates whether the carrier will pick up the package, and what fee is charged, if any.
type AvailableCarrierWillPickUpOption struct {
	CarrierWillPickUpOption CarrierWillPickUpOption `json:"CarrierWillPickUpOption"`
	Charge                  CurrencyAmount          `json:"Charge"`
}

type _AvailableCarrierWillPickUpOption AvailableCarrierWillPickUpOption

// NewAvailableCarrierWillPickUpOption instantiates a new AvailableCarrierWillPickUpOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableCarrierWillPickUpOption(carrierWillPickUpOption CarrierWillPickUpOption, charge CurrencyAmount) *AvailableCarrierWillPickUpOption {
	this := AvailableCarrierWillPickUpOption{}
	this.CarrierWillPickUpOption = carrierWillPickUpOption
	this.Charge = charge
	return &this
}

// NewAvailableCarrierWillPickUpOptionWithDefaults instantiates a new AvailableCarrierWillPickUpOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableCarrierWillPickUpOptionWithDefaults() *AvailableCarrierWillPickUpOption {
	this := AvailableCarrierWillPickUpOption{}
	return &this
}

// GetCarrierWillPickUpOption returns the CarrierWillPickUpOption field value
func (o *AvailableCarrierWillPickUpOption) GetCarrierWillPickUpOption() CarrierWillPickUpOption {
	if o == nil {
		var ret CarrierWillPickUpOption
		return ret
	}

	return o.CarrierWillPickUpOption
}

// GetCarrierWillPickUpOptionOk returns a tuple with the CarrierWillPickUpOption field value
// and a boolean to check if the value has been set.
func (o *AvailableCarrierWillPickUpOption) GetCarrierWillPickUpOptionOk() (*CarrierWillPickUpOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierWillPickUpOption, true
}

// SetCarrierWillPickUpOption sets field value
func (o *AvailableCarrierWillPickUpOption) SetCarrierWillPickUpOption(v CarrierWillPickUpOption) {
	o.CarrierWillPickUpOption = v
}

// GetCharge returns the Charge field value
func (o *AvailableCarrierWillPickUpOption) GetCharge() CurrencyAmount {
	if o == nil {
		var ret CurrencyAmount
		return ret
	}

	return o.Charge
}

// GetChargeOk returns a tuple with the Charge field value
// and a boolean to check if the value has been set.
func (o *AvailableCarrierWillPickUpOption) GetChargeOk() (*CurrencyAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Charge, true
}

// SetCharge sets field value
func (o *AvailableCarrierWillPickUpOption) SetCharge(v CurrencyAmount) {
	o.Charge = v
}

func (o AvailableCarrierWillPickUpOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CarrierWillPickUpOption"] = o.CarrierWillPickUpOption
	toSerialize["Charge"] = o.Charge
	return toSerialize, nil
}

type NullableAvailableCarrierWillPickUpOption struct {
	value *AvailableCarrierWillPickUpOption
	isSet bool
}

func (v NullableAvailableCarrierWillPickUpOption) Get() *AvailableCarrierWillPickUpOption {
	return v.value
}

func (v *NullableAvailableCarrierWillPickUpOption) Set(val *AvailableCarrierWillPickUpOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableCarrierWillPickUpOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableCarrierWillPickUpOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableCarrierWillPickUpOption(val *AvailableCarrierWillPickUpOption) *NullableAvailableCarrierWillPickUpOption {
	return &NullableAvailableCarrierWillPickUpOption{value: val, isSet: true}
}

func (v NullableAvailableCarrierWillPickUpOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAvailableCarrierWillPickUpOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
