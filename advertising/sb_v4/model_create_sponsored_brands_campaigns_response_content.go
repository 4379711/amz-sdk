package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsCampaignsResponseContent{}

// CreateSponsoredBrandsCampaignsResponseContent struct for CreateSponsoredBrandsCampaignsResponseContent
type CreateSponsoredBrandsCampaignsResponseContent struct {
	Campaigns *BulkCampaignOperationResponse `json:"campaigns,omitempty"`
}

// NewCreateSponsoredBrandsCampaignsResponseContent instantiates a new CreateSponsoredBrandsCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsCampaignsResponseContent() *CreateSponsoredBrandsCampaignsResponseContent {
	this := CreateSponsoredBrandsCampaignsResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsCampaignsResponseContentWithDefaults instantiates a new CreateSponsoredBrandsCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsCampaignsResponseContentWithDefaults() *CreateSponsoredBrandsCampaignsResponseContent {
	this := CreateSponsoredBrandsCampaignsResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsCampaignsResponseContent) GetCampaigns() BulkCampaignOperationResponse {
	if o == nil || IsNil(o.Campaigns) {
		var ret BulkCampaignOperationResponse
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsCampaignsResponseContent) GetCampaignsOk() (*BulkCampaignOperationResponse, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsCampaignsResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given BulkCampaignOperationResponse and assigns it to the Campaigns field.
func (o *CreateSponsoredBrandsCampaignsResponseContent) SetCampaigns(v BulkCampaignOperationResponse) {
	o.Campaigns = &v
}

func (o CreateSponsoredBrandsCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsCampaignsResponseContent struct {
	value *CreateSponsoredBrandsCampaignsResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsCampaignsResponseContent) Get() *CreateSponsoredBrandsCampaignsResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsCampaignsResponseContent) Set(val *CreateSponsoredBrandsCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsCampaignsResponseContent(val *CreateSponsoredBrandsCampaignsResponseContent) *NullableCreateSponsoredBrandsCampaignsResponseContent {
	return &NullableCreateSponsoredBrandsCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
