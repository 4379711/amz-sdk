package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListPackingOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPackingOptionsResponse{}

// ListPackingOptionsResponse The `listPlacementOptions` response.
type ListPackingOptionsResponse struct {
	// List of packing options.
	PackingOptions []PackingOption `json:"packingOptions"`
	Pagination     *Pagination     `json:"pagination,omitempty"`
}

type _ListPackingOptionsResponse ListPackingOptionsResponse

// NewListPackingOptionsResponse instantiates a new ListPackingOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPackingOptionsResponse(packingOptions []PackingOption) *ListPackingOptionsResponse {
	this := ListPackingOptionsResponse{}
	this.PackingOptions = packingOptions
	return &this
}

// NewListPackingOptionsResponseWithDefaults instantiates a new ListPackingOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPackingOptionsResponseWithDefaults() *ListPackingOptionsResponse {
	this := ListPackingOptionsResponse{}
	return &this
}

// GetPackingOptions returns the PackingOptions field value
func (o *ListPackingOptionsResponse) GetPackingOptions() []PackingOption {
	if o == nil {
		var ret []PackingOption
		return ret
	}

	return o.PackingOptions
}

// GetPackingOptionsOk returns a tuple with the PackingOptions field value
// and a boolean to check if the value has been set.
func (o *ListPackingOptionsResponse) GetPackingOptionsOk() ([]PackingOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackingOptions, true
}

// SetPackingOptions sets field value
func (o *ListPackingOptionsResponse) SetPackingOptions(v []PackingOption) {
	o.PackingOptions = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListPackingOptionsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPackingOptionsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListPackingOptionsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListPackingOptionsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListPackingOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["packingOptions"] = o.PackingOptions
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListPackingOptionsResponse struct {
	value *ListPackingOptionsResponse
	isSet bool
}

func (v NullableListPackingOptionsResponse) Get() *ListPackingOptionsResponse {
	return v.value
}

func (v *NullableListPackingOptionsResponse) Set(val *ListPackingOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPackingOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPackingOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPackingOptionsResponse(val *ListPackingOptionsResponse) *NullableListPackingOptionsResponse {
	return &NullableListPackingOptionsResponse{value: val, isSet: true}
}

func (v NullableListPackingOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListPackingOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
