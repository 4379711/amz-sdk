package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the CompetitivePricingType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompetitivePricingType{}

// CompetitivePricingType Competitive pricing information for the item.
type CompetitivePricingType struct {
	// A list of competitive pricing information.
	CompetitivePrices []CompetitivePriceType `json:"CompetitivePrices"`
	// The number of active offer listings for the item that was submitted. The listing count is returned by condition, one for each listing condition value that is returned.
	NumberOfOfferListings []OfferListingCountType `json:"NumberOfOfferListings"`
	TradeInValue          *MoneyType              `json:"TradeInValue,omitempty"`
}

type _CompetitivePricingType CompetitivePricingType

// NewCompetitivePricingType instantiates a new CompetitivePricingType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompetitivePricingType(competitivePrices []CompetitivePriceType, numberOfOfferListings []OfferListingCountType) *CompetitivePricingType {
	this := CompetitivePricingType{}
	this.CompetitivePrices = competitivePrices
	this.NumberOfOfferListings = numberOfOfferListings
	return &this
}

// NewCompetitivePricingTypeWithDefaults instantiates a new CompetitivePricingType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompetitivePricingTypeWithDefaults() *CompetitivePricingType {
	this := CompetitivePricingType{}
	return &this
}

// GetCompetitivePrices returns the CompetitivePrices field value
func (o *CompetitivePricingType) GetCompetitivePrices() []CompetitivePriceType {
	if o == nil {
		var ret []CompetitivePriceType
		return ret
	}

	return o.CompetitivePrices
}

// GetCompetitivePricesOk returns a tuple with the CompetitivePrices field value
// and a boolean to check if the value has been set.
func (o *CompetitivePricingType) GetCompetitivePricesOk() ([]CompetitivePriceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompetitivePrices, true
}

// SetCompetitivePrices sets field value
func (o *CompetitivePricingType) SetCompetitivePrices(v []CompetitivePriceType) {
	o.CompetitivePrices = v
}

// GetNumberOfOfferListings returns the NumberOfOfferListings field value
func (o *CompetitivePricingType) GetNumberOfOfferListings() []OfferListingCountType {
	if o == nil {
		var ret []OfferListingCountType
		return ret
	}

	return o.NumberOfOfferListings
}

// GetNumberOfOfferListingsOk returns a tuple with the NumberOfOfferListings field value
// and a boolean to check if the value has been set.
func (o *CompetitivePricingType) GetNumberOfOfferListingsOk() ([]OfferListingCountType, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberOfOfferListings, true
}

// SetNumberOfOfferListings sets field value
func (o *CompetitivePricingType) SetNumberOfOfferListings(v []OfferListingCountType) {
	o.NumberOfOfferListings = v
}

// GetTradeInValue returns the TradeInValue field value if set, zero value otherwise.
func (o *CompetitivePricingType) GetTradeInValue() MoneyType {
	if o == nil || IsNil(o.TradeInValue) {
		var ret MoneyType
		return ret
	}
	return *o.TradeInValue
}

// GetTradeInValueOk returns a tuple with the TradeInValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompetitivePricingType) GetTradeInValueOk() (*MoneyType, bool) {
	if o == nil || IsNil(o.TradeInValue) {
		return nil, false
	}
	return o.TradeInValue, true
}

// HasTradeInValue returns a boolean if a field has been set.
func (o *CompetitivePricingType) HasTradeInValue() bool {
	if o != nil && !IsNil(o.TradeInValue) {
		return true
	}

	return false
}

// SetTradeInValue gets a reference to the given MoneyType and assigns it to the TradeInValue field.
func (o *CompetitivePricingType) SetTradeInValue(v MoneyType) {
	o.TradeInValue = &v
}

func (o CompetitivePricingType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CompetitivePrices"] = o.CompetitivePrices
	toSerialize["NumberOfOfferListings"] = o.NumberOfOfferListings
	if !IsNil(o.TradeInValue) {
		toSerialize["TradeInValue"] = o.TradeInValue
	}
	return toSerialize, nil
}

type NullableCompetitivePricingType struct {
	value *CompetitivePricingType
	isSet bool
}

func (v NullableCompetitivePricingType) Get() *CompetitivePricingType {
	return v.value
}

func (v *NullableCompetitivePricingType) Set(val *CompetitivePricingType) {
	v.value = val
	v.isSet = true
}

func (v NullableCompetitivePricingType) IsSet() bool {
	return v.isSet
}

func (v *NullableCompetitivePricingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompetitivePricingType(val *CompetitivePricingType) *NullableCompetitivePricingType {
	return &NullableCompetitivePricingType{value: val, isSet: true}
}

func (v NullableCompetitivePricingType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCompetitivePricingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
