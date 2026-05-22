package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ListShipmentContentUpdatePreviewsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListShipmentContentUpdatePreviewsResponse{}

// ListShipmentContentUpdatePreviewsResponse The `ListShipmentContentUpdatePreviews` response.
type ListShipmentContentUpdatePreviewsResponse struct {
	// A list of content update previews in a shipment.
	ContentUpdatePreviews []ContentUpdatePreview `json:"contentUpdatePreviews"`
	Pagination            *Pagination            `json:"pagination,omitempty"`
}

type _ListShipmentContentUpdatePreviewsResponse ListShipmentContentUpdatePreviewsResponse

// NewListShipmentContentUpdatePreviewsResponse instantiates a new ListShipmentContentUpdatePreviewsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListShipmentContentUpdatePreviewsResponse(contentUpdatePreviews []ContentUpdatePreview) *ListShipmentContentUpdatePreviewsResponse {
	this := ListShipmentContentUpdatePreviewsResponse{}
	this.ContentUpdatePreviews = contentUpdatePreviews
	return &this
}

// NewListShipmentContentUpdatePreviewsResponseWithDefaults instantiates a new ListShipmentContentUpdatePreviewsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListShipmentContentUpdatePreviewsResponseWithDefaults() *ListShipmentContentUpdatePreviewsResponse {
	this := ListShipmentContentUpdatePreviewsResponse{}
	return &this
}

// GetContentUpdatePreviews returns the ContentUpdatePreviews field value
func (o *ListShipmentContentUpdatePreviewsResponse) GetContentUpdatePreviews() []ContentUpdatePreview {
	if o == nil {
		var ret []ContentUpdatePreview
		return ret
	}

	return o.ContentUpdatePreviews
}

// GetContentUpdatePreviewsOk returns a tuple with the ContentUpdatePreviews field value
// and a boolean to check if the value has been set.
func (o *ListShipmentContentUpdatePreviewsResponse) GetContentUpdatePreviewsOk() ([]ContentUpdatePreview, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentUpdatePreviews, true
}

// SetContentUpdatePreviews sets field value
func (o *ListShipmentContentUpdatePreviewsResponse) SetContentUpdatePreviews(v []ContentUpdatePreview) {
	o.ContentUpdatePreviews = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListShipmentContentUpdatePreviewsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShipmentContentUpdatePreviewsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListShipmentContentUpdatePreviewsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ListShipmentContentUpdatePreviewsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ListShipmentContentUpdatePreviewsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentUpdatePreviews"] = o.ContentUpdatePreviews
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListShipmentContentUpdatePreviewsResponse struct {
	value *ListShipmentContentUpdatePreviewsResponse
	isSet bool
}

func (v NullableListShipmentContentUpdatePreviewsResponse) Get() *ListShipmentContentUpdatePreviewsResponse {
	return v.value
}

func (v *NullableListShipmentContentUpdatePreviewsResponse) Set(val *ListShipmentContentUpdatePreviewsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListShipmentContentUpdatePreviewsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListShipmentContentUpdatePreviewsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListShipmentContentUpdatePreviewsResponse(val *ListShipmentContentUpdatePreviewsResponse) *NullableListShipmentContentUpdatePreviewsResponse {
	return &NullableListShipmentContentUpdatePreviewsResponse{value: val, isSet: true}
}

func (v NullableListShipmentContentUpdatePreviewsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListShipmentContentUpdatePreviewsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
