package orders_20260101

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the SearchOrdersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchOrdersResponse{}

// SearchOrdersResponse A list of orders.
type SearchOrdersResponse struct {
	// An array containing all orders that match the search criteria.
	Orders     []Order     `json:"orders"`
	Pagination *Pagination `json:"pagination,omitempty"`
	// Only orders updated before the specified time are returned. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
	LastUpdatedBefore *time.Time `json:"lastUpdatedBefore,omitempty"`
	// Only orders placed before the specified time are returned. The date must be in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) format.
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
}

// NewSearchOrdersResponse instantiates a new SearchOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchOrdersResponse(orders []Order) *SearchOrdersResponse {
	this := SearchOrdersResponse{}
	this.Orders = orders
	return &this
}

// NewSearchOrdersResponseWithDefaults instantiates a new SearchOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchOrdersResponseWithDefaults() *SearchOrdersResponse {
	this := SearchOrdersResponse{}
	return &this
}

// GetOrders returns the Orders field value
func (o *SearchOrdersResponse) GetOrders() []Order {
	if o == nil {
		var ret []Order
		return ret
	}

	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value
// and a boolean to check if the value has been set.
func (o *SearchOrdersResponse) GetOrdersOk() ([]Order, bool) {
	if o == nil {
		return nil, false
	}
	return o.Orders, true
}

// SetOrders sets field value
func (o *SearchOrdersResponse) SetOrders(v []Order) {
	o.Orders = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *SearchOrdersResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrdersResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *SearchOrdersResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *SearchOrdersResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetLastUpdatedBefore returns the LastUpdatedBefore field value if set, zero value otherwise.
func (o *SearchOrdersResponse) GetLastUpdatedBefore() time.Time {
	if o == nil || IsNil(o.LastUpdatedBefore) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedBefore
}

// GetLastUpdatedBeforeOk returns a tuple with the LastUpdatedBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrdersResponse) GetLastUpdatedBeforeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedBefore) {
		return nil, false
	}
	return o.LastUpdatedBefore, true
}

// HasLastUpdatedBefore returns a boolean if a field has been set.
func (o *SearchOrdersResponse) HasLastUpdatedBefore() bool {
	if o != nil && !IsNil(o.LastUpdatedBefore) {
		return true
	}

	return false
}

// SetLastUpdatedBefore gets a reference to the given time.Time and assigns it to the LastUpdatedBefore field.
func (o *SearchOrdersResponse) SetLastUpdatedBefore(v time.Time) {
	o.LastUpdatedBefore = &v
}

// GetCreatedBefore returns the CreatedBefore field value if set, zero value otherwise.
func (o *SearchOrdersResponse) GetCreatedBefore() time.Time {
	if o == nil || IsNil(o.CreatedBefore) {
		var ret time.Time
		return ret
	}
	return *o.CreatedBefore
}

// GetCreatedBeforeOk returns a tuple with the CreatedBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrdersResponse) GetCreatedBeforeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedBefore) {
		return nil, false
	}
	return o.CreatedBefore, true
}

// HasCreatedBefore returns a boolean if a field has been set.
func (o *SearchOrdersResponse) HasCreatedBefore() bool {
	if o != nil && !IsNil(o.CreatedBefore) {
		return true
	}

	return false
}

// SetCreatedBefore gets a reference to the given time.Time and assigns it to the CreatedBefore field.
func (o *SearchOrdersResponse) SetCreatedBefore(v time.Time) {
	o.CreatedBefore = &v
}

func (o SearchOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orders"] = o.Orders
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.LastUpdatedBefore) {
		toSerialize["lastUpdatedBefore"] = o.LastUpdatedBefore
	}
	if !IsNil(o.CreatedBefore) {
		toSerialize["createdBefore"] = o.CreatedBefore
	}
	return toSerialize, nil
}

type NullableSearchOrdersResponse struct {
	value *SearchOrdersResponse
	isSet bool
}

func (v NullableSearchOrdersResponse) Get() *SearchOrdersResponse {
	return v.value
}

func (v *NullableSearchOrdersResponse) Set(val *SearchOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchOrdersResponse(val *SearchOrdersResponse) *NullableSearchOrdersResponse {
	return &NullableSearchOrdersResponse{value: val, isSet: true}
}

func (v NullableSearchOrdersResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSearchOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
