package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the Pagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pagination{}

// Pagination When a request produces a response that exceeds the `pageSize`, pagination occurs. This means the response is divided into individual pages. To retrieve the next page or the previous page, you must pass the `nextToken` value or the `previousToken` value as the `pageToken` parameter in the next request. When you receive the last page, there will be no `nextToken` key in the pagination object.
type Pagination struct {
	// A token that can be used to fetch the next page.
	NextToken *string `json:"nextToken,omitempty"`
	// A token that can be used to fetch the previous page.
	PreviousToken *string `json:"previousToken,omitempty"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination() *Pagination {
	this := Pagination{}
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *Pagination) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *Pagination) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *Pagination) SetNextToken(v string) {
	o.NextToken = &v
}

// GetPreviousToken returns the PreviousToken field value if set, zero value otherwise.
func (o *Pagination) GetPreviousToken() string {
	if o == nil || IsNil(o.PreviousToken) {
		var ret string
		return ret
	}
	return *o.PreviousToken
}

// GetPreviousTokenOk returns a tuple with the PreviousToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetPreviousTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousToken) {
		return nil, false
	}
	return o.PreviousToken, true
}

// HasPreviousToken returns a boolean if a field has been set.
func (o *Pagination) HasPreviousToken() bool {
	if o != nil && !IsNil(o.PreviousToken) {
		return true
	}

	return false
}

// SetPreviousToken gets a reference to the given string and assigns it to the PreviousToken field.
func (o *Pagination) SetPreviousToken(v string) {
	o.PreviousToken = &v
}

func (o Pagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.PreviousToken) {
		toSerialize["previousToken"] = o.PreviousToken
	}
	return toSerialize, nil
}

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
