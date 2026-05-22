package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the ListOffersRequestSort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOffersRequestSort{}

// ListOffersRequestSort Use these parameters to sort the response.
type ListOffersRequestSort struct {
	Order SortOrder         `json:"order"`
	Key   ListOffersSortKey `json:"key"`
}

type _ListOffersRequestSort ListOffersRequestSort

// NewListOffersRequestSort instantiates a new ListOffersRequestSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOffersRequestSort(order SortOrder, key ListOffersSortKey) *ListOffersRequestSort {
	this := ListOffersRequestSort{}
	this.Order = order
	this.Key = key
	return &this
}

// NewListOffersRequestSortWithDefaults instantiates a new ListOffersRequestSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOffersRequestSortWithDefaults() *ListOffersRequestSort {
	this := ListOffersRequestSort{}
	return &this
}

// GetOrder returns the Order field value
func (o *ListOffersRequestSort) GetOrder() SortOrder {
	if o == nil {
		var ret SortOrder
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ListOffersRequestSort) GetOrderOk() (*SortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ListOffersRequestSort) SetOrder(v SortOrder) {
	o.Order = v
}

// GetKey returns the Key field value
func (o *ListOffersRequestSort) GetKey() ListOffersSortKey {
	if o == nil {
		var ret ListOffersSortKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ListOffersRequestSort) GetKeyOk() (*ListOffersSortKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ListOffersRequestSort) SetKey(v ListOffersSortKey) {
	o.Key = v
}

func (o ListOffersRequestSort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order"] = o.Order
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableListOffersRequestSort struct {
	value *ListOffersRequestSort
	isSet bool
}

func (v NullableListOffersRequestSort) Get() *ListOffersRequestSort {
	return v.value
}

func (v *NullableListOffersRequestSort) Set(val *ListOffersRequestSort) {
	v.value = val
	v.isSet = true
}

func (v NullableListOffersRequestSort) IsSet() bool {
	return v.isSet
}

func (v *NullableListOffersRequestSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOffersRequestSort(val *ListOffersRequestSort) *NullableListOffersRequestSort {
	return &NullableListOffersRequestSort{value: val, isSet: true}
}

func (v NullableListOffersRequestSort) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListOffersRequestSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
