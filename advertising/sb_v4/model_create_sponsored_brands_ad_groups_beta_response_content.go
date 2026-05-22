package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateSponsoredBrandsAdGroupsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSponsoredBrandsAdGroupsBetaResponseContent{}

// CreateSponsoredBrandsAdGroupsBetaResponseContent struct for CreateSponsoredBrandsAdGroupsBetaResponseContent
type CreateSponsoredBrandsAdGroupsBetaResponseContent struct {
	AdGroups *BulkAdGroupOperationResponse `json:"adGroups,omitempty"`
}

// NewCreateSponsoredBrandsAdGroupsBetaResponseContent instantiates a new CreateSponsoredBrandsAdGroupsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSponsoredBrandsAdGroupsBetaResponseContent() *CreateSponsoredBrandsAdGroupsBetaResponseContent {
	this := CreateSponsoredBrandsAdGroupsBetaResponseContent{}
	return &this
}

// NewCreateSponsoredBrandsAdGroupsBetaResponseContentWithDefaults instantiates a new CreateSponsoredBrandsAdGroupsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSponsoredBrandsAdGroupsBetaResponseContentWithDefaults() *CreateSponsoredBrandsAdGroupsBetaResponseContent {
	this := CreateSponsoredBrandsAdGroupsBetaResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value if set, zero value otherwise.
func (o *CreateSponsoredBrandsAdGroupsBetaResponseContent) GetAdGroups() BulkAdGroupOperationResponse {
	if o == nil || IsNil(o.AdGroups) {
		var ret BulkAdGroupOperationResponse
		return ret
	}
	return *o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSponsoredBrandsAdGroupsBetaResponseContent) GetAdGroupsOk() (*BulkAdGroupOperationResponse, bool) {
	if o == nil || IsNil(o.AdGroups) {
		return nil, false
	}
	return o.AdGroups, true
}

// HasAdGroups returns a boolean if a field has been set.
func (o *CreateSponsoredBrandsAdGroupsBetaResponseContent) HasAdGroups() bool {
	if o != nil && !IsNil(o.AdGroups) {
		return true
	}

	return false
}

// SetAdGroups gets a reference to the given BulkAdGroupOperationResponse and assigns it to the AdGroups field.
func (o *CreateSponsoredBrandsAdGroupsBetaResponseContent) SetAdGroups(v BulkAdGroupOperationResponse) {
	o.AdGroups = &v
}

func (o CreateSponsoredBrandsAdGroupsBetaResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroups) {
		toSerialize["adGroups"] = o.AdGroups
	}
	return toSerialize, nil
}

type NullableCreateSponsoredBrandsAdGroupsBetaResponseContent struct {
	value *CreateSponsoredBrandsAdGroupsBetaResponseContent
	isSet bool
}

func (v NullableCreateSponsoredBrandsAdGroupsBetaResponseContent) Get() *CreateSponsoredBrandsAdGroupsBetaResponseContent {
	return v.value
}

func (v *NullableCreateSponsoredBrandsAdGroupsBetaResponseContent) Set(val *CreateSponsoredBrandsAdGroupsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSponsoredBrandsAdGroupsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSponsoredBrandsAdGroupsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSponsoredBrandsAdGroupsBetaResponseContent(val *CreateSponsoredBrandsAdGroupsBetaResponseContent) *NullableCreateSponsoredBrandsAdGroupsBetaResponseContent {
	return &NullableCreateSponsoredBrandsAdGroupsBetaResponseContent{value: val, isSet: true}
}

func (v NullableCreateSponsoredBrandsAdGroupsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateSponsoredBrandsAdGroupsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
