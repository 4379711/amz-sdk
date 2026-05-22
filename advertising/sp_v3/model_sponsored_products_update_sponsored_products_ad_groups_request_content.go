package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent{}

// SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent struct for SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent
type SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent struct {
	// An array of adGroups with updated values.
	AdGroups []SponsoredProductsUpdateAdGroup `json:"adGroups"`
}

type _SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent

// NewSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent instantiates a new SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent(adGroups []SponsoredProductsUpdateAdGroup) *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent {
	this := SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) GetAdGroups() []SponsoredProductsUpdateAdGroup {
	if o == nil {
		var ret []SponsoredProductsUpdateAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) GetAdGroupsOk() ([]SponsoredProductsUpdateAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) SetAdGroups(v []SponsoredProductsUpdateAdGroup) {
	o.AdGroups = v
}

func (o SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent struct {
	value *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) Get() *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) Set(val *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent(val *SponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) *NullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
