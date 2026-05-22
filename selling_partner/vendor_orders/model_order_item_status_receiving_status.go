package vendor_orders

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the OrderItemStatusReceivingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemStatusReceivingStatus{}

// OrderItemStatusReceivingStatus Item receive status at the buyer's warehouse.
type OrderItemStatusReceivingStatus struct {
	// Receive status of the line item.
	ReceiveStatus    *string       `json:"receiveStatus,omitempty"`
	ReceivedQuantity *ItemQuantity `json:"receivedQuantity,omitempty"`
	// The date when the most recent item was received at the buyer's warehouse. Must be in ISO-8601 date/time format.
	LastReceiveDate *time.Time `json:"lastReceiveDate,omitempty"`
}

// NewOrderItemStatusReceivingStatus instantiates a new OrderItemStatusReceivingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemStatusReceivingStatus() *OrderItemStatusReceivingStatus {
	this := OrderItemStatusReceivingStatus{}
	return &this
}

// NewOrderItemStatusReceivingStatusWithDefaults instantiates a new OrderItemStatusReceivingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemStatusReceivingStatusWithDefaults() *OrderItemStatusReceivingStatus {
	this := OrderItemStatusReceivingStatus{}
	return &this
}

// GetReceiveStatus returns the ReceiveStatus field value if set, zero value otherwise.
func (o *OrderItemStatusReceivingStatus) GetReceiveStatus() string {
	if o == nil || IsNil(o.ReceiveStatus) {
		var ret string
		return ret
	}
	return *o.ReceiveStatus
}

// GetReceiveStatusOk returns a tuple with the ReceiveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemStatusReceivingStatus) GetReceiveStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveStatus) {
		return nil, false
	}
	return o.ReceiveStatus, true
}

// HasReceiveStatus returns a boolean if a field has been set.
func (o *OrderItemStatusReceivingStatus) HasReceiveStatus() bool {
	if o != nil && !IsNil(o.ReceiveStatus) {
		return true
	}

	return false
}

// SetReceiveStatus gets a reference to the given string and assigns it to the ReceiveStatus field.
func (o *OrderItemStatusReceivingStatus) SetReceiveStatus(v string) {
	o.ReceiveStatus = &v
}

// GetReceivedQuantity returns the ReceivedQuantity field value if set, zero value otherwise.
func (o *OrderItemStatusReceivingStatus) GetReceivedQuantity() ItemQuantity {
	if o == nil || IsNil(o.ReceivedQuantity) {
		var ret ItemQuantity
		return ret
	}
	return *o.ReceivedQuantity
}

// GetReceivedQuantityOk returns a tuple with the ReceivedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemStatusReceivingStatus) GetReceivedQuantityOk() (*ItemQuantity, bool) {
	if o == nil || IsNil(o.ReceivedQuantity) {
		return nil, false
	}
	return o.ReceivedQuantity, true
}

// HasReceivedQuantity returns a boolean if a field has been set.
func (o *OrderItemStatusReceivingStatus) HasReceivedQuantity() bool {
	if o != nil && !IsNil(o.ReceivedQuantity) {
		return true
	}

	return false
}

// SetReceivedQuantity gets a reference to the given ItemQuantity and assigns it to the ReceivedQuantity field.
func (o *OrderItemStatusReceivingStatus) SetReceivedQuantity(v ItemQuantity) {
	o.ReceivedQuantity = &v
}

// GetLastReceiveDate returns the LastReceiveDate field value if set, zero value otherwise.
func (o *OrderItemStatusReceivingStatus) GetLastReceiveDate() time.Time {
	if o == nil || IsNil(o.LastReceiveDate) {
		var ret time.Time
		return ret
	}
	return *o.LastReceiveDate
}

// GetLastReceiveDateOk returns a tuple with the LastReceiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItemStatusReceivingStatus) GetLastReceiveDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastReceiveDate) {
		return nil, false
	}
	return o.LastReceiveDate, true
}

// HasLastReceiveDate returns a boolean if a field has been set.
func (o *OrderItemStatusReceivingStatus) HasLastReceiveDate() bool {
	if o != nil && !IsNil(o.LastReceiveDate) {
		return true
	}

	return false
}

// SetLastReceiveDate gets a reference to the given time.Time and assigns it to the LastReceiveDate field.
func (o *OrderItemStatusReceivingStatus) SetLastReceiveDate(v time.Time) {
	o.LastReceiveDate = &v
}

func (o OrderItemStatusReceivingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReceiveStatus) {
		toSerialize["receiveStatus"] = o.ReceiveStatus
	}
	if !IsNil(o.ReceivedQuantity) {
		toSerialize["receivedQuantity"] = o.ReceivedQuantity
	}
	if !IsNil(o.LastReceiveDate) {
		toSerialize["lastReceiveDate"] = o.LastReceiveDate
	}
	return toSerialize, nil
}

type NullableOrderItemStatusReceivingStatus struct {
	value *OrderItemStatusReceivingStatus
	isSet bool
}

func (v NullableOrderItemStatusReceivingStatus) Get() *OrderItemStatusReceivingStatus {
	return v.value
}

func (v *NullableOrderItemStatusReceivingStatus) Set(val *OrderItemStatusReceivingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemStatusReceivingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemStatusReceivingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemStatusReceivingStatus(val *OrderItemStatusReceivingStatus) *NullableOrderItemStatusReceivingStatus {
	return &NullableOrderItemStatusReceivingStatus{value: val, isSet: true}
}

func (v NullableOrderItemStatusReceivingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderItemStatusReceivingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
