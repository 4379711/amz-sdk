package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the ListOfferMetricsRequestPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOfferMetricsRequestPagination{}

// ListOfferMetricsRequestPagination Use these parameters to paginate through the response.
type ListOfferMetricsRequestPagination struct {
	// The maximum number of results to return in the response.
	Limit int64 `json:"limit"`
	// The offset from which to retrieve the number of results specified by the `limit` value. The first result is at offset 0.
	Offset int64 `json:"offset"`
}

type _ListOfferMetricsRequestPagination ListOfferMetricsRequestPagination

// NewListOfferMetricsRequestPagination instantiates a new ListOfferMetricsRequestPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOfferMetricsRequestPagination(limit int64, offset int64) *ListOfferMetricsRequestPagination {
	this := ListOfferMetricsRequestPagination{}
	this.Limit = limit
	this.Offset = offset
	return &this
}

// NewListOfferMetricsRequestPaginationWithDefaults instantiates a new ListOfferMetricsRequestPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOfferMetricsRequestPaginationWithDefaults() *ListOfferMetricsRequestPagination {
	this := ListOfferMetricsRequestPagination{}
	return &this
}

// GetLimit returns the Limit field value
func (o *ListOfferMetricsRequestPagination) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListOfferMetricsRequestPagination) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListOfferMetricsRequestPagination) SetLimit(v int64) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *ListOfferMetricsRequestPagination) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListOfferMetricsRequestPagination) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListOfferMetricsRequestPagination) SetOffset(v int64) {
	o.Offset = v
}

func (o ListOfferMetricsRequestPagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	return toSerialize, nil
}

type NullableListOfferMetricsRequestPagination struct {
	value *ListOfferMetricsRequestPagination
	isSet bool
}

func (v NullableListOfferMetricsRequestPagination) Get() *ListOfferMetricsRequestPagination {
	return v.value
}

func (v *NullableListOfferMetricsRequestPagination) Set(val *ListOfferMetricsRequestPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableListOfferMetricsRequestPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableListOfferMetricsRequestPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOfferMetricsRequestPagination(val *ListOfferMetricsRequestPagination) *NullableListOfferMetricsRequestPagination {
	return &NullableListOfferMetricsRequestPagination{value: val, isSet: true}
}

func (v NullableListOfferMetricsRequestPagination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListOfferMetricsRequestPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
