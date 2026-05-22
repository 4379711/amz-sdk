package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateOrUpdateDynamicBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateOrUpdateDynamicBidding{}

// SponsoredProductsCreateOrUpdateDynamicBidding Specifies bidding controls. DynamicBidding is optional for both Create and Update requests. For Create Campaign requests, if you don't specify dynamicBidding, default strategy of `LEGACY_FOR_SALES` will be applied.
type SponsoredProductsCreateOrUpdateDynamicBidding struct {
	// Specifies Shopper Cohorts based bid adjustment controls. `shopperCohortBidding` is optional for both Create and Update requests. You can enable this control to adjust your bid based on the shopper cohorts. The percentage value set is the percentage of the original bid including any other bid adjustments such as `placementBidding`. For example, a `placementBidding` with 50% adjustment on a $1.00 bid would increase the bid to $1.50, and a `shopperCohortBidding` with 100% adjustment would further increase the bid to $3.00.
	ShopperCohortBidding []SponsoredProductsShopperCohortBidding         `json:"shopperCohortBidding,omitempty"`
	PlacementBidding     []SponsoredProductsPlacementBidding             `json:"placementBidding,omitempty"`
	Strategy             *SponsoredProductsCreateOrUpdateBiddingStrategy `json:"strategy,omitempty"`
}

// NewSponsoredProductsCreateOrUpdateDynamicBidding instantiates a new SponsoredProductsCreateOrUpdateDynamicBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateOrUpdateDynamicBidding() *SponsoredProductsCreateOrUpdateDynamicBidding {
	this := SponsoredProductsCreateOrUpdateDynamicBidding{}
	return &this
}

// NewSponsoredProductsCreateOrUpdateDynamicBiddingWithDefaults instantiates a new SponsoredProductsCreateOrUpdateDynamicBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateOrUpdateDynamicBiddingWithDefaults() *SponsoredProductsCreateOrUpdateDynamicBidding {
	this := SponsoredProductsCreateOrUpdateDynamicBidding{}
	return &this
}

// GetShopperCohortBidding returns the ShopperCohortBidding field value if set, zero value otherwise.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) GetShopperCohortBidding() []SponsoredProductsShopperCohortBidding {
	if o == nil || IsNil(o.ShopperCohortBidding) {
		var ret []SponsoredProductsShopperCohortBidding
		return ret
	}
	return o.ShopperCohortBidding
}

// GetShopperCohortBiddingOk returns a tuple with the ShopperCohortBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) GetShopperCohortBiddingOk() ([]SponsoredProductsShopperCohortBidding, bool) {
	if o == nil || IsNil(o.ShopperCohortBidding) {
		return nil, false
	}
	return o.ShopperCohortBidding, true
}

// HasShopperCohortBidding returns a boolean if a field has been set.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) HasShopperCohortBidding() bool {
	if o != nil && !IsNil(o.ShopperCohortBidding) {
		return true
	}

	return false
}

// SetShopperCohortBidding gets a reference to the given []SponsoredProductsShopperCohortBidding and assigns it to the ShopperCohortBidding field.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) SetShopperCohortBidding(v []SponsoredProductsShopperCohortBidding) {
	o.ShopperCohortBidding = v
}

// GetPlacementBidding returns the PlacementBidding field value if set, zero value otherwise.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) GetPlacementBidding() []SponsoredProductsPlacementBidding {
	if o == nil || IsNil(o.PlacementBidding) {
		var ret []SponsoredProductsPlacementBidding
		return ret
	}
	return o.PlacementBidding
}

// GetPlacementBiddingOk returns a tuple with the PlacementBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) GetPlacementBiddingOk() ([]SponsoredProductsPlacementBidding, bool) {
	if o == nil || IsNil(o.PlacementBidding) {
		return nil, false
	}
	return o.PlacementBidding, true
}

// HasPlacementBidding returns a boolean if a field has been set.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) HasPlacementBidding() bool {
	if o != nil && !IsNil(o.PlacementBidding) {
		return true
	}

	return false
}

// SetPlacementBidding gets a reference to the given []SponsoredProductsPlacementBidding and assigns it to the PlacementBidding field.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) SetPlacementBidding(v []SponsoredProductsPlacementBidding) {
	o.PlacementBidding = v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) GetStrategy() SponsoredProductsCreateOrUpdateBiddingStrategy {
	if o == nil || IsNil(o.Strategy) {
		var ret SponsoredProductsCreateOrUpdateBiddingStrategy
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) GetStrategyOk() (*SponsoredProductsCreateOrUpdateBiddingStrategy, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given SponsoredProductsCreateOrUpdateBiddingStrategy and assigns it to the Strategy field.
func (o *SponsoredProductsCreateOrUpdateDynamicBidding) SetStrategy(v SponsoredProductsCreateOrUpdateBiddingStrategy) {
	o.Strategy = &v
}

func (o SponsoredProductsCreateOrUpdateDynamicBidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShopperCohortBidding) {
		toSerialize["shopperCohortBidding"] = o.ShopperCohortBidding
	}
	if !IsNil(o.PlacementBidding) {
		toSerialize["placementBidding"] = o.PlacementBidding
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateOrUpdateDynamicBidding struct {
	value *SponsoredProductsCreateOrUpdateDynamicBidding
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateDynamicBidding) Get() *SponsoredProductsCreateOrUpdateDynamicBidding {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateDynamicBidding) Set(val *SponsoredProductsCreateOrUpdateDynamicBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateDynamicBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateDynamicBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateDynamicBidding(val *SponsoredProductsCreateOrUpdateDynamicBidding) *NullableSponsoredProductsCreateOrUpdateDynamicBidding {
	return &NullableSponsoredProductsCreateOrUpdateDynamicBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateDynamicBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateDynamicBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
