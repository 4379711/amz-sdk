package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNameFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNameFilter{}

// SponsoredProductsNameFilter Filter entities by name
type SponsoredProductsNameFilter struct {
	QueryTermMatchType *SponsoredProductsQueryTermMatchType `json:"queryTermMatchType,omitempty"`
	Include            []string                             `json:"include,omitempty"`
}

// NewSponsoredProductsNameFilter instantiates a new SponsoredProductsNameFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNameFilter() *SponsoredProductsNameFilter {
	this := SponsoredProductsNameFilter{}
	return &this
}

// NewSponsoredProductsNameFilterWithDefaults instantiates a new SponsoredProductsNameFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNameFilterWithDefaults() *SponsoredProductsNameFilter {
	this := SponsoredProductsNameFilter{}
	return &this
}

// GetQueryTermMatchType returns the QueryTermMatchType field value if set, zero value otherwise.
func (o *SponsoredProductsNameFilter) GetQueryTermMatchType() SponsoredProductsQueryTermMatchType {
	if o == nil || IsNil(o.QueryTermMatchType) {
		var ret SponsoredProductsQueryTermMatchType
		return ret
	}
	return *o.QueryTermMatchType
}

// GetQueryTermMatchTypeOk returns a tuple with the QueryTermMatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNameFilter) GetQueryTermMatchTypeOk() (*SponsoredProductsQueryTermMatchType, bool) {
	if o == nil || IsNil(o.QueryTermMatchType) {
		return nil, false
	}
	return o.QueryTermMatchType, true
}

// HasQueryTermMatchType returns a boolean if a field has been set.
func (o *SponsoredProductsNameFilter) HasQueryTermMatchType() bool {
	if o != nil && !IsNil(o.QueryTermMatchType) {
		return true
	}

	return false
}

// SetQueryTermMatchType gets a reference to the given SponsoredProductsQueryTermMatchType and assigns it to the QueryTermMatchType field.
func (o *SponsoredProductsNameFilter) SetQueryTermMatchType(v SponsoredProductsQueryTermMatchType) {
	o.QueryTermMatchType = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *SponsoredProductsNameFilter) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNameFilter) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *SponsoredProductsNameFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *SponsoredProductsNameFilter) SetInclude(v []string) {
	o.Include = v
}

func (o SponsoredProductsNameFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueryTermMatchType) {
		toSerialize["queryTermMatchType"] = o.QueryTermMatchType
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNameFilter struct {
	value *SponsoredProductsNameFilter
	isSet bool
}

func (v NullableSponsoredProductsNameFilter) Get() *SponsoredProductsNameFilter {
	return v.value
}

func (v *NullableSponsoredProductsNameFilter) Set(val *SponsoredProductsNameFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNameFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNameFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNameFilter(val *SponsoredProductsNameFilter) *NullableSponsoredProductsNameFilter {
	return &NullableSponsoredProductsNameFilter{value: val, isSet: true}
}

func (v NullableSponsoredProductsNameFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNameFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
