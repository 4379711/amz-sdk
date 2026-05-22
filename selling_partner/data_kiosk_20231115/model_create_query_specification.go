package data_kiosk_20231115

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateQuerySpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateQuerySpecification{}

// CreateQuerySpecification Information required to create the query.
type CreateQuerySpecification struct {
	// The GraphQL query to submit. A query must be at most 8000 characters after unnecessary whitespace is removed.
	Query string `json:"query"`
	// A token to fetch a certain page of query results when there are multiple pages of query results available. The value of this token must be fetched from the `pagination.nextToken` field of the `Query` object, and the `query` field for this object must also be set to the `query` field of the same `Query` object. A `Query` object can be retrieved from either the `getQueries` or `getQuery` operation. In the absence of this token value, the first page of query results will be requested.
	PaginationToken *string `json:"paginationToken,omitempty"`
}

type _CreateQuerySpecification CreateQuerySpecification

// NewCreateQuerySpecification instantiates a new CreateQuerySpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateQuerySpecification(query string) *CreateQuerySpecification {
	this := CreateQuerySpecification{}
	this.Query = query
	return &this
}

// NewCreateQuerySpecificationWithDefaults instantiates a new CreateQuerySpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateQuerySpecificationWithDefaults() *CreateQuerySpecification {
	this := CreateQuerySpecification{}
	return &this
}

// GetQuery returns the Query field value
func (o *CreateQuerySpecification) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CreateQuerySpecification) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CreateQuerySpecification) SetQuery(v string) {
	o.Query = v
}

// GetPaginationToken returns the PaginationToken field value if set, zero value otherwise.
func (o *CreateQuerySpecification) GetPaginationToken() string {
	if o == nil || IsNil(o.PaginationToken) {
		var ret string
		return ret
	}
	return *o.PaginationToken
}

// GetPaginationTokenOk returns a tuple with the PaginationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateQuerySpecification) GetPaginationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PaginationToken) {
		return nil, false
	}
	return o.PaginationToken, true
}

// HasPaginationToken returns a boolean if a field has been set.
func (o *CreateQuerySpecification) HasPaginationToken() bool {
	if o != nil && !IsNil(o.PaginationToken) {
		return true
	}

	return false
}

// SetPaginationToken gets a reference to the given string and assigns it to the PaginationToken field.
func (o *CreateQuerySpecification) SetPaginationToken(v string) {
	o.PaginationToken = &v
}

func (o CreateQuerySpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	if !IsNil(o.PaginationToken) {
		toSerialize["paginationToken"] = o.PaginationToken
	}
	return toSerialize, nil
}

type NullableCreateQuerySpecification struct {
	value *CreateQuerySpecification
	isSet bool
}

func (v NullableCreateQuerySpecification) Get() *CreateQuerySpecification {
	return v.value
}

func (v *NullableCreateQuerySpecification) Set(val *CreateQuerySpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateQuerySpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateQuerySpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateQuerySpecification(val *CreateQuerySpecification) *NullableCreateQuerySpecification {
	return &NullableCreateQuerySpecification{value: val, isSet: true}
}

func (v NullableCreateQuerySpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateQuerySpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
