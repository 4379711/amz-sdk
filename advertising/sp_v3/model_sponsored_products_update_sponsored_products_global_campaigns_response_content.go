package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent struct for SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent
type SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkGlobalCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent(campaigns SponsoredProductsBulkGlobalCampaignOperationResponse) *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkGlobalCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkGlobalCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkGlobalCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent(val *SponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
