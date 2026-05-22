package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListShipmentPalletsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListShipmentPalletsResponse{}

// ListShipmentPalletsResponse The `listShipmentPallets` response.
type ListShipmentPalletsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// The pallets in a shipment.
	Pallets []Pallet `json:"pallets"`
}

type _ListShipmentPalletsResponse ListShipmentPalletsResponse

// NewListShipmentPalletsResponse instantiates a new ListShipmentPalletsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListShipmentPalletsResponse(pallets []Pallet) *ListShipmentPalletsResponse {
	this := ListShipmentPalletsResponse{}
	this.Pallets = pallets
	return &this
}

// NewListShipmentPalletsResponseWithDefaults instantiates a new ListShipmentPalletsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListShipmentPalletsResponseWithDefaults() *ListShipmentPalletsResponse {
	this := ListShipmentPalletsResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListShipmentPalletsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShipmentPalletsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListShipmentPalletsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListShipmentPalletsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetPallets returns the Pallets field value
func (o *ListShipmentPalletsResponse) GetPallets() []Pallet {
	if o == nil {
		var ret []Pallet
		return ret
	}

	return o.Pallets
}

// GetPalletsOk returns a tuple with the Pallets field value
// and a boolean to check if the value has been set.
func (o *ListShipmentPalletsResponse) GetPalletsOk() ([]Pallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pallets, true
}

// SetPallets sets field value
func (o *ListShipmentPalletsResponse) SetPallets(v []Pallet) {
	o.Pallets = v
}

func (o ListShipmentPalletsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["pallets"] = o.Pallets
	return toSerialize, nil
}

type NullableListShipmentPalletsResponse struct {
	value *ListShipmentPalletsResponse
	isSet bool
}

func (v NullableListShipmentPalletsResponse) Get() *ListShipmentPalletsResponse {
	return v.value
}

func (v *NullableListShipmentPalletsResponse) Set(val *ListShipmentPalletsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListShipmentPalletsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListShipmentPalletsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListShipmentPalletsResponse(val *ListShipmentPalletsResponse) *NullableListShipmentPalletsResponse {
	return &NullableListShipmentPalletsResponse{value: val, isSet: true}
}

func (v NullableListShipmentPalletsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListShipmentPalletsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
