package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent struct for SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent
type SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkGlobalCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent(campaigns SponsoredProductsBulkGlobalCampaignOperationResponse) *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkGlobalCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkGlobalCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkGlobalCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent(val *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
