package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the ListOffersRequestPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOffersRequestPagination{}

// ListOffersRequestPagination Use these parameters to paginate through the response.
type ListOffersRequestPagination struct {
	// The maximum number of results to return in the response.
	Limit int64 `json:"limit"`
	// The offset from which to retrieve the number of results specified by the `limit` value. The first result is at offset 0.
	Offset int64 `json:"offset"`
}

type _ListOffersRequestPagination ListOffersRequestPagination

// NewListOffersRequestPagination instantiates a new ListOffersRequestPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOffersRequestPagination(limit int64, offset int64) *ListOffersRequestPagination {
	this := ListOffersRequestPagination{}
	this.Limit = limit
	this.Offset = offset
	return &this
}

// NewListOffersRequestPaginationWithDefaults instantiates a new ListOffersRequestPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOffersRequestPaginationWithDefaults() *ListOffersRequestPagination {
	this := ListOffersRequestPagination{}
	return &this
}

// GetLimit returns the Limit field value
func (o *ListOffersRequestPagination) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListOffersRequestPagination) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListOffersRequestPagination) SetLimit(v int64) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *ListOffersRequestPagination) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListOffersRequestPagination) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListOffersRequestPagination) SetOffset(v int64) {
	o.Offset = v
}

func (o ListOffersRequestPagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	return toSerialize, nil
}

type NullableListOffersRequestPagination struct {
	value *ListOffersRequestPagination
	isSet bool
}

func (v NullableListOffersRequestPagination) Get() *ListOffersRequestPagination {
	return v.value
}

func (v *NullableListOffersRequestPagination) Set(val *ListOffersRequestPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableListOffersRequestPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableListOffersRequestPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOffersRequestPagination(val *ListOffersRequestPagination) *NullableListOffersRequestPagination {
	return &NullableListOffersRequestPagination{value: val, isSet: true}
}

func (v NullableListOffersRequestPagination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListOffersRequestPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
