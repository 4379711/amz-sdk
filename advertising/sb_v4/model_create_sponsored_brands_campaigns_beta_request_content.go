package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsCampaignsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsCampaignsBetaRequestContent{}

// CreateSponsoredBrandsCampaignsBetaRequestContent struct for CreateSponsoredBrandsCampaignsBetaRequestContent
type CreateSponsoredBrandsCampaignsBetaRequestContent struct {
	Campaigns []CreateCampaign `json:"campaigns"`
}

type _CreateSponsoredBrandsCampaignsBetaRequestContent CreateSponsoredBrandsCampaignsBetaRequestContent

// NewCreateSponsoredBrandsCampaignsBetaRequestContent instantiates a new CreateSponsoredBrandsCampaignsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsCampaignsBetaRequestContent(campaigns []CreateCampaign) *CreateSponsoredBrandsCampaignsBetaRequestContent {
	this := CreateSponsoredBrandsCampaignsBetaRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewCreateSponsoredBrandsCampaignsBetaRequestContentWithDefaults instantiates a new CreateSponsoredBrandsCampaignsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsCampaignsBetaRequestContentWithDefaults() *CreateSponsoredBrandsCampaignsBetaRequestContent {
	this := CreateSponsoredBrandsCampaignsBetaRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *CreateSponsoredBrandsCampaignsBetaRequestContent) GetCampaigns() []CreateCampaign {
	if o == nil {
		var ret []CreateCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsCampaignsBetaRequestContent) GetCampaignsOk() ([]CreateCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *CreateSponsoredBrandsCampaignsBetaRequestContent) SetCampaigns(v []CreateCampaign) {
	o.Campaigns = v
}

func (o CreateSponsoredBrandsCampaignsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsCampaignsBetaRequestContent struct {
	value *CreateSponsoredBrandsCampaignsBetaRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsCampaignsBetaRequestContent) Get() *CreateSponsoredBrandsCampaignsBetaRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsCampaignsBetaRequestContent) Set(val *CreateSponsoredBrandsCampaignsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsCampaignsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsCampaignsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsCampaignsBetaRequestContent(val *CreateSponsoredBrandsCampaignsBetaRequestContent) *NullableCreateSponsoredBrandsCampaignsBetaRequestContent {
	return &NullableCreateSponsoredBrandsCampaignsBetaRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsCampaignsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsCampaignsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
