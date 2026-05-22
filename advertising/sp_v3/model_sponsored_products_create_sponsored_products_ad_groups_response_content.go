package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent{}

// SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent struct for SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent
type SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent struct {
	AdGroups SponsoredProductsBulkAdGroupOperationResponse `json:"adGroups"`
}

type _SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent

// NewSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent(adGroups SponsoredProductsBulkAdGroupOperationResponse) *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsAdGroupsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) GetAdGroups() SponsoredProductsBulkAdGroupOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkAdGroupOperationResponse
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) GetAdGroupsOk() (*SponsoredProductsBulkAdGroupOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) SetAdGroups(v SponsoredProductsBulkAdGroupOperationResponse) {
	o.AdGroups = v
}

func (o SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent(val *SponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
