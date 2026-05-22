package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsEntityStateFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsEntityStateFilter{}

// SponsoredProductsEntityStateFilter Filter entities by state. To filter live entities, only 'ENABLED', 'PAUSED' and 'ARCHIVED' can be used
type SponsoredProductsEntityStateFilter struct {
	Include []SponsoredProductsEntityState `json:"include"`
}

type _SponsoredProductsEntityStateFilter SponsoredProductsEntityStateFilter

// NewSponsoredProductsEntityStateFilter instantiates a new SponsoredProductsEntityStateFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsEntityStateFilter(include []SponsoredProductsEntityState) *SponsoredProductsEntityStateFilter {
	this := SponsoredProductsEntityStateFilter{}
	this.Include = include
	return &this
}

// NewSponsoredProductsEntityStateFilterWithDefaults instantiates a new SponsoredProductsEntityStateFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsEntityStateFilterWithDefaults() *SponsoredProductsEntityStateFilter {
	this := SponsoredProductsEntityStateFilter{}
	return &this
}

// GetInclude returns the Include field value
func (o *SponsoredProductsEntityStateFilter) GetInclude() []SponsoredProductsEntityState {
	if o == nil {
		var ret []SponsoredProductsEntityState
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsEntityStateFilter) GetIncludeOk() ([]SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include, true
}

// SetInclude sets field value
func (o *SponsoredProductsEntityStateFilter) SetInclude(v []SponsoredProductsEntityState) {
	o.Include = v
}

func (o SponsoredProductsEntityStateFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include"] = o.Include
	return toSerialize, nil
}

type NullableSponsoredProductsEntityStateFilter struct {
	value *SponsoredProductsEntityStateFilter
	isSet bool
}

func (v NullableSponsoredProductsEntityStateFilter) Get() *SponsoredProductsEntityStateFilter {
	return v.value
}

func (v *NullableSponsoredProductsEntityStateFilter) Set(val *SponsoredProductsEntityStateFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsEntityStateFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsEntityStateFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsEntityStateFilter(val *SponsoredProductsEntityStateFilter) *NullableSponsoredProductsEntityStateFilter {
	return &NullableSponsoredProductsEntityStateFilter{value: val, isSet: true}
}

func (v NullableSponsoredProductsEntityStateFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsEntityStateFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
