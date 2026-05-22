package vendor_invoices

import (
	"github.com/bytedance/sonic"
)

// checks if the ChargeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeDetails{}

// ChargeDetails Monetary and tax details of the charge.
type ChargeDetails struct {
	// Type of the charge applied.
	Type string `json:"type"`
	// Description of the charge.
	Description  *string `json:"description,omitempty"`
	ChargeAmount Money   `json:"chargeAmount"`
	// Tax amount details applied on this charge.
	TaxDetails []TaxDetails `json:"taxDetails,omitempty"`
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

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChargeDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChargeDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChargeDetails) SetDescription(v string) {
	o.Description = &v
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
func (o *ChargeDetails) GetTaxDetails() []TaxDetails {
	if o == nil || IsNil(o.TaxDetails) {
		var ret []TaxDetails
		return ret
	}
	return o.TaxDetails
}

// GetTaxDetailsOk returns a tuple with the TaxDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeDetails) GetTaxDetailsOk() ([]TaxDetails, bool) {
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

// SetTaxDetails gets a reference to the given []TaxDetails and assigns it to the TaxDetails field.
func (o *ChargeDetails) SetTaxDetails(v []TaxDetails) {
	o.TaxDetails = v
}

func (o ChargeDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
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
