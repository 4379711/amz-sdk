package vendor_direct_fulfillment_payments_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ChargeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeDetails{}

// ChargeDetails Monetary and tax details of the charge.
type ChargeDetails struct {
	// Type of charge applied.
	Type         string `json:"type"`
	ChargeAmount Money  `json:"chargeAmount"`
	// Individual tax details per line item.
	TaxDetails []TaxDetail `json:"taxDetails,omitempty"`
}

type _ChargeDetails ChargeDetails

// NewChargeDetails instantiates a new ChargeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeDetails(type_ string, chargeAmount Money) *ChargeDetails {
	this := ChargeDetails{}
	this.Type = type_
	this.ChargeAmount = chargeAmount
	return &this
}

// NewChargeDetailsWithDefaults instantiates a new ChargeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeDetailsWithDefaults() *ChargeDetails {
	this := ChargeDetails{}
	return &this
}

// GetType returns the Type field value
func (o *ChargeDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChargeDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChargeDetails) SetType(v string) {
	o.Type = v
}

// GetChargeAmount returns the ChargeAmount field value
func (o *ChargeDetails) GetChargeAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.ChargeAmount
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value
// and a boolean to check if the value has been set.
func (o *ChargeDetails) GetChargeAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChargeAmount, true
}

// SetChargeAmount sets field value
func (o *ChargeDetails) SetChargeAmount(v Money) {
	o.ChargeAmount = v
}

// GetTaxDetails returns the TaxDetails field value if set, zero value otherwise.
func (o *ChargeDetails) GetTaxDetails() []TaxDetail {
	if o == nil || IsNil(o.TaxDetails) {
		var ret []TaxDetail
		return ret
	}
	return o.TaxDetails
}

// GetTaxDetailsOk returns a tuple with the TaxDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeDetails) GetTaxDetailsOk() ([]TaxDetail, bool) {
	if o == nil || IsNil(o.TaxDetails) {
		return nil, false
	}
	return o.TaxDetails, true
}

// HasTaxDetails returns a boolean if a field has been set.
func (o *ChargeDetails) HasTaxDetails() bool {
	if o != nil && !IsNil(o.TaxDetails) {
		return true
	}

	return false
}

// SetTaxDetails gets a reference to the given []TaxDetail and assigns it to the TaxDetails field.
func (o *ChargeDetails) SetTaxDetails(v []TaxDetail) {
	o.TaxDetails = v
}

func (o ChargeDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["chargeAmount"] = o.ChargeAmount
	if !IsNil(o.TaxDetails) {
		toSerialize["taxDetails"] = o.TaxDetails
	}
	return toSerialize, nil
}

type NullableChargeDetails struct {
	value *ChargeDetails
	isSet bool
}

func (v NullableChargeDetails) Get() *ChargeDetails {
	return v.value
}

func (v *NullableChargeDetails) Set(val *ChargeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeDetails(val *ChargeDetails) *NullableChargeDetails {
	return &NullableChargeDetails{value: val, isSet: true}
}

func (v NullableChargeDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableChargeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
