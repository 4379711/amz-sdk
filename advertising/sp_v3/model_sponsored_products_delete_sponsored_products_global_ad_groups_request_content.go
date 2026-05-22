package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent struct for SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent
type SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent struct {
	AdGroupIdFilter *SponsoredProductsObjectIdFilter `json:"adGroupIdFilter,omitempty"`
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent() *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent{}
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent{}
	return &this
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) GetAdGroupIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) SetAdGroupIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent(val *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
