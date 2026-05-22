package vendor_direct_fulfillment_inventory_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemDetails{}

// ItemDetails Updated inventory details for an item.
type ItemDetails struct {
	// The buyer selected product identification of the item. Either buyerProductIdentifier or vendorProductIdentifier should be submitted.
	BuyerProductIdentifier *string `json:"buyerProductIdentifier,omitempty"`
	// The vendor selected product identification of the item. Either buyerProductIdentifier or vendorProductIdentifier should be submitted.
	VendorProductIdentifier *string      `json:"vendorProductIdentifier,omitempty"`
	AvailableQuantity       ItemQuantity `json:"availableQuantity"`
	// When true, the item is permanently unavailable.
	IsObsolete *bool `json:"isObsolete,omitempty"`
}

type _ItemDetails ItemDetails

// NewItemDetails instantiates a new ItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemDetails(availableQuantity ItemQuantity) *ItemDetails {
	this := ItemDetails{}
	this.AvailableQuantity = availableQuantity
	return &this
}

// NewItemDetailsWithDefaults instantiates a new ItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemDetailsWithDefaults() *ItemDetails {
	this := ItemDetails{}
	return &this
}

// GetBuyerProductIdentifier returns the BuyerProductIdentifier field value if set, zero value otherwise.
func (o *ItemDetails) GetBuyerProductIdentifier() string {
	if o == nil || IsNil(o.BuyerProductIdentifier) {
		var ret string
		return ret
	}
	return *o.BuyerProductIdentifier
}

// GetBuyerProductIdentifierOk returns a tuple with the BuyerProductIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDetails) GetBuyerProductIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerProductIdentifier) {
		return nil, false
	}
	return o.BuyerProductIdentifier, true
}

// HasBuyerProductIdentifier returns a boolean if a field has been set.
func (o *ItemDetails) HasBuyerProductIdentifier() bool {
	if o != nil && !IsNil(o.BuyerProductIdentifier) {
		return true
	}

	return false
}

// SetBuyerProductIdentifier gets a reference to the given string and assigns it to the BuyerProductIdentifier field.
func (o *ItemDetails) SetBuyerProductIdentifier(v string) {
	o.BuyerProductIdentifier = &v
}

// GetVendorProductIdentifier returns the VendorProductIdentifier field value if set, zero value otherwise.
func (o *ItemDetails) GetVendorProductIdentifier() string {
	if o == nil || IsNil(o.VendorProductIdentifier) {
		var ret string
		return ret
	}
	return *o.VendorProductIdentifier
}

// GetVendorProductIdentifierOk returns a tuple with the VendorProductIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDetails) GetVendorProductIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.VendorProductIdentifier) {
		return nil, false
	}
	return o.VendorProductIdentifier, true
}

// HasVendorProductIdentifier returns a boolean if a field has been set.
func (o *ItemDetails) HasVendorProductIdentifier() bool {
	if o != nil && !IsNil(o.VendorProductIdentifier) {
		return true
	}

	return false
}

// SetVendorProductIdentifier gets a reference to the given string and assigns it to the VendorProductIdentifier field.
func (o *ItemDetails) SetVendorProductIdentifier(v string) {
	o.VendorProductIdentifier = &v
}

// GetAvailableQuantity returns the AvailableQuantity field value
func (o *ItemDetails) GetAvailableQuantity() ItemQuantity {
	if o == nil {
		var ret ItemQuantity
		return ret
	}

	return o.AvailableQuantity
}

// GetAvailableQuantityOk returns a tuple with the AvailableQuantity field value
// and a boolean to check if the value has been set.
func (o *ItemDetails) GetAvailableQuantityOk() (*ItemQuantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableQuantity, true
}

// SetAvailableQuantity sets field value
func (o *ItemDetails) SetAvailableQuantity(v ItemQuantity) {
	o.AvailableQuantity = v
}

// GetIsObsolete returns the IsObsolete field value if set, zero value otherwise.
func (o *ItemDetails) GetIsObsolete() bool {
	if o == nil || IsNil(o.IsObsolete) {
		var ret bool
		return ret
	}
	return *o.IsObsolete
}

// GetIsObsoleteOk returns a tuple with the IsObsolete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemDetails) GetIsObsoleteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsObsolete) {
		return nil, false
	}
	return o.IsObsolete, true
}

// HasIsObsolete returns a boolean if a field has been set.
func (o *ItemDetails) HasIsObsolete() bool {
	if o != nil && !IsNil(o.IsObsolete) {
		return true
	}

	return false
}

// SetIsObsolete gets a reference to the given bool and assigns it to the IsObsolete field.
func (o *ItemDetails) SetIsObsolete(v bool) {
	o.IsObsolete = &v
}

func (o ItemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyerProductIdentifier) {
		toSerialize["buyerProductIdentifier"] = o.BuyerProductIdentifier
	}
	if !IsNil(o.VendorProductIdentifier) {
		toSerialize["vendorProductIdentifier"] = o.VendorProductIdentifier
	}
	toSerialize["availableQuantity"] = o.AvailableQuantity
	if !IsNil(o.IsObsolete) {
		toSerialize["isObsolete"] = o.IsObsolete
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
