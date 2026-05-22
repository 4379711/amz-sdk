package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCopySponsoredProductsCampaignsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCopySponsoredProductsCampaignsResponseContent{}

// SponsoredProductsCopySponsoredProductsCampaignsResponseContent struct for SponsoredProductsCopySponsoredProductsCampaignsResponseContent
type SponsoredProductsCopySponsoredProductsCampaignsResponseContent struct {
	CopyCampaignResponses []SponsoredProductsCopyCampaignResponse `json:"copyCampaignResponses"`
}

type _SponsoredProductsCopySponsoredProductsCampaignsResponseContent SponsoredProductsCopySponsoredProductsCampaignsResponseContent

// NewSponsoredProductsCopySponsoredProductsCampaignsResponseContent instantiates a new SponsoredProductsCopySponsoredProductsCampaignsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCopySponsoredProductsCampaignsResponseContent(copyCampaignResponses []SponsoredProductsCopyCampaignResponse) *SponsoredProductsCopySponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsCopySponsoredProductsCampaignsResponseContent{}
	this.CopyCampaignResponses = copyCampaignResponses
	return &this
}

// NewSponsoredProductsCopySponsoredProductsCampaignsResponseContentWithDefaults instantiates a new SponsoredProductsCopySponsoredProductsCampaignsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCopySponsoredProductsCampaignsResponseContentWithDefaults() *SponsoredProductsCopySponsoredProductsCampaignsResponseContent {
	this := SponsoredProductsCopySponsoredProductsCampaignsResponseContent{}
	return &this
}

// GetCopyCampaignResponses returns the CopyCampaignResponses field value
func (o *SponsoredProductsCopySponsoredProductsCampaignsResponseContent) GetCopyCampaignResponses() []SponsoredProductsCopyCampaignResponse {
	if o == nil {
		var ret []SponsoredProductsCopyCampaignResponse
		return ret
	}

	return o.CopyCampaignResponses
}

// GetCopyCampaignResponsesOk returns a tuple with the CopyCampaignResponses field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopySponsoredProductsCampaignsResponseContent) GetCopyCampaignResponsesOk() ([]SponsoredProductsCopyCampaignResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.CopyCampaignResponses, true
}

// SetCopyCampaignResponses sets field value
func (o *SponsoredProductsCopySponsoredProductsCampaignsResponseContent) SetCopyCampaignResponses(v []SponsoredProductsCopyCampaignResponse) {
	o.CopyCampaignResponses = v
}

func (o SponsoredProductsCopySponsoredProductsCampaignsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["copyCampaignResponses"] = o.CopyCampaignResponses
	return toSerialize, nil
}

type NullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent struct {
	value *SponsoredProductsCopySponsoredProductsCampaignsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent) Get() *SponsoredProductsCopySponsoredProductsCampaignsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent) Set(val *SponsoredProductsCopySponsoredProductsCampaignsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent(val *SponsoredProductsCopySponsoredProductsCampaignsResponseContent) *NullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent {
	return &NullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCopySponsoredProductsCampaignsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
