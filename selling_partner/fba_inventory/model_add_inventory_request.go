package fba_inventory

import (
	"github.com/bytedance/sonic"
)

// checks if the AddInventoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddInventoryRequest{}

// AddInventoryRequest The object with the list of Inventory to be added
type AddInventoryRequest struct {
	// List of Inventory to be added
	InventoryItems []InventoryItem `json:"inventoryItems,omitempty"`
}

// NewAddInventoryRequest instantiates a new AddInventoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddInventoryRequest() *AddInventoryRequest {
	this := AddInventoryRequest{}
	return &this
}

// NewAddInventoryRequestWithDefaults instantiates a new AddInventoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddInventoryRequestWithDefaults() *AddInventoryRequest {
	this := AddInventoryRequest{}
	return &this
}

// GetInventoryItems returns the InventoryItems field value if set, zero value otherwise.
func (o *AddInventoryRequest) GetInventoryItems() []InventoryItem {
	if o == nil || IsNil(o.InventoryItems) {
		var ret []InventoryItem
		return ret
	}
	return o.InventoryItems
}

// GetInventoryItemsOk returns a tuple with the InventoryItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddInventoryRequest) GetInventoryItemsOk() ([]InventoryItem, bool) {
	if o == nil || IsNil(o.InventoryItems) {
		return nil, false
	}
	return o.InventoryItems, true
}

// HasInventoryItems returns a boolean if a field has been set.
func (o *AddInventoryRequest) HasInventoryItems() bool {
	if o != nil && !IsNil(o.InventoryItems) {
		return true
	}

	return false
}

// SetInventoryItems gets a reference to the given []InventoryItem and assigns it to the InventoryItems field.
func (o *AddInventoryRequest) SetInventoryItems(v []InventoryItem) {
	o.InventoryItems = v
}

func (o AddInventoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InventoryItems) {
		toSerialize["inventoryItems"] = o.InventoryItems
	}
	return toSerialize, nil
}

type NullableAddInventoryRequest struct {
	value *AddInventoryRequest
	isSet bool
}

func (v NullableAddInventoryRequest) Get() *AddInventoryRequest {
	return v.value
}

func (v *NullableAddInventoryRequest) Set(val *AddInventoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddInventoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddInventoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddInventoryRequest(val *AddInventoryRequest) *NullableAddInventoryRequest {
	return &NullableAddInventoryRequest{value: val, isSet: true}
}

func (v NullableAddInventoryRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAddInventoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
