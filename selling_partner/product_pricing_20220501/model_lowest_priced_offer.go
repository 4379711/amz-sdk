package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the LowestPricedOffer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LowestPricedOffer{}

// LowestPricedOffer Describes the lowest priced offers for the specified item condition and offer type.
type LowestPricedOffer struct {
	LowestPricedOffersInput LowestPricedOffersInput `json:"lowestPricedOffersInput"`
	// A list of up to 20 lowest priced offers that match the criteria specified in `lowestPricedOffersInput`.
	Offers []Offer `json:"offers"`
}

type _LowestPricedOffer LowestPricedOffer

// NewLowestPricedOffer instantiates a new LowestPricedOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLowestPricedOffer(lowestPricedOffersInput LowestPricedOffersInput, offers []Offer) *LowestPricedOffer {
	this := LowestPricedOffer{}
	this.LowestPricedOffersInput = lowestPricedOffersInput
	this.Offers = offers
	return &this
}

// NewLowestPricedOfferWithDefaults instantiates a new LowestPricedOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLowestPricedOfferWithDefaults() *LowestPricedOffer {
	this := LowestPricedOffer{}
	return &this
}

// GetLowestPricedOffersInput returns the LowestPricedOffersInput field value
func (o *LowestPricedOffer) GetLowestPricedOffersInput() LowestPricedOffersInput {
	if o == nil {
		var ret LowestPricedOffersInput
		return ret
	}

	return o.LowestPricedOffersInput
}

// GetLowestPricedOffersInputOk returns a tuple with the LowestPricedOffersInput field value
// and a boolean to check if the value has been set.
func (o *LowestPricedOffer) GetLowestPricedOffersInputOk() (*LowestPricedOffersInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LowestPricedOffersInput, true
}

// SetLowestPricedOffersInput sets field value
func (o *LowestPricedOffer) SetLowestPricedOffersInput(v LowestPricedOffersInput) {
	o.LowestPricedOffersInput = v
}

// GetOffers returns the Offers field value
func (o *LowestPricedOffer) GetOffers() []Offer {
	if o == nil {
		var ret []Offer
		return ret
	}

	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value
// and a boolean to check if the value has been set.
func (o *LowestPricedOffer) GetOffersOk() ([]Offer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offers, true
}

// SetOffers sets field value
func (o *LowestPricedOffer) SetOffers(v []Offer) {
	o.Offers = v
}

func (o LowestPricedOffer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lowestPricedOffersInput"] = o.LowestPricedOffersInput
	toSerialize["offers"] = o.Offers
	return toSerialize, nil
}

type NullableLowestPricedOffer struct {
	value *LowestPricedOffer
	isSet bool
}

func (v NullableLowestPricedOffer) Get() *LowestPricedOffer {
	return v.value
}

func (v *NullableLowestPricedOffer) Set(val *LowestPricedOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableLowestPricedOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableLowestPricedOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLowestPricedOffer(val *LowestPricedOffer) *NullableLowestPricedOffer {
	return &NullableLowestPricedOffer{value: val, isSet: true}
}

func (v NullableLowestPricedOffer) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLowestPricedOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
