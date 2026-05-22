package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the CarrierAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CarrierAccount{}

// CarrierAccount Carrier Account details used to fetch rates.
type CarrierAccount struct {
	// Identifier for the seller's carrier account.
	CarrierAccountId string `json:"carrierAccountId"`
	// The carrier identifier for the offering, provided by the carrier.
	CarrierId string `json:"carrierId"`
}

type _CarrierAccount CarrierAccount

// NewCarrierAccount instantiates a new CarrierAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierAccount(carrierAccountId string, carrierId string) *CarrierAccount {
	this := CarrierAccount{}
	this.CarrierAccountId = carrierAccountId
	this.CarrierId = carrierId
	return &this
}

// NewCarrierAccountWithDefaults instantiates a new CarrierAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierAccountWithDefaults() *CarrierAccount {
	this := CarrierAccount{}
	return &this
}

// GetCarrierAccountId returns the CarrierAccountId field value
func (o *CarrierAccount) GetCarrierAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value
// and a boolean to check if the value has been set.
func (o *CarrierAccount) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierAccountId, true
}

// SetCarrierAccountId sets field value
func (o *CarrierAccount) SetCarrierAccountId(v string) {
	o.CarrierAccountId = v
}

// GetCarrierId returns the CarrierId field value
func (o *CarrierAccount) GetCarrierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierId
}

// GetCarrierIdOk returns a tuple with the CarrierId field value
// and a boolean to check if the value has been set.
func (o *CarrierAccount) GetCarrierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierId, true
}

// SetCarrierId sets field value
func (o *CarrierAccount) SetCarrierId(v string) {
	o.CarrierId = v
}

func (o CarrierAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["carrierAccountId"] = o.CarrierAccountId
	toSerialize["carrierId"] = o.CarrierId
	return toSerialize, nil
}

type NullableCarrierAccount struct {
	value *CarrierAccount
	isSet bool
}

func (v NullableCarrierAccount) Get() *CarrierAccount {
	return v.value
}

func (v *NullableCarrierAccount) Set(val *CarrierAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierAccount(val *CarrierAccount) *NullableCarrierAccount {
	return &NullableCarrierAccount{value: val, isSet: true}
}

func (v NullableCarrierAccount) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCarrierAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
