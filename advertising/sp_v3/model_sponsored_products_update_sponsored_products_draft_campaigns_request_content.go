package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent struct for SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent
type SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent struct {
	// An array of drafts with updated values.
	Campaigns []SponsoredProductsUpdateDraftCampaign `json:"campaigns"`
}

type _SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent(campaigns []SponsoredProductsUpdateDraftCampaign) *SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) GetCampaigns() []SponsoredProductsUpdateDraftCampaign {
	if o == nil {
		var ret []SponsoredProductsUpdateDraftCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) GetCampaignsOk() ([]SponsoredProductsUpdateDraftCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) SetCampaigns(v []SponsoredProductsUpdateDraftCampaign) {
	o.Campaigns = v
}

func (o SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent(val *SponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
