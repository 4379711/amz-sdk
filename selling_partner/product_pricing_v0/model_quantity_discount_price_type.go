package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the QuantityDiscountPriceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuantityDiscountPriceType{}

// QuantityDiscountPriceType Contains pricing information that includes special pricing when buying in bulk.
type QuantityDiscountPriceType struct {
	// Indicates at what quantity this price becomes active.
	QuantityTier         int32                `json:"quantityTier"`
	QuantityDiscountType QuantityDiscountType `json:"quantityDiscountType"`
	ListingPrice         MoneyType            `json:"listingPrice"`
}

type _QuantityDiscountPriceType QuantityDiscountPriceType

// NewQuantityDiscountPriceType instantiates a new QuantityDiscountPriceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuantityDiscountPriceType(quantityTier int32, quantityDiscountType QuantityDiscountType, listingPrice MoneyType) *QuantityDiscountPriceType {
	this := QuantityDiscountPriceType{}
	this.QuantityTier = quantityTier
	this.QuantityDiscountType = quantityDiscountType
	this.ListingPrice = listingPrice
	return &this
}

// NewQuantityDiscountPriceTypeWithDefaults instantiates a new QuantityDiscountPriceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuantityDiscountPriceTypeWithDefaults() *QuantityDiscountPriceType {
	this := QuantityDiscountPriceType{}
	return &this
}

// GetQuantityTier returns the QuantityTier field value
func (o *QuantityDiscountPriceType) GetQuantityTier() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QuantityTier
}

// GetQuantityTierOk returns a tuple with the QuantityTier field value
// and a boolean to check if the value has been set.
func (o *QuantityDiscountPriceType) GetQuantityTierOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuantityTier, true
}

// SetQuantityTier sets field value
func (o *QuantityDiscountPriceType) SetQuantityTier(v int32) {
	o.QuantityTier = v
}

// GetQuantityDiscountType returns the QuantityDiscountType field value
func (o *QuantityDiscountPriceType) GetQuantityDiscountType() QuantityDiscountType {
	if o == nil {
		var ret QuantityDiscountType
		return ret
	}

	return o.QuantityDiscountType
}

// GetQuantityDiscountTypeOk returns a tuple with the QuantityDiscountType field value
// and a boolean to check if the value has been set.
func (o *QuantityDiscountPriceType) GetQuantityDiscountTypeOk() (*QuantityDiscountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuantityDiscountType, true
}

// SetQuantityDiscountType sets field value
func (o *QuantityDiscountPriceType) SetQuantityDiscountType(v QuantityDiscountType) {
	o.QuantityDiscountType = v
}

// GetListingPrice returns the ListingPrice field value
func (o *QuantityDiscountPriceType) GetListingPrice() MoneyType {
	if o == nil {
		var ret MoneyType
		return ret
	}

	return o.ListingPrice
}

// GetListingPriceOk returns a tuple with the ListingPrice field value
// and a boolean to check if the value has been set.
func (o *QuantityDiscountPriceType) GetListingPriceOk() (*MoneyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingPrice, true
}

// SetListingPrice sets field value
func (o *QuantityDiscountPriceType) SetListingPrice(v MoneyType) {
	o.ListingPrice = v
}

func (o QuantityDiscountPriceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quantityTier"] = o.QuantityTier
	toSerialize["quantityDiscountType"] = o.QuantityDiscountType
	toSerialize["listingPrice"] = o.ListingPrice
	return toSerialize, nil
}

type NullableQuantityDiscountPriceType struct {
	value *QuantityDiscountPriceType
	isSet bool
}

func (v NullableQuantityDiscountPriceType) Get() *QuantityDiscountPriceType {
	return v.value
}

func (v *NullableQuantityDiscountPriceType) Set(val *QuantityDiscountPriceType) {
	v.value = val
	v.isSet = true
}

func (v NullableQuantityDiscountPriceType) IsSet() bool {
	return v.isSet
}

func (v *NullableQuantityDiscountPriceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuantityDiscountPriceType(val *QuantityDiscountPriceType) *NullableQuantityDiscountPriceType {
	return &NullableQuantityDiscountPriceType{value: val, isSet: true}
}

func (v NullableQuantityDiscountPriceType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableQuantityDiscountPriceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
