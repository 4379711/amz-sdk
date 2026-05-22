package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListPackingGroupBoxesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPackingGroupBoxesResponse{}

// ListPackingGroupBoxesResponse The `listPackingGroupBoxes` response.
type ListPackingGroupBoxesResponse struct {
	// Provides the information about the list of boxes in the packing group.
	Boxes      []Box       `json:"boxes"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type _ListPackingGroupBoxesResponse ListPackingGroupBoxesResponse

// NewListPackingGroupBoxesResponse instantiates a new ListPackingGroupBoxesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPackingGroupBoxesResponse(boxes []Box) *ListPackingGroupBoxesResponse {
	this := ListPackingGroupBoxesResponse{}
	this.Boxes = boxes
	return &this
}

// NewListPackingGroupBoxesResponseWithDefaults instantiates a new ListPackingGroupBoxesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPackingGroupBoxesResponseWithDefaults() *ListPackingGroupBoxesResponse {
	this := ListPackingGroupBoxesResponse{}
	return &this
}

// GetBoxes returns the Boxes field value
func (o *ListPackingGroupBoxesResponse) GetBoxes() []Box {
	if o == nil {
		var ret []Box
		return ret
	}

	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value
// and a boolean to check if the value has been set.
func (o *ListPackingGroupBoxesResponse) GetBoxesOk() ([]Box, bool) {
	if o == nil {
		return nil, false
	}
	return o.Boxes, true
}

// SetBoxes sets field value
func (o *ListPackingGroupBoxesResponse) SetBoxes(v []Box) {
	o.Boxes = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListPackingGroupBoxesResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPackingGroupBoxesResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListPackingGroupBoxesResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListPackingGroupBoxesResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListPackingGroupBoxesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boxes"] = o.Boxes
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListPackingGroupBoxesResponse struct {
	value *ListPackingGroupBoxesResponse
	isSet bool
}

func (v NullableListPackingGroupBoxesResponse) Get() *ListPackingGroupBoxesResponse {
	return v.value
}

func (v *NullableListPackingGroupBoxesResponse) Set(val *ListPackingGroupBoxesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPackingGroupBoxesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPackingGroupBoxesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPackingGroupBoxesResponse(val *ListPackingGroupBoxesResponse) *NullableListPackingGroupBoxesResponse {
	return &NullableListPackingGroupBoxesResponse{value: val, isSet: true}
}

func (v NullableListPackingGroupBoxesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListPackingGroupBoxesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
