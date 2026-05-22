package vendor_orders

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the OrderedQuantityDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderedQuantityDetails{}

// OrderedQuantityDetails Details of item quantity ordered.
type OrderedQuantityDetails struct {
	// The date when the line item quantity was updated by buyer. Must be in ISO-8601 date/time format.
	UpdatedDate       *time.Time    `json:"updatedDate,omitempty"`
	OrderedQuantity   *ItemQuantity `json:"orderedQuantity,omitempty"`
	CancelledQuantity *ItemQuantity `json:"cancelledQuantity,omitempty"`
}

// NewOrderedQuantityDetails instantiates a new OrderedQuantityDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderedQuantityDetails() *OrderedQuantityDetails {
	this := OrderedQuantityDetails{}
	return &this
}

// NewOrderedQuantityDetailsWithDefaults instantiates a new OrderedQuantityDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderedQuantityDetailsWithDefaults() *OrderedQuantityDetails {
	this := OrderedQuantityDetails{}
	return &this
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *OrderedQuantityDetails) GetUpdatedDate() time.Time {
	if o == nil || IsNil(o.UpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderedQuantityDetails) GetUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedDate) {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *OrderedQuantityDetails) HasUpdatedDate() bool {
	if o != nil && !IsNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given time.Time and assigns it to the UpdatedDate field.
func (o *OrderedQuantityDetails) SetUpdatedDate(v time.Time) {
	o.UpdatedDate = &v
}

// GetOrderedQuantity returns the OrderedQuantity field value if set, zero value otherwise.
func (o *OrderedQuantityDetails) GetOrderedQuantity() ItemQuantity {
	if o == nil || IsNil(o.OrderedQuantity) {
		var ret ItemQuantity
		return ret
	}
	return *o.OrderedQuantity
}

// GetOrderedQuantityOk returns a tuple with the OrderedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderedQuantityDetails) GetOrderedQuantityOk() (*ItemQuantity, bool) {
	if o == nil || IsNil(o.OrderedQuantity) {
		return nil, false
	}
	return o.OrderedQuantity, true
}

// HasOrderedQuantity returns a boolean if a field has been set.
func (o *OrderedQuantityDetails) HasOrderedQuantity() bool {
	if o != nil && !IsNil(o.OrderedQuantity) {
		return true
	}

	return false
}

// SetOrderedQuantity gets a reference to the given ItemQuantity and assigns it to the OrderedQuantity field.
func (o *OrderedQuantityDetails) SetOrderedQuantity(v ItemQuantity) {
	o.OrderedQuantity = &v
}

// GetCancelledQuantity returns the CancelledQuantity field value if set, zero value otherwise.
func (o *OrderedQuantityDetails) GetCancelledQuantity() ItemQuantity {
	if o == nil || IsNil(o.CancelledQuantity) {
		var ret ItemQuantity
		return ret
	}
	return *o.CancelledQuantity
}

// GetCancelledQuantityOk returns a tuple with the CancelledQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderedQuantityDetails) GetCancelledQuantityOk() (*ItemQuantity, bool) {
	if o == nil || IsNil(o.CancelledQuantity) {
		return nil, false
	}
	return o.CancelledQuantity, true
}

// HasCancelledQuantity returns a boolean if a field has been set.
func (o *OrderedQuantityDetails) HasCancelledQuantity() bool {
	if o != nil && !IsNil(o.CancelledQuantity) {
		return true
	}

	return false
}

// SetCancelledQuantity gets a reference to the given ItemQuantity and assigns it to the CancelledQuantity field.
func (o *OrderedQuantityDetails) SetCancelledQuantity(v ItemQuantity) {
	o.CancelledQuantity = &v
}

func (o OrderedQuantityDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpdatedDate) {
		toSerialize["updatedDate"] = o.UpdatedDate
	}
	if !IsNil(o.OrderedQuantity) {
		toSerialize["orderedQuantity"] = o.OrderedQuantity
	}
	if !IsNil(o.CancelledQuantity) {
		toSerialize["cancelledQuantity"] = o.CancelledQuantity
	}
	return toSerialize, nil
}

type NullableOrderedQuantityDetails struct {
	value *OrderedQuantityDetails
	isSet bool
}

func (v NullableOrderedQuantityDetails) Get() *OrderedQuantityDetails {
	return v.value
}

func (v *NullableOrderedQuantityDetails) Set(val *OrderedQuantityDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderedQuantityDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderedQuantityDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderedQuantityDetails(val *OrderedQuantityDetails) *NullableOrderedQuantityDetails {
	return &NullableOrderedQuantityDetails{value: val, isSet: true}
}

func (v NullableOrderedQuantityDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderedQuantityDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
