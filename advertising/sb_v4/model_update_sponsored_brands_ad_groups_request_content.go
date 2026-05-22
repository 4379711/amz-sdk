package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsAdGroupsRequestContent{}

// UpdateSponsoredBrandsAdGroupsRequestContent struct for UpdateSponsoredBrandsAdGroupsRequestContent
type UpdateSponsoredBrandsAdGroupsRequestContent struct {
	AdGroups []UpdateAdGroup `json:"adGroups"`
}

type _UpdateSponsoredBrandsAdGroupsRequestContent UpdateSponsoredBrandsAdGroupsRequestContent

// NewUpdateSponsoredBrandsAdGroupsRequestContent instantiates a new UpdateSponsoredBrandsAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsAdGroupsRequestContent(adGroups []UpdateAdGroup) *UpdateSponsoredBrandsAdGroupsRequestContent {
	this := UpdateSponsoredBrandsAdGroupsRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewUpdateSponsoredBrandsAdGroupsRequestContentWithDefaults instantiates a new UpdateSponsoredBrandsAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsAdGroupsRequestContentWithDefaults() *UpdateSponsoredBrandsAdGroupsRequestContent {
	this := UpdateSponsoredBrandsAdGroupsRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *UpdateSponsoredBrandsAdGroupsRequestContent) GetAdGroups() []UpdateAdGroup {
	if o == nil {
		var ret []UpdateAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsAdGroupsRequestContent) GetAdGroupsOk() ([]UpdateAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *UpdateSponsoredBrandsAdGroupsRequestContent) SetAdGroups(v []UpdateAdGroup) {
	o.AdGroups = v
}

func (o UpdateSponsoredBrandsAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsAdGroupsRequestContent struct {
	value *UpdateSponsoredBrandsAdGroupsRequestContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsAdGroupsRequestContent) Get() *UpdateSponsoredBrandsAdGroupsRequestContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsAdGroupsRequestContent) Set(val *UpdateSponsoredBrandsAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsAdGroupsRequestContent(val *UpdateSponsoredBrandsAdGroupsRequestContent) *NullableUpdateSponsoredBrandsAdGroupsRequestContent {
	return &NullableUpdateSponsoredBrandsAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
