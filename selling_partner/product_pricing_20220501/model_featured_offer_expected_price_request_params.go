package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the FeaturedOfferExpectedPriceRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturedOfferExpectedPriceRequestParams{}

// FeaturedOfferExpectedPriceRequestParams The parameters for an individual request.
type FeaturedOfferExpectedPriceRequestParams struct {
	// A marketplace identifier. Specifies the marketplace for which data is returned.
	MarketplaceId string `json:"marketplaceId"`
	// The seller SKU of the item.
	Sku     string   `json:"sku"`
	Segment *Segment `json:"segment,omitempty"`
}

type _FeaturedOfferExpectedPriceRequestParams FeaturedOfferExpectedPriceRequestParams

// NewFeaturedOfferExpectedPriceRequestParams instantiates a new FeaturedOfferExpectedPriceRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturedOfferExpectedPriceRequestParams(marketplaceId string, sku string) *FeaturedOfferExpectedPriceRequestParams {
	this := FeaturedOfferExpectedPriceRequestParams{}
	this.MarketplaceId = marketplaceId
	this.Sku = sku
	return &this
}

// NewFeaturedOfferExpectedPriceRequestParamsWithDefaults instantiates a new FeaturedOfferExpectedPriceRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturedOfferExpectedPriceRequestParamsWithDefaults() *FeaturedOfferExpectedPriceRequestParams {
	this := FeaturedOfferExpectedPriceRequestParams{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *FeaturedOfferExpectedPriceRequestParams) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *FeaturedOfferExpectedPriceRequestParams) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *FeaturedOfferExpectedPriceRequestParams) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetSku returns the Sku field value
func (o *FeaturedOfferExpectedPriceRequestParams) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *FeaturedOfferExpectedPriceRequestParams) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *FeaturedOfferExpectedPriceRequestParams) SetSku(v string) {
	o.Sku = v
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *FeaturedOfferExpectedPriceRequestParams) GetSegment() Segment {
	if o == nil || IsNil(o.Segment) {
		var ret Segment
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturedOfferExpectedPriceRequestParams) GetSegmentOk() (*Segment, bool) {
	if o == nil || IsNil(o.Segment) {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *FeaturedOfferExpectedPriceRequestParams) HasSegment() bool {
	if o != nil && !IsNil(o.Segment) {
		return true
	}

	return false
}

// SetSegment gets a reference to the given Segment and assigns it to the Segment field.
func (o *FeaturedOfferExpectedPriceRequestParams) SetSegment(v Segment) {
	o.Segment = &v
}

func (o FeaturedOfferExpectedPriceRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	toSerialize["sku"] = o.Sku
	if !IsNil(o.Segment) {
		toSerialize["segment"] = o.Segment
	}
	return toSerialize, nil
}

type NullableFeaturedOfferExpectedPriceRequestParams struct {
	value *FeaturedOfferExpectedPriceRequestParams
	isSet bool
}

func (v NullableFeaturedOfferExpectedPriceRequestParams) Get() *FeaturedOfferExpectedPriceRequestParams {
	return v.value
}

func (v *NullableFeaturedOfferExpectedPriceRequestParams) Set(val *FeaturedOfferExpectedPriceRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturedOfferExpectedPriceRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturedOfferExpectedPriceRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturedOfferExpectedPriceRequestParams(val *FeaturedOfferExpectedPriceRequestParams) *NullableFeaturedOfferExpectedPriceRequestParams {
	return &NullableFeaturedOfferExpectedPriceRequestParams{value: val, isSet: true}
}

func (v NullableFeaturedOfferExpectedPriceRequestParams) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeaturedOfferExpectedPriceRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
