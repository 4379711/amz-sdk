package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent{}

// SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent struct for SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent
type SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkDraftCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent

// NewSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent(campaigns SponsoredProductsBulkDraftCampaignOperationResponse) *SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkDraftCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkDraftCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkDraftCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent(val *SponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
