package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent{}

// SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent struct for SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent
type SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent struct {
	TargetIdFilter SponsoredProductsObjectIdFilter `json:"targetIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent

// NewSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent(targetIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent{}
	this.TargetIdFilter = targetIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent{}
	return &this
}

// GetTargetIdFilter returns the TargetIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) GetTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.TargetIdFilter
}

// GetTargetIdFilterOk returns a tuple with the TargetIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) GetTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetIdFilter, true
}

// SetTargetIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) SetTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetIdFilter"] = o.TargetIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent(val *SponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
