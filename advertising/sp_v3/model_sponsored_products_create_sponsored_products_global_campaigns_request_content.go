package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent{}

// SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent struct for SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent
type SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent struct {
	// An array of campaigns.
	Campaigns []SponsoredProductsCreateGlobalCampaign `json:"campaigns"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent

// NewSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent(campaigns []SponsoredProductsCreateGlobalCampaign) *SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) GetCampaigns() []SponsoredProductsCreateGlobalCampaign {
	if o == nil {
		var ret []SponsoredProductsCreateGlobalCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) GetCampaignsOk() ([]SponsoredProductsCreateGlobalCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) SetCampaigns(v []SponsoredProductsCreateGlobalCampaign) {
	o.Campaigns = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent(val *SponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
