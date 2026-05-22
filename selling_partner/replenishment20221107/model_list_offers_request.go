package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the ListOffersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOffersRequest{}

// ListOffersRequest The request body for the `listOffers` operation.
type ListOffersRequest struct {
	Pagination ListOffersRequestPagination `json:"pagination"`
	Filters    ListOffersRequestFilters    `json:"filters"`
	Sort       *ListOffersRequestSort      `json:"sort,omitempty"`
}

type _ListOffersRequest ListOffersRequest

// NewListOffersRequest instantiates a new ListOffersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOffersRequest(pagination ListOffersRequestPagination, filters ListOffersRequestFilters) *ListOffersRequest {
	this := ListOffersRequest{}
	this.Pagination = pagination
	this.Filters = filters
	return &this
}

// NewListOffersRequestWithDefaults instantiates a new ListOffersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOffersRequestWithDefaults() *ListOffersRequest {
	this := ListOffersRequest{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *ListOffersRequest) GetPagination() ListOffersRequestPagination {
	if o == nil {
		var ret ListOffersRequestPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListOffersRequest) GetPaginationOk() (*ListOffersRequestPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListOffersRequest) SetPagination(v ListOffersRequestPagination) {
	o.Pagination = v
}

// GetFilters returns the Filters field value
func (o *ListOffersRequest) GetFilters() ListOffersRequestFilters {
	if o == nil {
		var ret ListOffersRequestFilters
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ListOffersRequest) GetFiltersOk() (*ListOffersRequestFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *ListOffersRequest) SetFilters(v ListOffersRequestFilters) {
	o.Filters = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ListOffersRequest) GetSort() ListOffersRequestSort {
	if o == nil || IsNil(o.Sort) {
		var ret ListOffersRequestSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersRequest) GetSortOk() (*ListOffersRequestSort, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ListOffersRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given ListOffersRequestSort and assigns it to the Sort field.
func (o *ListOffersRequest) SetSort(v ListOffersRequestSort) {
	o.Sort = &v
}

func (o ListOffersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["filters"] = o.Filters
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	return toSerialize, nil
}

type NullableListOffersRequest struct {
	value *ListOffersRequest
	isSet bool
}

func (v NullableListOffersRequest) Get() *ListOffersRequest {
	return v.value
}

func (v *NullableListOffersRequest) Set(val *ListOffersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListOffersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListOffersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOffersRequest(val *ListOffersRequest) *NullableListOffersRequest {
	return &NullableListOffersRequest{value: val, isSet: true}
}

func (v NullableListOffersRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListOffersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
