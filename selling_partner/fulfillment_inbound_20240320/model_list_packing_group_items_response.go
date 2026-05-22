package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListPackingGroupItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPackingGroupItemsResponse{}

// ListPackingGroupItemsResponse The `listPackingGroupItems` response.
type ListPackingGroupItemsResponse struct {
	// Provides the information about the list of items in the packing group.
	Items      []Item      `json:"items"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type _ListPackingGroupItemsResponse ListPackingGroupItemsResponse

// NewListPackingGroupItemsResponse instantiates a new ListPackingGroupItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPackingGroupItemsResponse(items []Item) *ListPackingGroupItemsResponse {
	this := ListPackingGroupItemsResponse{}
	this.Items = items
	return &this
}

// NewListPackingGroupItemsResponseWithDefaults instantiates a new ListPackingGroupItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPackingGroupItemsResponseWithDefaults() *ListPackingGroupItemsResponse {
	this := ListPackingGroupItemsResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ListPackingGroupItemsResponse) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListPackingGroupItemsResponse) GetItemsOk() ([]Item, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListPackingGroupItemsResponse) SetItems(v []Item) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListPackingGroupItemsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPackingGroupItemsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListPackingGroupItemsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListPackingGroupItemsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListPackingGroupItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListPackingGroupItemsResponse struct {
	value *ListPackingGroupItemsResponse
	isSet bool
}

func (v NullableListPackingGroupItemsResponse) Get() *ListPackingGroupItemsResponse {
	return v.value
}

func (v *NullableListPackingGroupItemsResponse) Set(val *ListPackingGroupItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPackingGroupItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPackingGroupItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPackingGroupItemsResponse(val *ListPackingGroupItemsResponse) *NullableListPackingGroupItemsResponse {
	return &NullableListPackingGroupItemsResponse{value: val, isSet: true}
}

func (v NullableListPackingGroupItemsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListPackingGroupItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
