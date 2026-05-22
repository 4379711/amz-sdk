package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the Price type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Price{}

// Price The schema for item's price information, including listing price, shipping price, and Amazon Points.
type Price struct {
	ListingPrice  MoneyType  `json:"listingPrice"`
	ShippingPrice *MoneyType `json:"shippingPrice,omitempty"`
	Points        *Points    `json:"points,omitempty"`
}

type _Price Price

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice(listingPrice MoneyType) *Price {
	this := Price{}
	this.ListingPrice = listingPrice
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetListingPrice returns the ListingPrice field value
func (o *Price) GetListingPrice() MoneyType {
	if o == nil {
		var ret MoneyType
		return ret
	}

	return o.ListingPrice
}

// GetListingPriceOk returns a tuple with the ListingPrice field value
// and a boolean to check if the value has been set.
func (o *Price) GetListingPriceOk() (*MoneyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingPrice, true
}

// SetListingPrice sets field value
func (o *Price) SetListingPrice(v MoneyType) {
	o.ListingPrice = v
}

// GetShippingPrice returns the ShippingPrice field value if set, zero value otherwise.
func (o *Price) GetShippingPrice() MoneyType {
	if o == nil || IsNil(o.ShippingPrice) {
		var ret MoneyType
		return ret
	}
	return *o.ShippingPrice
}

// GetShippingPriceOk returns a tuple with the ShippingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetShippingPriceOk() (*MoneyType, bool) {
	if o == nil || IsNil(o.ShippingPrice) {
		return nil, false
	}
	return o.ShippingPrice, true
}

// HasShippingPrice returns a boolean if a field has been set.
func (o *Price) HasShippingPrice() bool {
	if o != nil && !IsNil(o.ShippingPrice) {
		return true
	}

	return false
}

// SetShippingPrice gets a reference to the given MoneyType and assigns it to the ShippingPrice field.
func (o *Price) SetShippingPrice(v MoneyType) {
	o.ShippingPrice = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *Price) GetPoints() Points {
	if o == nil || IsNil(o.Points) {
		var ret Points
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetPointsOk() (*Points, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *Price) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given Points and assigns it to the Points field.
func (o *Price) SetPoints(v Points) {
	o.Points = &v
}

func (o Price) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["listingPrice"] = o.ListingPrice
	if !IsNil(o.ShippingPrice) {
		toSerialize["shippingPrice"] = o.ShippingPrice
	}
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	return toSerialize, nil
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
