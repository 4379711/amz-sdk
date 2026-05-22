package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListShipmentItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListShipmentItemsResponse{}

// ListShipmentItemsResponse The `listShipmentItems` response.
type ListShipmentItemsResponse struct {
	// The items in a shipment.
	Items      []Item      `json:"items"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type _ListShipmentItemsResponse ListShipmentItemsResponse

// NewListShipmentItemsResponse instantiates a new ListShipmentItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListShipmentItemsResponse(items []Item) *ListShipmentItemsResponse {
	this := ListShipmentItemsResponse{}
	this.Items = items
	return &this
}

// NewListShipmentItemsResponseWithDefaults instantiates a new ListShipmentItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListShipmentItemsResponseWithDefaults() *ListShipmentItemsResponse {
	this := ListShipmentItemsResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ListShipmentItemsResponse) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListShipmentItemsResponse) GetItemsOk() ([]Item, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListShipmentItemsResponse) SetItems(v []Item) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListShipmentItemsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShipmentItemsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListShipmentItemsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListShipmentItemsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListShipmentItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListShipmentItemsResponse struct {
	value *ListShipmentItemsResponse
	isSet bool
}

func (v NullableListShipmentItemsResponse) Get() *ListShipmentItemsResponse {
	return v.value
}

func (v *NullableListShipmentItemsResponse) Set(val *ListShipmentItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListShipmentItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListShipmentItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListShipmentItemsResponse(val *ListShipmentItemsResponse) *NullableListShipmentItemsResponse {
	return &NullableListShipmentItemsResponse{value: val, isSet: true}
}

func (v NullableListShipmentItemsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListShipmentItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
