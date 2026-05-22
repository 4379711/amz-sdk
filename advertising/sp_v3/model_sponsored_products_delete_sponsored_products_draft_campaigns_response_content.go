package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent struct for SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent
type SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkDraftCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent(campaigns SponsoredProductsBulkDraftCampaignOperationResponse) *SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkDraftCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkDraftCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkDraftCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent(val *SponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
