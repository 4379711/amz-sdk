package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the NameFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameFilter{}

// NameFilter Filter entities by name
type NameFilter struct {
	QueryTermMatchType *QueryTermMatchType `json:"queryTermMatchType,omitempty"`
	Include            []string            `json:"include,omitempty"`
}

// NewNameFilter instantiates a new NameFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameFilter() *NameFilter {
	this := NameFilter{}
	return &this
}

// NewNameFilterWithDefaults instantiates a new NameFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameFilterWithDefaults() *NameFilter {
	this := NameFilter{}
	return &this
}

// GetQueryTermMatchType returns the QueryTermMatchType field value if set, zero value otherwise.
func (o *NameFilter) GetQueryTermMatchType() QueryTermMatchType {
	if o == nil || IsNil(o.QueryTermMatchType) {
		var ret QueryTermMatchType
		return ret
	}
	return *o.QueryTermMatchType
}

// GetQueryTermMatchTypeOk returns a tuple with the QueryTermMatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameFilter) GetQueryTermMatchTypeOk() (*QueryTermMatchType, bool) {
	if o == nil || IsNil(o.QueryTermMatchType) {
		return nil, false
	}
	return o.QueryTermMatchType, true
}

// HasQueryTermMatchType returns a boolean if a field has been set.
func (o *NameFilter) HasQueryTermMatchType() bool {
	if o != nil && !IsNil(o.QueryTermMatchType) {
		return true
	}

	return false
}

// SetQueryTermMatchType gets a reference to the given QueryTermMatchType and assigns it to the QueryTermMatchType field.
func (o *NameFilter) SetQueryTermMatchType(v QueryTermMatchType) {
	o.QueryTermMatchType = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *NameFilter) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameFilter) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *NameFilter) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *NameFilter) SetInclude(v []string) {
	o.Include = v
}

func (o NameFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueryTermMatchType) {
		toSerialize["queryTermMatchType"] = o.QueryTermMatchType
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	return toSerialize, nil
}

type NullableNameFilter struct {
	value *NameFilter
	isSet bool
}

func (v NullableNameFilter) Get() *NameFilter {
	return v.value
}

func (v *NullableNameFilter) Set(val *NameFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableNameFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableNameFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameFilter(val *NameFilter) *NullableNameFilter {
	return &NullableNameFilter{value: val, isSet: true}
}

func (v NullableNameFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableNameFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
