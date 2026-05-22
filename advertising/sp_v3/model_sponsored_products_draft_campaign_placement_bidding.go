package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignPlacementBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignPlacementBidding{}

// SponsoredProductsDraftCampaignPlacementBidding struct for SponsoredProductsDraftCampaignPlacementBidding
type SponsoredProductsDraftCampaignPlacementBidding struct {
	Percentage *int32                      `json:"percentage,omitempty"`
	Placement  *SponsoredProductsPlacement `json:"placement,omitempty"`
}

// NewSponsoredProductsDraftCampaignPlacementBidding instantiates a new SponsoredProductsDraftCampaignPlacementBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignPlacementBidding() *SponsoredProductsDraftCampaignPlacementBidding {
	this := SponsoredProductsDraftCampaignPlacementBidding{}
	return &this
}

// NewSponsoredProductsDraftCampaignPlacementBiddingWithDefaults instantiates a new SponsoredProductsDraftCampaignPlacementBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignPlacementBiddingWithDefaults() *SponsoredProductsDraftCampaignPlacementBidding {
	this := SponsoredProductsDraftCampaignPlacementBidding{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignPlacementBidding) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignPlacementBidding) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignPlacementBidding) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *SponsoredProductsDraftCampaignPlacementBidding) SetPercentage(v int32) {
	o.Percentage = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignPlacementBidding) GetPlacement() SponsoredProductsPlacement {
	if o == nil || IsNil(o.Placement) {
		var ret SponsoredProductsPlacement
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignPlacementBidding) GetPlacementOk() (*SponsoredProductsPlacement, bool) {
	if o == nil || IsNil(o.Placement) {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignPlacementBidding) HasPlacement() bool {
	if o != nil && !IsNil(o.Placement) {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given SponsoredProductsPlacement and assigns it to the Placement field.
func (o *SponsoredProductsDraftCampaignPlacementBidding) SetPlacement(v SponsoredProductsPlacement) {
	o.Placement = &v
}

func (o SponsoredProductsDraftCampaignPlacementBidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.Placement) {
		toSerialize["placement"] = o.Placement
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignPlacementBidding struct {
	value *SponsoredProductsDraftCampaignPlacementBidding
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignPlacementBidding) Get() *SponsoredProductsDraftCampaignPlacementBidding {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignPlacementBidding) Set(val *SponsoredProductsDraftCampaignPlacementBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignPlacementBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignPlacementBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignPlacementBidding(val *SponsoredProductsDraftCampaignPlacementBidding) *NullableSponsoredProductsDraftCampaignPlacementBidding {
	return &NullableSponsoredProductsDraftCampaignPlacementBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignPlacementBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignPlacementBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
