package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct for SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent
type SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct {
	NegativeTargetIdFilter *SponsoredProductsObjectIdFilter `json:"negativeTargetIdFilter,omitempty"`
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent{}
	return &this
}

// GetNegativeTargetIdFilter returns the NegativeTargetIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.NegativeTargetIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.NegativeTargetIdFilter
}

// GetNegativeTargetIdFilterOk returns a tuple with the NegativeTargetIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) GetNegativeTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.NegativeTargetIdFilter) {
		return nil, false
	}
	return o.NegativeTargetIdFilter, true
}

// HasNegativeTargetIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) HasNegativeTargetIdFilter() bool {
	if o != nil && !IsNil(o.NegativeTargetIdFilter) {
		return true
	}

	return false
}

// SetNegativeTargetIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the NegativeTargetIdFilter field.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) SetNegativeTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.NegativeTargetIdFilter = &v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NegativeTargetIdFilter) {
		toSerialize["negativeTargetIdFilter"] = o.NegativeTargetIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent(val *SponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalNegativeTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
