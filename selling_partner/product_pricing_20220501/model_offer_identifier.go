package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the OfferIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferIdentifier{}

// OfferIdentifier Identifies an offer from a particular seller for a specified ASIN.
type OfferIdentifier struct {
	// A marketplace identifier. Specifies the marketplace for which data is returned.
	MarketplaceId string `json:"marketplaceId"`
	// The seller identifier for the offer.
	SellerId *string `json:"sellerId,omitempty"`
	// The seller SKU of the item. This will only be present for the target offer, which belongs to the requesting seller.
	Sku *string `json:"sku,omitempty"`
	// The ASIN of the item.
	Asin            string           `json:"asin"`
	FulfillmentType *FulfillmentType `json:"fulfillmentType,omitempty"`
}

type _OfferIdentifier OfferIdentifier

// NewOfferIdentifier instantiates a new OfferIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferIdentifier(marketplaceId string, asin string) *OfferIdentifier {
	this := OfferIdentifier{}
	this.MarketplaceId = marketplaceId
	this.Asin = asin
	return &this
}

// NewOfferIdentifierWithDefaults instantiates a new OfferIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferIdentifierWithDefaults() *OfferIdentifier {
	this := OfferIdentifier{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *OfferIdentifier) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *OfferIdentifier) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *OfferIdentifier) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetSellerId returns the SellerId field value if set, zero value otherwise.
func (o *OfferIdentifier) GetSellerId() string {
	if o == nil || IsNil(o.SellerId) {
		var ret string
		return ret
	}
	return *o.SellerId
}

// GetSellerIdOk returns a tuple with the SellerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferIdentifier) GetSellerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SellerId) {
		return nil, false
	}
	return o.SellerId, true
}

// HasSellerId returns a boolean if a field has been set.
func (o *OfferIdentifier) HasSellerId() bool {
	if o != nil && !IsNil(o.SellerId) {
		return true
	}

	return false
}

// SetSellerId gets a reference to the given string and assigns it to the SellerId field.
func (o *OfferIdentifier) SetSellerId(v string) {
	o.SellerId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *OfferIdentifier) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferIdentifier) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *OfferIdentifier) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *OfferIdentifier) SetSku(v string) {
	o.Sku = &v
}

// GetAsin returns the Asin field value
func (o *OfferIdentifier) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *OfferIdentifier) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *OfferIdentifier) SetAsin(v string) {
	o.Asin = v
}

// GetFulfillmentType returns the FulfillmentType field value if set, zero value otherwise.
func (o *OfferIdentifier) GetFulfillmentType() FulfillmentType {
	if o == nil || IsNil(o.FulfillmentType) {
		var ret FulfillmentType
		return ret
	}
	return *o.FulfillmentType
}

// GetFulfillmentTypeOk returns a tuple with the FulfillmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferIdentifier) GetFulfillmentTypeOk() (*FulfillmentType, bool) {
	if o == nil || IsNil(o.FulfillmentType) {
		return nil, false
	}
	return o.FulfillmentType, true
}

// HasFulfillmentType returns a boolean if a field has been set.
func (o *OfferIdentifier) HasFulfillmentType() bool {
	if o != nil && !IsNil(o.FulfillmentType) {
		return true
	}

	return false
}

// SetFulfillmentType gets a reference to the given FulfillmentType and assigns it to the FulfillmentType field.
func (o *OfferIdentifier) SetFulfillmentType(v FulfillmentType) {
	o.FulfillmentType = &v
}

func (o OfferIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	if !IsNil(o.SellerId) {
		toSerialize["sellerId"] = o.SellerId
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	toSerialize["asin"] = o.Asin
	if !IsNil(o.FulfillmentType) {
		toSerialize["fulfillmentType"] = o.FulfillmentType
	}
	return toSerialize, nil
}

type NullableOfferIdentifier struct {
	value *OfferIdentifier
	isSet bool
}

func (v NullableOfferIdentifier) Get() *OfferIdentifier {
	return v.value
}

func (v *NullableOfferIdentifier) Set(val *OfferIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferIdentifier(val *OfferIdentifier) *NullableOfferIdentifier {
	return &NullableOfferIdentifier{value: val, isSet: true}
}

func (v NullableOfferIdentifier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOfferIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
