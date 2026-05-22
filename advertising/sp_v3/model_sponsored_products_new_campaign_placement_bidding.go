package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNewCampaignPlacementBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNewCampaignPlacementBidding{}

// SponsoredProductsNewCampaignPlacementBidding The product placement.
type SponsoredProductsNewCampaignPlacementBidding struct {
	// The bidding placement percentage.
	Percentage int32 `json:"percentage"`
	// The bidding placement. One of PLACEMENT_TOP, PLACEMENT_PRODUCT_PAGE, PLACEMENT_REST_OF_SEARCH.
	Placement string `json:"placement"`
}

type _SponsoredProductsNewCampaignPlacementBidding SponsoredProductsNewCampaignPlacementBidding

// NewSponsoredProductsNewCampaignPlacementBidding instantiates a new SponsoredProductsNewCampaignPlacementBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNewCampaignPlacementBidding(percentage int32, placement string) *SponsoredProductsNewCampaignPlacementBidding {
	this := SponsoredProductsNewCampaignPlacementBidding{}
	this.Percentage = percentage
	this.Placement = placement
	return &this
}

// NewSponsoredProductsNewCampaignPlacementBiddingWithDefaults instantiates a new SponsoredProductsNewCampaignPlacementBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNewCampaignPlacementBiddingWithDefaults() *SponsoredProductsNewCampaignPlacementBidding {
	this := SponsoredProductsNewCampaignPlacementBidding{}
	return &this
}

// GetPercentage returns the Percentage field value
func (o *SponsoredProductsNewCampaignPlacementBidding) GetPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNewCampaignPlacementBidding) GetPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value
func (o *SponsoredProductsNewCampaignPlacementBidding) SetPercentage(v int32) {
	o.Percentage = v
}

// GetPlacement returns the Placement field value
func (o *SponsoredProductsNewCampaignPlacementBidding) GetPlacement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNewCampaignPlacementBidding) GetPlacementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Placement, true
}

// SetPlacement sets field value
func (o *SponsoredProductsNewCampaignPlacementBidding) SetPlacement(v string) {
	o.Placement = v
}

func (o SponsoredProductsNewCampaignPlacementBidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["percentage"] = o.Percentage
	toSerialize["placement"] = o.Placement
	return toSerialize, nil
}

type NullableSponsoredProductsNewCampaignPlacementBidding struct {
	value *SponsoredProductsNewCampaignPlacementBidding
	isSet bool
}

func (v NullableSponsoredProductsNewCampaignPlacementBidding) Get() *SponsoredProductsNewCampaignPlacementBidding {
	return v.value
}

func (v *NullableSponsoredProductsNewCampaignPlacementBidding) Set(val *SponsoredProductsNewCampaignPlacementBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNewCampaignPlacementBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNewCampaignPlacementBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNewCampaignPlacementBidding(val *SponsoredProductsNewCampaignPlacementBidding) *NullableSponsoredProductsNewCampaignPlacementBidding {
	return &NullableSponsoredProductsNewCampaignPlacementBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsNewCampaignPlacementBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNewCampaignPlacementBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
