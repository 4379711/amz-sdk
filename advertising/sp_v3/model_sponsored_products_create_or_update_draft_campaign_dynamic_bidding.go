package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding{}

// SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding struct for SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding
type SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding struct {
	ShopperCohortBidding []SponsoredProductsDraftCampaignShopperCohortBidding `json:"shopperCohortBidding,omitempty"`
	PlacementBidding     []SponsoredProductsDraftCampaignPlacementBidding     `json:"placementBidding,omitempty"`
	Strategy             SponsoredProductsCreateOrUpdateBiddingStrategy       `json:"strategy"`
}

type _SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding

// NewSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding instantiates a new SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding(strategy SponsoredProductsCreateOrUpdateBiddingStrategy) *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding {
	this := SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding{}
	this.Strategy = strategy
	return &this
}

// NewSponsoredProductsCreateOrUpdateDraftCampaignDynamicBiddingWithDefaults instantiates a new SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateOrUpdateDraftCampaignDynamicBiddingWithDefaults() *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding {
	this := SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding{}
	return &this
}

// GetShopperCohortBidding returns the ShopperCohortBidding field value if set, zero value otherwise.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) GetShopperCohortBidding() []SponsoredProductsDraftCampaignShopperCohortBidding {
	if o == nil || IsNil(o.ShopperCohortBidding) {
		var ret []SponsoredProductsDraftCampaignShopperCohortBidding
		return ret
	}
	return o.ShopperCohortBidding
}

// GetShopperCohortBiddingOk returns a tuple with the ShopperCohortBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) GetShopperCohortBiddingOk() ([]SponsoredProductsDraftCampaignShopperCohortBidding, bool) {
	if o == nil || IsNil(o.ShopperCohortBidding) {
		return nil, false
	}
	return o.ShopperCohortBidding, true
}

// HasShopperCohortBidding returns a boolean if a field has been set.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) HasShopperCohortBidding() bool {
	if o != nil && !IsNil(o.ShopperCohortBidding) {
		return true
	}

	return false
}

// SetShopperCohortBidding gets a reference to the given []SponsoredProductsDraftCampaignShopperCohortBidding and assigns it to the ShopperCohortBidding field.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) SetShopperCohortBidding(v []SponsoredProductsDraftCampaignShopperCohortBidding) {
	o.ShopperCohortBidding = v
}

// GetPlacementBidding returns the PlacementBidding field value if set, zero value otherwise.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) GetPlacementBidding() []SponsoredProductsDraftCampaignPlacementBidding {
	if o == nil || IsNil(o.PlacementBidding) {
		var ret []SponsoredProductsDraftCampaignPlacementBidding
		return ret
	}
	return o.PlacementBidding
}

// GetPlacementBiddingOk returns a tuple with the PlacementBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) GetPlacementBiddingOk() ([]SponsoredProductsDraftCampaignPlacementBidding, bool) {
	if o == nil || IsNil(o.PlacementBidding) {
		return nil, false
	}
	return o.PlacementBidding, true
}

// HasPlacementBidding returns a boolean if a field has been set.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) HasPlacementBidding() bool {
	if o != nil && !IsNil(o.PlacementBidding) {
		return true
	}

	return false
}

// SetPlacementBidding gets a reference to the given []SponsoredProductsDraftCampaignPlacementBidding and assigns it to the PlacementBidding field.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) SetPlacementBidding(v []SponsoredProductsDraftCampaignPlacementBidding) {
	o.PlacementBidding = v
}

// GetStrategy returns the Strategy field value
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) GetStrategy() SponsoredProductsCreateOrUpdateBiddingStrategy {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateBiddingStrategy
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) GetStrategyOk() (*SponsoredProductsCreateOrUpdateBiddingStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) SetStrategy(v SponsoredProductsCreateOrUpdateBiddingStrategy) {
	o.Strategy = v
}

func (o SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding struct {
	value *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) Get() *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) Set(val *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding(val *SponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) *NullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding {
	return &NullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateDraftCampaignDynamicBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
