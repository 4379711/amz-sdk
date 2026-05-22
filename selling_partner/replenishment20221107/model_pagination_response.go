package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the PaginationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationResponse{}

// PaginationResponse Use these parameters to paginate through the response.
type PaginationResponse struct {
	// Total number of results matching the given filter criteria.
	TotalResults *int64 `json:"totalResults,omitempty"`
}

// NewPaginationResponse instantiates a new PaginationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationResponse() *PaginationResponse {
	this := PaginationResponse{}
	return &this
}

// NewPaginationResponseWithDefaults instantiates a new PaginationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationResponseWithDefaults() *PaginationResponse {
	this := PaginationResponse{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *PaginationResponse) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationResponse) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *PaginationResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *PaginationResponse) SetTotalResults(v int64) {
	o.TotalResults = &v
}

func (o PaginationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	return toSerialize, nil
}

type NullablePaginationResponse struct {
	value *PaginationResponse
	isSet bool
}

func (v NullablePaginationResponse) Get() *PaginationResponse {
	return v.value
}

func (v *NullablePaginationResponse) Set(val *PaginationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationResponse(val *PaginationResponse) *NullablePaginationResponse {
	return &NullablePaginationResponse{value: val, isSet: true}
}

func (v NullablePaginationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePaginationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
