package vendor_invoices

import (
	"github.com/bytedance/sonic"
)

// checks if the TaxRegistrationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRegistrationDetails{}

// TaxRegistrationDetails Tax registration details of the entity.
type TaxRegistrationDetails struct {
	// The tax registration type for the entity.
	TaxRegistrationType string `json:"taxRegistrationType"`
	// The tax registration number for the entity. For example, VAT ID, Consumption Tax ID.
	TaxRegistrationNumber string `json:"taxRegistrationNumber"`
}

type _TaxRegistrationDetails TaxRegistrationDetails

// NewTaxRegistrationDetails instantiates a new TaxRegistrationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRegistrationDetails(taxRegistrationType string, taxRegistrationNumber string) *TaxRegistrationDetails {
	this := TaxRegistrationDetails{}
	this.TaxRegistrationType = taxRegistrationType
	this.TaxRegistrationNumber = taxRegistrationNumber
	return &this
}

// NewTaxRegistrationDetailsWithDefaults instantiates a new TaxRegistrationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRegistrationDetailsWithDefaults() *TaxRegistrationDetails {
	this := TaxRegistrationDetails{}
	return &this
}

// GetTaxRegistrationType returns the TaxRegistrationType field value
func (o *TaxRegistrationDetails) GetTaxRegistrationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxRegistrationType
}

// GetTaxRegistrationTypeOk returns a tuple with the TaxRegistrationType field value
// and a boolean to check if the value has been set.
func (o *TaxRegistrationDetails) GetTaxRegistrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRegistrationType, true
}

// SetTaxRegistrationType sets field value
func (o *TaxRegistrationDetails) SetTaxRegistrationType(v string) {
	o.TaxRegistrationType = v
}

// GetTaxRegistrationNumber returns the TaxRegistrationNumber field value
func (o *TaxRegistrationDetails) GetTaxRegistrationNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxRegistrationNumber
}

// GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field value
// and a boolean to check if the value has been set.
func (o *TaxRegistrationDetails) GetTaxRegistrationNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRegistrationNumber, true
}

// SetTaxRegistrationNumber sets field value
func (o *TaxRegistrationDetails) SetTaxRegistrationNumber(v string) {
	o.TaxRegistrationNumber = v
}

func (o TaxRegistrationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["taxRegistrationType"] = o.TaxRegistrationType
	toSerialize["taxRegistrationNumber"] = o.TaxRegistrationNumber
	return toSerialize, nil
}

type NullableTaxRegistrationDetails struct {
	value *TaxRegistrationDetails
	isSet bool
}

func (v NullableTaxRegistrationDetails) Get() *TaxRegistrationDetails {
	return v.value
}

func (v *NullableTaxRegistrationDetails) Set(val *TaxRegistrationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRegistrationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRegistrationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRegistrationDetails(val *TaxRegistrationDetails) *NullableTaxRegistrationDetails {
	return &NullableTaxRegistrationDetails{value: val, isSet: true}
}

func (v NullableTaxRegistrationDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaxRegistrationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
