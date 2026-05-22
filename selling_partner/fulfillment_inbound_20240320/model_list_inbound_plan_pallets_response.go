package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListInboundPlanPalletsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInboundPlanPalletsResponse{}

// ListInboundPlanPalletsResponse The `listInboundPlanPallets` response.
type ListInboundPlanPalletsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// The pallets in an inbound plan.
	Pallets []Pallet `json:"pallets"`
}

type _ListInboundPlanPalletsResponse ListInboundPlanPalletsResponse

// NewListInboundPlanPalletsResponse instantiates a new ListInboundPlanPalletsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInboundPlanPalletsResponse(pallets []Pallet) *ListInboundPlanPalletsResponse {
	this := ListInboundPlanPalletsResponse{}
	this.Pallets = pallets
	return &this
}

// NewListInboundPlanPalletsResponseWithDefaults instantiates a new ListInboundPlanPalletsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInboundPlanPalletsResponseWithDefaults() *ListInboundPlanPalletsResponse {
	this := ListInboundPlanPalletsResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListInboundPlanPalletsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInboundPlanPalletsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListInboundPlanPalletsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListInboundPlanPalletsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetPallets returns the Pallets field value
func (o *ListInboundPlanPalletsResponse) GetPallets() []Pallet {
	if o == nil {
		var ret []Pallet
		return ret
	}

	return o.Pallets
}

// GetPalletsOk returns a tuple with the Pallets field value
// and a boolean to check if the value has been set.
func (o *ListInboundPlanPalletsResponse) GetPalletsOk() ([]Pallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pallets, true
}

// SetPallets sets field value
func (o *ListInboundPlanPalletsResponse) SetPallets(v []Pallet) {
	o.Pallets = v
}

func (o ListInboundPlanPalletsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["pallets"] = o.Pallets
	return toSerialize, nil
}

type NullableListInboundPlanPalletsResponse struct {
	value *ListInboundPlanPalletsResponse
	isSet bool
}

func (v NullableListInboundPlanPalletsResponse) Get() *ListInboundPlanPalletsResponse {
	return v.value
}

func (v *NullableListInboundPlanPalletsResponse) Set(val *ListInboundPlanPalletsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListInboundPlanPalletsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListInboundPlanPalletsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInboundPlanPalletsResponse(val *ListInboundPlanPalletsResponse) *NullableListInboundPlanPalletsResponse {
	return &NullableListInboundPlanPalletsResponse{value: val, isSet: true}
}

func (v NullableListInboundPlanPalletsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListInboundPlanPalletsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
