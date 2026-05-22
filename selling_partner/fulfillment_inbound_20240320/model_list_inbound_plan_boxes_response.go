package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListInboundPlanBoxesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInboundPlanBoxesResponse{}

// ListInboundPlanBoxesResponse The `listInboundPlanBoxes` response.
type ListInboundPlanBoxesResponse struct {
	// A list of boxes in an inbound plan.
	Boxes      []Box       `json:"boxes"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type _ListInboundPlanBoxesResponse ListInboundPlanBoxesResponse

// NewListInboundPlanBoxesResponse instantiates a new ListInboundPlanBoxesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInboundPlanBoxesResponse(boxes []Box) *ListInboundPlanBoxesResponse {
	this := ListInboundPlanBoxesResponse{}
	this.Boxes = boxes
	return &this
}

// NewListInboundPlanBoxesResponseWithDefaults instantiates a new ListInboundPlanBoxesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInboundPlanBoxesResponseWithDefaults() *ListInboundPlanBoxesResponse {
	this := ListInboundPlanBoxesResponse{}
	return &this
}

// GetBoxes returns the Boxes field value
func (o *ListInboundPlanBoxesResponse) GetBoxes() []Box {
	if o == nil {
		var ret []Box
		return ret
	}

	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value
// and a boolean to check if the value has been set.
func (o *ListInboundPlanBoxesResponse) GetBoxesOk() ([]Box, bool) {
	if o == nil {
		return nil, false
	}
	return o.Boxes, true
}

// SetBoxes sets field value
func (o *ListInboundPlanBoxesResponse) SetBoxes(v []Box) {
	o.Boxes = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListInboundPlanBoxesResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInboundPlanBoxesResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListInboundPlanBoxesResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListInboundPlanBoxesResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListInboundPlanBoxesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boxes"] = o.Boxes
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListInboundPlanBoxesResponse struct {
	value *ListInboundPlanBoxesResponse
	isSet bool
}

func (v NullableListInboundPlanBoxesResponse) Get() *ListInboundPlanBoxesResponse {
	return v.value
}

func (v *NullableListInboundPlanBoxesResponse) Set(val *ListInboundPlanBoxesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListInboundPlanBoxesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListInboundPlanBoxesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInboundPlanBoxesResponse(val *ListInboundPlanBoxesResponse) *NullableListInboundPlanBoxesResponse {
	return &NullableListInboundPlanBoxesResponse{value: val, isSet: true}
}

func (v NullableListInboundPlanBoxesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListInboundPlanBoxesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
