package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListPlacementOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPlacementOptionsResponse{}

// ListPlacementOptionsResponse The `listPlacementOptions` response.
type ListPlacementOptionsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Placement options generated for the inbound plan.
	PlacementOptions []PlacementOption `json:"placementOptions"`
}

type _ListPlacementOptionsResponse ListPlacementOptionsResponse

// NewListPlacementOptionsResponse instantiates a new ListPlacementOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPlacementOptionsResponse(placementOptions []PlacementOption) *ListPlacementOptionsResponse {
	this := ListPlacementOptionsResponse{}
	this.PlacementOptions = placementOptions
	return &this
}

// NewListPlacementOptionsResponseWithDefaults instantiates a new ListPlacementOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPlacementOptionsResponseWithDefaults() *ListPlacementOptionsResponse {
	this := ListPlacementOptionsResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListPlacementOptionsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlacementOptionsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListPlacementOptionsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListPlacementOptionsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetPlacementOptions returns the PlacementOptions field value
func (o *ListPlacementOptionsResponse) GetPlacementOptions() []PlacementOption {
	if o == nil {
		var ret []PlacementOption
		return ret
	}

	return o.PlacementOptions
}

// GetPlacementOptionsOk returns a tuple with the PlacementOptions field value
// and a boolean to check if the value has been set.
func (o *ListPlacementOptionsResponse) GetPlacementOptionsOk() ([]PlacementOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlacementOptions, true
}

// SetPlacementOptions sets field value
func (o *ListPlacementOptionsResponse) SetPlacementOptions(v []PlacementOption) {
	o.PlacementOptions = v
}

func (o ListPlacementOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["placementOptions"] = o.PlacementOptions
	return toSerialize, nil
}

type NullableListPlacementOptionsResponse struct {
	value *ListPlacementOptionsResponse
	isSet bool
}

func (v NullableListPlacementOptionsResponse) Get() *ListPlacementOptionsResponse {
	return v.value
}

func (v *NullableListPlacementOptionsResponse) Set(val *ListPlacementOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPlacementOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPlacementOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPlacementOptionsResponse(val *ListPlacementOptionsResponse) *NullableListPlacementOptionsResponse {
	return &NullableListPlacementOptionsResponse{value: val, isSet: true}
}

func (v NullableListPlacementOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListPlacementOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
