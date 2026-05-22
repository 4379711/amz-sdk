package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsCampaignsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsCampaignsBetaResponseContent{}

// DeleteSponsoredBrandsCampaignsBetaResponseContent struct for DeleteSponsoredBrandsCampaignsBetaResponseContent
type DeleteSponsoredBrandsCampaignsBetaResponseContent struct {
	Campaigns *BulkCampaignOperationResponse `json:"campaigns,omitempty"`
}

// NewDeleteSponsoredBrandsCampaignsBetaResponseContent instantiates a new DeleteSponsoredBrandsCampaignsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsCampaignsBetaResponseContent() *DeleteSponsoredBrandsCampaignsBetaResponseContent {
	this := DeleteSponsoredBrandsCampaignsBetaResponseContent{}
	return &this
}

// NewDeleteSponsoredBrandsCampaignsBetaResponseContentWithDefaults instantiates a new DeleteSponsoredBrandsCampaignsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsCampaignsBetaResponseContentWithDefaults() *DeleteSponsoredBrandsCampaignsBetaResponseContent {
	this := DeleteSponsoredBrandsCampaignsBetaResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsCampaignsBetaResponseContent) GetCampaigns() BulkCampaignOperationResponse {
	if o == nil || IsNil(o.Campaigns) {
		var ret BulkCampaignOperationResponse
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsCampaignsBetaResponseContent) GetCampaignsOk() (*BulkCampaignOperationResponse, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsCampaignsBetaResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given BulkCampaignOperationResponse and assigns it to the Campaigns field.
func (o *DeleteSponsoredBrandsCampaignsBetaResponseContent) SetCampaigns(v BulkCampaignOperationResponse) {
	o.Campaigns = &v
}

func (o DeleteSponsoredBrandsCampaignsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsCampaignsBetaResponseContent struct {
	value *DeleteSponsoredBrandsCampaignsBetaResponseContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsCampaignsBetaResponseContent) Get() *DeleteSponsoredBrandsCampaignsBetaResponseContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsCampaignsBetaResponseContent) Set(val *DeleteSponsoredBrandsCampaignsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsCampaignsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsCampaignsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsCampaignsBetaResponseContent(val *DeleteSponsoredBrandsCampaignsBetaResponseContent) *NullableDeleteSponsoredBrandsCampaignsBetaResponseContent {
	return &NullableDeleteSponsoredBrandsCampaignsBetaResponseContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsCampaignsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsCampaignsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
