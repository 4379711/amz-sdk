package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsAdGroupsRequestContent{}

// CreateSponsoredBrandsAdGroupsRequestContent struct for CreateSponsoredBrandsAdGroupsRequestContent
type CreateSponsoredBrandsAdGroupsRequestContent struct {
	AdGroups []CreateAdGroup `json:"adGroups"`
}

type _CreateSponsoredBrandsAdGroupsRequestContent CreateSponsoredBrandsAdGroupsRequestContent

// NewCreateSponsoredBrandsAdGroupsRequestContent instantiates a new CreateSponsoredBrandsAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsAdGroupsRequestContent(adGroups []CreateAdGroup) *CreateSponsoredBrandsAdGroupsRequestContent {
	this := CreateSponsoredBrandsAdGroupsRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewCreateSponsoredBrandsAdGroupsRequestContentWithDefaults instantiates a new CreateSponsoredBrandsAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsAdGroupsRequestContentWithDefaults() *CreateSponsoredBrandsAdGroupsRequestContent {
	this := CreateSponsoredBrandsAdGroupsRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *CreateSponsoredBrandsAdGroupsRequestContent) GetAdGroups() []CreateAdGroup {
	if o == nil {
		var ret []CreateAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsAdGroupsRequestContent) GetAdGroupsOk() ([]CreateAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *CreateSponsoredBrandsAdGroupsRequestContent) SetAdGroups(v []CreateAdGroup) {
	o.AdGroups = v
}

func (o CreateSponsoredBrandsAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsAdGroupsRequestContent struct {
	value *CreateSponsoredBrandsAdGroupsRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsAdGroupsRequestContent) Get() *CreateSponsoredBrandsAdGroupsRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsAdGroupsRequestContent) Set(val *CreateSponsoredBrandsAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsAdGroupsRequestContent(val *CreateSponsoredBrandsAdGroupsRequestContent) *NullableCreateSponsoredBrandsAdGroupsRequestContent {
	return &NullableCreateSponsoredBrandsAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
