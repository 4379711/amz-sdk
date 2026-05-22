package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent struct for SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent
type SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent struct {
	AdGroupIdFilter SponsoredProductsObjectIdFilter `json:"adGroupIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent

// NewSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent(adGroupIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent{}
	this.AdGroupIdFilter = adGroupIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent{}
	return &this
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupIdFilter, true
}

// SetAdGroupIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent(val *SponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
