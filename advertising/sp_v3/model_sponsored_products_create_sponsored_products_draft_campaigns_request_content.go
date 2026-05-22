package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent{}

// SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent struct for SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent
type SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent struct {
	// An array of drafts.
	Campaigns []SponsoredProductsCreateDraftCampaign `json:"campaigns"`
}

type _SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent

// NewSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent(campaigns []SponsoredProductsCreateDraftCampaign) *SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) GetCampaigns() []SponsoredProductsCreateDraftCampaign {
	if o == nil {
		var ret []SponsoredProductsCreateDraftCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) GetCampaignsOk() ([]SponsoredProductsCreateDraftCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) SetCampaigns(v []SponsoredProductsCreateDraftCampaign) {
	o.Campaigns = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent(val *SponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
