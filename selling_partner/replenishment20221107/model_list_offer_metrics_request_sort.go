package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the ListOfferMetricsRequestSort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOfferMetricsRequestSort{}

// ListOfferMetricsRequestSort Use these parameters to sort the response.
type ListOfferMetricsRequestSort struct {
	Order SortOrder               `json:"order"`
	Key   ListOfferMetricsSortKey `json:"key"`
}

type _ListOfferMetricsRequestSort ListOfferMetricsRequestSort

// NewListOfferMetricsRequestSort instantiates a new ListOfferMetricsRequestSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOfferMetricsRequestSort(order SortOrder, key ListOfferMetricsSortKey) *ListOfferMetricsRequestSort {
	this := ListOfferMetricsRequestSort{}
	this.Order = order
	this.Key = key
	return &this
}

// NewListOfferMetricsRequestSortWithDefaults instantiates a new ListOfferMetricsRequestSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOfferMetricsRequestSortWithDefaults() *ListOfferMetricsRequestSort {
	this := ListOfferMetricsRequestSort{}
	return &this
}

// GetOrder returns the Order field value
func (o *ListOfferMetricsRequestSort) GetOrder() SortOrder {
	if o == nil {
		var ret SortOrder
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ListOfferMetricsRequestSort) GetOrderOk() (*SortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ListOfferMetricsRequestSort) SetOrder(v SortOrder) {
	o.Order = v
}

// GetKey returns the Key field value
func (o *ListOfferMetricsRequestSort) GetKey() ListOfferMetricsSortKey {
	if o == nil {
		var ret ListOfferMetricsSortKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ListOfferMetricsRequestSort) GetKeyOk() (*ListOfferMetricsSortKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ListOfferMetricsRequestSort) SetKey(v ListOfferMetricsSortKey) {
	o.Key = v
}

func (o ListOfferMetricsRequestSort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order"] = o.Order
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableListOfferMetricsRequestSort struct {
	value *ListOfferMetricsRequestSort
	isSet bool
}

func (v NullableListOfferMetricsRequestSort) Get() *ListOfferMetricsRequestSort {
	return v.value
}

func (v *NullableListOfferMetricsRequestSort) Set(val *ListOfferMetricsRequestSort) {
	v.value = val
	v.isSet = true
}

func (v NullableListOfferMetricsRequestSort) IsSet() bool {
	return v.isSet
}

func (v *NullableListOfferMetricsRequestSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOfferMetricsRequestSort(val *ListOfferMetricsRequestSort) *NullableListOfferMetricsRequestSort {
	return &NullableListOfferMetricsRequestSort{value: val, isSet: true}
}

func (v NullableListOfferMetricsRequestSort) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListOfferMetricsRequestSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
