package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent{}

// SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent struct for SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent
type SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent struct {
	AdGroups SponsoredProductsBulkGlobalAdGroupOperationResponse `json:"adGroups"`
}

type _SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent

// NewSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent(adGroups SponsoredProductsBulkGlobalAdGroupOperationResponse) *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContentWithDefaults instantiates a new SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContentWithDefaults() *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent {
	this := SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) GetAdGroups() SponsoredProductsBulkGlobalAdGroupOperationResponse {
	if o == nil {
		var ret SponsoredProductsBulkGlobalAdGroupOperationResponse
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) GetAdGroupsOk() (*SponsoredProductsBulkGlobalAdGroupOperationResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) SetAdGroups(v SponsoredProductsBulkGlobalAdGroupOperationResponse) {
	o.AdGroups = v
}

func (o SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent struct {
	value *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) Get() *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) Set(val *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent(val *SponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) *NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent {
	return &NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsUpdateSponsoredProductsGlobalAdGroupsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
