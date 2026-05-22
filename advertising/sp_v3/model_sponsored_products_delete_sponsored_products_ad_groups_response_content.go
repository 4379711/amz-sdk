package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent struct for SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent
type SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent struct {
	AdGroups SponsoredProductsBulkAdGroupOperationResponse `json:"adGroups"`
}

type _SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent(adGroups SponsoredProductsBulkAdGroupOperationResponse) *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) GetAdGroups() SponsoredProductsBulkAdGroupOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkAdGroupOperationResponse
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) GetAdGroupsOk() (*SponsoredProductsBulkAdGroupOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) SetAdGroups(v SponsoredProductsBulkAdGroupOperationResponse) {
	o.AdGroups = v
}

func (o SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent(val *SponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
