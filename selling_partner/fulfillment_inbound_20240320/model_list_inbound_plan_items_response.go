package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListInboundPlanItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInboundPlanItemsResponse{}

// ListInboundPlanItemsResponse The `listInboundPlanItems` response.
type ListInboundPlanItemsResponse struct {
	// The items in an inbound plan.
	Items      []Item      `json:"items"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type _ListInboundPlanItemsResponse ListInboundPlanItemsResponse

// NewListInboundPlanItemsResponse instantiates a new ListInboundPlanItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInboundPlanItemsResponse(items []Item) *ListInboundPlanItemsResponse {
	this := ListInboundPlanItemsResponse{}
	this.Items = items
	return &this
}

// NewListInboundPlanItemsResponseWithDefaults instantiates a new ListInboundPlanItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInboundPlanItemsResponseWithDefaults() *ListInboundPlanItemsResponse {
	this := ListInboundPlanItemsResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ListInboundPlanItemsResponse) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListInboundPlanItemsResponse) GetItemsOk() ([]Item, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListInboundPlanItemsResponse) SetItems(v []Item) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListInboundPlanItemsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInboundPlanItemsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListInboundPlanItemsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListInboundPlanItemsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListInboundPlanItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListInboundPlanItemsResponse struct {
	value *ListInboundPlanItemsResponse
	isSet bool
}

func (v NullableListInboundPlanItemsResponse) Get() *ListInboundPlanItemsResponse {
	return v.value
}

func (v *NullableListInboundPlanItemsResponse) Set(val *ListInboundPlanItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListInboundPlanItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListInboundPlanItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInboundPlanItemsResponse(val *ListInboundPlanItemsResponse) *NullableListInboundPlanItemsResponse {
	return &NullableListInboundPlanItemsResponse{value: val, isSet: true}
}

func (v NullableListInboundPlanItemsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListInboundPlanItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
