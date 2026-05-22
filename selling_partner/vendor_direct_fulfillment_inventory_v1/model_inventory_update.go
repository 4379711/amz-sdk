package vendor_direct_fulfillment_inventory_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the InventoryUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryUpdate{}

// InventoryUpdate Inventory details required to update some or all items for the requested warehouse.
type InventoryUpdate struct {
	SellingParty PartyIdentification `json:"sellingParty"`
	// When true, this request contains a full feed. Otherwise, this request contains a partial feed. When sending a full feed, you must send information about all items in the warehouse. Any items not in the full feed are updated as not available. When sending a partial feed, only include the items that need an update to inventory. The status of other items will remain unchanged.
	IsFullUpdate bool `json:"isFullUpdate"`
	// A list of inventory items with updated details, including quantity available.
	Items []ItemDetails `json:"items"`
}

type _InventoryUpdate InventoryUpdate

// NewInventoryUpdate instantiates a new InventoryUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryUpdate(sellingParty PartyIdentification, isFullUpdate bool, items []ItemDetails) *InventoryUpdate {
	this := InventoryUpdate{}
	this.SellingParty = sellingParty
	this.IsFullUpdate = isFullUpdate
	this.Items = items
	return &this
}

// NewInventoryUpdateWithDefaults instantiates a new InventoryUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryUpdateWithDefaults() *InventoryUpdate {
	this := InventoryUpdate{}
	return &this
}

// GetSellingParty returns the SellingParty field value
func (o *InventoryUpdate) GetSellingParty() PartyIdentification {
	if o == nil {
		var ret PartyIdentification
		return ret
	}

	return o.SellingParty
}

// GetSellingPartyOk returns a tuple with the SellingParty field value
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetSellingPartyOk() (*PartyIdentification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingParty, true
}

// SetSellingParty sets field value
func (o *InventoryUpdate) SetSellingParty(v PartyIdentification) {
	o.SellingParty = v
}

// GetIsFullUpdate returns the IsFullUpdate field value
func (o *InventoryUpdate) GetIsFullUpdate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFullUpdate
}

// GetIsFullUpdateOk returns a tuple with the IsFullUpdate field value
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetIsFullUpdateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFullUpdate, true
}

// SetIsFullUpdate sets field value
func (o *InventoryUpdate) SetIsFullUpdate(v bool) {
	o.IsFullUpdate = v
}

// GetItems returns the Items field value
func (o *InventoryUpdate) GetItems() []ItemDetails {
	if o == nil {
		var ret []ItemDetails
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetItemsOk() ([]ItemDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InventoryUpdate) SetItems(v []ItemDetails) {
	o.Items = v
}

func (o InventoryUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellingParty"] = o.SellingParty
	toSerialize["isFullUpdate"] = o.IsFullUpdate
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableInventoryUpdate struct {
	value *InventoryUpdate
	isSet bool
}

func (v NullableInventoryUpdate) Get() *InventoryUpdate {
	return v.value
}

func (v *NullableInventoryUpdate) Set(val *InventoryUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryUpdate(val *InventoryUpdate) *NullableInventoryUpdate {
	return &NullableInventoryUpdate{value: val, isSet: true}
}

func (v NullableInventoryUpdate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInventoryUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
