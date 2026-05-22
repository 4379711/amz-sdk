package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsObjectIdFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsObjectIdFilter{}

// SponsoredProductsObjectIdFilter Filter entities by the list of objectIds
type SponsoredProductsObjectIdFilter struct {
	Include []string `json:"include"`
}

type _SponsoredProductsObjectIdFilter SponsoredProductsObjectIdFilter

// NewSponsoredProductsObjectIdFilter instantiates a new SponsoredProductsObjectIdFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsObjectIdFilter(include []string) *SponsoredProductsObjectIdFilter {
	this := SponsoredProductsObjectIdFilter{}
	this.Include = include
	return &this
}

// NewSponsoredProductsObjectIdFilterWithDefaults instantiates a new SponsoredProductsObjectIdFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsObjectIdFilterWithDefaults() *SponsoredProductsObjectIdFilter {
	this := SponsoredProductsObjectIdFilter{}
	return &this
}

// GetInclude returns the Include field value
func (o *SponsoredProductsObjectIdFilter) GetInclude() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsObjectIdFilter) GetIncludeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include, true
}

// SetInclude sets field value
func (o *SponsoredProductsObjectIdFilter) SetInclude(v []string) {
	o.Include = v
}

func (o SponsoredProductsObjectIdFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include"] = o.Include
	return toSerialize, nil
}

type NullableSponsoredProductsObjectIdFilter struct {
	value *SponsoredProductsObjectIdFilter
	isSet bool
}

func (v NullableSponsoredProductsObjectIdFilter) Get() *SponsoredProductsObjectIdFilter {
	return v.value
}

func (v *NullableSponsoredProductsObjectIdFilter) Set(val *SponsoredProductsObjectIdFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsObjectIdFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsObjectIdFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsObjectIdFilter(val *SponsoredProductsObjectIdFilter) *NullableSponsoredProductsObjectIdFilter {
	return &NullableSponsoredProductsObjectIdFilter{value: val, isSet: true}
}

func (v NullableSponsoredProductsObjectIdFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsObjectIdFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
