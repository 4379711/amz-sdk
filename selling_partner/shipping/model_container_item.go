package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the ContainerItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerItem{}

// ContainerItem Item in the container.
type ContainerItem struct {
	// The quantity of the items of this type in the container.
	Quantity   float32  `json:"quantity"`
	UnitPrice  Currency `json:"unitPrice"`
	UnitWeight Weight   `json:"unitWeight"`
	// A descriptive title of the item.
	Title string `json:"title"`
}

type _ContainerItem ContainerItem

// NewContainerItem instantiates a new ContainerItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerItem(quantity float32, unitPrice Currency, unitWeight Weight, title string) *ContainerItem {
	this := ContainerItem{}
	this.Quantity = quantity
	this.UnitPrice = unitPrice
	this.UnitWeight = unitWeight
	this.Title = title
	return &this
}

// NewContainerItemWithDefaults instantiates a new ContainerItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerItemWithDefaults() *ContainerItem {
	this := ContainerItem{}
	return &this
}

// GetQuantity returns the Quantity field value
func (o *ContainerItem) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ContainerItem) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ContainerItem) SetQuantity(v float32) {
	o.Quantity = v
}

// GetUnitPrice returns the UnitPrice field value
func (o *ContainerItem) GetUnitPrice() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value
// and a boolean to check if the value has been set.
func (o *ContainerItem) GetUnitPriceOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPrice, true
}

// SetUnitPrice sets field value
func (o *ContainerItem) SetUnitPrice(v Currency) {
	o.UnitPrice = v
}

// GetUnitWeight returns the UnitWeight field value
func (o *ContainerItem) GetUnitWeight() Weight {
	if o == nil {
		var ret Weight
		return ret
	}

	return o.UnitWeight
}

// GetUnitWeightOk returns a tuple with the UnitWeight field value
// and a boolean to check if the value has been set.
func (o *ContainerItem) GetUnitWeightOk() (*Weight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitWeight, true
}

// SetUnitWeight sets field value
func (o *ContainerItem) SetUnitWeight(v Weight) {
	o.UnitWeight = v
}

// GetTitle returns the Title field value
func (o *ContainerItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ContainerItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ContainerItem) SetTitle(v string) {
	o.Title = v
}

func (o ContainerItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quantity"] = o.Quantity
	toSerialize["unitPrice"] = o.UnitPrice
	toSerialize["unitWeight"] = o.UnitWeight
	toSerialize["title"] = o.Title
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
