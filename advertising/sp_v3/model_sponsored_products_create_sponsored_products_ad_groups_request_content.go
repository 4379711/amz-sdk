package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent{}

// SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent struct for SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent
type SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent struct {
	// An array of adGroups.
	AdGroups []SponsoredProductsCreateAdGroup `json:"adGroups"`
}

type _SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent

// NewSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent instantiates a new SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent(adGroups []SponsoredProductsCreateAdGroup) *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent{}
	this.AdGroups = adGroups
	return &this
}

// NewSponsoredProductsCreateSponsoredProductsAdGroupsRequestContentWithDefaults instantiates a new SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateSponsoredProductsAdGroupsRequestContentWithDefaults() *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent {
	this := SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent{}
	return &this
}

// GetAdGroups returns the AdGroups field value
func (o *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) GetAdGroups() []SponsoredProductsCreateAdGroup {
	if o == nil {
		var ret []SponsoredProductsCreateAdGroup
		return ret
	}

	return o.AdGroups
}

// GetAdGroupsOk returns a tuple with the AdGroups field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) GetAdGroupsOk() ([]SponsoredProductsCreateAdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdGroups, true
}

// SetAdGroups sets field value
func (o *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) SetAdGroups(v []SponsoredProductsCreateAdGroup) {
	o.AdGroups = v
}

func (o SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adGroups"] = o.AdGroups
	return toSerialize, nil
}

type NullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent struct {
	value *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) Get() *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) Set(val *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent(val *SponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) *NullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent {
	return &NullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateSponsoredProductsAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
