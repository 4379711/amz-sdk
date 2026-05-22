package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent{}

// SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent struct for SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent
type SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkGlobalCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent

// NewSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent(campaigns SponsoredProductsBulkGlobalCampaignOperationResponse) *SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkGlobalCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkGlobalCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkGlobalCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent(val *SponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
