package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent{}

// SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent struct for SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent
type SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent struct {
	TargetIdFilter SponsoredProductsObjectIdFilter `json:"targetIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent

// NewSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent(targetIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent{}
	this.TargetIdFilter = targetIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent{}
	return &this
}

// GetTargetIdFilter returns the TargetIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) GetTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.TargetIdFilter
}

// GetTargetIdFilterOk returns a tuple with the TargetIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) GetTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetIdFilter, true
}

// SetTargetIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) SetTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetIdFilter"] = o.TargetIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent(val *SponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
