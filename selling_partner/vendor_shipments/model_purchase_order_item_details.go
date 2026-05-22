package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseOrderItemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseOrderItemDetails{}

// PurchaseOrderItemDetails Item details for be provided for every item in shipment at either the item or carton or pallet level, whichever is appropriate.
type PurchaseOrderItemDetails struct {
	MaximumRetailPrice *Money `json:"maximumRetailPrice,omitempty"`
}

// NewPurchaseOrderItemDetails instantiates a new PurchaseOrderItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseOrderItemDetails() *PurchaseOrderItemDetails {
	this := PurchaseOrderItemDetails{}
	return &this
}

// NewPurchaseOrderItemDetailsWithDefaults instantiates a new PurchaseOrderItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseOrderItemDetailsWithDefaults() *PurchaseOrderItemDetails {
	this := PurchaseOrderItemDetails{}
	return &this
}

// GetMaximumRetailPrice returns the MaximumRetailPrice field value if set, zero value otherwise.
func (o *PurchaseOrderItemDetails) GetMaximumRetailPrice() Money {
	if o == nil || IsNil(o.MaximumRetailPrice) {
		var ret Money
		return ret
	}
	return *o.MaximumRetailPrice
}

// GetMaximumRetailPriceOk returns a tuple with the MaximumRetailPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrderItemDetails) GetMaximumRetailPriceOk() (*Money, bool) {
	if o == nil || IsNil(o.MaximumRetailPrice) {
		return nil, false
	}
	return o.MaximumRetailPrice, true
}

// HasMaximumRetailPrice returns a boolean if a field has been set.
func (o *PurchaseOrderItemDetails) HasMaximumRetailPrice() bool {
	if o != nil && !IsNil(o.MaximumRetailPrice) {
		return true
	}

	return false
}

// SetMaximumRetailPrice gets a reference to the given Money and assigns it to the MaximumRetailPrice field.
func (o *PurchaseOrderItemDetails) SetMaximumRetailPrice(v Money) {
	o.MaximumRetailPrice = &v
}

func (o PurchaseOrderItemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaximumRetailPrice) {
		toSerialize["maximumRetailPrice"] = o.MaximumRetailPrice
	}
	return toSerialize, nil
}

type NullablePurchaseOrderItemDetails struct {
	value *PurchaseOrderItemDetails
	isSet bool
}

func (v NullablePurchaseOrderItemDetails) Get() *PurchaseOrderItemDetails {
	return v.value
}

func (v *NullablePurchaseOrderItemDetails) Set(val *PurchaseOrderItemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderItemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderItemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderItemDetails(val *PurchaseOrderItemDetails) *NullablePurchaseOrderItemDetails {
	return &NullablePurchaseOrderItemDetails{value: val, isSet: true}
}

func (v NullablePurchaseOrderItemDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseOrderItemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
