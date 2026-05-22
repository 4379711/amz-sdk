package vendor_direct_fulfillment_orders_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the TaxDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxDetails{}

// TaxDetails The tax details for the order. _Note:_ Amazon calculates tax on the list price (Amazon retail price).
type TaxDetails struct {
	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation.
	TaxRate       *string `json:"taxRate,omitempty"`
	TaxAmount     Money   `json:"taxAmount"`
	TaxableAmount *Money  `json:"taxableAmount,omitempty"`
	// Tax type.
	Type *string `json:"type,omitempty"`
}

type _TaxDetails TaxDetails

// NewTaxDetails instantiates a new TaxDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxDetails(taxAmount Money) *TaxDetails {
	this := TaxDetails{}
	this.TaxAmount = taxAmount
	return &this
}

// NewTaxDetailsWithDefaults instantiates a new TaxDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxDetailsWithDefaults() *TaxDetails {
	this := TaxDetails{}
	return &this
}

// GetTaxRate returns the TaxRate field value if set, zero value otherwise.
func (o *TaxDetails) GetTaxRate() string {
	if o == nil || IsNil(o.TaxRate) {
		var ret string
		return ret
	}
	return *o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDetails) GetTaxRateOk() (*string, bool) {
	if o == nil || IsNil(o.TaxRate) {
		return nil, false
	}
	return o.TaxRate, true
}

// HasTaxRate returns a boolean if a field has been set.
func (o *TaxDetails) HasTaxRate() bool {
	if o != nil && !IsNil(o.TaxRate) {
		return true
	}

	return false
}

// SetTaxRate gets a reference to the given string and assigns it to the TaxRate field.
func (o *TaxDetails) SetTaxRate(v string) {
	o.TaxRate = &v
}

// GetTaxAmount returns the TaxAmount field value
func (o *TaxDetails) GetTaxAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value
// and a boolean to check if the value has been set.
func (o *TaxDetails) GetTaxAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxAmount, true
}

// SetTaxAmount sets field value
func (o *TaxDetails) SetTaxAmount(v Money) {
	o.TaxAmount = v
}

// GetTaxableAmount returns the TaxableAmount field value if set, zero value otherwise.
func (o *TaxDetails) GetTaxableAmount() Money {
	if o == nil || IsNil(o.TaxableAmount) {
		var ret Money
		return ret
	}
	return *o.TaxableAmount
}

// GetTaxableAmountOk returns a tuple with the TaxableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDetails) GetTaxableAmountOk() (*Money, bool) {
	if o == nil || IsNil(o.TaxableAmount) {
		return nil, false
	}
	return o.TaxableAmount, true
}

// HasTaxableAmount returns a boolean if a field has been set.
func (o *TaxDetails) HasTaxableAmount() bool {
	if o != nil && !IsNil(o.TaxableAmount) {
		return true
	}

	return false
}

// SetTaxableAmount gets a reference to the given Money and assigns it to the TaxableAmount field.
func (o *TaxDetails) SetTaxableAmount(v Money) {
	o.TaxableAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaxDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaxDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaxDetails) SetType(v string) {
	o.Type = &v
}

func (o TaxDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxRate) {
		toSerialize["taxRate"] = o.TaxRate
	}
	toSerialize["taxAmount"] = o.TaxAmount
	if !IsNil(o.TaxableAmount) {
		toSerialize["taxableAmount"] = o.TaxableAmount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTaxDetails struct {
	value *TaxDetails
	isSet bool
}

func (v NullableTaxDetails) Get() *TaxDetails {
	return v.value
}

func (v *NullableTaxDetails) Set(val *TaxDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxDetails(val *TaxDetails) *NullableTaxDetails {
	return &NullableTaxDetails{value: val, isSet: true}
}

func (v NullableTaxDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaxDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
