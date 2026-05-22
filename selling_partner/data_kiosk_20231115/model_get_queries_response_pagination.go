package data_kiosk_20231115

import (
	"github.com/bytedance/sonic"
)

// checks if the GetQueriesResponsePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQueriesResponsePagination{}

// GetQueriesResponsePagination When a request has results that are not included in this response, pagination occurs. This means the results are divided into pages. To retrieve the next page, you must pass the `nextToken` as the `paginationToken` query parameter in the subsequent `getQueries` request. All other parameters must be provided with the same values that were provided with the request that generated this token, with the exception of `pageSize` which can be modified between calls to `getQueries`. When there are no more pages to fetch, the `nextToken` field will be absent.
type GetQueriesResponsePagination struct {
	// A token that can be used to fetch the next page of results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewGetQueriesResponsePagination instantiates a new GetQueriesResponsePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQueriesResponsePagination() *GetQueriesResponsePagination {
	this := GetQueriesResponsePagination{}
	return &this
}

// NewGetQueriesResponsePaginationWithDefaults instantiates a new GetQueriesResponsePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQueriesResponsePaginationWithDefaults() *GetQueriesResponsePagination {
	this := GetQueriesResponsePagination{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetQueriesResponsePagination) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQueriesResponsePagination) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetQueriesResponsePagination) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetQueriesResponsePagination) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GetQueriesResponsePagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGetQueriesResponsePagination struct {
	value *GetQueriesResponsePagination
	isSet bool
}

func (v NullableGetQueriesResponsePagination) Get() *GetQueriesResponsePagination {
	return v.value
}

func (v *NullableGetQueriesResponsePagination) Set(val *GetQueriesResponsePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQueriesResponsePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQueriesResponsePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQueriesResponsePagination(val *GetQueriesResponsePagination) *NullableGetQueriesResponsePagination {
	return &NullableGetQueriesResponsePagination{value: val, isSet: true}
}

func (v NullableGetQueriesResponsePagination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetQueriesResponsePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
