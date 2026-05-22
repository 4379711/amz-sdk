package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsCampaignsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsCampaignsBetaRequestContent{}

// DeleteSponsoredBrandsCampaignsBetaRequestContent struct for DeleteSponsoredBrandsCampaignsBetaRequestContent
type DeleteSponsoredBrandsCampaignsBetaRequestContent struct {
	CampaignIdFilter *ObjectIdFilter `json:"campaignIdFilter,omitempty"`
}

// NewDeleteSponsoredBrandsCampaignsBetaRequestContent instantiates a new DeleteSponsoredBrandsCampaignsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsCampaignsBetaRequestContent() *DeleteSponsoredBrandsCampaignsBetaRequestContent {
	this := DeleteSponsoredBrandsCampaignsBetaRequestContent{}
	return &this
}

// NewDeleteSponsoredBrandsCampaignsBetaRequestContentWithDefaults instantiates a new DeleteSponsoredBrandsCampaignsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsCampaignsBetaRequestContentWithDefaults() *DeleteSponsoredBrandsCampaignsBetaRequestContent {
	this := DeleteSponsoredBrandsCampaignsBetaRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsCampaignsBetaRequestContent) GetCampaignIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsCampaignsBetaRequestContent) GetCampaignIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsCampaignsBetaRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given ObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *DeleteSponsoredBrandsCampaignsBetaRequestContent) SetCampaignIdFilter(v ObjectIdFilter) {
	o.CampaignIdFilter = &v
}

func (o DeleteSponsoredBrandsCampaignsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsCampaignsBetaRequestContent struct {
	value *DeleteSponsoredBrandsCampaignsBetaRequestContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsCampaignsBetaRequestContent) Get() *DeleteSponsoredBrandsCampaignsBetaRequestContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsCampaignsBetaRequestContent) Set(val *DeleteSponsoredBrandsCampaignsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsCampaignsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsCampaignsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsCampaignsBetaRequestContent(val *DeleteSponsoredBrandsCampaignsBetaRequestContent) *NullableDeleteSponsoredBrandsCampaignsBetaRequestContent {
	return &NullableDeleteSponsoredBrandsCampaignsBetaRequestContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsCampaignsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsCampaignsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
