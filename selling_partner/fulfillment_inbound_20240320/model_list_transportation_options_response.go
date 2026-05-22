package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListTransportationOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTransportationOptionsResponse{}

// ListTransportationOptionsResponse The `listTransportationOptions` response.
type ListTransportationOptionsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Transportation options generated for the placement option.
	TransportationOptions []TransportationOption `json:"transportationOptions"`
}

type _ListTransportationOptionsResponse ListTransportationOptionsResponse

// NewListTransportationOptionsResponse instantiates a new ListTransportationOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransportationOptionsResponse(transportationOptions []TransportationOption) *ListTransportationOptionsResponse {
	this := ListTransportationOptionsResponse{}
	this.TransportationOptions = transportationOptions
	return &this
}

// NewListTransportationOptionsResponseWithDefaults instantiates a new ListTransportationOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransportationOptionsResponseWithDefaults() *ListTransportationOptionsResponse {
	this := ListTransportationOptionsResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListTransportationOptionsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransportationOptionsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListTransportationOptionsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListTransportationOptionsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetTransportationOptions returns the TransportationOptions field value
func (o *ListTransportationOptionsResponse) GetTransportationOptions() []TransportationOption {
	if o == nil {
		var ret []TransportationOption
		return ret
	}

	return o.TransportationOptions
}

// GetTransportationOptionsOk returns a tuple with the TransportationOptions field value
// and a boolean to check if the value has been set.
func (o *ListTransportationOptionsResponse) GetTransportationOptionsOk() ([]TransportationOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransportationOptions, true
}

// SetTransportationOptions sets field value
func (o *ListTransportationOptionsResponse) SetTransportationOptions(v []TransportationOption) {
	o.TransportationOptions = v
}

func (o ListTransportationOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["transportationOptions"] = o.TransportationOptions
	return toSerialize, nil
}

type NullableListTransportationOptionsResponse struct {
	value *ListTransportationOptionsResponse
	isSet bool
}

func (v NullableListTransportationOptionsResponse) Get() *ListTransportationOptionsResponse {
	return v.value
}

func (v *NullableListTransportationOptionsResponse) Set(val *ListTransportationOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransportationOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransportationOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransportationOptionsResponse(val *ListTransportationOptionsResponse) *NullableListTransportationOptionsResponse {
	return &NullableListTransportationOptionsResponse{value: val, isSet: true}
}

func (v NullableListTransportationOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListTransportationOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
