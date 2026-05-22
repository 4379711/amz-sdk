package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDynamicBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDynamicBidding{}

// SponsoredProductsDynamicBidding struct for SponsoredProductsDynamicBidding
type SponsoredProductsDynamicBidding struct {
	// Specifies Shopper Cohorts based bid adjustment controls. `shopperCohortBidding` is optional for both Create and Update requests. You can enable this control to adjust your bid based on the shopper cohorts. The percentage value set is the percentage of the original bid including any other bid adjustments such as `placementBidding`. For example, a `placementBidding` with 50% adjustment on a $1.00 bid would increase the bid to $1.50, and a `shopperCohortBidding` with 100% adjustment would further increase the bid to $3.00.
	ShopperCohortBidding []SponsoredProductsShopperCohortBidding `json:"shopperCohortBidding,omitempty"`
	PlacementBidding     []SponsoredProductsPlacementBidding     `json:"placementBidding,omitempty"`
	Strategy             SponsoredProductsBiddingStrategy        `json:"strategy"`
}

type _SponsoredProductsDynamicBidding SponsoredProductsDynamicBidding

// NewSponsoredProductsDynamicBidding instantiates a new SponsoredProductsDynamicBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDynamicBidding(strategy SponsoredProductsBiddingStrategy) *SponsoredProductsDynamicBidding {
	this := SponsoredProductsDynamicBidding{}
	this.Strategy = strategy
	return &this
}

// NewSponsoredProductsDynamicBiddingWithDefaults instantiates a new SponsoredProductsDynamicBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDynamicBiddingWithDefaults() *SponsoredProductsDynamicBidding {
	this := SponsoredProductsDynamicBidding{}
	return &this
}

// GetShopperCohortBidding returns the ShopperCohortBidding field value if set, zero value otherwise.
func (o *SponsoredProductsDynamicBidding) GetShopperCohortBidding() []SponsoredProductsShopperCohortBidding {
	if o == nil || IsNil(o.ShopperCohortBidding) {
		var ret []SponsoredProductsShopperCohortBidding
		return ret
	}
	return o.ShopperCohortBidding
}

// GetShopperCohortBiddingOk returns a tuple with the ShopperCohortBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDynamicBidding) GetShopperCohortBiddingOk() ([]SponsoredProductsShopperCohortBidding, bool) {
	if o == nil || IsNil(o.ShopperCohortBidding) {
		return nil, false
	}
	return o.ShopperCohortBidding, true
}

// HasShopperCohortBidding returns a boolean if a field has been set.
func (o *SponsoredProductsDynamicBidding) HasShopperCohortBidding() bool {
	if o != nil && !IsNil(o.ShopperCohortBidding) {
		return true
	}

	return false
}

// SetShopperCohortBidding gets a reference to the given []SponsoredProductsShopperCohortBidding and assigns it to the ShopperCohortBidding field.
func (o *SponsoredProductsDynamicBidding) SetShopperCohortBidding(v []SponsoredProductsShopperCohortBidding) {
	o.ShopperCohortBidding = v
}

// GetPlacementBidding returns the PlacementBidding field value if set, zero value otherwise.
func (o *SponsoredProductsDynamicBidding) GetPlacementBidding() []SponsoredProductsPlacementBidding {
	if o == nil || IsNil(o.PlacementBidding) {
		var ret []SponsoredProductsPlacementBidding
		return ret
	}
	return o.PlacementBidding
}

// GetPlacementBiddingOk returns a tuple with the PlacementBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDynamicBidding) GetPlacementBiddingOk() ([]SponsoredProductsPlacementBidding, bool) {
	if o == nil || IsNil(o.PlacementBidding) {
		return nil, false
	}
	return o.PlacementBidding, true
}

// HasPlacementBidding returns a boolean if a field has been set.
func (o *SponsoredProductsDynamicBidding) HasPlacementBidding() bool {
	if o != nil && !IsNil(o.PlacementBidding) {
		return true
	}

	return false
}

// SetPlacementBidding gets a reference to the given []SponsoredProductsPlacementBidding and assigns it to the PlacementBidding field.
func (o *SponsoredProductsDynamicBidding) SetPlacementBidding(v []SponsoredProductsPlacementBidding) {
	o.PlacementBidding = v
}

// GetStrategy returns the Strategy field value
func (o *SponsoredProductsDynamicBidding) GetStrategy() SponsoredProductsBiddingStrategy {
	if o == nil {
		var ret SponsoredProductsBiddingStrategy
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDynamicBidding) GetStrategyOk() (*SponsoredProductsBiddingStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *SponsoredProductsDynamicBidding) SetStrategy(v SponsoredProductsBiddingStrategy) {
	o.Strategy = v
}

func (o SponsoredProductsDynamicBidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShopperCohortBidding) {
		toSerialize["shopperCohortBidding"] = o.ShopperCohortBidding
	}
	if !IsNil(o.PlacementBidding) {
		toSerialize["placementBidding"] = o.PlacementBidding
	}
	toSerialize["strategy"] = o.Strategy
	return toSerialize, nil
}

type NullableSponsoredProductsDynamicBidding struct {
	value *SponsoredProductsDynamicBidding
	isSet bool
}

func (v NullableSponsoredProductsDynamicBidding) Get() *SponsoredProductsDynamicBidding {
	return v.value
}

func (v *NullableSponsoredProductsDynamicBidding) Set(val *SponsoredProductsDynamicBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDynamicBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDynamicBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDynamicBidding(val *SponsoredProductsDynamicBidding) *NullableSponsoredProductsDynamicBidding {
	return &NullableSponsoredProductsDynamicBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsDynamicBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDynamicBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
