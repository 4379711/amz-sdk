package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemProcurement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemProcurement{}

// ItemProcurement The vendor procurement information for the listings item.
type ItemProcurement struct {
	CostPrice Money `json:"costPrice"`
}

type _ItemProcurement ItemProcurement

// NewItemProcurement instantiates a new ItemProcurement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemProcurement(costPrice Money) *ItemProcurement {
	this := ItemProcurement{}
	this.CostPrice = costPrice
	return &this
}

// NewItemProcurementWithDefaults instantiates a new ItemProcurement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemProcurementWithDefaults() *ItemProcurement {
	this := ItemProcurement{}
	return &this
}

// GetCostPrice returns the CostPrice field value
func (o *ItemProcurement) GetCostPrice() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.CostPrice
}

// GetCostPriceOk returns a tuple with the CostPrice field value
// and a boolean to check if the value has been set.
func (o *ItemProcurement) GetCostPriceOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostPrice, true
}

// SetCostPrice sets field value
func (o *ItemProcurement) SetCostPrice(v Money) {
	o.CostPrice = v
}

func (o ItemProcurement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["costPrice"] = o.CostPrice
	return toSerialize, nil
}

type NullableItemProcurement struct {
	value *ItemProcurement
	isSet bool
}

func (v NullableItemProcurement) Get() *ItemProcurement {
	return v.value
}

func (v *NullableItemProcurement) Set(val *ItemProcurement) {
	v.value = val
	v.isSet = true
}

func (v NullableItemProcurement) IsSet() bool {
	return v.isSet
}

func (v *NullableItemProcurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemProcurement(val *ItemProcurement) *NullableItemProcurement {
	return &NullableItemProcurement{value: val, isSet: true}
}

func (v NullableItemProcurement) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemProcurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
