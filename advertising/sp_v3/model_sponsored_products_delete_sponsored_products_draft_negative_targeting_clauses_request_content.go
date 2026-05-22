package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent{}

// SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent struct for SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent
type SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent struct {
	NegativeTargetIdFilter SponsoredProductsObjectIdFilter `json:"negativeTargetIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent

// NewSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent(negativeTargetIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent{}
	this.NegativeTargetIdFilter = negativeTargetIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetIdFilter returns the NegativeTargetIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.NegativeTargetIdFilter
}

// GetNegativeTargetIdFilterOk returns a tuple with the NegativeTargetIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetIdFilter, true
}

// SetNegativeTargetIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) SetNegativeTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeTargetIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetIdFilter"] = o.NegativeTargetIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent(val *SponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsDraftNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
