package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignDynamicBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignDynamicBidding{}

// SponsoredProductsDraftCampaignDynamicBidding struct for SponsoredProductsDraftCampaignDynamicBidding
type SponsoredProductsDraftCampaignDynamicBidding struct {
	ShopperCohortBidding []SponsoredProductsDraftCampaignShopperCohortBidding `json:"shopperCohortBidding,omitempty"`
	PlacementBidding     []SponsoredProductsDraftCampaignPlacementBidding     `json:"placementBidding,omitempty"`
	Strategy             SponsoredProductsBiddingStrategy                     `json:"strategy"`
}

type _SponsoredProductsDraftCampaignDynamicBidding SponsoredProductsDraftCampaignDynamicBidding

// NewSponsoredProductsDraftCampaignDynamicBidding instantiates a new SponsoredProductsDraftCampaignDynamicBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignDynamicBidding(strategy SponsoredProductsBiddingStrategy) *SponsoredProductsDraftCampaignDynamicBidding {
	this := SponsoredProductsDraftCampaignDynamicBidding{}
	this.Strategy = strategy
	return &this
}

// NewSponsoredProductsDraftCampaignDynamicBiddingWithDefaults instantiates a new SponsoredProductsDraftCampaignDynamicBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignDynamicBiddingWithDefaults() *SponsoredProductsDraftCampaignDynamicBidding {
	this := SponsoredProductsDraftCampaignDynamicBidding{}
	return &this
}

// GetShopperCohortBidding returns the ShopperCohortBidding field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignDynamicBidding) GetShopperCohortBidding() []SponsoredProductsDraftCampaignShopperCohortBidding {
	if o == nil || IsNil(o.ShopperCohortBidding) {
		var ret []SponsoredProductsDraftCampaignShopperCohortBidding
		return ret
	}
	return o.ShopperCohortBidding
}

// GetShopperCohortBiddingOk returns a tuple with the ShopperCohortBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignDynamicBidding) GetShopperCohortBiddingOk() ([]SponsoredProductsDraftCampaignShopperCohortBidding, bool) {
	if o == nil || IsNil(o.ShopperCohortBidding) {
		return nil, false
	}
	return o.ShopperCohortBidding, true
}

// HasShopperCohortBidding returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignDynamicBidding) HasShopperCohortBidding() bool {
	if o != nil && !IsNil(o.ShopperCohortBidding) {
		return true
	}

	return false
}

// SetShopperCohortBidding gets a reference to the given []SponsoredProductsDraftCampaignShopperCohortBidding and assigns it to the ShopperCohortBidding field.
func (o *SponsoredProductsDraftCampaignDynamicBidding) SetShopperCohortBidding(v []SponsoredProductsDraftCampaignShopperCohortBidding) {
	o.ShopperCohortBidding = v
}

// GetPlacementBidding returns the PlacementBidding field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignDynamicBidding) GetPlacementBidding() []SponsoredProductsDraftCampaignPlacementBidding {
	if o == nil || IsNil(o.PlacementBidding) {
		var ret []SponsoredProductsDraftCampaignPlacementBidding
		return ret
	}
	return o.PlacementBidding
}

// GetPlacementBiddingOk returns a tuple with the PlacementBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignDynamicBidding) GetPlacementBiddingOk() ([]SponsoredProductsDraftCampaignPlacementBidding, bool) {
	if o == nil || IsNil(o.PlacementBidding) {
		return nil, false
	}
	return o.PlacementBidding, true
}

// HasPlacementBidding returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignDynamicBidding) HasPlacementBidding() bool {
	if o != nil && !IsNil(o.PlacementBidding) {
		return true
	}

	return false
}

// SetPlacementBidding gets a reference to the given []SponsoredProductsDraftCampaignPlacementBidding and assigns it to the PlacementBidding field.
func (o *SponsoredProductsDraftCampaignDynamicBidding) SetPlacementBidding(v []SponsoredProductsDraftCampaignPlacementBidding) {
	o.PlacementBidding = v
}

// GetStrategy returns the Strategy field value
func (o *SponsoredProductsDraftCampaignDynamicBidding) GetStrategy() SponsoredProductsBiddingStrategy {
	if o == nil {
		var ret SponsoredProductsBiddingStrategy
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignDynamicBidding) GetStrategyOk() (*SponsoredProductsBiddingStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *SponsoredProductsDraftCampaignDynamicBidding) SetStrategy(v SponsoredProductsBiddingStrategy) {
	o.Strategy = v
}

func (o SponsoredProductsDraftCampaignDynamicBidding) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftCampaignDynamicBidding struct {
	value *SponsoredProductsDraftCampaignDynamicBidding
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignDynamicBidding) Get() *SponsoredProductsDraftCampaignDynamicBidding {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignDynamicBidding) Set(val *SponsoredProductsDraftCampaignDynamicBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignDynamicBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignDynamicBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignDynamicBidding(val *SponsoredProductsDraftCampaignDynamicBidding) *NullableSponsoredProductsDraftCampaignDynamicBidding {
	return &NullableSponsoredProductsDraftCampaignDynamicBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignDynamicBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignDynamicBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
