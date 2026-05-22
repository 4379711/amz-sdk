package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteAllSPTargetsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteAllSPTargetsRequestContent{}

// SponsoredProductsDeleteAllSPTargetsRequestContent struct for SponsoredProductsDeleteAllSPTargetsRequestContent
type SponsoredProductsDeleteAllSPTargetsRequestContent struct {
	TargetIdFilter SponsoredProductsObjectIdFilter `json:"targetIdFilter"`
}

type _SponsoredProductsDeleteAllSPTargetsRequestContent SponsoredProductsDeleteAllSPTargetsRequestContent

// NewSponsoredProductsDeleteAllSPTargetsRequestContent instantiates a new SponsoredProductsDeleteAllSPTargetsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteAllSPTargetsRequestContent(targetIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsDeleteAllSPTargetsRequestContent {
	this := SponsoredProductsDeleteAllSPTargetsRequestContent{}
	this.TargetIdFilter = targetIdFilter
	return &this
}

// NewSponsoredProductsDeleteAllSPTargetsRequestContentWithDefaults instantiates a new SponsoredProductsDeleteAllSPTargetsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteAllSPTargetsRequestContentWithDefaults() *SponsoredProductsDeleteAllSPTargetsRequestContent {
	this := SponsoredProductsDeleteAllSPTargetsRequestContent{}
	return &this
}

// GetTargetIdFilter returns the TargetIdFilter field value
func (o *SponsoredProductsDeleteAllSPTargetsRequestContent) GetTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.TargetIdFilter
}

// GetTargetIdFilterOk returns a tuple with the TargetIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteAllSPTargetsRequestContent) GetTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetIdFilter, true
}

// SetTargetIdFilter sets field value
func (o *SponsoredProductsDeleteAllSPTargetsRequestContent) SetTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetIdFilter = v
}

func (o SponsoredProductsDeleteAllSPTargetsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetIdFilter"] = o.TargetIdFilter
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteAllSPTargetsRequestContent struct {
	value *SponsoredProductsDeleteAllSPTargetsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteAllSPTargetsRequestContent) Get() *SponsoredProductsDeleteAllSPTargetsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteAllSPTargetsRequestContent) Set(val *SponsoredProductsDeleteAllSPTargetsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteAllSPTargetsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteAllSPTargetsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteAllSPTargetsRequestContent(val *SponsoredProductsDeleteAllSPTargetsRequestContent) *NullableSponsoredProductsDeleteAllSPTargetsRequestContent {
	return &NullableSponsoredProductsDeleteAllSPTargetsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteAllSPTargetsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteAllSPTargetsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
