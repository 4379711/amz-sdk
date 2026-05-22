package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsCampaignsResponseContent{}

// UpdateSponsoredBrandsCampaignsResponseContent struct for UpdateSponsoredBrandsCampaignsResponseContent
type UpdateSponsoredBrandsCampaignsResponseContent struct {
	Campaigns *BulkCampaignOperationResponse `json:"campaigns,omitempty"`
}

// NewUpdateSponsoredBrandsCampaignsResponseContent instantiates a new UpdateSponsoredBrandsCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsCampaignsResponseContent() *UpdateSponsoredBrandsCampaignsResponseContent {
	this := UpdateSponsoredBrandsCampaignsResponseContent{}
	return &this
}

// NewUpdateSponsoredBrandsCampaignsResponseContentWithDefaults instantiates a new UpdateSponsoredBrandsCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsCampaignsResponseContentWithDefaults() *UpdateSponsoredBrandsCampaignsResponseContent {
	this := UpdateSponsoredBrandsCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *UpdateSponsoredBrandsCampaignsResponseContent) GetCampaigns() BulkCampaignOperationResponse {
	if o == nil || IsNil(o.Campaigns) {
		var ret BulkCampaignOperationResponse
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsCampaignsResponseContent) GetCampaignsOk() (*BulkCampaignOperationResponse, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *UpdateSponsoredBrandsCampaignsResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given BulkCampaignOperationResponse and assigns it to the Campaigns field.
func (o *UpdateSponsoredBrandsCampaignsResponseContent) SetCampaigns(v BulkCampaignOperationResponse) {
	o.Campaigns = &v
}

func (o UpdateSponsoredBrandsCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsCampaignsResponseContent struct {
	value *UpdateSponsoredBrandsCampaignsResponseContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsCampaignsResponseContent) Get() *UpdateSponsoredBrandsCampaignsResponseContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsCampaignsResponseContent) Set(val *UpdateSponsoredBrandsCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsCampaignsResponseContent(val *UpdateSponsoredBrandsCampaignsResponseContent) *NullableUpdateSponsoredBrandsCampaignsResponseContent {
	return &NullableUpdateSponsoredBrandsCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
