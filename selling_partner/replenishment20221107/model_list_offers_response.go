package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the ListOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOffersResponse{}

// ListOffersResponse The response schema for the `listOffers` operation.
type ListOffersResponse struct {
	// A list of offers.
	Offers     []ListOffersResponseOffer `json:"offers,omitempty"`
	Pagination *PaginationResponse       `json:"pagination,omitempty"`
}

// NewListOffersResponse instantiates a new ListOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOffersResponse() *ListOffersResponse {
	this := ListOffersResponse{}
	return &this
}

// NewListOffersResponseWithDefaults instantiates a new ListOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOffersResponseWithDefaults() *ListOffersResponse {
	this := ListOffersResponse{}
	return &this
}

// GetOffers returns the Offers field value if set, zero value otherwise.
func (o *ListOffersResponse) GetOffers() []ListOffersResponseOffer {
	if o == nil || IsNil(o.Offers) {
		var ret []ListOffersResponseOffer
		return ret
	}
	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersResponse) GetOffersOk() ([]ListOffersResponseOffer, bool) {
	if o == nil || IsNil(o.Offers) {
		return nil, false
	}
	return o.Offers, true
}

// HasOffers returns a boolean if a field has been set.
func (o *ListOffersResponse) HasOffers() bool {
	if o != nil && !IsNil(o.Offers) {
		return true
	}

	return false
}

// SetOffers gets a reference to the given []ListOffersResponseOffer and assigns it to the Offers field.
func (o *ListOffersResponse) SetOffers(v []ListOffersResponseOffer) {
	o.Offers = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListOffersResponse) GetPagination() PaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOffersResponse) GetPaginationOk() (*PaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListOffersResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationResponse and assigns it to the Pagination field.
func (o *ListOffersResponse) SetPagination(v PaginationResponse) {
	o.Pagination = &v
}

func (o ListOffersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offers) {
		toSerialize["offers"] = o.Offers
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListOffersResponse struct {
	value *ListOffersResponse
	isSet bool
}

func (v NullableListOffersResponse) Get() *ListOffersResponse {
	return v.value
}

func (v *NullableListOffersResponse) Set(val *ListOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOffersResponse(val *ListOffersResponse) *NullableListOffersResponse {
	return &NullableListOffersResponse{value: val, isSet: true}
}

func (v NullableListOffersResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
