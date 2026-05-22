package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsCampaignsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsCampaignsBetaRequestContent{}

// UpdateSponsoredBrandsCampaignsBetaRequestContent struct for UpdateSponsoredBrandsCampaignsBetaRequestContent
type UpdateSponsoredBrandsCampaignsBetaRequestContent struct {
	Campaigns []UpdateCampaign `json:"campaigns"`
}

type _UpdateSponsoredBrandsCampaignsBetaRequestContent UpdateSponsoredBrandsCampaignsBetaRequestContent

// NewUpdateSponsoredBrandsCampaignsBetaRequestContent instantiates a new UpdateSponsoredBrandsCampaignsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsCampaignsBetaRequestContent(campaigns []UpdateCampaign) *UpdateSponsoredBrandsCampaignsBetaRequestContent {
	this := UpdateSponsoredBrandsCampaignsBetaRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewUpdateSponsoredBrandsCampaignsBetaRequestContentWithDefaults instantiates a new UpdateSponsoredBrandsCampaignsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsCampaignsBetaRequestContentWithDefaults() *UpdateSponsoredBrandsCampaignsBetaRequestContent {
	this := UpdateSponsoredBrandsCampaignsBetaRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *UpdateSponsoredBrandsCampaignsBetaRequestContent) GetCampaigns() []UpdateCampaign {
	if o == nil {
		var ret []UpdateCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsCampaignsBetaRequestContent) GetCampaignsOk() ([]UpdateCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *UpdateSponsoredBrandsCampaignsBetaRequestContent) SetCampaigns(v []UpdateCampaign) {
	o.Campaigns = v
}

func (o UpdateSponsoredBrandsCampaignsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsCampaignsBetaRequestContent struct {
	value *UpdateSponsoredBrandsCampaignsBetaRequestContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsCampaignsBetaRequestContent) Get() *UpdateSponsoredBrandsCampaignsBetaRequestContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsCampaignsBetaRequestContent) Set(val *UpdateSponsoredBrandsCampaignsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsCampaignsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsCampaignsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsCampaignsBetaRequestContent(val *UpdateSponsoredBrandsCampaignsBetaRequestContent) *NullableUpdateSponsoredBrandsCampaignsBetaRequestContent {
	return &NullableUpdateSponsoredBrandsCampaignsBetaRequestContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsCampaignsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsCampaignsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
