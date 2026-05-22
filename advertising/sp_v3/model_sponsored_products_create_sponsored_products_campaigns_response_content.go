package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsCampaignsResponseContent{}

// SponsoredProductsCreateSponsoredProductsCampaignsResponseContent struct for SponsoredProductsCreateSponsoredProductsCampaignsResponseContent
type SponsoredProductsCreateSponsoredProductsCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsCreateSponsoredProductsCampaignsResponseContent SponsoredProductsCreateSponsoredProductsCampaignsResponseContent

// NewSponsoredProductsCreateSponsoredProductsCampaignsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsCampaignsResponseContent(campaigns SponsoredProductsBulkCampaignOperationResponse) *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsCampaignsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsCreateSponsoredProductsCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent(val *SponsoredProductsCreateSponsoredProductsCampaignsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
