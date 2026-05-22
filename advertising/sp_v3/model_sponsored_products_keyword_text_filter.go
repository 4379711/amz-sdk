package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsKeywordTextFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsKeywordTextFilter{}

// SponsoredProductsKeywordTextFilter Filter by keywordText
type SponsoredProductsKeywordTextFilter struct {
	QueryTermMatchType *SponsoredProductsQueryTermMatchType `json:"queryTermMatchType,omitempty"`
	Include            []string                             `json:"include,omitempty"`
}

// NewSponsoredProductsKeywordTextFilter instantiates a new SponsoredProductsKeywordTextFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsKeywordTextFilter() *SponsoredProductsKeywordTextFilter {
	this := SponsoredProductsKeywordTextFilter{}
	return &this
}

// NewSponsoredProductsKeywordTextFilterWithDefaults instantiates a new SponsoredProductsKeywordTextFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsKeywordTextFilterWithDefaults() *SponsoredProductsKeywordTextFilter {
	this := SponsoredProductsKeywordTextFilter{}
	return &this
}

// GetQueryTermMatchType returns the QueryTermMatchType field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordTextFilter) GetQueryTermMatchType() SponsoredProductsQueryTermMatchType {
	if o == nil || IsNil(o.QueryTermMatchType) {
		var ret SponsoredProductsQueryTermMatchType
		return ret
	}
	return *o.QueryTermMatchType
}

// GetQueryTermMatchTypeOk returns a tuple with the QueryTermMatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordTextFilter) GetQueryTermMatchTypeOk() (*SponsoredProductsQueryTermMatchType, bool) {
	if o == nil || IsNil(o.QueryTermMatchType) {
		return nil, false
	}
	return o.QueryTermMatchType, true
}

// HasQueryTermMatchType returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordTextFilter) HasQueryTermMatchType() bool {
	if o != nil && !IsNil(o.QueryTermMatchType) {
		return true
	}

	return false
}

// SetQueryTermMatchType gets a reference to the given SponsoredProductsQueryTermMatchType and assigns it to the QueryTermMatchType field.
func (o *SponsoredProductsKeywordTextFilter) SetQueryTermMatchType(v SponsoredProductsQueryTermMatchType) {
	o.QueryTermMatchType = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordTextFilter) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordTextFilter) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordTextFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *SponsoredProductsKeywordTextFilter) SetInclude(v []string) {
	o.Include = v
}

func (o SponsoredProductsKeywordTextFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueryTermMatchType) {
		toSerialize["queryTermMatchType"] = o.QueryTermMatchType
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableSponsoredProductsKeywordTextFilter struct {
	value *SponsoredProductsKeywordTextFilter
	isSet bool
}

func (v NullableSponsoredProductsKeywordTextFilter) Get() *SponsoredProductsKeywordTextFilter {
	return v.value
}

func (v *NullableSponsoredProductsKeywordTextFilter) Set(val *SponsoredProductsKeywordTextFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordTextFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordTextFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordTextFilter(val *SponsoredProductsKeywordTextFilter) *NullableSponsoredProductsKeywordTextFilter {
	return &NullableSponsoredProductsKeywordTextFilter{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordTextFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordTextFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
