package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListInboundPlansResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInboundPlansResponse{}

// ListInboundPlansResponse The `listInboundPlans` response.
type ListInboundPlansResponse struct {
	// A list of inbound plans with minimal information.
	InboundPlans []InboundPlanSummary `json:"inboundPlans,omitempty"`
	Pagination   *Pagination          `json:"pagination,omitempty"`
}

// NewListInboundPlansResponse instantiates a new ListInboundPlansResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInboundPlansResponse() *ListInboundPlansResponse {
	this := ListInboundPlansResponse{}
	return &this
}

// NewListInboundPlansResponseWithDefaults instantiates a new ListInboundPlansResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInboundPlansResponseWithDefaults() *ListInboundPlansResponse {
	this := ListInboundPlansResponse{}
	return &this
}

// GetInboundPlans returns the InboundPlans field value if set, zero value otherwise.
func (o *ListInboundPlansResponse) GetInboundPlans() []InboundPlanSummary {
	if o == nil || IsNil(o.InboundPlans) {
		var ret []InboundPlanSummary
		return ret
	}
	return o.InboundPlans
}

// GetInboundPlansOk returns a tuple with the InboundPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInboundPlansResponse) GetInboundPlansOk() ([]InboundPlanSummary, bool) {
	if o == nil || IsNil(o.InboundPlans) {
		return nil, false
	}
	return o.InboundPlans, true
}

// HasInboundPlans returns a boolean if a field has been set.
func (o *ListInboundPlansResponse) HasInboundPlans() bool {
	if o != nil && !IsNil(o.InboundPlans) {
		return true
	}

	return false
}

// SetInboundPlans gets a reference to the given []InboundPlanSummary and assigns it to the InboundPlans field.
func (o *ListInboundPlansResponse) SetInboundPlans(v []InboundPlanSummary) {
	o.InboundPlans = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListInboundPlansResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInboundPlansResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListInboundPlansResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListInboundPlansResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListInboundPlansResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InboundPlans) {
		toSerialize["inboundPlans"] = o.InboundPlans
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListInboundPlansResponse struct {
	value *ListInboundPlansResponse
	isSet bool
}

func (v NullableListInboundPlansResponse) Get() *ListInboundPlansResponse {
	return v.value
}

func (v *NullableListInboundPlansResponse) Set(val *ListInboundPlansResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListInboundPlansResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListInboundPlansResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInboundPlansResponse(val *ListInboundPlansResponse) *NullableListInboundPlansResponse {
	return &NullableListInboundPlansResponse{value: val, isSet: true}
}

func (v NullableListInboundPlansResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListInboundPlansResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
