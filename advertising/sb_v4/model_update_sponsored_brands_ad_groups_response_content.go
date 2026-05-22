package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsAdGroupsResponseContent{}

// UpdateSponsoredBrandsAdGroupsResponseContent struct for UpdateSponsoredBrandsAdGroupsResponseContent
type UpdateSponsoredBrandsAdGroupsResponseContent struct {
	AdGroups *BulkAdGroupOperationResponse `json:"adGroups,omitempty"`
}

// NewUpdateSponsoredBrandsAdGroupsResponseContent instantiates a new UpdateSponsoredBrandsAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsAdGroupsResponseContent() *UpdateSponsoredBrandsAdGroupsResponseContent {
	this := UpdateSponsoredBrandsAdGroupsResponseContent{}
	return &this
}

// NewUpdateSponsoredBrandsAdGroupsResponseContentWithDefaults instantiates a new UpdateSponsoredBrandsAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsAdGroupsResponseContentWithDefaults() *UpdateSponsoredBrandsAdGroupsResponseContent {
	this := UpdateSponsoredBrandsAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *UpdateSponsoredBrandsAdGroupsResponseContent) GetAdGroups() BulkAdGroupOperationResponse {
	if o == nil || IsNil(o.AdGroups) {
		var ret BulkAdGroupOperationResponse
		return ret
	}
	return *o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsAdGroupsResponseContent) GetAdGroupsOk() (*BulkAdGroupOperationResponse, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *UpdateSponsoredBrandsAdGroupsResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given BulkAdGroupOperationResponse and assigns it to the AdGroups field.
func (o *UpdateSponsoredBrandsAdGroupsResponseContent) SetAdGroups(v BulkAdGroupOperationResponse) {
	o.AdGroups = &v
}

func (o UpdateSponsoredBrandsAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroups) {
		toSerialize["adGroups"] = o.AdGroups
	}
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsAdGroupsResponseContent struct {
	value *UpdateSponsoredBrandsAdGroupsResponseContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsAdGroupsResponseContent) Get() *UpdateSponsoredBrandsAdGroupsResponseContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsAdGroupsResponseContent) Set(val *UpdateSponsoredBrandsAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsAdGroupsResponseContent(val *UpdateSponsoredBrandsAdGroupsResponseContent) *NullableUpdateSponsoredBrandsAdGroupsResponseContent {
	return &NullableUpdateSponsoredBrandsAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
