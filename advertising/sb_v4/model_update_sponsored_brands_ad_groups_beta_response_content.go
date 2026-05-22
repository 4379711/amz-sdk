package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSponsoredBrandsAdGroupsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSponsoredBrandsAdGroupsBetaResponseContent{}

// UpdateSponsoredBrandsAdGroupsBetaResponseContent struct for UpdateSponsoredBrandsAdGroupsBetaResponseContent
type UpdateSponsoredBrandsAdGroupsBetaResponseContent struct {
	AdGroups *BulkAdGroupOperationResponse `json:"adGroups,omitempty"`
}

// NewUpdateSponsoredBrandsAdGroupsBetaResponseContent instantiates a new UpdateSponsoredBrandsAdGroupsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSponsoredBrandsAdGroupsBetaResponseContent() *UpdateSponsoredBrandsAdGroupsBetaResponseContent {
	this := UpdateSponsoredBrandsAdGroupsBetaResponseContent{}
	return &this
}

// NewUpdateSponsoredBrandsAdGroupsBetaResponseContentWithDefaults instantiates a new UpdateSponsoredBrandsAdGroupsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSponsoredBrandsAdGroupsBetaResponseContentWithDefaults() *UpdateSponsoredBrandsAdGroupsBetaResponseContent {
	this := UpdateSponsoredBrandsAdGroupsBetaResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *UpdateSponsoredBrandsAdGroupsBetaResponseContent) GetAdGroups() BulkAdGroupOperationResponse {
	if o == nil || IsNil(o.AdGroups) {
		var ret BulkAdGroupOperationResponse
		return ret
	}
	return *o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSponsoredBrandsAdGroupsBetaResponseContent) GetAdGroupsOk() (*BulkAdGroupOperationResponse, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *UpdateSponsoredBrandsAdGroupsBetaResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given BulkAdGroupOperationResponse and assigns it to the AdGroups field.
func (o *UpdateSponsoredBrandsAdGroupsBetaResponseContent) SetAdGroups(v BulkAdGroupOperationResponse) {
	o.AdGroups = &v
}

func (o UpdateSponsoredBrandsAdGroupsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroups) {
		toSerialize["adGroups"] = o.AdGroups
	}
	return toSerialize, nil
}

type NullableUpdateSponsoredBrandsAdGroupsBetaResponseContent struct {
	value *UpdateSponsoredBrandsAdGroupsBetaResponseContent
	isSet bool
}

func (v NullableUpdateSponsoredBrandsAdGroupsBetaResponseContent) Get() *UpdateSponsoredBrandsAdGroupsBetaResponseContent {
	return v.value
}

func (v *NullableUpdateSponsoredBrandsAdGroupsBetaResponseContent) Set(val *UpdateSponsoredBrandsAdGroupsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSponsoredBrandsAdGroupsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSponsoredBrandsAdGroupsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSponsoredBrandsAdGroupsBetaResponseContent(val *UpdateSponsoredBrandsAdGroupsBetaResponseContent) *NullableUpdateSponsoredBrandsAdGroupsBetaResponseContent {
	return &NullableUpdateSponsoredBrandsAdGroupsBetaResponseContent{value: val, isSet: true}
}

func (v NullableUpdateSponsoredBrandsAdGroupsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSponsoredBrandsAdGroupsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
