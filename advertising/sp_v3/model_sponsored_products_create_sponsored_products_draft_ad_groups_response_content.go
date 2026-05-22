package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent{}

// SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent struct for SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent
type SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent struct {
	AdGroups SponsoredProductsBulkDraftAdGroupOperationResponse `json:"adGroups"`
}

type _SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent

// NewSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent(adGroups SponsoredProductsBulkDraftAdGroupOperationResponse) *SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) GetAdGroups() SponsoredProductsBulkDraftAdGroupOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkDraftAdGroupOperationResponse
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) GetAdGroupsOk() (*SponsoredProductsBulkDraftAdGroupOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) SetAdGroups(v SponsoredProductsBulkDraftAdGroupOperationResponse) {
	o.AdGroups = v
}

func (o SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent(val *SponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsDraftAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
