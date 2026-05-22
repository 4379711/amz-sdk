package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent struct for SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent
type SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent struct {
	AdGroups SponsoredProductsBulkGlobalAdGroupOperationResponse `json:"adGroups"`
}

type _SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent

// NewSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent(adGroups SponsoredProductsBulkGlobalAdGroupOperationResponse) *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) GetAdGroups() SponsoredProductsBulkGlobalAdGroupOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalAdGroupOperationResponse
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) GetAdGroupsOk() (*SponsoredProductsBulkGlobalAdGroupOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) SetAdGroups(v SponsoredProductsBulkGlobalAdGroupOperationResponse) {
	o.AdGroups = v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent(val *SponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
