package vendor_shipments

import (
	"github.com/bytedance/sonic"
)

// checks if the ContainerItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerItem{}

// ContainerItem Carton/Pallet level details for the item.
type ContainerItem struct {
	// The reference number for the item. Please provide the itemSequenceNumber from the 'items' segment to refer to that item's details here.
	ItemReference   string       `json:"itemReference"`
	ShippedQuantity ItemQuantity `json:"shippedQuantity"`
	ItemDetails     *ItemDetails `json:"itemDetails,omitempty"`
}

type _ContainerItem ContainerItem

// NewContainerItem instantiates a new ContainerItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerItem(itemReference string, shippedQuantity ItemQuantity) *ContainerItem {
	this := ContainerItem{}
	this.ItemReference = itemReference
	this.ShippedQuantity = shippedQuantity
	return &this
}

// NewContainerItemWithDefaults instantiates a new ContainerItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerItemWithDefaults() *ContainerItem {
	this := ContainerItem{}
	return &this
}

// GetItemReference returns the ItemReference field value
func (o *ContainerItem) GetItemReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemReference
}

// GetItemReferenceOk returns a tuple with the ItemReference field value
// and a boolean to check if the value has been set.
func (o *ContainerItem) GetItemReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemReference, true
}

// SetItemReference sets field value
func (o *ContainerItem) SetItemReference(v string) {
	o.ItemReference = v
}

// GetShippedQuantity returns the ShippedQuantity field value
func (o *ContainerItem) GetShippedQuantity() ItemQuantity {
	if o == nil {
		var ret ItemQuantity
		return ret
	}

	return o.ShippedQuantity
}

// GetShippedQuantityOk returns a tuple with the ShippedQuantity field value
// and a boolean to check if the value has been set.
func (o *ContainerItem) GetShippedQuantityOk() (*ItemQuantity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippedQuantity, true
}

// SetShippedQuantity sets field value
func (o *ContainerItem) SetShippedQuantity(v ItemQuantity) {
	o.ShippedQuantity = v
}

// GetItemDetails returns the ItemDetails field value if set, zero value otherwise.
func (o *ContainerItem) GetItemDetails() ItemDetails {
	if o == nil || IsNil(o.ItemDetails) {
		var ret ItemDetails
		return ret
	}
	return *o.ItemDetails
}

// GetItemDetailsOk returns a tuple with the ItemDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerItem) GetItemDetailsOk() (*ItemDetails, bool) {
	if o == nil || IsNil(o.ItemDetails) {
		return nil, false
	}
	return o.ItemDetails, true
}

// HasItemDetails returns a boolean if a field has been set.
func (o *ContainerItem) HasItemDetails() bool {
	if o != nil && !IsNil(o.ItemDetails) {
		return true
	}

	return false
}

// SetItemDetails gets a reference to the given ItemDetails and assigns it to the ItemDetails field.
func (o *ContainerItem) SetItemDetails(v ItemDetails) {
	o.ItemDetails = &v
}

func (o ContainerItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["itemReference"] = o.ItemReference
	toSerialize["shippedQuantity"] = o.ShippedQuantity
	if !IsNil(o.ItemDetails) {
		toSerialize["itemDetails"] = o.ItemDetails
	}
	return toSerialize, nil
}

type NullableContainerItem struct {
	value *ContainerItem
	isSet bool
}

func (v NullableContainerItem) Get() *ContainerItem {
	return v.value
}

func (v *NullableContainerItem) Set(val *ContainerItem) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerItem) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerItem(val *ContainerItem) *NullableContainerItem {
	return &NullableContainerItem{value: val, isSet: true}
}

func (v NullableContainerItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContainerItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
