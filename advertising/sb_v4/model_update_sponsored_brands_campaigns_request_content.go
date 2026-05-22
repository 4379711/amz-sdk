package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsCampaignsRequestContent{}

// UpdateSponsoredBrandsCampaignsRequestContent struct for UpdateSponsoredBrandsCampaignsRequestContent
type UpdateSponsoredBrandsCampaignsRequestContent struct {
	Campaigns []UpdateCampaign `json:"campaigns"`
}

type _UpdateSponsoredBrandsCampaignsRequestContent UpdateSponsoredBrandsCampaignsRequestContent

// NewUpdateSponsoredBrandsCampaignsRequestContent instantiates a new UpdateSponsoredBrandsCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsCampaignsRequestContent(campaigns []UpdateCampaign) *UpdateSponsoredBrandsCampaignsRequestContent {
	this := UpdateSponsoredBrandsCampaignsRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewUpdateSponsoredBrandsCampaignsRequestContentWithDefaults instantiates a new UpdateSponsoredBrandsCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsCampaignsRequestContentWithDefaults() *UpdateSponsoredBrandsCampaignsRequestContent {
	this := UpdateSponsoredBrandsCampaignsRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *UpdateSponsoredBrandsCampaignsRequestContent) GetCampaigns() []UpdateCampaign {
	if o == nil {
		var ret []UpdateCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsCampaignsRequestContent) GetCampaignsOk() ([]UpdateCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *UpdateSponsoredBrandsCampaignsRequestContent) SetCampaigns(v []UpdateCampaign) {
	o.Campaigns = v
}

func (o UpdateSponsoredBrandsCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsCampaignsRequestContent struct {
	value *UpdateSponsoredBrandsCampaignsRequestContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsCampaignsRequestContent) Get() *UpdateSponsoredBrandsCampaignsRequestContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsCampaignsRequestContent) Set(val *UpdateSponsoredBrandsCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsCampaignsRequestContent(val *UpdateSponsoredBrandsCampaignsRequestContent) *NullableUpdateSponsoredBrandsCampaignsRequestContent {
	return &NullableUpdateSponsoredBrandsCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
