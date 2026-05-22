package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent{}

// SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent struct for SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent
type SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkDraftCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent

// NewSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent instantiates a new SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent(campaigns SponsoredProductsBulkDraftCampaignOperationResponse) *SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContentWithDefaults() *SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkDraftCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkDraftCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkDraftCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent struct {
	value *SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) Get() *SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) Set(val *SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent(val *SponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) *NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent {
	return &NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsPromoteSponsoredProductsDraftCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
