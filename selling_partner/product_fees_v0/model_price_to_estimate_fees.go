package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the PriceToEstimateFees type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceToEstimateFees{}

// PriceToEstimateFees Price information for an item, used to estimate fees.
type PriceToEstimateFees struct {
	ListingPrice MoneyType  `json:"ListingPrice"`
	Shipping     *MoneyType `json:"Shipping,omitempty"`
	Points       *Points    `json:"Points,omitempty"`
}

type _PriceToEstimateFees PriceToEstimateFees

// NewPriceToEstimateFees instantiates a new PriceToEstimateFees object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceToEstimateFees(listingPrice MoneyType) *PriceToEstimateFees {
	this := PriceToEstimateFees{}
	this.ListingPrice = listingPrice
	return &this
}

// NewPriceToEstimateFeesWithDefaults instantiates a new PriceToEstimateFees object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceToEstimateFeesWithDefaults() *PriceToEstimateFees {
	this := PriceToEstimateFees{}
	return &this
}

// GetListingPrice returns the ListingPrice field value
func (o *PriceToEstimateFees) GetListingPrice() MoneyType {
	if o == nil {
		var ret MoneyType
		return ret
	}

	return o.ListingPrice
}

// GetListingPriceOk returns a tuple with the ListingPrice field value
// and a boolean to check if the value has been set.
func (o *PriceToEstimateFees) GetListingPriceOk() (*MoneyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingPrice, true
}

// SetListingPrice sets field value
func (o *PriceToEstimateFees) SetListingPrice(v MoneyType) {
	o.ListingPrice = v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *PriceToEstimateFees) GetShipping() MoneyType {
	if o == nil || IsNil(o.Shipping) {
		var ret MoneyType
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceToEstimateFees) GetShippingOk() (*MoneyType, bool) {
	if o == nil || IsNil(o.Shipping) {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *PriceToEstimateFees) HasShipping() bool {
	if o != nil && !IsNil(o.Shipping) {
		return true
	}

	return false
}

// SetShipping gets a reference to the given MoneyType and assigns it to the Shipping field.
func (o *PriceToEstimateFees) SetShipping(v MoneyType) {
	o.Shipping = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *PriceToEstimateFees) GetPoints() Points {
	if o == nil || IsNil(o.Points) {
		var ret Points
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceToEstimateFees) GetPointsOk() (*Points, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *PriceToEstimateFees) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given Points and assigns it to the Points field.
func (o *PriceToEstimateFees) SetPoints(v Points) {
	o.Points = &v
}

func (o PriceToEstimateFees) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ListingPrice"] = o.ListingPrice
	if !IsNil(o.Shipping) {
		toSerialize["Shipping"] = o.Shipping
	}
	if !IsNil(o.Points) {
		toSerialize["Points"] = o.Points
	}
	return toSerialize, nil
}

type NullablePriceToEstimateFees struct {
	value *PriceToEstimateFees
	isSet bool
}

func (v NullablePriceToEstimateFees) Get() *PriceToEstimateFees {
	return v.value
}

func (v *NullablePriceToEstimateFees) Set(val *PriceToEstimateFees) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceToEstimateFees) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceToEstimateFees) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceToEstimateFees(val *PriceToEstimateFees) *NullablePriceToEstimateFees {
	return &NullablePriceToEstimateFees{value: val, isSet: true}
}

func (v NullablePriceToEstimateFees) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePriceToEstimateFees) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
