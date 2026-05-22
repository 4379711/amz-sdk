package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent struct for SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent
type SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent struct {
	CampaignIdFilter *SponsoredProductsObjectIdFilter `json:"campaignIdFilter,omitempty"`
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent{}
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = &v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent(val *SponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
