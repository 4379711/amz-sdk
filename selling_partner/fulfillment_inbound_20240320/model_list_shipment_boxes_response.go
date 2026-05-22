package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListShipmentBoxesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListShipmentBoxesResponse{}

// ListShipmentBoxesResponse The `listShipmentBoxes` response.
type ListShipmentBoxesResponse struct {
	// A list of boxes in a shipment.
	Boxes      []Box       `json:"boxes"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type _ListShipmentBoxesResponse ListShipmentBoxesResponse

// NewListShipmentBoxesResponse instantiates a new ListShipmentBoxesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListShipmentBoxesResponse(boxes []Box) *ListShipmentBoxesResponse {
	this := ListShipmentBoxesResponse{}
	this.Boxes = boxes
	return &this
}

// NewListShipmentBoxesResponseWithDefaults instantiates a new ListShipmentBoxesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListShipmentBoxesResponseWithDefaults() *ListShipmentBoxesResponse {
	this := ListShipmentBoxesResponse{}
	return &this
}

// GetBoxes returns the Boxes field value
func (o *ListShipmentBoxesResponse) GetBoxes() []Box {
	if o == nil {
		var ret []Box
		return ret
	}

	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value
// and a boolean to check if the value has been set.
func (o *ListShipmentBoxesResponse) GetBoxesOk() ([]Box, bool) {
	if o == nil {
		return nil, false
	}
	return o.Boxes, true
}

// SetBoxes sets field value
func (o *ListShipmentBoxesResponse) SetBoxes(v []Box) {
	o.Boxes = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListShipmentBoxesResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShipmentBoxesResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListShipmentBoxesResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListShipmentBoxesResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListShipmentBoxesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boxes"] = o.Boxes
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListShipmentBoxesResponse struct {
	value *ListShipmentBoxesResponse
	isSet bool
}

func (v NullableListShipmentBoxesResponse) Get() *ListShipmentBoxesResponse {
	return v.value
}

func (v *NullableListShipmentBoxesResponse) Set(val *ListShipmentBoxesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListShipmentBoxesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListShipmentBoxesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListShipmentBoxesResponse(val *ListShipmentBoxesResponse) *NullableListShipmentBoxesResponse {
	return &NullableListShipmentBoxesResponse{value: val, isSet: true}
}

func (v NullableListShipmentBoxesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListShipmentBoxesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
