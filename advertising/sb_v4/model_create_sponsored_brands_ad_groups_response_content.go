package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsAdGroupsResponseContent{}

// CreateSponsoredBrandsAdGroupsResponseContent struct for CreateSponsoredBrandsAdGroupsResponseContent
type CreateSponsoredBrandsAdGroupsResponseContent struct {
	AdGroups *BulkAdGroupOperationResponse `json:"adGroups,omitempty"`
}

// NewCreateSponsoredBrandsAdGroupsResponseContent instantiates a new CreateSponsoredBrandsAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsAdGroupsResponseContent() *CreateSponsoredBrandsAdGroupsResponseContent {
	this := CreateSponsoredBrandsAdGroupsResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsAdGroupsResponseContentWithDefaults instantiates a new CreateSponsoredBrandsAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsAdGroupsResponseContentWithDefaults() *CreateSponsoredBrandsAdGroupsResponseContent {
	this := CreateSponsoredBrandsAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsAdGroupsResponseContent) GetAdGroups() BulkAdGroupOperationResponse {
	if o == nil || IsNil(o.AdGroups) {
		var ret BulkAdGroupOperationResponse
		return ret
	}
	return *o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsAdGroupsResponseContent) GetAdGroupsOk() (*BulkAdGroupOperationResponse, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsAdGroupsResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given BulkAdGroupOperationResponse and assigns it to the AdGroups field.
func (o *CreateSponsoredBrandsAdGroupsResponseContent) SetAdGroups(v BulkAdGroupOperationResponse) {
	o.AdGroups = &v
}

func (o CreateSponsoredBrandsAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroups) {
		toSerialize["adGroups"] = o.AdGroups
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsAdGroupsResponseContent struct {
	value *CreateSponsoredBrandsAdGroupsResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsAdGroupsResponseContent) Get() *CreateSponsoredBrandsAdGroupsResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsAdGroupsResponseContent) Set(val *CreateSponsoredBrandsAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsAdGroupsResponseContent(val *CreateSponsoredBrandsAdGroupsResponseContent) *NullableCreateSponsoredBrandsAdGroupsResponseContent {
	return &NullableCreateSponsoredBrandsAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
