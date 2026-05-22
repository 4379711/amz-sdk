package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsCampaignsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsCampaignsBetaResponseContent{}

// CreateSponsoredBrandsCampaignsBetaResponseContent struct for CreateSponsoredBrandsCampaignsBetaResponseContent
type CreateSponsoredBrandsCampaignsBetaResponseContent struct {
	Campaigns *BulkCampaignOperationResponse `json:"campaigns,omitempty"`
}

// NewCreateSponsoredBrandsCampaignsBetaResponseContent instantiates a new CreateSponsoredBrandsCampaignsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsCampaignsBetaResponseContent() *CreateSponsoredBrandsCampaignsBetaResponseContent {
	this := CreateSponsoredBrandsCampaignsBetaResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsCampaignsBetaResponseContentWithDefaults instantiates a new CreateSponsoredBrandsCampaignsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsCampaignsBetaResponseContentWithDefaults() *CreateSponsoredBrandsCampaignsBetaResponseContent {
	this := CreateSponsoredBrandsCampaignsBetaResponseContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsCampaignsBetaResponseContent) GetCampaigns() BulkCampaignOperationResponse {
	if o == nil || IsNil(o.Campaigns) {
		var ret BulkCampaignOperationResponse
		return ret
	}
	return *o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsCampaignsBetaResponseContent) GetCampaignsOk() (*BulkCampaignOperationResponse, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsCampaignsBetaResponseContent) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given BulkCampaignOperationResponse and assigns it to the Campaigns field.
func (o *CreateSponsoredBrandsCampaignsBetaResponseContent) SetCampaigns(v BulkCampaignOperationResponse) {
	o.Campaigns = &v
}

func (o CreateSponsoredBrandsCampaignsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaigns) {
		toSerialize["campaigns"] = o.Campaigns
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsCampaignsBetaResponseContent struct {
	value *CreateSponsoredBrandsCampaignsBetaResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsCampaignsBetaResponseContent) Get() *CreateSponsoredBrandsCampaignsBetaResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsCampaignsBetaResponseContent) Set(val *CreateSponsoredBrandsCampaignsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsCampaignsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsCampaignsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsCampaignsBetaResponseContent(val *CreateSponsoredBrandsCampaignsBetaResponseContent) *NullableCreateSponsoredBrandsCampaignsBetaResponseContent {
	return &NullableCreateSponsoredBrandsCampaignsBetaResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsCampaignsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsCampaignsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
