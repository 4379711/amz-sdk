package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent struct for SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent
type SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent struct {
	DraftAdGroups SponsoredProductsBulkDraftAdGroupOperationResponse `json:"draftAdGroups"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent(draftAdGroups SponsoredProductsBulkDraftAdGroupOperationResponse) *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent{}
	this.DraftAdGroups = draftAdGroups
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent{}
	return &this
}

// GetDraftAdGroups returns the DraftAdGroups field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) GetDraftAdGroups() SponsoredProductsBulkDraftAdGroupOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftAdGroupOperationResponse
		return ret
	}

	return o.DraftAdGroups
}

// GetDraftAdGroupsOk returns a tuple with the DraftAdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) GetDraftAdGroupsOk() (*SponsoredProductsBulkDraftAdGroupOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DraftAdGroups, true
}

// SetDraftAdGroups sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) SetDraftAdGroups(v SponsoredProductsBulkDraftAdGroupOperationResponse) {
	o.DraftAdGroups = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["draftAdGroups"] = o.DraftAdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent(val *SponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
