package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent{}

// SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent struct for SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent
type SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent struct {
	NegativeTargetIdFilter SponsoredProductsObjectIdFilter `json:"negativeTargetIdFilter"`
}

type _SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent

// NewSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent(negativeTargetIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent{}
	this.NegativeTargetIdFilter = negativeTargetIdFilter
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetIdFilter returns the NegativeTargetIdFilter field value
func (o *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.NegativeTargetIdFilter
}

// GetNegativeTargetIdFilterOk returns a tuple with the NegativeTargetIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NegativeTargetIdFilter, true
}

// SetNegativeTargetIdFilter sets field value
func (o *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) SetNegativeTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeTargetIdFilter = v
}

func (o SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["negativeTargetIdFilter"] = o.NegativeTargetIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent(val *SponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
