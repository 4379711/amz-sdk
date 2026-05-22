package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsCampaignsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsCampaignsBetaResponseContent{}

// UpdateSponsoredBrandsCampaignsBetaResponseContent struct for UpdateSponsoredBrandsCampaignsBetaResponseContent
type UpdateSponsoredBrandsCampaignsBetaResponseContent struct {
	Campaigns *BulkCampaignOperationResponse `json:"campaigns,omitempty"`
}

// NewUpdateSponsoredBrandsCampaignsBetaResponseContent instantiates a new UpdateSponsoredBrandsCampaignsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsCampaignsBetaResponseContent() *UpdateSponsoredBrandsCampaignsBetaResponseContent {
	this := UpdateSponsoredBrandsCampaignsBetaResponseContent{}
	return &this
}

// NewUpdateSponsoredBrandsCampaignsBetaResponseContentWithDefaults instantiates a new UpdateSponsoredBrandsCampaignsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsCampaignsBetaResponseContentWithDefaults() *UpdateSponsoredBrandsCampaignsBetaResponseContent {
	this := UpdateSponsoredBrandsCampaignsBetaResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *UpdateSponsoredBrandsCampaignsBetaResponseContent) GetCampaigns() BulkCampaignOperationResponse {
	if o == nil || IsNil(o.Campaigns) {
		var ret BulkCampaignOperationResponse
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsCampaignsBetaResponseContent) GetCampaignsOk() (*BulkCampaignOperationResponse, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *UpdateSponsoredBrandsCampaignsBetaResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given BulkCampaignOperationResponse and assigns it to the Campaigns field.
func (o *UpdateSponsoredBrandsCampaignsBetaResponseContent) SetCampaigns(v BulkCampaignOperationResponse) {
	o.Campaigns = &v
}

func (o UpdateSponsoredBrandsCampaignsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsCampaignsBetaResponseContent struct {
	value *UpdateSponsoredBrandsCampaignsBetaResponseContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsCampaignsBetaResponseContent) Get() *UpdateSponsoredBrandsCampaignsBetaResponseContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsCampaignsBetaResponseContent) Set(val *UpdateSponsoredBrandsCampaignsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsCampaignsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsCampaignsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsCampaignsBetaResponseContent(val *UpdateSponsoredBrandsCampaignsBetaResponseContent) *NullableUpdateSponsoredBrandsCampaignsBetaResponseContent {
	return &NullableUpdateSponsoredBrandsCampaignsBetaResponseContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsCampaignsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsCampaignsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
