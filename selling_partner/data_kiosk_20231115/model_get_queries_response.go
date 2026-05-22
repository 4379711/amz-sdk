package data_kiosk_20231115

import (
	"github.com/bytedance/sonic"
)

// checks if the GetQueriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQueriesResponse{}

// GetQueriesResponse The response for the `getQueries` operation.
type GetQueriesResponse struct {
	// A list of queries.
	Queries    []Query                       `json:"queries"`
	Pagination *GetQueriesResponsePagination `json:"pagination,omitempty"`
}

type _GetQueriesResponse GetQueriesResponse

// NewGetQueriesResponse instantiates a new GetQueriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQueriesResponse(queries []Query) *GetQueriesResponse {
	this := GetQueriesResponse{}
	this.Queries = queries
	return &this
}

// NewGetQueriesResponseWithDefaults instantiates a new GetQueriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQueriesResponseWithDefaults() *GetQueriesResponse {
	this := GetQueriesResponse{}
	return &this
}

// GetQueries returns the Queries field value
func (o *GetQueriesResponse) GetQueries() []Query {
	if o == nil {
		var ret []Query
		return ret
	}

	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *GetQueriesResponse) GetQueriesOk() ([]Query, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queries, true
}

// SetQueries sets field value
func (o *GetQueriesResponse) SetQueries(v []Query) {
	o.Queries = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetQueriesResponse) GetPagination() GetQueriesResponsePagination {
	if o == nil || IsNil(o.Pagination) {
		var ret GetQueriesResponsePagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQueriesResponse) GetPaginationOk() (*GetQueriesResponsePagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetQueriesResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given GetQueriesResponsePagination and assigns it to the Pagination field.
func (o *GetQueriesResponse) SetPagination(v GetQueriesResponsePagination) {
	o.Pagination = &v
}

func (o GetQueriesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queries"] = o.Queries
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableGetQueriesResponse struct {
	value *GetQueriesResponse
	isSet bool
}

func (v NullableGetQueriesResponse) Get() *GetQueriesResponse {
	return v.value
}

func (v *NullableGetQueriesResponse) Set(val *GetQueriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQueriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQueriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQueriesResponse(val *GetQueriesResponse) *NullableGetQueriesResponse {
	return &NullableGetQueriesResponse{value: val, isSet: true}
}

func (v NullableGetQueriesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetQueriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
