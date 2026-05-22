package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignShopperCohortBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignShopperCohortBidding{}

// SponsoredProductsDraftCampaignShopperCohortBidding struct for SponsoredProductsDraftCampaignShopperCohortBidding
type SponsoredProductsDraftCampaignShopperCohortBidding struct {
	ShopperCohortType SponsoredProductsShopperCohortType `json:"shopperCohortType"`
	Percentage        *int32                             `json:"percentage,omitempty"`
	// Required when \"AUDIENCE_SEGMENT\" is used for shopperCohortType.
	AudienceSegments []SponsoredProductsAudienceSegment `json:"audienceSegments,omitempty"`
}

type _SponsoredProductsDraftCampaignShopperCohortBidding SponsoredProductsDraftCampaignShopperCohortBidding

// NewSponsoredProductsDraftCampaignShopperCohortBidding instantiates a new SponsoredProductsDraftCampaignShopperCohortBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignShopperCohortBidding(shopperCohortType SponsoredProductsShopperCohortType) *SponsoredProductsDraftCampaignShopperCohortBidding {
	this := SponsoredProductsDraftCampaignShopperCohortBidding{}
	this.ShopperCohortType = shopperCohortType
	return &this
}

// NewSponsoredProductsDraftCampaignShopperCohortBiddingWithDefaults instantiates a new SponsoredProductsDraftCampaignShopperCohortBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignShopperCohortBiddingWithDefaults() *SponsoredProductsDraftCampaignShopperCohortBidding {
	this := SponsoredProductsDraftCampaignShopperCohortBidding{}
	return &this
}

// GetShopperCohortType returns the ShopperCohortType field value
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) GetShopperCohortType() SponsoredProductsShopperCohortType {
	if o == nil {
		var ret SponsoredProductsShopperCohortType
		return ret
	}

	return o.ShopperCohortType
}

// GetShopperCohortTypeOk returns a tuple with the ShopperCohortType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) GetShopperCohortTypeOk() (*SponsoredProductsShopperCohortType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperCohortType, true
}

// SetShopperCohortType sets field value
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) SetShopperCohortType(v SponsoredProductsShopperCohortType) {
	o.ShopperCohortType = v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) SetPercentage(v int32) {
	o.Percentage = &v
}

// GetAudienceSegments returns the AudienceSegments field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) GetAudienceSegments() []SponsoredProductsAudienceSegment {
	if o == nil || IsNil(o.AudienceSegments) {
		var ret []SponsoredProductsAudienceSegment
		return ret
	}
	return o.AudienceSegments
}

// GetAudienceSegmentsOk returns a tuple with the AudienceSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) GetAudienceSegmentsOk() ([]SponsoredProductsAudienceSegment, bool) {
	if o == nil || IsNil(o.AudienceSegments) {
		return nil, false
	}
	return o.AudienceSegments, true
}

// HasAudienceSegments returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) HasAudienceSegments() bool {
	if o != nil && !IsNil(o.AudienceSegments) {
		return true
	}

	return false
}

// SetAudienceSegments gets a reference to the given []SponsoredProductsAudienceSegment and assigns it to the AudienceSegments field.
func (o *SponsoredProductsDraftCampaignShopperCohortBidding) SetAudienceSegments(v []SponsoredProductsAudienceSegment) {
	o.AudienceSegments = v
}

func (o SponsoredProductsDraftCampaignShopperCohortBidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shopperCohortType"] = o.ShopperCohortType
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.AudienceSegments) {
		toSerialize["audienceSegments"] = o.AudienceSegments
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignShopperCohortBidding struct {
	value *SponsoredProductsDraftCampaignShopperCohortBidding
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignShopperCohortBidding) Get() *SponsoredProductsDraftCampaignShopperCohortBidding {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignShopperCohortBidding) Set(val *SponsoredProductsDraftCampaignShopperCohortBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignShopperCohortBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignShopperCohortBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignShopperCohortBidding(val *SponsoredProductsDraftCampaignShopperCohortBidding) *NullableSponsoredProductsDraftCampaignShopperCohortBidding {
	return &NullableSponsoredProductsDraftCampaignShopperCohortBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignShopperCohortBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignShopperCohortBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
