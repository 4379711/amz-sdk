package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsCampaignsRequestContent{}

// DeleteSponsoredBrandsCampaignsRequestContent struct for DeleteSponsoredBrandsCampaignsRequestContent
type DeleteSponsoredBrandsCampaignsRequestContent struct {
	CampaignIdFilter *ObjectIdFilter `json:"campaignIdFilter,omitempty"`
}

// NewDeleteSponsoredBrandsCampaignsRequestContent instantiates a new DeleteSponsoredBrandsCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsCampaignsRequestContent() *DeleteSponsoredBrandsCampaignsRequestContent {
	this := DeleteSponsoredBrandsCampaignsRequestContent{}
	return &this
}

// NewDeleteSponsoredBrandsCampaignsRequestContentWithDefaults instantiates a new DeleteSponsoredBrandsCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsCampaignsRequestContentWithDefaults() *DeleteSponsoredBrandsCampaignsRequestContent {
	this := DeleteSponsoredBrandsCampaignsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsCampaignsRequestContent) GetCampaignIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsCampaignsRequestContent) GetCampaignIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsCampaignsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given ObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *DeleteSponsoredBrandsCampaignsRequestContent) SetCampaignIdFilter(v ObjectIdFilter) {
	o.CampaignIdFilter = &v
}

func (o DeleteSponsoredBrandsCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsCampaignsRequestContent struct {
	value *DeleteSponsoredBrandsCampaignsRequestContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsCampaignsRequestContent) Get() *DeleteSponsoredBrandsCampaignsRequestContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsCampaignsRequestContent) Set(val *DeleteSponsoredBrandsCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsCampaignsRequestContent(val *DeleteSponsoredBrandsCampaignsRequestContent) *NullableDeleteSponsoredBrandsCampaignsRequestContent {
	return &NullableDeleteSponsoredBrandsCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
