package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsAdGroupsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsAdGroupsBetaRequestContent{}

// UpdateSponsoredBrandsAdGroupsBetaRequestContent struct for UpdateSponsoredBrandsAdGroupsBetaRequestContent
type UpdateSponsoredBrandsAdGroupsBetaRequestContent struct {
	AdGroups []UpdateAdGroup `json:"adGroups"`
}

type _UpdateSponsoredBrandsAdGroupsBetaRequestContent UpdateSponsoredBrandsAdGroupsBetaRequestContent

// NewUpdateSponsoredBrandsAdGroupsBetaRequestContent instantiates a new UpdateSponsoredBrandsAdGroupsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsAdGroupsBetaRequestContent(adGroups []UpdateAdGroup) *UpdateSponsoredBrandsAdGroupsBetaRequestContent {
	this := UpdateSponsoredBrandsAdGroupsBetaRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewUpdateSponsoredBrandsAdGroupsBetaRequestContentWithDefaults instantiates a new UpdateSponsoredBrandsAdGroupsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsAdGroupsBetaRequestContentWithDefaults() *UpdateSponsoredBrandsAdGroupsBetaRequestContent {
	this := UpdateSponsoredBrandsAdGroupsBetaRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *UpdateSponsoredBrandsAdGroupsBetaRequestContent) GetAdGroups() []UpdateAdGroup {
	if o == nil {
		var ret []UpdateAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsAdGroupsBetaRequestContent) GetAdGroupsOk() ([]UpdateAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *UpdateSponsoredBrandsAdGroupsBetaRequestContent) SetAdGroups(v []UpdateAdGroup) {
	o.AdGroups = v
}

func (o UpdateSponsoredBrandsAdGroupsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsAdGroupsBetaRequestContent struct {
	value *UpdateSponsoredBrandsAdGroupsBetaRequestContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsAdGroupsBetaRequestContent) Get() *UpdateSponsoredBrandsAdGroupsBetaRequestContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsAdGroupsBetaRequestContent) Set(val *UpdateSponsoredBrandsAdGroupsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsAdGroupsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsAdGroupsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsAdGroupsBetaRequestContent(val *UpdateSponsoredBrandsAdGroupsBetaRequestContent) *NullableUpdateSponsoredBrandsAdGroupsBetaRequestContent {
	return &NullableUpdateSponsoredBrandsAdGroupsBetaRequestContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsAdGroupsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsAdGroupsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
