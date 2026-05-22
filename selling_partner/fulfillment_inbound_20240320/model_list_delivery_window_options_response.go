package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListDeliveryWindowOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDeliveryWindowOptionsResponse{}

// ListDeliveryWindowOptionsResponse The `listDeliveryWindowOptions` response.
type ListDeliveryWindowOptionsResponse struct {
	// Delivery window options generated for the placement option.
	DeliveryWindowOptions []DeliveryWindowOption `json:"deliveryWindowOptions"`
	Pagination            *Pagination            `json:"pagination,omitempty"`
}

type _ListDeliveryWindowOptionsResponse ListDeliveryWindowOptionsResponse

// NewListDeliveryWindowOptionsResponse instantiates a new ListDeliveryWindowOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDeliveryWindowOptionsResponse(deliveryWindowOptions []DeliveryWindowOption) *ListDeliveryWindowOptionsResponse {
	this := ListDeliveryWindowOptionsResponse{}
	this.DeliveryWindowOptions = deliveryWindowOptions
	return &this
}

// NewListDeliveryWindowOptionsResponseWithDefaults instantiates a new ListDeliveryWindowOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDeliveryWindowOptionsResponseWithDefaults() *ListDeliveryWindowOptionsResponse {
	this := ListDeliveryWindowOptionsResponse{}
	return &this
}

// GetDeliveryWindowOptions returns the DeliveryWindowOptions field value
func (o *ListDeliveryWindowOptionsResponse) GetDeliveryWindowOptions() []DeliveryWindowOption {
	if o == nil {
		var ret []DeliveryWindowOption
		return ret
	}

	return o.DeliveryWindowOptions
}

// GetDeliveryWindowOptionsOk returns a tuple with the DeliveryWindowOptions field value
// and a boolean to check if the value has been set.
func (o *ListDeliveryWindowOptionsResponse) GetDeliveryWindowOptionsOk() ([]DeliveryWindowOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryWindowOptions, true
}

// SetDeliveryWindowOptions sets field value
func (o *ListDeliveryWindowOptionsResponse) SetDeliveryWindowOptions(v []DeliveryWindowOption) {
	o.DeliveryWindowOptions = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListDeliveryWindowOptionsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDeliveryWindowOptionsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListDeliveryWindowOptionsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListDeliveryWindowOptionsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListDeliveryWindowOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliveryWindowOptions"] = o.DeliveryWindowOptions
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListDeliveryWindowOptionsResponse struct {
	value *ListDeliveryWindowOptionsResponse
	isSet bool
}

func (v NullableListDeliveryWindowOptionsResponse) Get() *ListDeliveryWindowOptionsResponse {
	return v.value
}

func (v *NullableListDeliveryWindowOptionsResponse) Set(val *ListDeliveryWindowOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDeliveryWindowOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDeliveryWindowOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDeliveryWindowOptionsResponse(val *ListDeliveryWindowOptionsResponse) *NullableListDeliveryWindowOptionsResponse {
	return &NullableListDeliveryWindowOptionsResponse{value: val, isSet: true}
}

func (v NullableListDeliveryWindowOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListDeliveryWindowOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
