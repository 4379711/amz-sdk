package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsCampaignsRequestContent{}

// SponsoredProductsCreateSponsoredProductsCampaignsRequestContent struct for SponsoredProductsCreateSponsoredProductsCampaignsRequestContent
type SponsoredProductsCreateSponsoredProductsCampaignsRequestContent struct {
	// An array of campaigns.
	Campaigns []SponsoredProductsCreateCampaign `json:"campaigns"`
}

type _SponsoredProductsCreateSponsoredProductsCampaignsRequestContent SponsoredProductsCreateSponsoredProductsCampaignsRequestContent

// NewSponsoredProductsCreateSponsoredProductsCampaignsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsCampaignsRequestContent(campaigns []SponsoredProductsCreateCampaign) *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignsRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsCampaignsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignsRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent) GetCampaigns() []SponsoredProductsCreateCampaign {
	if o == nil {
		var ret []SponsoredProductsCreateCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent) GetCampaignsOk() ([]SponsoredProductsCreateCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent) SetCampaigns(v []SponsoredProductsCreateCampaign) {
	o.Campaigns = v
}

func (o SponsoredProductsCreateSponsoredProductsCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent(val *SponsoredProductsCreateSponsoredProductsCampaignsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
