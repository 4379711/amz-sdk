package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent struct for SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent
type SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent struct {
	AdGroups SponsoredProductsBulkAdGroupOperationResponse `json:"adGroups"`
}

type _SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent(adGroups SponsoredProductsBulkAdGroupOperationResponse) *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) GetAdGroups() SponsoredProductsBulkAdGroupOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkAdGroupOperationResponse
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) GetAdGroupsOk() (*SponsoredProductsBulkAdGroupOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) SetAdGroups(v SponsoredProductsBulkAdGroupOperationResponse) {
	o.AdGroups = v
}

func (o SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent(val *SponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
