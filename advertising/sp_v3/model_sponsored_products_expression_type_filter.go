package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsExpressionTypeFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsExpressionTypeFilter{}

// SponsoredProductsExpressionTypeFilter Filter entities by ExpressionType
type SponsoredProductsExpressionTypeFilter struct {
	Include []SponsoredProductsExpressionType `json:"include"`
}

type _SponsoredProductsExpressionTypeFilter SponsoredProductsExpressionTypeFilter

// NewSponsoredProductsExpressionTypeFilter instantiates a new SponsoredProductsExpressionTypeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsExpressionTypeFilter(include []SponsoredProductsExpressionType) *SponsoredProductsExpressionTypeFilter {
	this := SponsoredProductsExpressionTypeFilter{}
	this.Include = include
	return &this
}

// NewSponsoredProductsExpressionTypeFilterWithDefaults instantiates a new SponsoredProductsExpressionTypeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsExpressionTypeFilterWithDefaults() *SponsoredProductsExpressionTypeFilter {
	this := SponsoredProductsExpressionTypeFilter{}
	return &this
}

// GetInclude returns the Include field value
func (o *SponsoredProductsExpressionTypeFilter) GetInclude() []SponsoredProductsExpressionType {
	if o == nil {
		var ret []SponsoredProductsExpressionType
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsExpressionTypeFilter) GetIncludeOk() ([]SponsoredProductsExpressionType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Include, true
}

// SetInclude sets field value
func (o *SponsoredProductsExpressionTypeFilter) SetInclude(v []SponsoredProductsExpressionType) {
	o.Include = v
}

func (o SponsoredProductsExpressionTypeFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include"] = o.Include
	return toSerialize, nil
}

type NullableSponsoredProductsExpressionTypeFilter struct {
	value *SponsoredProductsExpressionTypeFilter
	isSet bool
}

func (v NullableSponsoredProductsExpressionTypeFilter) Get() *SponsoredProductsExpressionTypeFilter {
	return v.value
}

func (v *NullableSponsoredProductsExpressionTypeFilter) Set(val *SponsoredProductsExpressionTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsExpressionTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsExpressionTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsExpressionTypeFilter(val *SponsoredProductsExpressionTypeFilter) *NullableSponsoredProductsExpressionTypeFilter {
	return &NullableSponsoredProductsExpressionTypeFilter{value: val, isSet: true}
}

func (v NullableSponsoredProductsExpressionTypeFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsExpressionTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
