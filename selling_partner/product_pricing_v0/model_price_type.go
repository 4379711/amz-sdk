package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the PriceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceType{}

// PriceType Schema for item's price information, including listing price, shipping price, and Amazon points.
type PriceType struct {
	LandedPrice  *MoneyType `json:"LandedPrice,omitempty"`
	ListingPrice MoneyType  `json:"ListingPrice"`
	Shipping     *MoneyType `json:"Shipping,omitempty"`
	Points       *Points    `json:"Points,omitempty"`
}

type _PriceType PriceType

// NewPriceType instantiates a new PriceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceType(listingPrice MoneyType) *PriceType {
	this := PriceType{}
	this.ListingPrice = listingPrice
	return &this
}

// NewPriceTypeWithDefaults instantiates a new PriceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceTypeWithDefaults() *PriceType {
	this := PriceType{}
	return &this
}

// GetLandedPrice returns the LandedPrice field value if set, zero value otherwise.
func (o *PriceType) GetLandedPrice() MoneyType {
	if o == nil || IsNil(o.LandedPrice) {
		var ret MoneyType
		return ret
	}
	return *o.LandedPrice
}

// GetLandedPriceOk returns a tuple with the LandedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceType) GetLandedPriceOk() (*MoneyType, bool) {
	if o == nil || IsNil(o.LandedPrice) {
		return nil, false
	}
	return o.LandedPrice, true
}

// HasLandedPrice returns a boolean if a field has been set.
func (o *PriceType) HasLandedPrice() bool {
	if o != nil && !IsNil(o.LandedPrice) {
		return true
	}

	return false
}

// SetLandedPrice gets a reference to the given MoneyType and assigns it to the LandedPrice field.
func (o *PriceType) SetLandedPrice(v MoneyType) {
	o.LandedPrice = &v
}

// GetListingPrice returns the ListingPrice field value
func (o *PriceType) GetListingPrice() MoneyType {
	if o == nil {
		var ret MoneyType
		return ret
	}

	return o.ListingPrice
}

// GetListingPriceOk returns a tuple with the ListingPrice field value
// and a boolean to check if the value has been set.
func (o *PriceType) GetListingPriceOk() (*MoneyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingPrice, true
}

// SetListingPrice sets field value
func (o *PriceType) SetListingPrice(v MoneyType) {
	o.ListingPrice = v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *PriceType) GetShipping() MoneyType {
	if o == nil || IsNil(o.Shipping) {
		var ret MoneyType
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceType) GetShippingOk() (*MoneyType, bool) {
	if o == nil || IsNil(o.Shipping) {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *PriceType) HasShipping() bool {
	if o != nil && !IsNil(o.Shipping) {
		return true
	}

	return false
}

// SetShipping gets a reference to the given MoneyType and assigns it to the Shipping field.
func (o *PriceType) SetShipping(v MoneyType) {
	o.Shipping = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *PriceType) GetPoints() Points {
	if o == nil || IsNil(o.Points) {
		var ret Points
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceType) GetPointsOk() (*Points, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *PriceType) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given Points and assigns it to the Points field.
func (o *PriceType) SetPoints(v Points) {
	o.Points = &v
}

func (o PriceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LandedPrice) {
		toSerialize["LandedPrice"] = o.LandedPrice
	}
	toSerialize["ListingPrice"] = o.ListingPrice
	if !IsNil(o.Shipping) {
		toSerialize["Shipping"] = o.Shipping
	}
	if !IsNil(o.Points) {
		toSerialize["Points"] = o.Points
	}
	return toSerialize, nil
}

type NullablePriceType struct {
	value *PriceType
	isSet bool
}

func (v NullablePriceType) Get() *PriceType {
	return v.value
}

func (v *NullablePriceType) Set(val *PriceType) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceType) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceType(val *PriceType) *NullablePriceType {
	return &NullablePriceType{value: val, isSet: true}
}

func (v NullablePriceType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePriceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
