package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsAdGroupsResponseContent{}

// DeleteSponsoredBrandsAdGroupsResponseContent struct for DeleteSponsoredBrandsAdGroupsResponseContent
type DeleteSponsoredBrandsAdGroupsResponseContent struct {
	AdGroups *BulkAdGroupOperationResponse `json:"adGroups,omitempty"`
}

// NewDeleteSponsoredBrandsAdGroupsResponseContent instantiates a new DeleteSponsoredBrandsAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsAdGroupsResponseContent() *DeleteSponsoredBrandsAdGroupsResponseContent {
	this := DeleteSponsoredBrandsAdGroupsResponseContent{}
	return &this
}

// NewDeleteSponsoredBrandsAdGroupsResponseContentWithDefaults instantiates a new DeleteSponsoredBrandsAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsAdGroupsResponseContentWithDefaults() *DeleteSponsoredBrandsAdGroupsResponseContent {
	this := DeleteSponsoredBrandsAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsAdGroupsResponseContent) GetAdGroups() BulkAdGroupOperationResponse {
	if o == nil || IsNil(o.AdGroups) {
		var ret BulkAdGroupOperationResponse
		return ret
	}
	return *o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsAdGroupsResponseContent) GetAdGroupsOk() (*BulkAdGroupOperationResponse, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsAdGroupsResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given BulkAdGroupOperationResponse and assigns it to the AdGroups field.
func (o *DeleteSponsoredBrandsAdGroupsResponseContent) SetAdGroups(v BulkAdGroupOperationResponse) {
	o.AdGroups = &v
}

func (o DeleteSponsoredBrandsAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroups) {
		toSerialize["adGroups"] = o.AdGroups
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsAdGroupsResponseContent struct {
	value *DeleteSponsoredBrandsAdGroupsResponseContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsAdGroupsResponseContent) Get() *DeleteSponsoredBrandsAdGroupsResponseContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsAdGroupsResponseContent) Set(val *DeleteSponsoredBrandsAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsAdGroupsResponseContent(val *DeleteSponsoredBrandsAdGroupsResponseContent) *NullableDeleteSponsoredBrandsAdGroupsResponseContent {
	return &NullableDeleteSponsoredBrandsAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
