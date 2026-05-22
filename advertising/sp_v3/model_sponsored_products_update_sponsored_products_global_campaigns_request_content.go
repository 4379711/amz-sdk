package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent struct for SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent
type SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent struct {
	// An array of campaigns with updated values.
	Campaigns []SponsoredProductsUpdateGlobalCampaign `json:"campaigns"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent(campaigns []SponsoredProductsUpdateGlobalCampaign) *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) GetCampaigns() []SponsoredProductsUpdateGlobalCampaign {
	if o == nil {
		var ret []SponsoredProductsUpdateGlobalCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) GetCampaignsOk() ([]SponsoredProductsUpdateGlobalCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) SetCampaigns(v []SponsoredProductsUpdateGlobalCampaign) {
	o.Campaigns = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent(val *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
