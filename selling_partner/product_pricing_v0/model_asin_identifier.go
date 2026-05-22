package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ASINIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ASINIdentifier{}

// ASINIdentifier Schema to identify an item by MarketPlaceId and ASIN.
type ASINIdentifier struct {
	// A marketplace identifier.
	MarketplaceId string `json:"MarketplaceId"`
	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN string `json:"ASIN"`
}

type _ASINIdentifier ASINIdentifier

// NewASINIdentifier instantiates a new ASINIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewASINIdentifier(marketplaceId string, aSIN string) *ASINIdentifier {
	this := ASINIdentifier{}
	this.MarketplaceId = marketplaceId
	this.ASIN = aSIN
	return &this
}

// NewASINIdentifierWithDefaults instantiates a new ASINIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewASINIdentifierWithDefaults() *ASINIdentifier {
	this := ASINIdentifier{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *ASINIdentifier) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *ASINIdentifier) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *ASINIdentifier) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetASIN returns the ASIN field value
func (o *ASINIdentifier) GetASIN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ASIN
}

// GetASINOk returns a tuple with the ASIN field value
// and a boolean to check if the value has been set.
func (o *ASINIdentifier) GetASINOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ASIN, true
}

// SetASIN sets field value
func (o *ASINIdentifier) SetASIN(v string) {
	o.ASIN = v
}

func (o ASINIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MarketplaceId"] = o.MarketplaceId
	toSerialize["ASIN"] = o.ASIN
	return toSerialize, nil
}

type NullableASINIdentifier struct {
	value *ASINIdentifier
	isSet bool
}

func (v NullableASINIdentifier) Get() *ASINIdentifier {
	return v.value
}

func (v *NullableASINIdentifier) Set(val *ASINIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableASINIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableASINIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableASINIdentifier(val *ASINIdentifier) *NullableASINIdentifier {
	return &NullableASINIdentifier{value: val, isSet: true}
}

func (v NullableASINIdentifier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableASINIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
