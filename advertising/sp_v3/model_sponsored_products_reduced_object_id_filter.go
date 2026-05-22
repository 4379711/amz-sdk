package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsReducedObjectIdFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsReducedObjectIdFilter{}

// SponsoredProductsReducedObjectIdFilter Filter entities by the list of objectIds
type SponsoredProductsReducedObjectIdFilter struct {
	Include []string `json:"include"`
}

type _SponsoredProductsReducedObjectIdFilter SponsoredProductsReducedObjectIdFilter

// NewSponsoredProductsReducedObjectIdFilter instantiates a new SponsoredProductsReducedObjectIdFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsReducedObjectIdFilter(include []string) *SponsoredProductsReducedObjectIdFilter {
	this := SponsoredProductsReducedObjectIdFilter{}
	this.Include = include
	return &this
}

// NewSponsoredProductsReducedObjectIdFilterWithDefaults instantiates a new SponsoredProductsReducedObjectIdFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsReducedObjectIdFilterWithDefaults() *SponsoredProductsReducedObjectIdFilter {
	this := SponsoredProductsReducedObjectIdFilter{}
	return &this
}

// GetInclude returns the Include field value
func (o *SponsoredProductsReducedObjectIdFilter) GetInclude() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsReducedObjectIdFilter) GetIncludeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include, true
}

// SetInclude sets field value
func (o *SponsoredProductsReducedObjectIdFilter) SetInclude(v []string) {
	o.Include = v
}

func (o SponsoredProductsReducedObjectIdFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include"] = o.Include
	return toSerialize, nil
}

type NullableSponsoredProductsReducedObjectIdFilter struct {
	value *SponsoredProductsReducedObjectIdFilter
	isSet bool
}

func (v NullableSponsoredProductsReducedObjectIdFilter) Get() *SponsoredProductsReducedObjectIdFilter {
	return v.value
}

func (v *NullableSponsoredProductsReducedObjectIdFilter) Set(val *SponsoredProductsReducedObjectIdFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsReducedObjectIdFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsReducedObjectIdFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsReducedObjectIdFilter(val *SponsoredProductsReducedObjectIdFilter) *NullableSponsoredProductsReducedObjectIdFilter {
	return &NullableSponsoredProductsReducedObjectIdFilter{value: val, isSet: true}
}

func (v NullableSponsoredProductsReducedObjectIdFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsReducedObjectIdFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
