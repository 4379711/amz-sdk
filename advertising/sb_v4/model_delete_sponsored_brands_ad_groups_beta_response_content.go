package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsAdGroupsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsAdGroupsBetaResponseContent{}

// DeleteSponsoredBrandsAdGroupsBetaResponseContent struct for DeleteSponsoredBrandsAdGroupsBetaResponseContent
type DeleteSponsoredBrandsAdGroupsBetaResponseContent struct {
	AdGroups *BulkAdGroupOperationResponse `json:"adGroups,omitempty"`
}

// NewDeleteSponsoredBrandsAdGroupsBetaResponseContent instantiates a new DeleteSponsoredBrandsAdGroupsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsAdGroupsBetaResponseContent() *DeleteSponsoredBrandsAdGroupsBetaResponseContent {
	this := DeleteSponsoredBrandsAdGroupsBetaResponseContent{}
	return &this
}

// NewDeleteSponsoredBrandsAdGroupsBetaResponseContentWithDefaults instantiates a new DeleteSponsoredBrandsAdGroupsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsAdGroupsBetaResponseContentWithDefaults() *DeleteSponsoredBrandsAdGroupsBetaResponseContent {
	this := DeleteSponsoredBrandsAdGroupsBetaResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsAdGroupsBetaResponseContent) GetAdGroups() BulkAdGroupOperationResponse {
	if o == nil || IsNil(o.AdGroups) {
		var ret BulkAdGroupOperationResponse
		return ret
	}
	return *o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsAdGroupsBetaResponseContent) GetAdGroupsOk() (*BulkAdGroupOperationResponse, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsAdGroupsBetaResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given BulkAdGroupOperationResponse and assigns it to the AdGroups field.
func (o *DeleteSponsoredBrandsAdGroupsBetaResponseContent) SetAdGroups(v BulkAdGroupOperationResponse) {
	o.AdGroups = &v
}

func (o DeleteSponsoredBrandsAdGroupsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroups) {
		toSerialize["adGroups"] = o.AdGroups
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsAdGroupsBetaResponseContent struct {
	value *DeleteSponsoredBrandsAdGroupsBetaResponseContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsAdGroupsBetaResponseContent) Get() *DeleteSponsoredBrandsAdGroupsBetaResponseContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsAdGroupsBetaResponseContent) Set(val *DeleteSponsoredBrandsAdGroupsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsAdGroupsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsAdGroupsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsAdGroupsBetaResponseContent(val *DeleteSponsoredBrandsAdGroupsBetaResponseContent) *NullableDeleteSponsoredBrandsAdGroupsBetaResponseContent {
	return &NullableDeleteSponsoredBrandsAdGroupsBetaResponseContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsAdGroupsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsAdGroupsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
