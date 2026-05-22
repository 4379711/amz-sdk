package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsCampaignsResponseContent{}

// DeleteSponsoredBrandsCampaignsResponseContent struct for DeleteSponsoredBrandsCampaignsResponseContent
type DeleteSponsoredBrandsCampaignsResponseContent struct {
	Campaigns *BulkCampaignOperationResponse `json:"campaigns,omitempty"`
}

// NewDeleteSponsoredBrandsCampaignsResponseContent instantiates a new DeleteSponsoredBrandsCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsCampaignsResponseContent() *DeleteSponsoredBrandsCampaignsResponseContent {
	this := DeleteSponsoredBrandsCampaignsResponseContent{}
	return &this
}

// NewDeleteSponsoredBrandsCampaignsResponseContentWithDefaults instantiates a new DeleteSponsoredBrandsCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsCampaignsResponseContentWithDefaults() *DeleteSponsoredBrandsCampaignsResponseContent {
	this := DeleteSponsoredBrandsCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsCampaignsResponseContent) GetCampaigns() BulkCampaignOperationResponse {
	if o == nil || IsNil(o.Campaigns) {
		var ret BulkCampaignOperationResponse
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsCampaignsResponseContent) GetCampaignsOk() (*BulkCampaignOperationResponse, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsCampaignsResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given BulkCampaignOperationResponse and assigns it to the Campaigns field.
func (o *DeleteSponsoredBrandsCampaignsResponseContent) SetCampaigns(v BulkCampaignOperationResponse) {
	o.Campaigns = &v
}

func (o DeleteSponsoredBrandsCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsCampaignsResponseContent struct {
	value *DeleteSponsoredBrandsCampaignsResponseContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsCampaignsResponseContent) Get() *DeleteSponsoredBrandsCampaignsResponseContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsCampaignsResponseContent) Set(val *DeleteSponsoredBrandsCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsCampaignsResponseContent(val *DeleteSponsoredBrandsCampaignsResponseContent) *NullableDeleteSponsoredBrandsCampaignsResponseContent {
	return &NullableDeleteSponsoredBrandsCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
