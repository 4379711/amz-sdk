package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNewCampaignDynamicBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNewCampaignDynamicBidding{}

// SponsoredProductsNewCampaignDynamicBidding Specifies bidding controls.
type SponsoredProductsNewCampaignDynamicBidding struct {
	// The product placement.
	PlacementBidding []SponsoredProductsNewCampaignPlacementBidding `json:"placementBidding,omitempty"`
	// One of LEGACY_FOR_SALES, AUTO_FOR_SALES, MANUAL, RULE_BASED.
	Strategy string `json:"strategy"`
}

type _SponsoredProductsNewCampaignDynamicBidding SponsoredProductsNewCampaignDynamicBidding

// NewSponsoredProductsNewCampaignDynamicBidding instantiates a new SponsoredProductsNewCampaignDynamicBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNewCampaignDynamicBidding(strategy string) *SponsoredProductsNewCampaignDynamicBidding {
	this := SponsoredProductsNewCampaignDynamicBidding{}
	this.Strategy = strategy
	return &this
}

// NewSponsoredProductsNewCampaignDynamicBiddingWithDefaults instantiates a new SponsoredProductsNewCampaignDynamicBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNewCampaignDynamicBiddingWithDefaults() *SponsoredProductsNewCampaignDynamicBidding {
	this := SponsoredProductsNewCampaignDynamicBidding{}
	return &this
}

// GetPlacementBidding returns the PlacementBidding field value if set, zero value otherwise.
func (o *SponsoredProductsNewCampaignDynamicBidding) GetPlacementBidding() []SponsoredProductsNewCampaignPlacementBidding {
	if o == nil || IsNil(o.PlacementBidding) {
		var ret []SponsoredProductsNewCampaignPlacementBidding
		return ret
	}
	return o.PlacementBidding
}

// GetPlacementBiddingOk returns a tuple with the PlacementBidding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNewCampaignDynamicBidding) GetPlacementBiddingOk() ([]SponsoredProductsNewCampaignPlacementBidding, bool) {
	if o == nil || IsNil(o.PlacementBidding) {
		return nil, false
	}
	return o.PlacementBidding, true
}

// HasPlacementBidding returns a boolean if a field has been set.
func (o *SponsoredProductsNewCampaignDynamicBidding) HasPlacementBidding() bool {
	if o != nil && !IsNil(o.PlacementBidding) {
		return true
	}

	return false
}

// SetPlacementBidding gets a reference to the given []SponsoredProductsNewCampaignPlacementBidding and assigns it to the PlacementBidding field.
func (o *SponsoredProductsNewCampaignDynamicBidding) SetPlacementBidding(v []SponsoredProductsNewCampaignPlacementBidding) {
	o.PlacementBidding = v
}

// GetStrategy returns the Strategy field value
func (o *SponsoredProductsNewCampaignDynamicBidding) GetStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNewCampaignDynamicBidding) GetStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *SponsoredProductsNewCampaignDynamicBidding) SetStrategy(v string) {
	o.Strategy = v
}

func (o SponsoredProductsNewCampaignDynamicBidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlacementBidding) {
		toSerialize["placementBidding"] = o.PlacementBidding
	}
	toSerialize["strategy"] = o.Strategy
	return toSerialize, nil
}

type NullableSponsoredProductsNewCampaignDynamicBidding struct {
	value *SponsoredProductsNewCampaignDynamicBidding
	isSet bool
}

func (v NullableSponsoredProductsNewCampaignDynamicBidding) Get() *SponsoredProductsNewCampaignDynamicBidding {
	return v.value
}

func (v *NullableSponsoredProductsNewCampaignDynamicBidding) Set(val *SponsoredProductsNewCampaignDynamicBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNewCampaignDynamicBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNewCampaignDynamicBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNewCampaignDynamicBidding(val *SponsoredProductsNewCampaignDynamicBidding) *NullableSponsoredProductsNewCampaignDynamicBidding {
	return &NullableSponsoredProductsNewCampaignDynamicBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsNewCampaignDynamicBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNewCampaignDynamicBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
