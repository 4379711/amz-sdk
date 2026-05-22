package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsCampaignsRequestContent{}

// CreateSponsoredBrandsCampaignsRequestContent struct for CreateSponsoredBrandsCampaignsRequestContent
type CreateSponsoredBrandsCampaignsRequestContent struct {
	Campaigns []CreateCampaign `json:"campaigns"`
}

type _CreateSponsoredBrandsCampaignsRequestContent CreateSponsoredBrandsCampaignsRequestContent

// NewCreateSponsoredBrandsCampaignsRequestContent instantiates a new CreateSponsoredBrandsCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsCampaignsRequestContent(campaigns []CreateCampaign) *CreateSponsoredBrandsCampaignsRequestContent {
	this := CreateSponsoredBrandsCampaignsRequestContent{}
	this.Campaigns = campaigns
	return &this
}

// NewCreateSponsoredBrandsCampaignsRequestContentWithDefaults instantiates a new CreateSponsoredBrandsCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsCampaignsRequestContentWithDefaults() *CreateSponsoredBrandsCampaignsRequestContent {
	this := CreateSponsoredBrandsCampaignsRequestContent{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *CreateSponsoredBrandsCampaignsRequestContent) GetCampaigns() []CreateCampaign {
	if o == nil {
		var ret []CreateCampaign
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsCampaignsRequestContent) GetCampaignsOk() ([]CreateCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *CreateSponsoredBrandsCampaignsRequestContent) SetCampaigns(v []CreateCampaign) {
	o.Campaigns = v
}

func (o CreateSponsoredBrandsCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsCampaignsRequestContent struct {
	value *CreateSponsoredBrandsCampaignsRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsCampaignsRequestContent) Get() *CreateSponsoredBrandsCampaignsRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsCampaignsRequestContent) Set(val *CreateSponsoredBrandsCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsCampaignsRequestContent(val *CreateSponsoredBrandsCampaignsRequestContent) *NullableCreateSponsoredBrandsCampaignsRequestContent {
	return &NullableCreateSponsoredBrandsCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
