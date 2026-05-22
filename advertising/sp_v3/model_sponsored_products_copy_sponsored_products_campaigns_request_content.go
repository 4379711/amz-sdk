package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCopySponsoredProductsCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCopySponsoredProductsCampaignsRequestContent{}

// SponsoredProductsCopySponsoredProductsCampaignsRequestContent struct for SponsoredProductsCopySponsoredProductsCampaignsRequestContent
type SponsoredProductsCopySponsoredProductsCampaignsRequestContent struct {
	// An array of campaigns.
	CopyCampaignsItems []SponsoredProductsCopyCampaign `json:"copyCampaignsItems"`
}

type _SponsoredProductsCopySponsoredProductsCampaignsRequestContent SponsoredProductsCopySponsoredProductsCampaignsRequestContent

// NewSponsoredProductsCopySponsoredProductsCampaignsRequestContent instantiates a new SponsoredProductsCopySponsoredProductsCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCopySponsoredProductsCampaignsRequestContent(copyCampaignsItems []SponsoredProductsCopyCampaign) *SponsoredProductsCopySponsoredProductsCampaignsRequestContent {
	this := SponsoredProductsCopySponsoredProductsCampaignsRequestContent{}
	this.CopyCampaignsItems = copyCampaignsItems
	return &this
}

// NewSponsoredProductsCopySponsoredProductsCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsCopySponsoredProductsCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCopySponsoredProductsCampaignsRequestContentWithDefaults() *SponsoredProductsCopySponsoredProductsCampaignsRequestContent {
	this := SponsoredProductsCopySponsoredProductsCampaignsRequestContent{}
	return &this
}

// GetCopyCampaignsItems returns the CopyCampaignsItems field value
func (o *SponsoredProductsCopySponsoredProductsCampaignsRequestContent) GetCopyCampaignsItems() []SponsoredProductsCopyCampaign {
	if o == nil {
		var ret []SponsoredProductsCopyCampaign
		return ret
	}

	return o.CopyCampaignsItems
}

// GetCopyCampaignsItemsOk returns a tuple with the CopyCampaignsItems field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopySponsoredProductsCampaignsRequestContent) GetCopyCampaignsItemsOk() ([]SponsoredProductsCopyCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.CopyCampaignsItems, true
}

// SetCopyCampaignsItems sets field value
func (o *SponsoredProductsCopySponsoredProductsCampaignsRequestContent) SetCopyCampaignsItems(v []SponsoredProductsCopyCampaign) {
	o.CopyCampaignsItems = v
}

func (o SponsoredProductsCopySponsoredProductsCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["copyCampaignsItems"] = o.CopyCampaignsItems
	return toSerialize, nil
}

type NullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent struct {
	value *SponsoredProductsCopySponsoredProductsCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent) Get() *SponsoredProductsCopySponsoredProductsCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent) Set(val *SponsoredProductsCopySponsoredProductsCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent(val *SponsoredProductsCopySponsoredProductsCampaignsRequestContent) *NullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent {
	return &NullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCopySponsoredProductsCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
