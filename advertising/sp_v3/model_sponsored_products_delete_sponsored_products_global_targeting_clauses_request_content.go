package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent{}

// SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent struct for SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent
type SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent struct {
	TargetIdFilter *SponsoredProductsObjectIdFilter `json:"targetIdFilter,omitempty"`
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent() *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent{}
	return &this
}

// NewSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContentWithDefaults() *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent {
	this := SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent{}
	return &this
}

// GetTargetIdFilter returns the TargetIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) GetTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.TargetIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.TargetIdFilter
}

// GetTargetIdFilterOk returns a tuple with the TargetIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) GetTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.TargetIdFilter) {
		return nil, false
	}
	return o.TargetIdFilter, true
}

// HasTargetIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) HasTargetIdFilter() bool {
	if o != nil && !IsNil(o.TargetIdFilter) {
		return true
	}

	return false
}

// SetTargetIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the TargetIdFilter field.
func (o *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) SetTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetIdFilter = &v
}

func (o SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetIdFilter) {
		toSerialize["targetIdFilter"] = o.TargetIdFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent struct {
	value *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) Get() *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) Set(val *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent(val *SponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) *NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent {
	return &NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteSponsoredProductsGlobalTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
