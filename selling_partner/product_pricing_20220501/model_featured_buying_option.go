package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the FeaturedBuyingOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturedBuyingOption{}

// FeaturedBuyingOption Describes a featured buying option, which includes a list of segmented featured offers for a particular item condition.
type FeaturedBuyingOption struct {
	// The buying option type for the featured offer. `buyingOptionType` represents the buying options that a customer receives on the detail page, such as `B2B`, `Fresh`, and `Subscribe n Save`. `buyingOptionType` currently supports `NEW` as a value.
	BuyingOptionType string `json:"buyingOptionType"`
	// A list of segmented featured offers for the current buying option type. A segment can be considered as a group of regional contexts that all have the same featured offer. A regional context is a combination of factors such as customer type, region, or postal code and buying option.
	SegmentedFeaturedOffers []SegmentedFeaturedOffer `json:"segmentedFeaturedOffers"`
}

type _FeaturedBuyingOption FeaturedBuyingOption

// NewFeaturedBuyingOption instantiates a new FeaturedBuyingOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturedBuyingOption(buyingOptionType string, segmentedFeaturedOffers []SegmentedFeaturedOffer) *FeaturedBuyingOption {
	this := FeaturedBuyingOption{}
	this.BuyingOptionType = buyingOptionType
	this.SegmentedFeaturedOffers = segmentedFeaturedOffers
	return &this
}

// NewFeaturedBuyingOptionWithDefaults instantiates a new FeaturedBuyingOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturedBuyingOptionWithDefaults() *FeaturedBuyingOption {
	this := FeaturedBuyingOption{}
	return &this
}

// GetBuyingOptionType returns the BuyingOptionType field value
func (o *FeaturedBuyingOption) GetBuyingOptionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuyingOptionType
}

// GetBuyingOptionTypeOk returns a tuple with the BuyingOptionType field value
// and a boolean to check if the value has been set.
func (o *FeaturedBuyingOption) GetBuyingOptionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuyingOptionType, true
}

// SetBuyingOptionType sets field value
func (o *FeaturedBuyingOption) SetBuyingOptionType(v string) {
	o.BuyingOptionType = v
}

// GetSegmentedFeaturedOffers returns the SegmentedFeaturedOffers field value
func (o *FeaturedBuyingOption) GetSegmentedFeaturedOffers() []SegmentedFeaturedOffer {
	if o == nil {
		var ret []SegmentedFeaturedOffer
		return ret
	}

	return o.SegmentedFeaturedOffers
}

// GetSegmentedFeaturedOffersOk returns a tuple with the SegmentedFeaturedOffers field value
// and a boolean to check if the value has been set.
func (o *FeaturedBuyingOption) GetSegmentedFeaturedOffersOk() ([]SegmentedFeaturedOffer, bool) {
	if o == nil {
		return nil, false
	}
	return o.SegmentedFeaturedOffers, true
}

// SetSegmentedFeaturedOffers sets field value
func (o *FeaturedBuyingOption) SetSegmentedFeaturedOffers(v []SegmentedFeaturedOffer) {
	o.SegmentedFeaturedOffers = v
}

func (o FeaturedBuyingOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buyingOptionType"] = o.BuyingOptionType
	toSerialize["segmentedFeaturedOffers"] = o.SegmentedFeaturedOffers
	return toSerialize, nil
}

type NullableFeaturedBuyingOption struct {
	value *FeaturedBuyingOption
	isSet bool
}

func (v NullableFeaturedBuyingOption) Get() *FeaturedBuyingOption {
	return v.value
}

func (v *NullableFeaturedBuyingOption) Set(val *FeaturedBuyingOption) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturedBuyingOption) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturedBuyingOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturedBuyingOption(val *FeaturedBuyingOption) *NullableFeaturedBuyingOption {
	return &NullableFeaturedBuyingOption{value: val, isSet: true}
}

func (v NullableFeaturedBuyingOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeaturedBuyingOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
