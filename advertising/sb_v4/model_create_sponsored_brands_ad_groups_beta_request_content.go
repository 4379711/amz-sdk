package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsAdGroupsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsAdGroupsBetaRequestContent{}

// CreateSponsoredBrandsAdGroupsBetaRequestContent struct for CreateSponsoredBrandsAdGroupsBetaRequestContent
type CreateSponsoredBrandsAdGroupsBetaRequestContent struct {
	AdGroups []CreateAdGroup `json:"adGroups"`
}

type _CreateSponsoredBrandsAdGroupsBetaRequestContent CreateSponsoredBrandsAdGroupsBetaRequestContent

// NewCreateSponsoredBrandsAdGroupsBetaRequestContent instantiates a new CreateSponsoredBrandsAdGroupsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsAdGroupsBetaRequestContent(adGroups []CreateAdGroup) *CreateSponsoredBrandsAdGroupsBetaRequestContent {
	this := CreateSponsoredBrandsAdGroupsBetaRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewCreateSponsoredBrandsAdGroupsBetaRequestContentWithDefaults instantiates a new CreateSponsoredBrandsAdGroupsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsAdGroupsBetaRequestContentWithDefaults() *CreateSponsoredBrandsAdGroupsBetaRequestContent {
	this := CreateSponsoredBrandsAdGroupsBetaRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *CreateSponsoredBrandsAdGroupsBetaRequestContent) GetAdGroups() []CreateAdGroup {
	if o == nil {
		var ret []CreateAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsAdGroupsBetaRequestContent) GetAdGroupsOk() ([]CreateAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *CreateSponsoredBrandsAdGroupsBetaRequestContent) SetAdGroups(v []CreateAdGroup) {
	o.AdGroups = v
}

func (o CreateSponsoredBrandsAdGroupsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsAdGroupsBetaRequestContent struct {
	value *CreateSponsoredBrandsAdGroupsBetaRequestContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsAdGroupsBetaRequestContent) Get() *CreateSponsoredBrandsAdGroupsBetaRequestContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsAdGroupsBetaRequestContent) Set(val *CreateSponsoredBrandsAdGroupsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsAdGroupsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsAdGroupsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsAdGroupsBetaRequestContent(val *CreateSponsoredBrandsAdGroupsBetaRequestContent) *NullableCreateSponsoredBrandsAdGroupsBetaRequestContent {
	return &NullableCreateSponsoredBrandsAdGroupsBetaRequestContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsAdGroupsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsAdGroupsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
