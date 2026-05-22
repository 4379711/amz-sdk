package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAsinFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAsinFilter{}

// SponsoredProductsAsinFilter struct for SponsoredProductsAsinFilter
type SponsoredProductsAsinFilter struct {
	QueryTermMatchType *SponsoredProductsQueryTermMatchType `json:"queryTermMatchType,omitempty"`
	Include            []string                             `json:"include,omitempty"`
}

// NewSponsoredProductsAsinFilter instantiates a new SponsoredProductsAsinFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAsinFilter() *SponsoredProductsAsinFilter {
	this := SponsoredProductsAsinFilter{}
	return &this
}

// NewSponsoredProductsAsinFilterWithDefaults instantiates a new SponsoredProductsAsinFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAsinFilterWithDefaults() *SponsoredProductsAsinFilter {
	this := SponsoredProductsAsinFilter{}
	return &this
}

// GetQueryTermMatchType returns the QueryTermMatchType field value if set, zero value otherwise.
func (o *SponsoredProductsAsinFilter) GetQueryTermMatchType() SponsoredProductsQueryTermMatchType {
	if o == nil || IsNil(o.QueryTermMatchType) {
		var ret SponsoredProductsQueryTermMatchType
		return ret
	}
	return *o.QueryTermMatchType
}

// GetQueryTermMatchTypeOk returns a tuple with the QueryTermMatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAsinFilter) GetQueryTermMatchTypeOk() (*SponsoredProductsQueryTermMatchType, bool) {
	if o == nil || IsNil(o.QueryTermMatchType) {
		return nil, false
	}
	return o.QueryTermMatchType, true
}

// HasQueryTermMatchType returns a boolean if a field has been set.
func (o *SponsoredProductsAsinFilter) HasQueryTermMatchType() bool {
	if o != nil && !IsNil(o.QueryTermMatchType) {
		return true
	}

	return false
}

// SetQueryTermMatchType gets a reference to the given SponsoredProductsQueryTermMatchType and assigns it to the QueryTermMatchType field.
func (o *SponsoredProductsAsinFilter) SetQueryTermMatchType(v SponsoredProductsQueryTermMatchType) {
	o.QueryTermMatchType = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *SponsoredProductsAsinFilter) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAsinFilter) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *SponsoredProductsAsinFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *SponsoredProductsAsinFilter) SetInclude(v []string) {
	o.Include = v
}

func (o SponsoredProductsAsinFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueryTermMatchType) {
		toSerialize["queryTermMatchType"] = o.QueryTermMatchType
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableSponsoredProductsAsinFilter struct {
	value *SponsoredProductsAsinFilter
	isSet bool
}

func (v NullableSponsoredProductsAsinFilter) Get() *SponsoredProductsAsinFilter {
	return v.value
}

func (v *NullableSponsoredProductsAsinFilter) Set(val *SponsoredProductsAsinFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAsinFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAsinFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAsinFilter(val *SponsoredProductsAsinFilter) *NullableSponsoredProductsAsinFilter {
	return &NullableSponsoredProductsAsinFilter{value: val, isSet: true}
}

func (v NullableSponsoredProductsAsinFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAsinFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
