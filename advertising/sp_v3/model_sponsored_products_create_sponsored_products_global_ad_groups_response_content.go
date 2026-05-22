package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent{}

// SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent struct for SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent
type SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent struct {
	AdGroups SponsoredProductsBulkGlobalAdGroupOperationResponse `json:"adGroups"`
}

type _SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent

// NewSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent instantiates a new SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent(adGroups SponsoredProductsBulkGlobalAdGroupOperationResponse) *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContentWithDefaults() *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent {
	this := SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) GetAdGroups() SponsoredProductsBulkGlobalAdGroupOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalAdGroupOperationResponse
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) GetAdGroupsOk() (*SponsoredProductsBulkGlobalAdGroupOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) SetAdGroups(v SponsoredProductsBulkGlobalAdGroupOperationResponse) {
	o.AdGroups = v
}

func (o SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent struct {
	value *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) Get() *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) Set(val *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent(val *SponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) *NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent {
	return &NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsGlobalAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
