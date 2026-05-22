package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent{}

// SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent struct for SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent
type SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent struct {
	// An array of draftAdGroups.
	AdGroups []SponsoredProductsCreateDraftAdGroup `json:"adGroups"`
}

type _SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent

// NewSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent(adGroups []SponsoredProductsCreateDraftAdGroup) *SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) GetAdGroups() []SponsoredProductsCreateDraftAdGroup {
	if o == nil {
		var ret []SponsoredProductsCreateDraftAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) GetAdGroupsOk() ([]SponsoredProductsCreateDraftAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) SetAdGroups(v []SponsoredProductsCreateDraftAdGroup) {
	o.AdGroups = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent(val *SponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
