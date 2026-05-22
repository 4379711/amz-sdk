package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the CustomPlacementInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomPlacementInput{}

// CustomPlacementInput Provide units going to the warehouse.
type CustomPlacementInput struct {
	// Items included while creating Inbound Plan.
	Items []ItemInput `json:"items"`
	// Warehouse Id.
	WarehouseId string `json:"warehouseId"`
}

type _CustomPlacementInput CustomPlacementInput

// NewCustomPlacementInput instantiates a new CustomPlacementInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPlacementInput(items []ItemInput, warehouseId string) *CustomPlacementInput {
	this := CustomPlacementInput{}
	this.Items = items
	this.WarehouseId = warehouseId
	return &this
}

// NewCustomPlacementInputWithDefaults instantiates a new CustomPlacementInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPlacementInputWithDefaults() *CustomPlacementInput {
	this := CustomPlacementInput{}
	return &this
}

// GetItems returns the Items field value
func (o *CustomPlacementInput) GetItems() []ItemInput {
	if o == nil {
		var ret []ItemInput
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CustomPlacementInput) GetItemsOk() ([]ItemInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CustomPlacementInput) SetItems(v []ItemInput) {
	o.Items = v
}

// GetWarehouseId returns the WarehouseId field value
func (o *CustomPlacementInput) GetWarehouseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value
// and a boolean to check if the value has been set.
func (o *CustomPlacementInput) GetWarehouseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarehouseId, true
}

// SetWarehouseId sets field value
func (o *CustomPlacementInput) SetWarehouseId(v string) {
	o.WarehouseId = v
}

func (o CustomPlacementInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["warehouseId"] = o.WarehouseId
	return toSerialize, nil
}

type NullableCustomPlacementInput struct {
	value *CustomPlacementInput
	isSet bool
}

func (v NullableCustomPlacementInput) Get() *CustomPlacementInput {
	return v.value
}

func (v *NullableCustomPlacementInput) Set(val *CustomPlacementInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPlacementInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPlacementInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPlacementInput(val *CustomPlacementInput) *NullableCustomPlacementInput {
	return &NullableCustomPlacementInput{value: val, isSet: true}
}

func (v NullableCustomPlacementInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomPlacementInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
