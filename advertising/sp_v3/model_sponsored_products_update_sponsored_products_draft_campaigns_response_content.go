package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent struct for SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent
type SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkDraftCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent(campaigns SponsoredProductsBulkDraftCampaignOperationResponse) *SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkDraftCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkDraftCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkDraftCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent(val *SponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsDraftCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
