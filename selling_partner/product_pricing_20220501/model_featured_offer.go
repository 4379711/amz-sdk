package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the FeaturedOffer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturedOffer{}

// FeaturedOffer Schema for `currentFeaturedOffer` or `competingFeaturedOffer`.
type FeaturedOffer struct {
	OfferIdentifier OfferIdentifier `json:"offerIdentifier"`
	Condition       *Condition      `json:"condition,omitempty"`
	Price           *Price          `json:"price,omitempty"`
}

type _FeaturedOffer FeaturedOffer

// NewFeaturedOffer instantiates a new FeaturedOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturedOffer(offerIdentifier OfferIdentifier) *FeaturedOffer {
	this := FeaturedOffer{}
	this.OfferIdentifier = offerIdentifier
	return &this
}

// NewFeaturedOfferWithDefaults instantiates a new FeaturedOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturedOfferWithDefaults() *FeaturedOffer {
	this := FeaturedOffer{}
	return &this
}

// GetOfferIdentifier returns the OfferIdentifier field value
func (o *FeaturedOffer) GetOfferIdentifier() OfferIdentifier {
	if o == nil {
		var ret OfferIdentifier
		return ret
	}

	return o.OfferIdentifier
}

// GetOfferIdentifierOk returns a tuple with the OfferIdentifier field value
// and a boolean to check if the value has been set.
func (o *FeaturedOffer) GetOfferIdentifierOk() (*OfferIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferIdentifier, true
}

// SetOfferIdentifier sets field value
func (o *FeaturedOffer) SetOfferIdentifier(v OfferIdentifier) {
	o.OfferIdentifier = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *FeaturedOffer) GetCondition() Condition {
	if o == nil || IsNil(o.Condition) {
		var ret Condition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturedOffer) GetConditionOk() (*Condition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *FeaturedOffer) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given Condition and assigns it to the Condition field.
func (o *FeaturedOffer) SetCondition(v Condition) {
	o.Condition = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *FeaturedOffer) GetPrice() Price {
	if o == nil || IsNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturedOffer) GetPriceOk() (*Price, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *FeaturedOffer) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *FeaturedOffer) SetPrice(v Price) {
	o.Price = &v
}

func (o FeaturedOffer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerIdentifier"] = o.OfferIdentifier
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableFeaturedOffer struct {
	value *FeaturedOffer
	isSet bool
}

func (v NullableFeaturedOffer) Get() *FeaturedOffer {
	return v.value
}

func (v *NullableFeaturedOffer) Set(val *FeaturedOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturedOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturedOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturedOffer(val *FeaturedOffer) *NullableFeaturedOffer {
	return &NullableFeaturedOffer{value: val, isSet: true}
}

func (v NullableFeaturedOffer) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeaturedOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
