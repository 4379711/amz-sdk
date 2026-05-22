package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the TaxDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxDetail{}

// TaxDetail Indicates the tax specifications associated with the shipment for customs compliance purposes in certain regions.
type TaxDetail struct {
	TaxType TaxType `json:"taxType"`
	// The shipper's tax registration number associated with the shipment for customs compliance purposes in certain regions.
	TaxRegistrationNumber string `json:"taxRegistrationNumber"`
}

type _TaxDetail TaxDetail

// NewTaxDetail instantiates a new TaxDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxDetail(taxType TaxType, taxRegistrationNumber string) *TaxDetail {
	this := TaxDetail{}
	this.TaxType = taxType
	this.TaxRegistrationNumber = taxRegistrationNumber
	return &this
}

// NewTaxDetailWithDefaults instantiates a new TaxDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxDetailWithDefaults() *TaxDetail {
	this := TaxDetail{}
	return &this
}

// GetTaxType returns the TaxType field value
func (o *TaxDetail) GetTaxType() TaxType {
	if o == nil {
		var ret TaxType
		return ret
	}

	return o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value
// and a boolean to check if the value has been set.
func (o *TaxDetail) GetTaxTypeOk() (*TaxType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxType, true
}

// SetTaxType sets field value
func (o *TaxDetail) SetTaxType(v TaxType) {
	o.TaxType = v
}

// GetTaxRegistrationNumber returns the TaxRegistrationNumber field value
func (o *TaxDetail) GetTaxRegistrationNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxRegistrationNumber
}

// GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field value
// and a boolean to check if the value has been set.
func (o *TaxDetail) GetTaxRegistrationNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRegistrationNumber, true
}

// SetTaxRegistrationNumber sets field value
func (o *TaxDetail) SetTaxRegistrationNumber(v string) {
	o.TaxRegistrationNumber = v
}

func (o TaxDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["taxType"] = o.TaxType
	toSerialize["taxRegistrationNumber"] = o.TaxRegistrationNumber
	return toSerialize, nil
}

type NullableTaxDetail struct {
	value *TaxDetail
	isSet bool
}

func (v NullableTaxDetail) Get() *TaxDetail {
	return v.value
}

func (v *NullableTaxDetail) Set(val *TaxDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxDetail(val *TaxDetail) *NullableTaxDetail {
	return &NullableTaxDetail{value: val, isSet: true}
}

func (v NullableTaxDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaxDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
