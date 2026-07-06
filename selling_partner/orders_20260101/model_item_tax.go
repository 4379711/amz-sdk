package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemTax type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemTax{}

// ItemTax Tax information for an order item.
type ItemTax struct {
	// A list of tax calculation breakdowns for the order item.
	TaxCalculationBreakdowns []ItemTaxCalculationBreakdown `json:"taxCalculationBreakdowns,omitempty"`
	// A list of tax collections for the order item.
	TaxCollections []ItemTaxCollection `json:"taxCollections,omitempty"`
}

// NewItemTax instantiates a new ItemTax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemTax() *ItemTax {
	this := ItemTax{}
	return &this
}

// NewItemTaxWithDefaults instantiates a new ItemTax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemTaxWithDefaults() *ItemTax {
	this := ItemTax{}
	return &this
}

// GetTaxCalculationBreakdowns returns the TaxCalculationBreakdowns field value if set, zero value otherwise.
func (o *ItemTax) GetTaxCalculationBreakdowns() []ItemTaxCalculationBreakdown {
	if o == nil || IsNil(o.TaxCalculationBreakdowns) {
		var ret []ItemTaxCalculationBreakdown
		return ret
	}
	return o.TaxCalculationBreakdowns
}

// GetTaxCalculationBreakdownsOk returns a tuple with the TaxCalculationBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemTax) GetTaxCalculationBreakdownsOk() ([]ItemTaxCalculationBreakdown, bool) {
	if o == nil || IsNil(o.TaxCalculationBreakdowns) {
		return nil, false
	}
	return o.TaxCalculationBreakdowns, true
}

// HasTaxCalculationBreakdowns returns a boolean if a field has been set.
func (o *ItemTax) HasTaxCalculationBreakdowns() bool {
	if o != nil && !IsNil(o.TaxCalculationBreakdowns) {
		return true
	}

	return false
}

// SetTaxCalculationBreakdowns gets a reference to the given []ItemTaxCalculationBreakdown and assigns it to the TaxCalculationBreakdowns field.
func (o *ItemTax) SetTaxCalculationBreakdowns(v []ItemTaxCalculationBreakdown) {
	o.TaxCalculationBreakdowns = v
}

// GetTaxCollections returns the TaxCollections field value if set, zero value otherwise.
func (o *ItemTax) GetTaxCollections() []ItemTaxCollection {
	if o == nil || IsNil(o.TaxCollections) {
		var ret []ItemTaxCollection
		return ret
	}
	return o.TaxCollections
}

// GetTaxCollectionsOk returns a tuple with the TaxCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemTax) GetTaxCollectionsOk() ([]ItemTaxCollection, bool) {
	if o == nil || IsNil(o.TaxCollections) {
		return nil, false
	}
	return o.TaxCollections, true
}

// HasTaxCollections returns a boolean if a field has been set.
func (o *ItemTax) HasTaxCollections() bool {
	if o != nil && !IsNil(o.TaxCollections) {
		return true
	}

	return false
}

// SetTaxCollections gets a reference to the given []ItemTaxCollection and assigns it to the TaxCollections field.
func (o *ItemTax) SetTaxCollections(v []ItemTaxCollection) {
	o.TaxCollections = v
}

func (o ItemTax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxCalculationBreakdowns) {
		toSerialize["taxCalculationBreakdowns"] = o.TaxCalculationBreakdowns
	}
	if !IsNil(o.TaxCollections) {
		toSerialize["taxCollections"] = o.TaxCollections
	}
	return toSerialize, nil
}

type NullableItemTax struct {
	value *ItemTax
	isSet bool
}

func (v NullableItemTax) Get() *ItemTax {
	return v.value
}

func (v *NullableItemTax) Set(val *ItemTax) {
	v.value = val
	v.isSet = true
}

func (v NullableItemTax) IsSet() bool {
	return v.isSet
}

func (v *NullableItemTax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemTax(val *ItemTax) *NullableItemTax {
	return &NullableItemTax{value: val, isSet: true}
}

func (v NullableItemTax) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemTax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
