package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent struct for SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent
type SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent struct {
	Campaigns SponsoredProductsBulkCampaignOperationResponse `json:"campaigns"`
}

type _SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent(campaigns SponsoredProductsBulkCampaignOperationResponse) *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent{}
	this.Campaigns = campaigns
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsCampaignsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) GetCampaigns() SponsoredProductsBulkCampaignOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkCampaignOperationResponse
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) GetCampaignsOk() (*SponsoredProductsBulkCampaignOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaigns, true
}

// SetCampaigns sets field value
func (o *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) SetCampaigns(v SponsoredProductsBulkCampaignOperationResponse) {
	o.Campaigns = v
}

func (o SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent(val *SponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
