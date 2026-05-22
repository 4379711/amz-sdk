package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the TaxRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRate{}

// TaxRate Contains the type and rate of tax.
type TaxRate struct {
	// Rate of cess tax.
	CessRate *float32 `json:"cessRate,omitempty"`
	// Rate of gst tax.
	GstRate *float32 `json:"gstRate,omitempty"`
	// Type of tax. Possible values: `CGST`, `SGST`, `IGST`, `TOTAL_TAX`.
	TaxType *string `json:"taxType,omitempty"`
}

// NewTaxRate instantiates a new TaxRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRate() *TaxRate {
	this := TaxRate{}
	return &this
}

// NewTaxRateWithDefaults instantiates a new TaxRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRateWithDefaults() *TaxRate {
	this := TaxRate{}
	return &this
}

// GetCessRate returns the CessRate field value if set, zero value otherwise.
func (o *TaxRate) GetCessRate() float32 {
	if o == nil || IsNil(o.CessRate) {
		var ret float32
		return ret
	}
	return *o.CessRate
}

// GetCessRateOk returns a tuple with the CessRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRate) GetCessRateOk() (*float32, bool) {
	if o == nil || IsNil(o.CessRate) {
		return nil, false
	}
	return o.CessRate, true
}

// HasCessRate returns a boolean if a field has been set.
func (o *TaxRate) HasCessRate() bool {
	if o != nil && !IsNil(o.CessRate) {
		return true
	}

	return false
}

// SetCessRate gets a reference to the given float32 and assigns it to the CessRate field.
func (o *TaxRate) SetCessRate(v float32) {
	o.CessRate = &v
}

// GetGstRate returns the GstRate field value if set, zero value otherwise.
func (o *TaxRate) GetGstRate() float32 {
	if o == nil || IsNil(o.GstRate) {
		var ret float32
		return ret
	}
	return *o.GstRate
}

// GetGstRateOk returns a tuple with the GstRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRate) GetGstRateOk() (*float32, bool) {
	if o == nil || IsNil(o.GstRate) {
		return nil, false
	}
	return o.GstRate, true
}

// HasGstRate returns a boolean if a field has been set.
func (o *TaxRate) HasGstRate() bool {
	if o != nil && !IsNil(o.GstRate) {
		return true
	}

	return false
}

// SetGstRate gets a reference to the given float32 and assigns it to the GstRate field.
func (o *TaxRate) SetGstRate(v float32) {
	o.GstRate = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *TaxRate) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRate) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *TaxRate) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *TaxRate) SetTaxType(v string) {
	o.TaxType = &v
}

func (o TaxRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CessRate) {
		toSerialize["cessRate"] = o.CessRate
	}
	if !IsNil(o.GstRate) {
		toSerialize["gstRate"] = o.GstRate
	}
	if !IsNil(o.TaxType) {
		toSerialize["taxType"] = o.TaxType
	}
	return toSerialize, nil
}

type NullableTaxRate struct {
	value *TaxRate
	isSet bool
}

func (v NullableTaxRate) Get() *TaxRate {
	return v.value
}

func (v *NullableTaxRate) Set(val *TaxRate) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRate) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRate(val *TaxRate) *NullableTaxRate {
	return &NullableTaxRate{value: val, isSet: true}
}

func (v NullableTaxRate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaxRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
