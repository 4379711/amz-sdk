package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemDetails{}

// ItemDetails Item details for be provided for every item in shipment at either the item or carton or pallet level, whichever is appropriate.
type ItemDetails struct {
	// The purchase order number for the shipment being confirmed. If the items in this shipment belong to multiple purchase order numbers that are in particular carton or pallet within the shipment, then provide the purchaseOrderNumber at the appropriate carton or pallet level. Formatting Notes: 8-character alpha-numeric code.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// The batch or lot number associates an item with information the manufacturer considers relevant for traceability of the trade item to which the Element String is applied. The data may refer to the trade item itself or to items contained. This field is mandatory for all perishable items.
	LotNumber          *string `json:"lotNumber,omitempty"`
	Expiry             *Expiry `json:"expiry,omitempty"`
	MaximumRetailPrice *Money  `json:"maximumRetailPrice,omitempty"`
	// Identification of the instructions on how specified item/carton/pallet should be handled.
	HandlingCode *string `json:"handlingCode,omitempty"`
}

// NewItemDetails instantiates a new ItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemDetails() *ItemDetails {
	this := ItemDetails{}
	return &this
}

// NewItemDetailsWithDefaults instantiates a new ItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemDetailsWithDefaults() *ItemDetails {
	this := ItemDetails{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise.
func (o *ItemDetails) GetPurchaseOrderNumber() string {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDetails) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		return nil, false
	}
	return o.PurchaseOrderNumber, true
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *ItemDetails) HasPurchaseOrderNumber() bool {
	if o != nil && !IsNil(o.PurchaseOrderNumber) {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given string and assigns it to the PurchaseOrderNumber field.
func (o *ItemDetails) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = &v
}

// GetLotNumber returns the LotNumber field value if set, zero value otherwise.
func (o *ItemDetails) GetLotNumber() string {
	if o == nil || IsNil(o.LotNumber) {
		var ret string
		return ret
	}
	return *o.LotNumber
}

// GetLotNumberOk returns a tuple with the LotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDetails) GetLotNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LotNumber) {
		return nil, false
	}
	return o.LotNumber, true
}

// HasLotNumber returns a boolean if a field has been set.
func (o *ItemDetails) HasLotNumber() bool {
	if o != nil && !IsNil(o.LotNumber) {
		return true
	}

	return false
}

// SetLotNumber gets a reference to the given string and assigns it to the LotNumber field.
func (o *ItemDetails) SetLotNumber(v string) {
	o.LotNumber = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *ItemDetails) GetExpiry() Expiry {
	if o == nil || IsNil(o.Expiry) {
		var ret Expiry
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDetails) GetExpiryOk() (*Expiry, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *ItemDetails) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given Expiry and assigns it to the Expiry field.
func (o *ItemDetails) SetExpiry(v Expiry) {
	o.Expiry = &v
}

// GetMaximumRetailPrice returns the MaximumRetailPrice field value if set, zero value otherwise.
func (o *ItemDetails) GetMaximumRetailPrice() Money {
	if o == nil || IsNil(o.MaximumRetailPrice) {
		var ret Money
		return ret
	}
	return *o.MaximumRetailPrice
}

// GetMaximumRetailPriceOk returns a tuple with the MaximumRetailPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDetails) GetMaximumRetailPriceOk() (*Money, bool) {
	if o == nil || IsNil(o.MaximumRetailPrice) {
		return nil, false
	}
	return o.MaximumRetailPrice, true
}

// HasMaximumRetailPrice returns a boolean if a field has been set.
func (o *ItemDetails) HasMaximumRetailPrice() bool {
	if o != nil && !IsNil(o.MaximumRetailPrice) {
		return true
	}

	return false
}

// SetMaximumRetailPrice gets a reference to the given Money and assigns it to the MaximumRetailPrice field.
func (o *ItemDetails) SetMaximumRetailPrice(v Money) {
	o.MaximumRetailPrice = &v
}

// GetHandlingCode returns the HandlingCode field value if set, zero value otherwise.
func (o *ItemDetails) GetHandlingCode() string {
	if o == nil || IsNil(o.HandlingCode) {
		var ret string
		return ret
	}
	return *o.HandlingCode
}

// GetHandlingCodeOk returns a tuple with the HandlingCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDetails) GetHandlingCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HandlingCode) {
		return nil, false
	}
	return o.HandlingCode, true
}

// HasHandlingCode returns a boolean if a field has been set.
func (o *ItemDetails) HasHandlingCode() bool {
	if o != nil && !IsNil(o.HandlingCode) {
		return true
	}

	return false
}

// SetHandlingCode gets a reference to the given string and assigns it to the HandlingCode field.
func (o *ItemDetails) SetHandlingCode(v string) {
	o.HandlingCode = &v
}

func (o ItemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PurchaseOrderNumber) {
		toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	}
	if !IsNil(o.LotNumber) {
		toSerialize["lotNumber"] = o.LotNumber
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.MaximumRetailPrice) {
		toSerialize["maximumRetailPrice"] = o.MaximumRetailPrice
	}
	if !IsNil(o.HandlingCode) {
		toSerialize["handlingCode"] = o.HandlingCode
	}
	return toSerialize, nil
}

type NullableItemDetails struct {
	value *ItemDetails
	isSet bool
}

func (v NullableItemDetails) Get() *ItemDetails {
	return v.value
}

func (v *NullableItemDetails) Set(val *ItemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableItemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableItemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemDetails(val *ItemDetails) *NullableItemDetails {
	return &NullableItemDetails{value: val, isSet: true}
}

func (v NullableItemDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
