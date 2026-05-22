package vendor_direct_fulfillment_orders_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the TaxItemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxItemDetails{}

// TaxItemDetails Total tax details for the line item.
type TaxItemDetails struct {
	// A list of tax line items.
	TaxLineItem []TaxDetails `json:"taxLineItem,omitempty"`
}

// NewTaxItemDetails instantiates a new TaxItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxItemDetails() *TaxItemDetails {
	this := TaxItemDetails{}
	return &this
}

// NewTaxItemDetailsWithDefaults instantiates a new TaxItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxItemDetailsWithDefaults() *TaxItemDetails {
	this := TaxItemDetails{}
	return &this
}

// GetTaxLineItem returns the TaxLineItem field value if set, zero value otherwise.
func (o *TaxItemDetails) GetTaxLineItem() []TaxDetails {
	if o == nil || IsNil(o.TaxLineItem) {
		var ret []TaxDetails
		return ret
	}
	return o.TaxLineItem
}

// GetTaxLineItemOk returns a tuple with the TaxLineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxItemDetails) GetTaxLineItemOk() ([]TaxDetails, bool) {
	if o == nil || IsNil(o.TaxLineItem) {
		return nil, false
	}
	return o.TaxLineItem, true
}

// HasTaxLineItem returns a boolean if a field has been set.
func (o *TaxItemDetails) HasTaxLineItem() bool {
	if o != nil && !IsNil(o.TaxLineItem) {
		return true
	}

	return false
}

// SetTaxLineItem gets a reference to the given []TaxDetails and assigns it to the TaxLineItem field.
func (o *TaxItemDetails) SetTaxLineItem(v []TaxDetails) {
	o.TaxLineItem = v
}

func (o TaxItemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxLineItem) {
		toSerialize["taxLineItem"] = o.TaxLineItem
	}
	return toSerialize, nil
}

type NullableTaxItemDetails struct {
	value *TaxItemDetails
	isSet bool
}

func (v NullableTaxItemDetails) Get() *TaxItemDetails {
	return v.value
}

func (v *NullableTaxItemDetails) Set(val *TaxItemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxItemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxItemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxItemDetails(val *TaxItemDetails) *NullableTaxItemDetails {
	return &NullableTaxItemDetails{value: val, isSet: true}
}

func (v NullableTaxItemDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTaxItemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
