package data_kiosk_20231115

import (
	"github.com/bytedance/sonic"
)

// checks if the QueryPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryPagination{}

// QueryPagination When a query produces results that are not included in the data document, pagination occurs. This means the results are divided into pages. To retrieve the next page, you must pass a `CreateQuerySpecification` object with `paginationToken` set to this object's `nextToken` and with `query` set to this object's `query` in the subsequent `createQuery` request. When there are no more pages to fetch, the `nextToken` field will be absent.
type QueryPagination struct {
	// A token that can be used to fetch the next page of results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewQueryPagination instantiates a new QueryPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPagination() *QueryPagination {
	this := QueryPagination{}
	return &this
}

// NewQueryPaginationWithDefaults instantiates a new QueryPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPaginationWithDefaults() *QueryPagination {
	this := QueryPagination{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *QueryPagination) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPagination) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *QueryPagination) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *QueryPagination) SetNextToken(v string) {
	o.NextToken = &v
}

func (o QueryPagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableQueryPagination struct {
	value *QueryPagination
	isSet bool
}

func (v NullableQueryPagination) Get() *QueryPagination {
	return v.value
}

func (v *NullableQueryPagination) Set(val *QueryPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPagination(val *QueryPagination) *NullableQueryPagination {
	return &NullableQueryPagination{value: val, isSet: true}
}

func (v NullableQueryPagination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableQueryPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
