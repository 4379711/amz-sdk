package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the CODSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CODSettings{}

// CODSettings The COD (Cash On Delivery) charges that you associate with a COD fulfillment order.
type CODSettings struct {
	// When true, this fulfillment order requires a COD (Cash On Delivery) payment.
	IsCodRequired     bool   `json:"isCodRequired"`
	CodCharge         *Money `json:"codCharge,omitempty"`
	CodChargeTax      *Money `json:"codChargeTax,omitempty"`
	ShippingCharge    *Money `json:"shippingCharge,omitempty"`
	ShippingChargeTax *Money `json:"shippingChargeTax,omitempty"`
}

type _CODSettings CODSettings

// NewCODSettings instantiates a new CODSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCODSettings(isCodRequired bool) *CODSettings {
	this := CODSettings{}
	this.IsCodRequired = isCodRequired
	return &this
}

// NewCODSettingsWithDefaults instantiates a new CODSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCODSettingsWithDefaults() *CODSettings {
	this := CODSettings{}
	return &this
}

// GetIsCodRequired returns the IsCodRequired field value
func (o *CODSettings) GetIsCodRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCodRequired
}

// GetIsCodRequiredOk returns a tuple with the IsCodRequired field value
// and a boolean to check if the value has been set.
func (o *CODSettings) GetIsCodRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCodRequired, true
}

// SetIsCodRequired sets field value
func (o *CODSettings) SetIsCodRequired(v bool) {
	o.IsCodRequired = v
}

// GetCodCharge returns the CodCharge field value if set, zero value otherwise.
func (o *CODSettings) GetCodCharge() Money {
	if o == nil || IsNil(o.CodCharge) {
		var ret Money
		return ret
	}
	return *o.CodCharge
}

// GetCodChargeOk returns a tuple with the CodCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CODSettings) GetCodChargeOk() (*Money, bool) {
	if o == nil || IsNil(o.CodCharge) {
		return nil, false
	}
	return o.CodCharge, true
}

// HasCodCharge returns a boolean if a field has been set.
func (o *CODSettings) HasCodCharge() bool {
	if o != nil && !IsNil(o.CodCharge) {
		return true
	}

	return false
}

// SetCodCharge gets a reference to the given Money and assigns it to the CodCharge field.
func (o *CODSettings) SetCodCharge(v Money) {
	o.CodCharge = &v
}

// GetCodChargeTax returns the CodChargeTax field value if set, zero value otherwise.
func (o *CODSettings) GetCodChargeTax() Money {
	if o == nil || IsNil(o.CodChargeTax) {
		var ret Money
		return ret
	}
	return *o.CodChargeTax
}

// GetCodChargeTaxOk returns a tuple with the CodChargeTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CODSettings) GetCodChargeTaxOk() (*Money, bool) {
	if o == nil || IsNil(o.CodChargeTax) {
		return nil, false
	}
	return o.CodChargeTax, true
}

// HasCodChargeTax returns a boolean if a field has been set.
func (o *CODSettings) HasCodChargeTax() bool {
	if o != nil && !IsNil(o.CodChargeTax) {
		return true
	}

	return false
}

// SetCodChargeTax gets a reference to the given Money and assigns it to the CodChargeTax field.
func (o *CODSettings) SetCodChargeTax(v Money) {
	o.CodChargeTax = &v
}

// GetShippingCharge returns the ShippingCharge field value if set, zero value otherwise.
func (o *CODSettings) GetShippingCharge() Money {
	if o == nil || IsNil(o.ShippingCharge) {
		var ret Money
		return ret
	}
	return *o.ShippingCharge
}

// GetShippingChargeOk returns a tuple with the ShippingCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CODSettings) GetShippingChargeOk() (*Money, bool) {
	if o == nil || IsNil(o.ShippingCharge) {
		return nil, false
	}
	return o.ShippingCharge, true
}

// HasShippingCharge returns a boolean if a field has been set.
func (o *CODSettings) HasShippingCharge() bool {
	if o != nil && !IsNil(o.ShippingCharge) {
		return true
	}

	return false
}

// SetShippingCharge gets a reference to the given Money and assigns it to the ShippingCharge field.
func (o *CODSettings) SetShippingCharge(v Money) {
	o.ShippingCharge = &v
}

// GetShippingChargeTax returns the ShippingChargeTax field value if set, zero value otherwise.
func (o *CODSettings) GetShippingChargeTax() Money {
	if o == nil || IsNil(o.ShippingChargeTax) {
		var ret Money
		return ret
	}
	return *o.ShippingChargeTax
}

// GetShippingChargeTaxOk returns a tuple with the ShippingChargeTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CODSettings) GetShippingChargeTaxOk() (*Money, bool) {
	if o == nil || IsNil(o.ShippingChargeTax) {
		return nil, false
	}
	return o.ShippingChargeTax, true
}

// HasShippingChargeTax returns a boolean if a field has been set.
func (o *CODSettings) HasShippingChargeTax() bool {
	if o != nil && !IsNil(o.ShippingChargeTax) {
		return true
	}

	return false
}

// SetShippingChargeTax gets a reference to the given Money and assigns it to the ShippingChargeTax field.
func (o *CODSettings) SetShippingChargeTax(v Money) {
	o.ShippingChargeTax = &v
}

func (o CODSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isCodRequired"] = o.IsCodRequired
	if !IsNil(o.CodCharge) {
		toSerialize["codCharge"] = o.CodCharge
	}
	if !IsNil(o.CodChargeTax) {
		toSerialize["codChargeTax"] = o.CodChargeTax
	}
	if !IsNil(o.ShippingCharge) {
		toSerialize["shippingCharge"] = o.ShippingCharge
	}
	if !IsNil(o.ShippingChargeTax) {
		toSerialize["shippingChargeTax"] = o.ShippingChargeTax
	}
	return toSerialize, nil
}

type NullableCODSettings struct {
	value *CODSettings
	isSet bool
}

func (v NullableCODSettings) Get() *CODSettings {
	return v.value
}

func (v *NullableCODSettings) Set(val *CODSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCODSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCODSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCODSettings(val *CODSettings) *NullableCODSettings {
	return &NullableCODSettings{value: val, isSet: true}
}

func (v NullableCODSettings) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCODSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
