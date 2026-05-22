package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent{}

// SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent struct for SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent
type SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent struct {
	// An array of adGroups.
	AdGroups []SponsoredProductsCreateGlobalAdGroup `json:"adGroups"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent

// NewSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent(adGroups []SponsoredProductsCreateGlobalAdGroup) *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) GetAdGroups() []SponsoredProductsCreateGlobalAdGroup {
	if o == nil {
		var ret []SponsoredProductsCreateGlobalAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) GetAdGroupsOk() ([]SponsoredProductsCreateGlobalAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) SetAdGroups(v []SponsoredProductsCreateGlobalAdGroup) {
	o.AdGroups = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent(val *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
