package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent struct for SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent
type SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent struct {
	AdGroupIdFilter SponsoredProductsObjectIdFilter `json:"adGroupIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent(adGroupIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent{}
	this.AdGroupIdFilter = adGroupIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent{}
	return &this
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupIdFilter, true
}

// SetAdGroupIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent(val *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
