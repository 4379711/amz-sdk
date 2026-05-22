package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent struct for SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent
type SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent(campaigns SponsoredProductsBulkCampaignOperationResponse) *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsCampaignsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent(val *SponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
