package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the CompetitiveSummaryResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompetitiveSummaryResponseBody{}

// CompetitiveSummaryResponseBody The `competitiveSummaryResponse` body for a requested ASIN and `marketplaceId`.
type CompetitiveSummaryResponseBody struct {
	// The ASIN of the item.
	Asin string `json:"asin"`
	// A marketplace identifier. Specifies the marketplace for which data is returned.
	MarketplaceId string `json:"marketplaceId"`
	// A list of featured buying options for the specified ASIN `marketplaceId` combination.
	FeaturedBuyingOptions []FeaturedBuyingOption `json:"featuredBuyingOptions,omitempty"`
	// A list of lowest priced offers for the specified ASIN `marketplaceId` combination.
	LowestPricedOffers []LowestPricedOffer `json:"lowestPricedOffers,omitempty"`
	// A list of reference prices for the specified ASIN `marketplaceId` combination.
	ReferencePrices []ReferencePrice `json:"referencePrices,omitempty"`
	// A list of error responses that are returned when a request is unsuccessful.
	Errors []Error `json:"errors,omitempty"`
}

type _CompetitiveSummaryResponseBody CompetitiveSummaryResponseBody

// NewCompetitiveSummaryResponseBody instantiates a new CompetitiveSummaryResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompetitiveSummaryResponseBody(asin string, marketplaceId string) *CompetitiveSummaryResponseBody {
	this := CompetitiveSummaryResponseBody{}
	this.Asin = asin
	this.MarketplaceId = marketplaceId
	return &this
}

// NewCompetitiveSummaryResponseBodyWithDefaults instantiates a new CompetitiveSummaryResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompetitiveSummaryResponseBodyWithDefaults() *CompetitiveSummaryResponseBody {
	this := CompetitiveSummaryResponseBody{}
	return &this
}

// GetAsin returns the Asin field value
func (o *CompetitiveSummaryResponseBody) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryResponseBody) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *CompetitiveSummaryResponseBody) SetAsin(v string) {
	o.Asin = v
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *CompetitiveSummaryResponseBody) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryResponseBody) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *CompetitiveSummaryResponseBody) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetFeaturedBuyingOptions returns the FeaturedBuyingOptions field value if set, zero value otherwise.
func (o *CompetitiveSummaryResponseBody) GetFeaturedBuyingOptions() []FeaturedBuyingOption {
	if o == nil || IsNil(o.FeaturedBuyingOptions) {
		var ret []FeaturedBuyingOption
		return ret
	}
	return o.FeaturedBuyingOptions
}

// GetFeaturedBuyingOptionsOk returns a tuple with the FeaturedBuyingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryResponseBody) GetFeaturedBuyingOptionsOk() ([]FeaturedBuyingOption, bool) {
	if o == nil || IsNil(o.FeaturedBuyingOptions) {
		return nil, false
	}
	return o.FeaturedBuyingOptions, true
}

// HasFeaturedBuyingOptions returns a boolean if a field has been set.
func (o *CompetitiveSummaryResponseBody) HasFeaturedBuyingOptions() bool {
	if o != nil && !IsNil(o.FeaturedBuyingOptions) {
		return true
	}

	return false
}

// SetFeaturedBuyingOptions gets a reference to the given []FeaturedBuyingOption and assigns it to the FeaturedBuyingOptions field.
func (o *CompetitiveSummaryResponseBody) SetFeaturedBuyingOptions(v []FeaturedBuyingOption) {
	o.FeaturedBuyingOptions = v
}

// GetLowestPricedOffers returns the LowestPricedOffers field value if set, zero value otherwise.
func (o *CompetitiveSummaryResponseBody) GetLowestPricedOffers() []LowestPricedOffer {
	if o == nil || IsNil(o.LowestPricedOffers) {
		var ret []LowestPricedOffer
		return ret
	}
	return o.LowestPricedOffers
}

// GetLowestPricedOffersOk returns a tuple with the LowestPricedOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryResponseBody) GetLowestPricedOffersOk() ([]LowestPricedOffer, bool) {
	if o == nil || IsNil(o.LowestPricedOffers) {
		return nil, false
	}
	return o.LowestPricedOffers, true
}

// HasLowestPricedOffers returns a boolean if a field has been set.
func (o *CompetitiveSummaryResponseBody) HasLowestPricedOffers() bool {
	if o != nil && !IsNil(o.LowestPricedOffers) {
		return true
	}

	return false
}

// SetLowestPricedOffers gets a reference to the given []LowestPricedOffer and assigns it to the LowestPricedOffers field.
func (o *CompetitiveSummaryResponseBody) SetLowestPricedOffers(v []LowestPricedOffer) {
	o.LowestPricedOffers = v
}

// GetReferencePrices returns the ReferencePrices field value if set, zero value otherwise.
func (o *CompetitiveSummaryResponseBody) GetReferencePrices() []ReferencePrice {
	if o == nil || IsNil(o.ReferencePrices) {
		var ret []ReferencePrice
		return ret
	}
	return o.ReferencePrices
}

// GetReferencePricesOk returns a tuple with the ReferencePrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryResponseBody) GetReferencePricesOk() ([]ReferencePrice, bool) {
	if o == nil || IsNil(o.ReferencePrices) {
		return nil, false
	}
	return o.ReferencePrices, true
}

// HasReferencePrices returns a boolean if a field has been set.
func (o *CompetitiveSummaryResponseBody) HasReferencePrices() bool {
	if o != nil && !IsNil(o.ReferencePrices) {
		return true
	}

	return false
}

// SetReferencePrices gets a reference to the given []ReferencePrice and assigns it to the ReferencePrices field.
func (o *CompetitiveSummaryResponseBody) SetReferencePrices(v []ReferencePrice) {
	o.ReferencePrices = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CompetitiveSummaryResponseBody) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompetitiveSummaryResponseBody) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CompetitiveSummaryResponseBody) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *CompetitiveSummaryResponseBody) SetErrors(v []Error) {
	o.Errors = v
}

func (o CompetitiveSummaryResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asin"] = o.Asin
	toSerialize["marketplaceId"] = o.MarketplaceId
	if !IsNil(o.FeaturedBuyingOptions) {
		toSerialize["featuredBuyingOptions"] = o.FeaturedBuyingOptions
	}
	if !IsNil(o.LowestPricedOffers) {
		toSerialize["lowestPricedOffers"] = o.LowestPricedOffers
	}
	if !IsNil(o.ReferencePrices) {
		toSerialize["referencePrices"] = o.ReferencePrices
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableCompetitiveSummaryResponseBody struct {
	value *CompetitiveSummaryResponseBody
	isSet bool
}

func (v NullableCompetitiveSummaryResponseBody) Get() *CompetitiveSummaryResponseBody {
	return v.value
}

func (v *NullableCompetitiveSummaryResponseBody) Set(val *CompetitiveSummaryResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCompetitiveSummaryResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCompetitiveSummaryResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompetitiveSummaryResponseBody(val *CompetitiveSummaryResponseBody) *NullableCompetitiveSummaryResponseBody {
	return &NullableCompetitiveSummaryResponseBody{value: val, isSet: true}
}

func (v NullableCompetitiveSummaryResponseBody) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCompetitiveSummaryResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
