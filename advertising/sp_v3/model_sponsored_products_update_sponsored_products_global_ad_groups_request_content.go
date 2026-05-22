package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent struct for SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent
type SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent struct {
	// An array of adGroups with updated values.
	AdGroups []SponsoredProductsUpdateGlobalAdGroup `json:"adGroups"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent(adGroups []SponsoredProductsUpdateGlobalAdGroup) *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) GetAdGroups() []SponsoredProductsUpdateGlobalAdGroup {
	if o == nil {
		var ret []SponsoredProductsUpdateGlobalAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) GetAdGroupsOk() ([]SponsoredProductsUpdateGlobalAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) SetAdGroups(v []SponsoredProductsUpdateGlobalAdGroup) {
	o.AdGroups = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent(val *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
