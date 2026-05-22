package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent struct for SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent
type SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent struct {
	// An array of campaigns with updated values. Note: targetingType cannot be updated
	Campaigns []SponsoredProductsUpdateCampaign `json:"campaigns"`
}

type _SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent(campaigns []SponsoredProductsUpdateCampaign) *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsCampaignsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) GetCampaigns() []SponsoredProductsUpdateCampaign {
	if o == nil {
		var ret []SponsoredProductsUpdateCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) GetCampaignsOk() ([]SponsoredProductsUpdateCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) SetCampaigns(v []SponsoredProductsUpdateCampaign) {
	o.Campaigns = v
}

func (o SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent(val *SponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
